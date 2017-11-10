package types

type Runtime_ScriptId string
type Runtime_RemoteObjectId string
type Runtime_UnserializableValue string
type Runtime_RemoteObject struct {
	Type			string				`json:"type"`// Object type.
	Subtype			*string				`json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
	ClassName		*string				`json:"className,omitempty"`// Object class (constructor) name. Specified for <code>object</code> type values only.
	Value			interface{}			`json:"value,omitempty"`// Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue	*Runtime_UnserializableValue	`json:"unserializableValue,omitempty"`// Primitive value which can not be JSON-stringified does not have <code>value</code>, but gets this property.
	Description		*string				`json:"description,omitempty"`// String representation of the object.
	ObjectId		*Runtime_RemoteObjectId		`json:"objectId,omitempty"`// Unique object identifier (for non-primitive values).
	Preview			*Runtime_ObjectPreview		`json:"preview,omitempty"`// Preview containing abbreviated property values. Specified for <code>object</code> type values only.
	CustomPreview		*Runtime_CustomPreview		`json:"customPreview,omitempty"`
}
type Runtime_CustomPreview struct {
	Header				string			`json:"header"`
	HasBody				bool			`json:"hasBody"`
	FormatterObjectId		Runtime_RemoteObjectId	`json:"formatterObjectId"`
	BindRemoteObjectFunctionId	Runtime_RemoteObjectId	`json:"bindRemoteObjectFunctionId"`
	ConfigObjectId			*Runtime_RemoteObjectId	`json:"configObjectId,omitempty"`
}
type Runtime_ObjectPreview struct {
	Type		string				`json:"type"`// Object type.
	Subtype		*string				`json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
	Description	*string				`json:"description,omitempty"`// String representation of the object.
	Overflow	bool				`json:"overflow"`// True iff some of the properties or entries of the original object did not fit.
	Properties	[]Runtime_PropertyPreview	`json:"properties"`// List of the properties.
	Entries		[]Runtime_EntryPreview		`json:"entries,omitempty"`// List of the entries. Specified for <code>map</code> and <code>set</code> subtype values only.
}
type Runtime_PropertyPreview struct {
	Name		string			`json:"name"`// Property name.
	Type		string			`json:"type"`// Object type. Accessor means that the property itself is an accessor property.
	Value		*string			`json:"value,omitempty"`// User-friendly property value string.
	ValuePreview	*Runtime_ObjectPreview	`json:"valuePreview,omitempty"`// Nested value preview.
	Subtype		*string			`json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
}
type Runtime_EntryPreview struct {
	Key	*Runtime_ObjectPreview	`json:"key,omitempty"`// Preview of the key. Specified for map-like collection entries.
	Value	Runtime_ObjectPreview	`json:"value"`// Preview of the value.
}
type Runtime_PropertyDescriptor struct {
	Name		string			`json:"name"`// Property name or symbol description.
	Value		*Runtime_RemoteObject	`json:"value,omitempty"`// The value associated with the property.
	Writable	*bool			`json:"writable,omitempty"`// True if the value associated with the property may be changed (data descriptors only).
	Get		*Runtime_RemoteObject	`json:"get,omitempty"`// A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
	Set		*Runtime_RemoteObject	`json:"set,omitempty"`// A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
	Configurable	bool			`json:"configurable"`// True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable	bool			`json:"enumerable"`// True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown	*bool			`json:"wasThrown,omitempty"`// True if the result was thrown during the evaluation.
	IsOwn		*bool			`json:"isOwn,omitempty"`// True if the property is owned for the object.
	Symbol		*Runtime_RemoteObject	`json:"symbol,omitempty"`// Property symbol object, if the property is of the <code>symbol</code> type.
}
type Runtime_InternalPropertyDescriptor struct {
	Name	string			`json:"name"`// Conventional property name.
	Value	*Runtime_RemoteObject	`json:"value,omitempty"`// The value associated with the property.
}
type Runtime_CallArgument struct {
	Value			interface{}			`json:"value,omitempty"`// Primitive value or serializable javascript object.
	UnserializableValue	*Runtime_UnserializableValue	`json:"unserializableValue,omitempty"`// Primitive value which can not be JSON-stringified.
	ObjectId		*Runtime_RemoteObjectId		`json:"objectId,omitempty"`// Remote object handle.
}
type Runtime_ExecutionContextId int
type Runtime_ExecutionContextDescription struct {
	Id	Runtime_ExecutionContextId	`json:"id"`// Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin	string				`json:"origin"`// Execution context origin.
	Name	string				`json:"name"`// Human readable name describing given context.
	AuxData	*map[string]interface{}		`json:"auxData,omitempty"`// Embedder-specific auxiliary data.
}
type Runtime_ExceptionDetails struct {
	ExceptionId		int				`json:"exceptionId"`// Exception id.
	Text			string				`json:"text"`// Exception text, which should be used together with exception object when available.
	LineNumber		int				`json:"lineNumber"`// Line number of the exception location (0-based).
	ColumnNumber		int				`json:"columnNumber"`// Column number of the exception location (0-based).
	ScriptId		*Runtime_ScriptId		`json:"scriptId,omitempty"`// Script ID of the exception location.
	Url			*string				`json:"url,omitempty"`// URL of the exception location, to be used when the script was not reported.
	StackTrace		*Runtime_StackTrace		`json:"stackTrace,omitempty"`// JavaScript stack trace if available.
	Exception		*Runtime_RemoteObject		`json:"exception,omitempty"`// Exception object if available.
	ExecutionContextId	*Runtime_ExecutionContextId	`json:"executionContextId,omitempty"`// Identifier of the context where exception happened.
}
type Runtime_Timestamp float32
type Runtime_CallFrame struct {
	FunctionName	string			`json:"functionName"`// JavaScript function name.
	ScriptId	Runtime_ScriptId	`json:"scriptId"`// JavaScript script id.
	Url		string			`json:"url"`// JavaScript script name or url.
	LineNumber	int			`json:"lineNumber"`// JavaScript script line number (0-based).
	ColumnNumber	int			`json:"columnNumber"`// JavaScript script column number (0-based).
}
type Runtime_StackTrace struct {
	Description		*string			`json:"description,omitempty"`// String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames		[]Runtime_CallFrame	`json:"callFrames"`// JavaScript function name.
	Parent			*Runtime_StackTrace	`json:"parent,omitempty"`// Asynchronous JavaScript stack trace that preceded this stack, if available.
	PromiseCreationFrame	*Runtime_CallFrame	`json:"promiseCreationFrame,omitempty"`// Creation frame of the Promise which produced the next synchronous trace when resolved, if available.
}
