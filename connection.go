package cri

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

// DefaultAddress for remote debugging protocol
const DefaultAddress = "127.0.0.1:9222"

// DefaultCommandTimeout specifies default duration to receive command response
const DefaultCommandTimeout = 10 * time.Second

const maxIntValue = 1<<31 - 1

// ConnectionOption defines a function type to set values of ConnectionOptions
type ConnectionOption func(*ConnectionOptions)

// ConnectionOptions defines connection parameters
type ConnectionOptions struct {
	// Address of remote devtools instance, default used is DefaultAddress
	Address string
	// TargetID of target to connect
	TargetID string
	// SocketAddress of target to connect
	SocketAddress string
	// TLSClientConfig specifies TLS configuration to use
	TLSConfig *tls.Config
	// CommandTimeout specifies duration to receive command response, default used is DefaultCommandTimeout
	CommandTimeout time.Duration
	// Error if provided, is called with errors from connection reader.
	Error func(err error)
}

// option iterates over all arguments to set final options
func (co *ConnectionOptions) option(opts ...ConnectionOption) {
	for i := 0; i < len(opts); i++ {
		opts[i](co)
	}
}

// SetAddress sets remote address for connection
func SetAddress(addr string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.Address = addr
	}
}

// SetSocketAddress sets target websocket connection address
func SetSocketAddress(addr string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.SocketAddress = addr
	}
}

// SetTargetID sets target for connection
func SetTargetID(targetID string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.TargetID = targetID
	}
}

// SetTLSConfig sets tls config for connection
func SetTLSConfig(config *tls.Config) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.TLSConfig = config
	}
}

// SetEventTimeout sets eventTimeout for connection
func SetEventTimeout(timeout time.Duration) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.EventTimeout = timeout
	}
}

// SetCommandTimeout sets commandTimeout for connection
func SetCommandTimeout(timeout time.Duration) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.CommandTimeout = timeout
	}
}

// SetLogger sets logging for connection
func SetError(fn func(err error)) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.Error = fn
	}
}

var (
	ErrConnectionClosed = errors.New("Connection closed")
)

type closeStore struct {
	sync.RWMutex
	closed bool
}

func (cs *closeStore) markClosed() {
	cs.Lock()
	defer cs.Unlock()
	cs.closed = true
}

func (cs *closeStore) isClosed() bool {
	cs.RLock()
	defer cs.RUnlock()
	return cs.closed
}

// Connection contains socket connection to remote target. Its safe to share
// same instance among multiple goroutines.
type Connection struct {
	addr, wsAddr             string
	tlsConfig                *tls.Config
	eventTimeout, cmdTimeout time.Duration

	conn *websocket.Conn

	reqChn chan commandRequest

	closeStore

	counter int32

	errorFn func(err error)
}

// NewConnection creates a connection to remote target. Connection options can be
// given in arguments using SetAddress, SetTargetID, SetSocketAddress etc.
// To see all options check ConnectionOptions.
//
// When no arguments are provided, connection is created using DefaultAddress.
// If TargetID is provided, connection looks for remote target and connects to it,
// incase the target is not found, first target returned by remote is used.
// If SocketAddress is provided then Address and TargetID are ignored.
func NewConnection(opts ...ConnectionOption) (*Connection, error) {
	opt := &ConnectionOptions{
		Address:        DefaultAddress,
		CommandTimeout: DefaultCommandTimeout,
		Error:          nopError,
	}
	opt.option(opts...)

	instance := &Connection{
		eventTimeout: opt.EventTimeout,
		cmdTimeout:   opt.CommandTimeout,
		tlsConfig:    opt.TLSConfig,
		reqChn:       make(chan commandRequest),
		errorFn:      opt.Error,
	}

	if len(opt.SocketAddress) != 0 {
		parsedAddr, addrParseErr := url.Parse(opt.SocketAddress)
		if addrParseErr != nil {
			return nil, addrParseErr
		}

		instance.wsAddr = parsedAddr.String()
		instance.addr = parsedAddr.Host
	} else {
		targets, targetsErr := GetTargets(opts...)
		if targetsErr != nil {
			return nil, targetsErr
		}

		wsAddr := targets[0].WebSocketDebuggerURL
		if len(opt.TargetID) != 0 {
			for i := range targets {
				if targets[i].ID == opt.TargetID {
					wsAddr = targets[i].WebSocketDebuggerURL
					break
				}
			}
		}

		instance.wsAddr = wsAddr
		instance.addr = opt.Address
	}

	dialer := &websocket.Dialer{TLSClientConfig: instance.tlsConfig}
	conn, res, resErr := dialer.Dial(instance.wsAddr, nil)
	if resErr != nil {
		return nil, resErr
	}
	if res.StatusCode != 101 {
		return nil, errors.New("Invalid status code")
	}

	instance.conn = conn

	go instance.runner()

	return instance, nil
}

