package cri

import (
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
const DefaultAddress = "http://127.0.0.1:9222"

const maxIntValue = 1<<31 - 1

type ConnectionOption func(*ConnectionOpts)

type ConnectionOpts struct {
	Addr       string
	SocketAddr string
}

func (co *ConnectionOpts) Option(opts ...ConnectionOption) {
	for i := 0; i < len(opts); i++ {
		opts[i](co)
	}
}

func SetAddress(addr string) ConnectionOption {
	return func(co *ConnectionOpts) {
		co.Addr = addr
	}
}

func SetSocketAddress(addr string) ConnectionOption {
	return func(co *ConnectionOpts) {
		co.SocketAddr = addr
	}
}

type eventRequest struct {
	Method string

	eventChn chan commandResponse
}

type eventsStore struct {
	sync.RWMutex
	events map[string]map[EventRequest]bool
}

func (es *eventsStore) addEvent(ereq EventRequest) {
	es.Lock()
	defer es.Unlock()
	if es.events == nil {
		es.events = make(map[string]map[EventRequest]bool)
	}
	if _, ok := es.events[ereq.Method]; !ok {
		es.events[ereq.Method] = make(map[EventRequest]bool)
	}
	es.events[ereq.Method][ereq] = true
}

func (es *eventsStore) deleteEvent(ereq EventRequest) {
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

func (es *eventsStore) forEvents(cmd string, fn func(EventRequest)) {
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
	addr, wsAddr string
	cmdTimeout   time.Duration

	conn   *websocket.Conn
	reqChn chan commandRequest

	responseMap syncmap.Map
	eventMap    eventsStore

	counter     int
	counterLock sync.RWMutex
}

func NewConnection(opts ...ConnectionOption) (*Connection, error) {
	opt := &ConnectionOpts{
		Addr: DefaultAddress,
	}
	opt.Option(opts...)

	if len(opt.SocketAddr) == 0 {
		if versionErr := checkVersion(opt.Addr); versionErr != nil {
			return nil, versionErr
		}

		wsAddr, wsAddrErr := getSocketAddress(opt.Addr)
		if wsAddrErr != nil {
			return nil, wsAddrErr
		}

		opt.SocketAddr = wsAddr
	}

	dialer := websocket.Dialer{}
	conn, res, resErr := dialer.Dial(opt.SocketAddr, nil)
	if resErr != nil {
		return nil, resErr
	}
	if res.StatusCode != 101 {
		return nil, errors.New("Invalid status code")
	}

	instance := &Connection{
		addr:       opt.Addr,
		wsAddr:     opt.SocketAddr,
		cmdTimeout: time.Second * 5,
		conn:       conn,
		reqChn:     make(chan commandRequest),
	}
	go instance.reader()
	go instance.writer()

	return instance, nil
}

type commandResponse struct {
	ID     int                    `json:"id"`
	Method string                 `json:"method"`
	Result map[string]interface{} `json:"result"`
	Params map[string]interface{} `json:"params"`
	Error  *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}

// commandRequest
type commandRequest struct {
	ID     int         `json:"id"`     // id of the command request, should be greater than 0
	Method string      `json:"method"` // method takes
	Params interface{} `json:"params"` // params contains parameters taken by command

	resChn     chan commandResponse
	errChn     chan error
	timeoutChn chan bool
}

// Send sends a command and associated parameters to the target and waits for the response
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

func (c *Connection) On(event string, closeChn chan struct{}) func(params interface{}) error {
	eve := EventRequest{
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
		case req := <-c.reqChn:
			if writeErr := c.conn.WriteJSON(req); writeErr != nil {
				req.errChn <- writeErr
			}
		}
	}
}

func (c *Connection) reader() {
	for {
		var data commandResponse
		if decodeErr := c.conn.ReadJSON(&data); decodeErr != nil {
			fmt.Println(decodeErr.Error())
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
			c.eventMap.forEvents(data.Method, func(val EventRequest) {
				val.eventChn <- data
			})
		}
	}
}

func checkVersion(addr string) error {
	var response struct {
		Browser              string `json:"Browser"`
		ProtocolVersion      string `json:"Protocol-Version"`
		UserAgent            string `json:"User-Agent"`
		V8Version            string `json:"V8-Version"`
		WebKitVersion        string `json:"WebKit-Version"`
		WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	}

	res, resErr := http.Get(addr + "/json/version")
	if resErr != nil {
		return resErr
	}
	if res.StatusCode != 200 {
		return errors.New("Invalid status code")
	}
	if decodeErr := json.NewDecoder(res.Body).Decode(&response); decodeErr != nil {
		return decodeErr
	}

	majorV, minorV := Version()
	if response.ProtocolVersion != majorV+"."+minorV {
		return errors.New("Version Mismatch")
	}

	return nil
}

func getSocketAddress(addr string) (string, error) {
	var response []struct {
		Description          string `json:"description"`
		DevtoolsFrontendURL  string `json:"devtoolsFrontendUrl"`
		ID                   string `json:"id"`
		Title                string `json:"title"`
		Type                 string `json:"type"`
		URL                  string `json:"url"`
		WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
	}

	res, resErr := http.Get(addr + "/json/list")
	if resErr != nil {
		return "", resErr
	}
	if res.StatusCode != 200 {
		return "", errors.New("Invalid status code")
	}
	if decodeErr := json.NewDecoder(res.Body).Decode(&response); decodeErr != nil {
		return "", decodeErr
	}
	if len(response) < 1 {
		return "", errors.New("No valid socket addresses")
	}

	return response[0].WebSocketDebuggerURL, nil
}
