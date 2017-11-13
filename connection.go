package cri

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

const (
	MaxIntValue    = 1<<31 - 1
	DefaultAddress = "http://127.0.0.1:9222"
)

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

type CommandResponse struct {
	ID     int                    `json:"id"`
	Method string                 `json:"method"`
	Result map[string]interface{} `json:"result"`
	Error  *struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}

type CommandRequest struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`

	resChn chan CommandResponse
	errChn chan error
}

type EventRequest struct {
	Method string

	eventChn chan CommandResponse
}

type Connection struct {
	addr, wsAddr string
	conn         *websocket.Conn

	counter     int
	counterLock sync.RWMutex
	reqChn      chan CommandRequest
	responseMap sync.Map
	eventMap    sync.Map
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
		addr:   opt.Addr,
		wsAddr: opt.SocketAddr,
		conn:   conn,
		reqChn: make(chan CommandRequest),
	}
	go instance.reader()
	go instance.writer()

	return instance, nil
}

func (c *Connection) Send(command string, request, response interface{}) error {
	cmd := CommandRequest{
		ID:     c.id(),
		Method: command,
		Params: request,
		errChn: make(chan error),
		resChn: make(chan CommandResponse),
	}
	defer close(cmd.errChn)
	defer close(cmd.resChn)

	return c.cmd(cmd, response)
}

func (c *Connection) On(event string, closeChn chan struct{}) func(params interface{}) error {
	eve := EventRequest{
		Method:   event,
		eventChn: make(chan CommandResponse),
	}
	eveMap, eok := c.eventMap.Load(eve.Method)
	if !eok {
		c.eventMap.Store(eve.Method, sync.Map{})
	}
	if val, ok := eveMap.(sync.Map); ok {
		val.Store(eve, true)
		eveMap = val
	} else {
		val = sync.Map{}
		val.Store(eve, true)
		eveMap = val
	}
	c.eventMap.Store(eve.Method, eveMap)
	defer func() {
		go func() {
			<-closeChn
			if eveMap, eok := c.eventMap.Load(eve.Method); eok {
				if rmap, rok := eveMap.(sync.Map); rok {
					rmap.Delete(eve)
				}
			}
		}()
	}()

	return func(params interface{}) error {
		res := <-eve.eventChn
		if params != nil {
			return mapstructure.Decode(res.Result, params)
		}
		return nil
	}
}

func (c *Connection) Close() error {
	return nil
}

func (c *Connection) id() int {
	c.counterLock.Lock()
	defer c.counterLock.Unlock()
	nextCount := c.counter + 1
	if nextCount == MaxIntValue {
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
				continue
			}

			req.errChn <- nil
			c.responseMap.Store(req.ID, req)
		}
	}
}

func (c *Connection) reader() {
	for {
		var data CommandResponse
		if decodeErr := c.conn.ReadJSON(&data); decodeErr != nil {
			fmt.Println(decodeErr.Error())
		}

		if data.ID > 0 {
			var err error

			if val, vok := c.responseMap.Load(data.ID); vok {
				if req, rok := val.(CommandRequest); rok {
					if data.Error != nil {
						err = errors.New(data.Error.Message)
					}
					req.resChn <- data
					req.errChn <- err
				}
			}

			c.responseMap.Delete(data.ID)
		} else if len(data.Method) > 0 {
			if val, vok := c.eventMap.Load(data.Method); vok {
				if eve, eok := val.(sync.Map); eok {
					eve.Range(func(key, val interface{}) bool {
						if kval, kok := key.(EventRequest); kok {
							kval.eventChn <- data
						}
						return true
					})
				}
			}
		}
	}
}

func (c *Connection) cmd(req CommandRequest, res interface{}) error {
	if req.Params == nil {
		req.Params = struct{}{}
	}

	c.reqChn <- req
	if reqErr := <-req.errChn; reqErr != nil {
		return reqErr
	}

	cmdRes := <-req.resChn
	if resErr := <-req.errChn; resErr != nil {
		return resErr
	}

	if res == nil {
		return nil
	}

	return mapstructure.Decode(cmdRes.Result, res)
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
