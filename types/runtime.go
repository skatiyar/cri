package types

type Runtime_ScriptId string
type Runtime_RemoteObjectId string
type Runtime_UnserializableValue string
type Runtime_RemoteObject struct {
	Type                string                       `json:"type"`
	Subtype             *string                      `json:"subtype,omitempty"`
	ClassName           *string                      `json:"className,omitempty"`
	Value               interface{}                  `json:"value,omitempty"`
	UnserializableValue *Runtime_UnserializableValue `json:"unserializableValue,omitempty"`
	Description         *string                      `json:"description,omitempty"`
	ObjectId            *Runtime_RemoteObjectId      `json:"objectId,omitempty"`
	Preview             *Runtime_ObjectPreview       `json:"preview,omitempty"`
	CustomPreview       *Runtime_CustomPreview       `json:"customPreview,omitempty"`
}
type Runtime_CustomPreview struct {
	Header                     string                  `json:"header"`
	HasBody                    bool                    `json:"hasBody"`
	FormatterObjectId          Runtime_RemoteObjectId  `json:"formatterObjectId"`
	BindRemoteObjectFunctionId Runtime_RemoteObjectId  `json:"bindRemoteObjectFunctionId"`
	ConfigObjectId             *Runtime_RemoteObjectId `json:"configObjectId,omitempty"`
}
type Runtime_ObjectPreview struct {
	Type        string                    `json:"type"`
	Subtype     *string                   `json:"subtype,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Overflow    bool                      `json:"overflow"`
	Properties  []Runtime_PropertyPreview `json:"properties"`
	Entries     []Runtime_EntryPreview    `json:"entries,omitempty"`
}
type Runtime_PropertyPreview struct {
	Name         string                 `json:"name"`
	Type         string                 `json:"type"`
	Value        *string                `json:"value,omitempty"`
	ValuePreview *Runtime_ObjectPreview `json:"valuePreview,omitempty"`
	Subtype      *string                `json:"subtype,omitempty"`
}
type Runtime_EntryPreview struct {
	Key   *Runtime_ObjectPreview `json:"key,omitempty"`
	Value Runtime_ObjectPreview  `json:"value"`
}
type Runtime_PropertyDescriptor struct {
	Name         string                `json:"name"`
	Value        *Runtime_RemoteObject `json:"value,omitempty"`
	Writable     *bool                 `json:"writable,omitempty"`
	Get          *Runtime_RemoteObject `json:"get,omitempty"`
	Set          *Runtime_RemoteObject `json:"set,omitempty"`
	Configurable bool                  `json:"configurable"`
	Enumerable   bool                  `json:"enumerable"`
	WasThrown    *bool                 `json:"wasThrown,omitempty"`
	IsOwn        *bool                 `json:"isOwn,omitempty"`
	Symbol       *Runtime_RemoteObject `json:"symbol,omitempty"`
}
type Runtime_InternalPropertyDescriptor struct {
	Name  string                `json:"name"`
	Value *Runtime_RemoteObject `json:"value,omitempty"`
}
type Runtime_CallArgument struct {
	Value               interface{}                  `json:"value,omitempty"`
	UnserializableValue *Runtime_UnserializableValue `json:"unserializableValue,omitempty"`
	ObjectId            *Runtime_RemoteObjectId      `json:"objectId,omitempty"`
}
type Runtime_ExecutionContextId int
type Runtime_ExecutionContextDescription struct {
	Id      Runtime_ExecutionContextId `json:"id"`
	Origin  string                     `json:"origin"`
	Name    string                     `json:"name"`
	AuxData *map[string]interface{}    `json:"auxData,omitempty"`
}
type Runtime_ExceptionDetails struct {
	ExceptionId        int                         `json:"exceptionId"`
	Text               string                      `json:"text"`
	LineNumber         int                         `json:"lineNumber"`
	ColumnNumber       int                         `json:"columnNumber"`
	ScriptId           *Runtime_ScriptId           `json:"scriptId,omitempty"`
	Url                *string                     `json:"url,omitempty"`
	StackTrace         *Runtime_StackTrace         `json:"stackTrace,omitempty"`
	Exception          *Runtime_RemoteObject       `json:"exception,omitempty"`
	ExecutionContextId *Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}
type Runtime_Timestamp float32
type Runtime_CallFrame struct {
	FunctionName string           `json:"functionName"`
	ScriptId     Runtime_ScriptId `json:"scriptId"`
	Url          string           `json:"url"`
	LineNumber   int              `json:"lineNumber"`
	ColumnNumber int              `json:"columnNumber"`
}
type Runtime_StackTrace struct {
	Description          *string             `json:"description,omitempty"`
	CallFrames           []Runtime_CallFrame `json:"callFrames"`
	Parent               *Runtime_StackTrace `json:"parent,omitempty"`
	PromiseCreationFrame *Runtime_CallFrame  `json:"promiseCreationFrame,omitempty"`
}
