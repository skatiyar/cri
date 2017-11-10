package types

type DOMDebugger_DOMBreakpointType string
type DOMDebugger_EventListener struct {
	Type            string                `json:"type"`
	UseCapture      bool                  `json:"useCapture"`
	Passive         bool                  `json:"passive"`
	Once            bool                  `json:"once"`
	ScriptId        Runtime_ScriptId      `json:"scriptId"`
	LineNumber      int                   `json:"lineNumber"`
	ColumnNumber    int                   `json:"columnNumber"`
	Handler         *Runtime_RemoteObject `json:"handler,omitempty"`
	OriginalHandler *Runtime_RemoteObject `json:"originalHandler,omitempty"`
	BackendNodeId   *DOM_BackendNodeId    `json:"backendNodeId,omitempty"`
}
