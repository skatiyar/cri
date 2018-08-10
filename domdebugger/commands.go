/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package domdebugger provides commands and events for DOMDebugger domain.
package domdebugger

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in DOMDebugger domain
const (
	GetEventListeners               = "DOMDebugger.getEventListeners"
	RemoveDOMBreakpoint             = "DOMDebugger.removeDOMBreakpoint"
	RemoveEventListenerBreakpoint   = "DOMDebugger.removeEventListenerBreakpoint"
	RemoveInstrumentationBreakpoint = "DOMDebugger.removeInstrumentationBreakpoint"
	RemoveXHRBreakpoint             = "DOMDebugger.removeXHRBreakpoint"
	SetDOMBreakpoint                = "DOMDebugger.setDOMBreakpoint"
	SetEventListenerBreakpoint      = "DOMDebugger.setEventListenerBreakpoint"
	SetInstrumentationBreakpoint    = "DOMDebugger.setInstrumentationBreakpoint"
	SetXHRBreakpoint                = "DOMDebugger.setXHRBreakpoint"
)

// DOM debugging allows setting breakpoints on particular DOM operations and events. JavaScript execution will stop on these operations as if there was a regular breakpoint set.
type DOMDebugger struct {
	conn cri.Connector
}

// New creates a DOMDebugger instance
func New(conn cri.Connector) *DOMDebugger {
	return &DOMDebugger{conn}
}

type GetEventListenersRequest struct {
	// Identifier of the object to return listeners for.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth *int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce *bool `json:"pierce,omitempty"`
}

type GetEventListenersResponse struct {
	// Array of relevant listeners.
	Listeners []types.DOMDebugger_EventListener `json:"listeners"`
}

// Returns event listeners of the given object.
func (obj *DOMDebugger) GetEventListeners(request *GetEventListenersRequest) (response GetEventListenersResponse, err error) {
	err = obj.conn.Send(GetEventListeners, request, &response)
	return
}

type RemoveDOMBreakpointRequest struct {
	// Identifier of the node to remove breakpoint from.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Type of the breakpoint to remove.
	Type types.DOMDebugger_DOMBreakpointType `json:"type"`
}

// Removes DOM breakpoint that was set using `setDOMBreakpoint`.
func (obj *DOMDebugger) RemoveDOMBreakpoint(request *RemoveDOMBreakpointRequest) (err error) {
	err = obj.conn.Send(RemoveDOMBreakpoint, request, nil)
	return
}

type RemoveEventListenerBreakpointRequest struct {
	// Event name.
	EventName string `json:"eventName"`
	// EventTarget interface name.
	// NOTE Experimental
	TargetName *string `json:"targetName,omitempty"`
}

// Removes breakpoint on particular DOM event.
func (obj *DOMDebugger) RemoveEventListenerBreakpoint(request *RemoveEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send(RemoveEventListenerBreakpoint, request, nil)
	return
}

type RemoveInstrumentationBreakpointRequest struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// Removes breakpoint on particular native event.
func (obj *DOMDebugger) RemoveInstrumentationBreakpoint(request *RemoveInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send(RemoveInstrumentationBreakpoint, request, nil)
	return
}

type RemoveXHRBreakpointRequest struct {
	// Resource URL substring.
	Url string `json:"url"`
}

// Removes breakpoint from XMLHttpRequest.
func (obj *DOMDebugger) RemoveXHRBreakpoint(request *RemoveXHRBreakpointRequest) (err error) {
	err = obj.conn.Send(RemoveXHRBreakpoint, request, nil)
	return
}

type SetDOMBreakpointRequest struct {
	// Identifier of the node to set breakpoint on.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Type of the operation to stop upon.
	Type types.DOMDebugger_DOMBreakpointType `json:"type"`
}

// Sets breakpoint on particular operation with DOM.
func (obj *DOMDebugger) SetDOMBreakpoint(request *SetDOMBreakpointRequest) (err error) {
	err = obj.conn.Send(SetDOMBreakpoint, request, nil)
	return
}

type SetEventListenerBreakpointRequest struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`
	// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any EventTarget.
	// NOTE Experimental
	TargetName *string `json:"targetName,omitempty"`
}

// Sets breakpoint on particular DOM event.
func (obj *DOMDebugger) SetEventListenerBreakpoint(request *SetEventListenerBreakpointRequest) (err error) {
	err = obj.conn.Send(SetEventListenerBreakpoint, request, nil)
	return
}

type SetInstrumentationBreakpointRequest struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// Sets breakpoint on particular native event.
func (obj *DOMDebugger) SetInstrumentationBreakpoint(request *SetInstrumentationBreakpointRequest) (err error) {
	err = obj.conn.Send(SetInstrumentationBreakpoint, request, nil)
	return
}

type SetXHRBreakpointRequest struct {
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	Url string `json:"url"`
}

// Sets breakpoint on XMLHttpRequest.
func (obj *DOMDebugger) SetXHRBreakpoint(request *SetXHRBreakpointRequest) (err error) {
	err = obj.conn.Send(SetXHRBreakpoint, request, nil)
	return
}