type commandResponse struct {
	ID     int32                  `json:"id"`     // id of the command request or 0 in case of event
	Method string                 `json:"method"` // command or event name
	Result map[string]interface{} `json:"result"` // parameters returned for command
	Error  *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"` // error returned for command
	Params map[string]interface{} `json:"params"` // parameters returned for event
}

type commandRequest struct {
	ID     int32       `json:"id"`     // id of the command request, should be greater than 0
	Method string      `json:"method"` // method takes the command to be sent
	Params interface{} `json:"params"` // params contains parameters taken by command

	resChn       chan commandResponse
	errChn       chan error
	timeoutTimer *time.Timer
}

// Send writes command and associated parameters to underlying connection.
// It waits for command response and decodes it in response argument.
// Timeout error is returned if response is not received.
func (c *Connection) Send(command string, request, response interface{}) error {
	if c.isClosed() {
		return errors.New("connection closed")
	}

	cmd := commandRequest{
		ID:           c.id(),
		Method:       command,
		Params:       request,
		resChn:       make(chan commandResponse, 1),
		errChn:       make(chan error, 1),
		timeoutTimer: time.NewTimer(c.cmdTimeout),
	}
	if cmd.Params == nil {
		cmd.Params = struct{}{}
	}

	c.reqChn <- cmd

	select {
	case cmdRes := <-cmd.resChn:
		cmd.timeoutTimer.Stop()
		if response == nil {
			return nil
		}
		return mapstructure.Decode(cmdRes.Result, response)
	case resErr := <-cmd.errChn:
		cmd.timeoutTimer.Stop()
		return resErr
	case <-cmd.timeoutTimer.C:
		return errors.New("command response timeout")
	}
}

type eventRequest struct {
	Method   string
	eventChn chan commandResponse
	errChn   chan error
}

// On listens for subscribed event. It takes event name and a non nil channel as arguments.
// It returns a function, which blocks current routine till channel is closed.
// When event is received, parameters are decoded in params argument or error is returned.
func (c *Connection) On(event string) (listen func(params interface{}) error, close func()) {
	eve := eventRequest{
		Method:   event,
		eventChn: make(chan commandResponse, 1),
		errChn:   make(chan error, 1),
	}
	c.events.addEvent(eve)

	listen = func(params interface{}) error {
		timeoutTimer := time.NewTimer(c.eventTimeout)
		select {
		case pRes := <-eve.eventChn:
			timeoutTimer.Stop()
			if params == nil {
				return nil
			}
			return mapstructure.Decode(pRes.Result, params)
		case eveErr := <-eve.errChn:
			timeoutTimer.Stop()
			return eveErr
		case <-timeoutTimer.C:
			return errors.New("event response timeout")
		}
	}
	close = func() {
		c.events.deleteEvent(eve)
	}
	return
}

// Close stops websocket connection.
func (c *Connection) Close() error {
	c.clean()
	if closeErr := c.conn.Close(); closeErr != nil {
		return closeErr
	}

	return nil
}

func (c *Connection) id() int32 {
	if atomic.CompareAndSwapInt32(&c.counter, maxIntValue, 1) {
		return 1
	} else {
		return atomic.AddInt32(&c.counter, 1)
	}
}

