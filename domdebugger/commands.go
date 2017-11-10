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
	NodeId	types.DOM_NodeId			`json:"nodeId"`// Identifier of the node to set breakpoint on.
	Type	types.DOMDebugger_DOMBreakpointType	`json:"type"`// Type of the operation to stop upon.
}

func (obj *DOMDebugger) SetDOMBreakpoint(request *SetDOMBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setDOMBreakpoint", request, nil)
	return
}

type RemoveDOMBreakpointRequest struct {
	NodeId	types.DOM_NodeId			`json:"nodeId"`// Identifier of the node to remove breakpoint from.
	Type	types.DOMDebugger_DOMBreakpointType	`json:"type"`// Type of the breakpoint to remove.
}

func (obj *DOMDebugger) RemoveDOMBreakpoint(request *RemoveDOMBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeDOMBreakpoint", request, nil)
	return
}

type SetEventListenerBreakpointRequest struct {
	EventName	string	`json:"eventName"`// DOM Event name to stop on (any DOM event will do).
	TargetName	*string	`json:"targetName,omitempty"`// EventTarget interface name to stop on. If equal to <code>"*"</code> or not provided, will stop on any EventTarget.
}

func (obj *DOMDebugger) SetEventListenerBreakpoint(request *SetEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setEventListenerBreakpoint", request, nil)
	return
}

type RemoveEventListenerBreakpointRequest struct {
	EventName	string	`json:"eventName"`// Event name.
	TargetName	*string	`json:"targetName,omitempty"`// EventTarget interface name.
}

func (obj *DOMDebugger) RemoveEventListenerBreakpoint(request *RemoveEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeEventListenerBreakpoint", request, nil)
	return
}

type SetInstrumentationBreakpointRequest struct {
	EventName string `json:"eventName"`// Instrumentation name to stop on.
}

func (obj *DOMDebugger) SetInstrumentationBreakpoint(request *SetInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setInstrumentationBreakpoint", request, nil)
	return
}

type RemoveInstrumentationBreakpointRequest struct {
	EventName string `json:"eventName"`// Instrumentation name to stop on.
}

func (obj *DOMDebugger) RemoveInstrumentationBreakpoint(request *RemoveInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeInstrumentationBreakpoint", request, nil)
	return
}

type SetXHRBreakpointRequest struct {
	Url string `json:"url"`// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
}

func (obj *DOMDebugger) SetXHRBreakpoint(request *SetXHRBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.setXHRBreakpoint", request, nil)
	return
}

type RemoveXHRBreakpointRequest struct {
	Url string `json:"url"`// Resource URL substring.
}

func (obj *DOMDebugger) RemoveXHRBreakpoint(request *RemoveXHRBreakpointRequest) (err error) {
	err = obj.conn.Send("DOMDebugger.removeXHRBreakpoint", request, nil)
	return
}

type GetEventListenersRequest struct {
	ObjectId	types.Runtime_RemoteObjectId	`json:"objectId"`// Identifier of the object to return listeners for.
	Depth		*int				`json:"depth,omitempty"`// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce		*bool				`json:"pierce,omitempty"`// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
}

func (obj *DOMDebugger) GetEventListeners(request *GetEventListenersRequest) (response struct {
	Listeners []types.DOMDebugger_EventListener `json:"listeners"`// Array of relevant listeners.
}, err error) {
	err = obj.conn.Send("DOMDebugger.getEventListeners", request, &response)
	return
}
