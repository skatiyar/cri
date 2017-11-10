package types

type Debugger_BreakpointId string
type Debugger_CallFrameId string
type Debugger_Location struct {
	ScriptId     Runtime_ScriptId `json:"scriptId"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber *int             `json:"columnNumber,omitempty"`
}
type Debugger_ScriptPosition struct {
	LineNumber   int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}
type Debugger_CallFrame struct {
	CallFrameId      Debugger_CallFrameId  `json:"callFrameId"`
	FunctionName     string                `json:"functionName"`
	FunctionLocation *Debugger_Location    `json:"functionLocation,omitempty"`
	Location         Debugger_Location     `json:"location"`
	Url              string                `json:"url"`
	ScopeChain       []Debugger_Scope      `json:"scopeChain"`
	This             Runtime_RemoteObject  `json:"this"`
	ReturnValue      *Runtime_RemoteObject `json:"returnValue,omitempty"`
}
type Debugger_Scope struct {
	Type          string               `json:"type"`
	Object        Runtime_RemoteObject `json:"object"`
	Name          *string              `json:"name,omitempty"`
	StartLocation *Debugger_Location   `json:"startLocation,omitempty"`
	EndLocation   *Debugger_Location   `json:"endLocation,omitempty"`
}
type Debugger_SearchMatch struct {
	LineNumber  float32 `json:"lineNumber"`
	LineContent string  `json:"lineContent"`
}
type Debugger_BreakLocation struct {
	ScriptId     Runtime_ScriptId `json:"scriptId"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber *int             `json:"columnNumber,omitempty"`
	Type         *string          `json:"type,omitempty"`
}