func (c *Connection) runner() {
	writeChn := make(chan commandRequest, 10)
	readChn := make(chan commandResponse, 10)

	go c.writer(writeChn)
	go c.reader(readChn)

	commands := make(map[int32]commandRequest)

	for {
		select {
		case req, ok := <-c.reqChn:
			if ok {
				commands[req.ID] = req
				writeChn <- req
			}
		case res, ok := <-readChn:
			if res.ID > 0 {
				if cmd, ok := commands[res.ID]; ok {
					if res.Error != nil {
						cmd.errChn <- errors.New(res.Error.Message)
					} else {
						cmd.resChn <- res
					}
					delete(commands, res.ID)
				}
			} else if len(res.Method) > 0 {
				c.events.forEvents(res.Method, func(val eventRequest) {
					val.eventChn <- res
				})
			}
		}
	}
}

func (c *Connection) writer(chn chan commandRequest) {
	for req := range chn {
		// NOTE writer doesn't throw error on connection close.
		if writeErr := c.conn.WriteJSON(req); writeErr != nil {
			req.errChn <- writeErr
		}
	}
}

func (c *Connection) reader(chn chan commandResponse) {
	for !c.isClosed() {
		var data commandResponse
		if decodeErr := c.conn.ReadJSON(&data); decodeErr != nil {
			if err, ok := decodeErr.(*websocket.CloseError); ok {
				c.errorFn(err)
				c.clean()
				return
			}
			c.errorFn(decodeErr)
		}

	}
}

func (c *Connection) clean() {
	c.markClosed()
	c.commands.Lock()
	for id, cmd := range c.commands.commands {
		cmd.errChn <- errors.New("connection closed")
		delete(c.commands.commands, id)
	}
	c.commands.Unlock()
	c.events.Lock()
	for id, list := range c.events.events {
		for eve := range list {
			eve.errChn <- errors.New("connection closed")
			delete(c.events.events[id], eve)
		}
	}
	c.events.Unlock()
}

// VersionResponse contains fields received in response to version query.
type VersionResponse struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// GetVersion fetches the protocol version of remote
func GetVersion(opts ...ConnectionOption) (VersionResponse, error) {
	opt := &ConnectionOptions{
		Address: DefaultAddress,
	}
	opt.option(opts...)

	client := &http.Client{}
	if opt.TLSConfig != nil {
		client.Transport = &http.Transport{
			TLSClientConfig: opt.TLSConfig,
		}
	}

	var data VersionResponse
	res, resErr := client.Get(getAddress(opt) + "/json/version")
	if resErr != nil {
		return data, resErr
	}
	if res.StatusCode != 200 {
		return data, errors.New("Invalid status code")
	}
	if decodeErr := json.NewDecoder(res.Body).Decode(&data); decodeErr != nil {
		return data, decodeErr
	}

	return data, nil
}

// Targets represent a list of connectable remotes
type Targets []struct {
	Description          string `json:"description"`
	DevtoolsFrontendURL  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// GetTargets retreives targets in remote connection
func GetTargets(opts ...ConnectionOption) (Targets, error) {
	opt := &ConnectionOptions{
		Address: DefaultAddress,
	}
	opt.option(opts...)

	client := &http.Client{}
	if opt.TLSConfig != nil {
		client.Transport = &http.Transport{
			TLSClientConfig: opt.TLSConfig,
		}
	}

	var data Targets
	res, resErr := client.Get(getAddress(opt) + "/json/list")
	if resErr != nil {
		return data, resErr
	}
	if res.StatusCode != 200 {
		return data, errors.New("Invalid status code")
	}
	if decodeErr := json.NewDecoder(res.Body).Decode(&data); decodeErr != nil {
		return data, decodeErr
	}
	if len(data) < 1 {
		return data, errors.New("No valid targets")
	}

	return data, nil
}

// getAddress adds https if tls config is present
func getAddress(opts *ConnectionOptions) string {
	if opts.TLSConfig != nil {
		return "https://" + opts.Address
	} else {
		return "http://" + opts.Address
	}
}

func nopError(err error) {}
