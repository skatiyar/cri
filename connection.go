package cri

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/sync/syncmap"
)

// DefaultAddress for remote debugging protocol
const DefaultAddress = "127.0.0.1:9222"

// DefaultEventTimeout specifies default duration to receive an event
const DefaultEventTimeout = 2 * time.Minute

// DefaultCommandTimeout specifies default duration to receive command response
const DefaultCommandTimeout = 2 * time.Minute

const maxIntValue = 1<<31 - 1

// ConnectionOption defines a function type to set values of ConnectionOptions
type ConnectionOption func(*ConnectionOptions)

// ConnectionOptions defines the connection parameters
type ConnectionOptions struct {
	// Address of the remote devtools instance, default used is DefaultAddress
	Address string
	// TargetID of the target to connect
	TargetID string
	// SocketAddress of the target to connect
	SocketAddress string
	// TLSClientConfig specifies the TLS configuration to use, default is nil
	TLSConfig *tls.Config
	// EventTimeout specifies the duration to receive an event, default used is DefaultEventTimeout
	EventTimeout time.Duration
	// CommandTimeout specifies the duration to receive command response, default used is DefaultCommandTimeout
	CommandTimeout time.Duration
}

// option iterates over all arguments to set the final options
func (co *ConnectionOptions) option(opts ...ConnectionOption) {
	for i := 0; i < len(opts); i++ {
		opts[i](co)
	}
}

// SetAddress sets the remote address for the connection
func SetAddress(addr string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.Address = addr
	}
}

// SetSocketAddress sets the target websocket connection address
func SetSocketAddress(addr string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.SocketAddress = addr
	}
}

// SetTargetID sets the target for connection
func SetTargetID(targetID string) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.TargetID = targetID
	}
}

// SetTLSConfig sets the tls config for the connection
func SetTLSConfig(config *tls.Config) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.TLSConfig = config
	}
}

// SetEventTimeout sets the eventTimeout for connection
func SetEventTimeout(timeout time.Duration) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.EventTimeout = timeout
	}
}

// SetCommandTimeout sets the commandTimeout for connection
func SetCommandTimeout(timeout time.Duration) ConnectionOption {
	return func(co *ConnectionOptions) {
		co.CommandTimeout = timeout
	}
}

type eventsStore struct {
	sync.RWMutex
	events map[string]map[eventRequest]bool
}

func (es *eventsStore) addEvent(ereq eventRequest) {
	es.Lock()
	defer es.Unlock()
	if es.events == nil {
		es.events = make(map[string]map[eventRequest]bool)
	}
	if _, ok := es.events[ereq.Method]; !ok {
		es.events[ereq.Method] = make(map[eventRequest]bool)
	}
	es.events[ereq.Method][ereq] = true
}

func (es *eventsStore) deleteEvent(ereq eventRequest) {
	es.Lock()
	defer es.Unlock()
	if es.events == nil {
		return
	}
	if _, ok := es.events[ereq.Method]; !ok {
		return
	}
	delete(es.events[ereq.Method], ereq)
}

func (es *eventsStore) forEvents(cmd string, fn func(eventRequest)) {
	es.RLock()
	defer es.RUnlock()
	if es.events == nil {
		return
	}
	if _, ok := es.events[cmd]; !ok {
		return
	}
	for eve := range es.events[cmd] {
		go fn(eve)
	}
}

// Connection represents a connection to remote target.
// Connection is thread safe can be called from multiple go routines.
// Connection can be used to send commands or listen to events from target.
// This satisfies the Connector interface required by protocols to work.
type Connection struct {
	addr, wsAddr             string
	tlsConfig                *tls.Config
	eventTimeout, cmdTimeout time.Duration

	conn     *websocket.Conn
	reqChn   chan commandRequest
	closeChn chan struct{}

	responseMap syncmap.Map
	eventMap    eventsStore

	counter     int
	counterLock sync.RWMutex
}

