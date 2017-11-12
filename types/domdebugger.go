/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// DOM breakpoint type.
type DOMDebugger_DOMBreakpointType string

// Object event listener.
type DOMDebugger_EventListener struct {
	// <code>EventListener</code>'s type.
	Type string `json:"type"`
	// <code>EventListener</code>'s useCapture.
	UseCapture bool `json:"useCapture"`
	// <code>EventListener</code>'s passive flag.
	Passive bool `json:"passive"`
	// <code>EventListener</code>'s once flag.
	Once bool `json:"once"`
	// Script id of the handler code.
	ScriptId Runtime_ScriptId `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber int `json:"columnNumber"`
	// Event handler function value.
	Handler *Runtime_RemoteObject `json:"handler,omitempty"`
	// Event original handler function value.
	OriginalHandler *Runtime_RemoteObject `json:"originalHandler,omitempty"`
	// Node the listener is added to (if any).
	BackendNodeId *DOM_BackendNodeId `json:"backendNodeId,omitempty"`
}
