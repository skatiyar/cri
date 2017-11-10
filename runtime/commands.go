package runtime

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Runtime struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Runtime {
	return &Runtime{conn}
}

type EvaluateRequest struct {
	Expression		string					`json:"expression"`// Expression to evaluate.
	ObjectGroup		*string					`json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
	IncludeCommandLineAPI	*bool					`json:"includeCommandLineAPI,omitempty"`// Determines whether Command Line API should be available during the evaluation.
	Silent			*bool					`json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	ContextId		*types.Runtime_ExecutionContextId	`json:"contextId,omitempty"`// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ReturnByValue		*bool					`json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview		*bool					`json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
	UserGesture		*bool					`json:"userGesture,omitempty"`// Whether execution should be treated as initiated by user in the UI.
	AwaitPromise		*bool					`json:"awaitPromise,omitempty"`// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
}

func (obj *Runtime) Evaluate(request *EvaluateRequest) (response struct {
	Result			types.Runtime_RemoteObject	`json:"result"`// Evaluation result.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Runtime.evaluate", request, &response)
	return
}

type AwaitPromiseRequest struct {
	PromiseObjectId	types.Runtime_RemoteObjectId	`json:"promiseObjectId"`// Identifier of the promise.
	ReturnByValue	*bool				`json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview	*bool				`json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
}

func (obj *Runtime) AwaitPromise(request *AwaitPromiseRequest) (response struct {
	Result			types.Runtime_RemoteObject	`json:"result"`// Promise result. Will contain rejected value if promise was rejected.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details if stack strace is available.
}, err error) {
	err = obj.conn.Send("Runtime.awaitPromise", request, &response)
	return
}

type CallFunctionOnRequest struct {
	FunctionDeclaration	string					`json:"functionDeclaration"`// Declaration of the function to call.
	ObjectId		*types.Runtime_RemoteObjectId		`json:"objectId,omitempty"`// Identifier of the object to call function on. Either objectId or executionContextId should be specified.
	Arguments		[]types.Runtime_CallArgument		`json:"arguments,omitempty"`// Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Silent			*bool					`json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	ReturnByValue		*bool					`json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview		*bool					`json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
	UserGesture		*bool					`json:"userGesture,omitempty"`// Whether execution should be treated as initiated by user in the UI.
	AwaitPromise		*bool					`json:"awaitPromise,omitempty"`// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
	ExecutionContextId	*types.Runtime_ExecutionContextId	`json:"executionContextId,omitempty"`// Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
	ObjectGroup		*string					`json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
}

func (obj *Runtime) CallFunctionOn(request *CallFunctionOnRequest) (response struct {
	Result			types.Runtime_RemoteObject	`json:"result"`// Call result.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Runtime.callFunctionOn", request, &response)
	return
}

type GetPropertiesRequest struct {
	ObjectId		types.Runtime_RemoteObjectId	`json:"objectId"`// Identifier of the object to return properties for.
	OwnProperties		*bool				`json:"ownProperties,omitempty"`// If true, returns properties belonging only to the element itself, not to its prototype chain.
	AccessorPropertiesOnly	*bool				`json:"accessorPropertiesOnly,omitempty"`// If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	GeneratePreview		*bool				`json:"generatePreview,omitempty"`// Whether preview should be generated for the results.
}

func (obj *Runtime) GetProperties(request *GetPropertiesRequest) (response struct {
	Result			[]types.Runtime_PropertyDescriptor		`json:"result"`// Object properties.
	InternalProperties	[]types.Runtime_InternalPropertyDescriptor	`json:"internalProperties,omitempty"`// Internal object properties (only of the element itself).
	ExceptionDetails	*types.Runtime_ExceptionDetails			`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Runtime.getProperties", request, &response)
	return
}

type ReleaseObjectRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`// Identifier of the object to release.
}

func (obj *Runtime) ReleaseObject(request *ReleaseObjectRequest) (err error) {
	err = obj.conn.Send("Runtime.releaseObject", request, nil)
	return
}

type ReleaseObjectGroupRequest struct {
	ObjectGroup string `json:"objectGroup"`// Symbolic object group name.
}

func (obj *Runtime) ReleaseObjectGroup(request *ReleaseObjectGroupRequest) (err error) {
	err = obj.conn.Send("Runtime.releaseObjectGroup", request, nil)
	return
}
func (obj *Runtime) RunIfWaitingForDebugger() (err error) {
	err = obj.conn.Send("Runtime.runIfWaitingForDebugger", nil, nil)
	return
}
func (obj *Runtime) Enable() (err error) {
	err = obj.conn.Send("Runtime.enable", nil, nil)
	return
}
func (obj *Runtime) Disable() (err error) {
	err = obj.conn.Send("Runtime.disable", nil, nil)
	return
}
func (obj *Runtime) DiscardConsoleEntries() (err error) {
	err = obj.conn.Send("Runtime.discardConsoleEntries", nil, nil)
	return
}

type SetCustomObjectFormatterEnabledRequest struct {
	Enabled bool `json:"enabled"`
}

func (obj *Runtime) SetCustomObjectFormatterEnabled(request *SetCustomObjectFormatterEnabledRequest) (err error) {
	err = obj.conn.Send("Runtime.setCustomObjectFormatterEnabled", request, nil)
	return
}

type CompileScriptRequest struct {
	Expression		string					`json:"expression"`// Expression to compile.
	SourceURL		string					`json:"sourceURL"`// Source url to be set for the script.
	PersistScript		bool					`json:"persistScript"`// Specifies whether the compiled script should be persisted.
	ExecutionContextId	*types.Runtime_ExecutionContextId	`json:"executionContextId,omitempty"`// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
}

func (obj *Runtime) CompileScript(request *CompileScriptRequest) (response struct {
	ScriptId		*types.Runtime_ScriptId		`json:"scriptId,omitempty"`// Id of the script.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Runtime.compileScript", request, &response)
	return
}

type RunScriptRequest struct {
	ScriptId		types.Runtime_ScriptId			`json:"scriptId"`// Id of the script to run.
	ExecutionContextId	*types.Runtime_ExecutionContextId	`json:"executionContextId,omitempty"`// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ObjectGroup		*string					`json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
	Silent			*bool					`json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	IncludeCommandLineAPI	*bool					`json:"includeCommandLineAPI,omitempty"`// Determines whether Command Line API should be available during the evaluation.
	ReturnByValue		*bool					`json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview		*bool					`json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
	AwaitPromise		*bool					`json:"awaitPromise,omitempty"`// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
}

func (obj *Runtime) RunScript(request *RunScriptRequest) (response struct {
	Result			types.Runtime_RemoteObject	`json:"result"`// Run result.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Runtime.runScript", request, &response)
	return
}

type QueryObjectsRequest struct {
	PrototypeObjectId types.Runtime_RemoteObjectId `json:"prototypeObjectId"`// Identifier of the prototype to return objects for.
}

func (obj *Runtime) QueryObjects(request *QueryObjectsRequest) (response struct {
	Objects types.Runtime_RemoteObject `json:"objects"`// Array with objects.
}, err error) {
	err = obj.conn.Send("Runtime.queryObjects", request, &response)
	return
}

type GlobalLexicalScopeNamesRequest struct {
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`// Specifies in which execution context to lookup global scope variables.
}

func (obj *Runtime) GlobalLexicalScopeNames(request *GlobalLexicalScopeNamesRequest) (response struct {
	Names []string `json:"names"`
}, err error) {
	err = obj.conn.Send("Runtime.globalLexicalScopeNames", request, &response)
	return
}