func NewConnection(opts ...ConnectionOption) (*Connection, error) {
	opt := &ConnectionOptions{
		Address:        DefaultAddress,
		EventTimeout:   DefaultEventTimeout,
		CommandTimeout: DefaultCommandTimeout,
	}
	opt.option(opts...)

	instance := &Connection{
		eventTimeout: opt.EventTimeout,
		cmdTimeout:   opt.CommandTimeout,
		tlsConfig:    opt.TLSConfig,
		reqChn:       make(chan commandRequest),
		closeChn:     make(chan struct{}),
	}

	if len(opt.SocketAddress) == 0 {
		instance.addr = opt.Address
		if versionErr := instance.isPackageCompatible(); versionErr != nil {
			return nil, versionErr
		}

		if len(opt.TargetID) == 0 {
			wsAddr, wsAddrErr := instance.getDefaultSocketAddress()
			if wsAddrErr != nil {
				return nil, wsAddrErr
			}

			opt.SocketAddress = wsAddr
		} else {
			wsAddr, wsAddrErr := instance.getSocketAddressByTarget(opt.TargetID)
			if wsAddrErr != nil {
				return nil, wsAddrErr
			}

			opt.SocketAddress = wsAddr
		}
	}

	dialer := &websocket.Dialer{TLSClientConfig: opt.TLSConfig}
	conn, res, resErr := dialer.Dial(opt.SocketAddress, nil)
	if resErr != nil {
		return nil, resErr
	}
	if res.StatusCode != 101 {
		return nil, errors.New("Invalid status code")
	}

	instance.conn = conn
	instance.addr = conn.RemoteAddr().String()

	go instance.reader()
	go instance.writer()

	return instance, nil
}

type commandResponse struct {
	ID     int                    `json:"id"`     // id of the command request or 0 in case of event
	Method string                 `json:"method"` // command or event name
	Result map[string]interface{} `json:"result"`
	Params map[string]interface{} `json:"params"`
	Error  *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}

type commandRequest struct {
	ID     int         `json:"id"`     // id of the command request, should be greater than 0
	Method string      `json:"method"` // method takes the command to be sent
	Params interface{} `json:"params"` // params contains parameters taken by command

	resChn     chan commandResponse
	errChn     chan error
	timeoutChn chan bool
}

// Send sends a command and associated parameters to the target and waits for the response.
// and decodes the respose back in location provided
func (c *Connection) Send(command string, request, response interface{}) error {
	cmd := commandRequest{
		ID:         c.id(),
		Method:     command,
		Params:     request,
		resChn:     make(chan commandResponse, 1),
		errChn:     make(chan error, 1),
		timeoutChn: make(chan bool, 1),
	}

	c.responseMap.Store(cmd.ID, cmd)
	defer func() {
		c.responseMap.Delete(cmd.ID)
	}()

	go func() {
		time.Sleep(c.cmdTimeout)
		cmd.timeoutChn <- true
	}()

	if cmd.Params == nil {
		cmd.Params = struct{}{}
	}

	c.reqChn <- cmd

	select {
	case cmdRes := <-cmd.resChn:
		if response == nil {
			return nil
		}
		return mapstructure.Decode(cmdRes.Result, response)
	case resErr := <-cmd.errChn:
		return resErr
	case <-cmd.timeoutChn:
		return errors.New("command response timeout")
	}
}

type eventRequest struct {
	Method string

	eventChn chan commandResponse
}

func (c *Connection) On(event string, closeChn chan struct{}) func(params interface{}) error {
	eve := eventRequest{
		Method:   event,
		eventChn: make(chan commandResponse, 1),
	}
	c.eventMap.addEvent(eve)

	defer func() {
		go func() {
			<-closeChn
			c.eventMap.deleteEvent(eve)
		}()
	}()

	return func(params interface{}) error {
		res := <-eve.eventChn
		if params != nil {
			return mapstructure.Decode(res.Params, params)
		}
		return nil
	}
}

