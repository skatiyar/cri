package types

type DOMDebugger_DOMBreakpointType string
type DOMDebugger_EventListener struct {
	Type		string			`json:"type"`// <code>EventListener</code>'s type.
	UseCapture	bool			`json:"useCapture"`// <code>EventListener</code>'s useCapture.
	Passive		bool			`json:"passive"`// <code>EventListener</code>'s passive flag.
	Once		bool			`json:"once"`// <code>EventListener</code>'s once flag.
	ScriptId	Runtime_ScriptId	`json:"scriptId"`// Script id of the handler code.
	LineNumber	int			`json:"lineNumber"`// Line number in the script (0-based).
	ColumnNumber	int			`json:"columnNumber"`// Column number in the script (0-based).
	Handler		*Runtime_RemoteObject	`json:"handler,omitempty"`// Event handler function value.
	OriginalHandler	*Runtime_RemoteObject	`json:"originalHandler,omitempty"`// Event original handler function value.
	BackendNodeId	*DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Node the listener is added to (if any).
}
