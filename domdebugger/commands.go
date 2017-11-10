package domdebugger

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type DOMDebugger struct {
	conn cri.Connector
}

func New(conn cri.Connector) *DOMDebugger {
	return &DOMDebugger{conn}
}

type SetDOMBreakpointRequest struct {
	NodeId types.DOM_NodeId                    `json:"nodeId"`
	Type   types.DOMDebugger_DOMBreakpointType `json:"type"`
}

func (obj *DOMDebugger) SetDOMBreakpoint(request *SetDOMBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setDOMBreakpoint", request, nil)
	return
}

type RemoveDOMBreakpointRequest struct {
	NodeId types.DOM_NodeId                    `json:"nodeId"`
	Type   types.DOMDebugger_DOMBreakpointType `json:"type"`
}

func (obj *DOMDebugger) RemoveDOMBreakpoint(request *RemoveDOMBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeDOMBreakpoint", request, nil)
	return
}

type SetEventListenerBreakpointRequest struct {
	EventName  string  `json:"eventName"`
	TargetName *string `json:"targetName,omitempty"`
}

func (obj *DOMDebugger) SetEventListenerBreakpoint(request *SetEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setEventListenerBreakpoint", request, nil)
	return
}

type RemoveEventListenerBreakpointRequest struct {
	EventName  string  `json:"eventName"`
	TargetName *string `json:"targetName,omitempty"`
}

func (obj *DOMDebugger) RemoveEventListenerBreakpoint(request *RemoveEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeEventListenerBreakpoint", request, nil)
	return
}

type SetInstrumentationBreakpointRequest struct {
	EventName string `json:"eventName"`
}

func (obj *DOMDebugger) SetInstrumentationBreakpoint(request *SetInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setInstrumentationBreakpoint", request, nil)
	return
}

type RemoveInstrumentationBreakpointRequest struct {
	EventName string `json:"eventName"`
}

func (obj *DOMDebugger) RemoveInstrumentationBreakpoint(request *RemoveInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeInstrumentationBreakpoint", request, nil)
	return
}

type SetXHRBreakpointRequest struct {
	Url string `json:"url"`
}

func (obj *DOMDebugger) SetXHRBreakpoint(request *SetXHRBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setXHRBreakpoint", request, nil)
	return
}

type RemoveXHRBreakpointRequest struct {
	Url string `json:"url"`
}

func (obj *DOMDebugger) RemoveXHRBreakpoint(request *RemoveXHRBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeXHRBreakpoint", request, nil)
	return
}

type GetEventListenersRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
	Depth    *int                         `json:"depth,omitempty"`
	Pierce   *bool                        `json:"pierce,omitempty"`
}

func (obj *DOMDebugger) GetEventListeners(request *GetEventListenersRequest) (response struct {
	Listeners []types.DOMDebugger_EventListener `json:"listeners"`
}, err error) {
	err = obj.conn.Send("DOMDebugger.getEventListeners", request, &response)
	return
}