func (c *Connection) Close() error {
	close(c.closeChn)
	if closeErr := c.conn.Close(); closeErr != nil {
		return closeErr
	}

	return nil
}

func (c *Connection) id() int {
	c.counterLock.Lock()
	defer c.counterLock.Unlock()
	nextCount := c.counter + 1
	if nextCount == maxIntValue {
		c.counter = 1
	} else {
		c.counter = nextCount
	}
	return nextCount
}

func (c *Connection) writer() {
	for {
		select {
		case <-c.closeChn:
			return
		case req := <-c.reqChn:
			if writeErr := c.conn.WriteJSON(req); writeErr != nil {
				req.errChn <- writeErr
			}
		}
	}
}

func (c *Connection) reader() {
	for {
		select {
		case <-c.closeChn:
			return
		default:
			var data commandResponse
			if decodeErr := c.conn.ReadJSON(&data); decodeErr != nil {
				fmt.Println("---", decodeErr.Error())
			}

			if data.ID > 0 {
				if val, vok := c.responseMap.Load(data.ID); vok {
					if req, rok := val.(commandRequest); rok {
						if data.Error != nil {
							req.errChn <- errors.New(data.Error.Message)
						} else {
							req.resChn <- data
						}
					}
				}
			} else if len(data.Method) > 0 {
				c.eventMap.forEvents(data.Method, func(val eventRequest) {
					val.eventChn <- data
				})
			}
		}
	}
}

type versionResponse struct {
	Browser              string `json:"Browser"`
	ProtocolVersion      string `json:"Protocol-Version"`
	UserAgent            string `json:"User-Agent"`
	V8Version            string `json:"V8-Version"`
	WebKitVersion        string `json:"WebKit-Version"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// getVersion fetches the protocol version of remote
func (c *Connection) getVersion() (versionResponse, error) {
	client := &http.Client{}
	if c.tlsConfig != nil {
		client.Transport = &http.Transport{
			TLSClientConfig: c.tlsConfig,
		}
	}

	var data versionResponse
	res, resErr := client.Get(c.getAddress() + "/json/version")
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

// target represents a connectable resource
type target struct {
	Description          string `json:"description"`
	DevtoolsFrontendURL  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// getTargetsList retreives targets in remote connection
func (c *Connection) getTargetsList() ([]target, error) {
	client := &http.Client{}
	if c.tlsConfig != nil {
		client.Transport = &http.Transport{
			TLSClientConfig: c.tlsConfig,
		}
	}

	var data []target
	res, resErr := client.Get(c.getAddress() + "/json/list")
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
		return data, errors.New("No valid socket addresses")
	}

	return data, nil
}

// getAddress adds https if tls config is present
func (c *Connection) getAddress() string {
	if c.tlsConfig != nil {
		return "https://" + c.addr
	} else {
		return "http://" + c.addr
	}
}

// getDefaultSocketAddress picks the first element in target array
func (c *Connection) getDefaultSocketAddress() (string, error) {
	targets, targetsErr := c.getTargetsList()
	if targetsErr != nil {
		return "", targetsErr
	}

	return targets[0].WebSocketDebuggerURL, nil
}

func (c *Connection) getSocketAddressByTarget(target string) (string, error) {
	targets, targetsErr := c.getTargetsList()
	if targetsErr != nil {
		return "", targetsErr
	}

	for i := 0; i < len(targets); i++ {
		if targets[i].ID == target {
			return targets[i].WebSocketDebuggerURL, nil
		}
	}

	return "", errors.New("target not found")
}

func (c *Connection) isPackageCompatible() error {
	verData, verErr := c.getVersion()
	if verErr != nil {
		return verErr
	}

	majorV, minorV := Version()
	if verData.ProtocolVersion != (majorV + "." + minorV) {
		return errors.New("Version Mismatch")
	}

	return nil
}
