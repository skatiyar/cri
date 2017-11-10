package types

type Debugger_BreakpointId string
type Debugger_CallFrameId string
type Debugger_Location struct {
	ScriptId	Runtime_ScriptId	`json:"scriptId"`// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber	int			`json:"lineNumber"`// Line number in the script (0-based).
	ColumnNumber	*int			`json:"columnNumber,omitempty"`// Column number in the script (0-based).
}
type Debugger_ScriptPosition struct {
	LineNumber	int	`json:"lineNumber"`
	ColumnNumber	int	`json:"columnNumber"`
}
type Debugger_CallFrame struct {
	CallFrameId		Debugger_CallFrameId	`json:"callFrameId"`// Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName		string			`json:"functionName"`// Name of the JavaScript function called on this call frame.
	FunctionLocation	*Debugger_Location	`json:"functionLocation,omitempty"`// Location in the source code.
	Location		Debugger_Location	`json:"location"`// Location in the source code.
	Url			string			`json:"url"`// JavaScript script name or url.
	ScopeChain		[]Debugger_Scope	`json:"scopeChain"`// Scope chain for this call frame.
	This			Runtime_RemoteObject	`json:"this"`// <code>this</code> object for this call frame.
	ReturnValue		*Runtime_RemoteObject	`json:"returnValue,omitempty"`// The value being returned, if the function is at return point.
}
type Debugger_Scope struct {
	Type		string			`json:"type"`// Scope type.
	Object		Runtime_RemoteObject	`json:"object"`// Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Name		*string			`json:"name,omitempty"`
	StartLocation	*Debugger_Location	`json:"startLocation,omitempty"`// Location in the source code where scope starts
	EndLocation	*Debugger_Location	`json:"endLocation,omitempty"`// Location in the source code where scope ends
}
type Debugger_SearchMatch struct {
	LineNumber	float32	`json:"lineNumber"`// Line number in resource content.
	LineContent	string	`json:"lineContent"`// Line with match content.
}
type Debugger_BreakLocation struct {
	ScriptId	Runtime_ScriptId	`json:"scriptId"`// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber	int			`json:"lineNumber"`// Line number in the script (0-based).
	ColumnNumber	*int			`json:"columnNumber,omitempty"`// Column number in the script (0-based).
	Type		*string			`json:"type,omitempty"`
}
