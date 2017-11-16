/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Unique script identifier.
type Runtime_ScriptId string

// Unique object identifier.
type Runtime_RemoteObjectId string

// Primitive value which cannot be JSON-stringified.
type Runtime_UnserializableValue string

// Mirror object referencing original JavaScript object.
type Runtime_RemoteObject struct {
	// Object type.
	Type string `json:"type"`
	// Object subtype hint. Specified for <code>object</code> type values only.
	Subtype *string `json:"subtype,omitempty"`
	// Object class (constructor) name. Specified for <code>object</code> type values only.
	ClassName *string `json:"className,omitempty"`
	// Remote object value in case of primitive values or JSON values (if it was requested).
	Value interface{} `json:"value,omitempty"`
	// Primitive value which can not be JSON-stringified does not have <code>value</code>, but gets this property.
	UnserializableValue *Runtime_UnserializableValue `json:"unserializableValue,omitempty"`
	// String representation of the object.
	Description *string `json:"description,omitempty"`
	// Unique object identifier (for non-primitive values).
	ObjectId *Runtime_RemoteObjectId `json:"objectId,omitempty"`
	// Preview containing abbreviated property values. Specified for <code>object</code> type values only.
	// NOTE Experimental
	Preview *Runtime_ObjectPreview `json:"preview,omitempty"`
	// NOTE Experimental
	CustomPreview *Runtime_CustomPreview `json:"customPreview,omitempty"`
}

type Runtime_CustomPreview struct {
	Header                     string                  `json:"header"`
	HasBody                    bool                    `json:"hasBody"`
	FormatterObjectId          Runtime_RemoteObjectId  `json:"formatterObjectId"`
	BindRemoteObjectFunctionId Runtime_RemoteObjectId  `json:"bindRemoteObjectFunctionId"`
	ConfigObjectId             *Runtime_RemoteObjectId `json:"configObjectId,omitempty"`
}

// Object containing abbreviated remote object value.
type Runtime_ObjectPreview struct {
	// Object type.
	Type string `json:"type"`
	// Object subtype hint. Specified for <code>object</code> type values only.
	Subtype *string `json:"subtype,omitempty"`
	// String representation of the object.
	Description *string `json:"description,omitempty"`
	// True iff some of the properties or entries of the original object did not fit.
	Overflow bool `json:"overflow"`
	// List of the properties.
	Properties []Runtime_PropertyPreview `json:"properties"`
	// List of the entries. Specified for <code>map</code> and <code>set</code> subtype values only.
	Entries []Runtime_EntryPreview `json:"entries,omitempty"`
}

type Runtime_PropertyPreview struct {
	// Property name.
	Name string `json:"name"`
	// Object type. Accessor means that the property itself is an accessor property.
	Type string `json:"type"`
	// User-friendly property value string.
	Value *string `json:"value,omitempty"`
	// Nested value preview.
	ValuePreview *Runtime_ObjectPreview `json:"valuePreview,omitempty"`
	// Object subtype hint. Specified for <code>object</code> type values only.
	Subtype *string `json:"subtype,omitempty"`
}

type Runtime_EntryPreview struct {
	// Preview of the key. Specified for map-like collection entries.
	Key *Runtime_ObjectPreview `json:"key,omitempty"`
	// Preview of the value.
	Value Runtime_ObjectPreview `json:"value"`
}

// Object property descriptor.
type Runtime_PropertyDescriptor struct {
	// Property name or symbol description.
	Name string `json:"name"`
	// The value associated with the property.
	Value *Runtime_RemoteObject `json:"value,omitempty"`
	// True if the value associated with the property may be changed (data descriptors only).
	Writable *bool `json:"writable,omitempty"`
	// A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
	Get *Runtime_RemoteObject `json:"get,omitempty"`
	// A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
	Set *Runtime_RemoteObject `json:"set,omitempty"`
	// True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Configurable bool `json:"configurable"`
	// True if this property shows up during enumeration of the properties on the corresponding object.
	Enumerable bool `json:"enumerable"`
	// True if the result was thrown during the evaluation.
	WasThrown *bool `json:"wasThrown,omitempty"`
	// True if the property is owned for the object.
	IsOwn *bool `json:"isOwn,omitempty"`
	// Property symbol object, if the property is of the <code>symbol</code> type.
	Symbol *Runtime_RemoteObject `json:"symbol,omitempty"`
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type Runtime_InternalPropertyDescriptor struct {
	// Conventional property name.
	Name string `json:"name"`
	// The value associated with the property.
	Value *Runtime_RemoteObject `json:"value,omitempty"`
}

// Represents function call argument. Either remote object id <code>objectId</code>, primitive <code>value</code>, unserializable primitive value or neither of (for undefined) them should be specified.
type Runtime_CallArgument struct {
	// Primitive value or serializable javascript object.
	Value interface{} `json:"value,omitempty"`
	// Primitive value which can not be JSON-stringified.
	UnserializableValue *Runtime_UnserializableValue `json:"unserializableValue,omitempty"`
	// Remote object handle.
	ObjectId *Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

// Id of an execution context.
type Runtime_ExecutionContextId int

// Description of an isolated world.
type Runtime_ExecutionContextDescription struct {
	// Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Id Runtime_ExecutionContextId `json:"id"`
	// Execution context origin.
	Origin string `json:"origin"`
	// Human readable name describing given context.
	Name string `json:"name"`
	// Embedder-specific auxiliary data.
	AuxData map[string]interface{} `json:"auxData,omitempty"`
}

// Detailed information about exception (or error) that was thrown during script compilation or execution.
type Runtime_ExceptionDetails struct {
	// Exception id.
	ExceptionId int `json:"exceptionId"`
	// Exception text, which should be used together with exception object when available.
	Text string `json:"text"`
	// Line number of the exception location (0-based).
	LineNumber int `json:"lineNumber"`
	// Column number of the exception location (0-based).
	ColumnNumber int `json:"columnNumber"`
	// Script ID of the exception location.
	ScriptId *Runtime_ScriptId `json:"scriptId,omitempty"`
	// URL of the exception location, to be used when the script was not reported.
	Url *string `json:"url,omitempty"`
	// JavaScript stack trace if available.
	StackTrace *Runtime_StackTrace `json:"stackTrace,omitempty"`
	// Exception object if available.
	Exception *Runtime_RemoteObject `json:"exception,omitempty"`
	// Identifier of the context where exception happened.
	ExecutionContextId *Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

// Number of milliseconds since epoch.
type Runtime_Timestamp float32

// Stack entry for runtime errors and assertions.
type Runtime_CallFrame struct {
	// JavaScript function name.
	FunctionName string `json:"functionName"`
	// JavaScript script id.
	ScriptId Runtime_ScriptId `json:"scriptId"`
	// JavaScript script name or url.
	Url string `json:"url"`
	// JavaScript script line number (0-based).
	LineNumber int `json:"lineNumber"`
	// JavaScript script column number (0-based).
	ColumnNumber int `json:"columnNumber"`
}

// Call frames for assertions or error messages.
type Runtime_StackTrace struct {
	// String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	Description *string `json:"description,omitempty"`
	// JavaScript function name.
	CallFrames []Runtime_CallFrame `json:"callFrames"`
	// Asynchronous JavaScript stack trace that preceded this stack, if available.
	Parent *Runtime_StackTrace `json:"parent,omitempty"`
	// Creation frame of the Promise which produced the next synchronous trace when resolved, if available.
	// NOTE Experimental
	PromiseCreationFrame *Runtime_CallFrame `json:"promiseCreationFrame,omitempty"`
}

type Runtime_AsyncTaskId string
