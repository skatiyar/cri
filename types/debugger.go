/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types


//Breakpoint identifier.
type Debugger_BreakpointId string

//Call frame identifier.
type Debugger_CallFrameId string

//Location in the source code.
type Debugger_Location struct {
	// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	ScriptId Runtime_ScriptId `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber *int `json:"columnNumber,omitempty"`
}

//Location in the source code.
type Debugger_ScriptPosition struct {
	LineNumber int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}

//JavaScript call frame. Array of call frames form the call stack.
type Debugger_CallFrame struct {
	// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	CallFrameId Debugger_CallFrameId `json:"callFrameId"`
	// Name of the JavaScript function called on this call frame.
	FunctionName string `json:"functionName"`
	// Location in the source code.
	// NOTE Experimental
	FunctionLocation *Debugger_Location `json:"functionLocation,omitempty"`
	// Location in the source code.
	Location Debugger_Location `json:"location"`
	// JavaScript script name or url.
	Url string `json:"url"`
	// Scope chain for this call frame.
	ScopeChain []Debugger_Scope `json:"scopeChain"`
	// <code>this</code> object for this call frame.
	This Runtime_RemoteObject `json:"this"`
	// The value being returned, if the function is at return point.
	ReturnValue *Runtime_RemoteObject `json:"returnValue,omitempty"`
}

//Scope description.
type Debugger_Scope struct {
	// Scope type.
	Type string `json:"type"`
	// Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Object Runtime_RemoteObject `json:"object"`
	Name *string `json:"name,omitempty"`
	// Location in the source code where scope starts
	StartLocation *Debugger_Location `json:"startLocation,omitempty"`
	// Location in the source code where scope ends
	EndLocation *Debugger_Location `json:"endLocation,omitempty"`
}

//Search match for resource.
type Debugger_SearchMatch struct {
	// Line number in resource content.
	LineNumber float32 `json:"lineNumber"`
	// Line with match content.
	LineContent string `json:"lineContent"`
}


type Debugger_BreakLocation struct {
	// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	ScriptId Runtime_ScriptId `json:"scriptId"`
	// Line number in the script (0-based).
	LineNumber int `json:"lineNumber"`
	// Column number in the script (0-based).
	ColumnNumber *int `json:"columnNumber,omitempty"`
	Type *string `json:"type,omitempty"`
}

