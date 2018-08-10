/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package runtime provides commands and events for Runtime domain.
package runtime

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Runtime domain
const (
	AwaitPromise                    = "Runtime.awaitPromise"
	CallFunctionOn                  = "Runtime.callFunctionOn"
	CompileScript                   = "Runtime.compileScript"
	Disable                         = "Runtime.disable"
	DiscardConsoleEntries           = "Runtime.discardConsoleEntries"
	Enable                          = "Runtime.enable"
	Evaluate                        = "Runtime.evaluate"
	GetIsolateId                    = "Runtime.getIsolateId"
	GetHeapUsage                    = "Runtime.getHeapUsage"
	GetProperties                   = "Runtime.getProperties"
	GlobalLexicalScopeNames         = "Runtime.globalLexicalScopeNames"
	QueryObjects                    = "Runtime.queryObjects"
	ReleaseObject                   = "Runtime.releaseObject"
	ReleaseObjectGroup              = "Runtime.releaseObjectGroup"
	RunIfWaitingForDebugger         = "Runtime.runIfWaitingForDebugger"
	RunScript                       = "Runtime.runScript"
	SetAsyncCallStackDepth          = "Runtime.setAsyncCallStackDepth"
	SetCustomObjectFormatterEnabled = "Runtime.setCustomObjectFormatterEnabled"
	SetMaxCallStackSizeToCapture    = "Runtime.setMaxCallStackSizeToCapture"
	TerminateExecution              = "Runtime.terminateExecution"
	AddBinding                      = "Runtime.addBinding"
	RemoveBinding                   = "Runtime.removeBinding"
)

// List of events in Runtime domain
const (
	BindingCalled             = "Runtime.bindingCalled"
	ConsoleAPICalled          = "Runtime.consoleAPICalled"
	ExceptionRevoked          = "Runtime.exceptionRevoked"
	ExceptionThrown           = "Runtime.exceptionThrown"
	ExecutionContextCreated   = "Runtime.executionContextCreated"
	ExecutionContextDestroyed = "Runtime.executionContextDestroyed"
	ExecutionContextsCleared  = "Runtime.executionContextsCleared"
	InspectRequested          = "Runtime.inspectRequested"
)

// Runtime domain exposes JavaScript runtime by means of remote evaluation and mirror objects. Evaluation results are returned as mirror object that expose object type, string representation and unique identifier that can be used for further object reference. Original objects are maintained in memory unless they are either explicitly released or are released along with the other objects in their object group.
type Runtime struct {
	conn cri.Connector
}

// New creates a Runtime instance
func New(conn cri.Connector) *Runtime {
	return &Runtime{conn}
}

type AwaitPromiseRequest struct {
	// Identifier of the promise.
	PromiseObjectId types.Runtime_RemoteObjectId `json:"promiseObjectId"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview *bool `json:"generatePreview,omitempty"`
}

type AwaitPromiseResponse struct {
	// Promise result. Will contain rejected value if promise was rejected.
	Result types.Runtime_RemoteObject `json:"result"`
	// Exception details if stack strace is available.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Add handler to promise with given promise object id.
func (obj *Runtime) AwaitPromise(request *AwaitPromiseRequest) (response AwaitPromiseResponse, err error) {
	err = obj.conn.Send(AwaitPromise, request, &response)
	return
}

type CallFunctionOnRequest struct {
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`
	// Identifier of the object to call function on. Either objectId or executionContextId should be specified.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
	// Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Arguments []types.Runtime_CallArgument `json:"arguments,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent *bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture *bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise *bool `json:"awaitPromise,omitempty"`
	// Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
	ObjectGroup *string `json:"objectGroup,omitempty"`
}

type CallFunctionOnResponse struct {
	// Call result.
	Result types.Runtime_RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
func (obj *Runtime) CallFunctionOn(request *CallFunctionOnRequest) (response CallFunctionOnResponse, err error) {
	err = obj.conn.Send(CallFunctionOn, request, &response)
	return
}

type CompileScriptRequest struct {
	// Expression to compile.
	Expression string `json:"expression"`
	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`
	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

type CompileScriptResponse struct {
	// Id of the script.
	ScriptId *types.Runtime_ScriptId `json:"scriptId,omitempty"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Compiles expression.
func (obj *Runtime) CompileScript(request *CompileScriptRequest) (response CompileScriptResponse, err error) {
	err = obj.conn.Send(CompileScript, request, &response)
	return
}

// Disables reporting of execution contexts creation.
func (obj *Runtime) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Discards collected exceptions and console API calls.
func (obj *Runtime) DiscardConsoleEntries() (err error) {
	err = obj.conn.Send(DiscardConsoleEntries, nil, nil)
	return
}

// Enables reporting of execution contexts creation by means of `executionContextCreated` event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (obj *Runtime) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type EvaluateRequest struct {
	// Expression to evaluate.
	Expression string `json:"expression"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent *bool `json:"silent,omitempty"`
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ContextId *types.Runtime_ExecutionContextId `json:"contextId,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture *bool `json:"userGesture,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise *bool `json:"awaitPromise,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	// NOTE Experimental
	ThrowOnSideEffect *bool `json:"throwOnSideEffect,omitempty"`
	// Terminate execution after timing out (number of milliseconds).
	// NOTE Experimental
	Timeout *types.Runtime_TimeDelta `json:"timeout,omitempty"`
}

type EvaluateResponse struct {
	// Evaluation result.
	Result types.Runtime_RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Evaluates expression on global object.
func (obj *Runtime) Evaluate(request *EvaluateRequest) (response EvaluateResponse, err error) {
	err = obj.conn.Send(Evaluate, request, &response)
	return
}

type GetIsolateIdResponse struct {
	// The isolate id.
	Id string `json:"id"`
}

// Returns the isolate id.
func (obj *Runtime) GetIsolateId() (response GetIsolateIdResponse, err error) {
	err = obj.conn.Send(GetIsolateId, nil, &response)
	return
}

type GetHeapUsageResponse struct {
	// Used heap size in bytes.
	UsedSize float32 `json:"usedSize"`
	// Allocated heap size in bytes.
	TotalSize float32 `json:"totalSize"`
}

// Returns the JavaScript heap usage. It is the total usage of the corresponding isolate not scoped to a particular Runtime.
func (obj *Runtime) GetHeapUsage() (response GetHeapUsageResponse, err error) {
	err = obj.conn.Send(GetHeapUsage, nil, &response)
	return
}

type GetPropertiesRequest struct {
	// Identifier of the object to return properties for.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
	// If true, returns properties belonging only to the element itself, not to its prototype chain.
	OwnProperties *bool `json:"ownProperties,omitempty"`
	// If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	// NOTE Experimental
	AccessorPropertiesOnly *bool `json:"accessorPropertiesOnly,omitempty"`
	// Whether preview should be generated for the results.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
}

type GetPropertiesResponse struct {
	// Object properties.
	Result []types.Runtime_PropertyDescriptor `json:"result"`
	// Internal object properties (only of the element itself).
	InternalProperties []types.Runtime_InternalPropertyDescriptor `json:"internalProperties,omitempty"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Returns properties of a given object. Object group of the result is inherited from the target object.
func (obj *Runtime) GetProperties(request *GetPropertiesRequest) (response GetPropertiesResponse, err error) {
	err = obj.conn.Send(GetProperties, request, &response)
	return
}

type GlobalLexicalScopeNamesRequest struct {
	// Specifies in which execution context to lookup global scope variables.
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

type GlobalLexicalScopeNamesResponse struct {
	Names []string `json:"names"`
}

// Returns all let, const and class variables from global scope.
func (obj *Runtime) GlobalLexicalScopeNames(request *GlobalLexicalScopeNamesRequest) (response GlobalLexicalScopeNamesResponse, err error) {
	err = obj.conn.Send(GlobalLexicalScopeNames, request, &response)
	return
}

type QueryObjectsRequest struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectId types.Runtime_RemoteObjectId `json:"prototypeObjectId"`
	// Symbolic group name that can be used to release the results.
	ObjectGroup *string `json:"objectGroup,omitempty"`
}

type QueryObjectsResponse struct {
	// Array with objects.
	Objects types.Runtime_RemoteObject `json:"objects"`
}

func (obj *Runtime) QueryObjects(request *QueryObjectsRequest) (response QueryObjectsResponse, err error) {
	err = obj.conn.Send(QueryObjects, request, &response)
	return
}

type ReleaseObjectRequest struct {
	// Identifier of the object to release.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

// Releases remote object with given id.
func (obj *Runtime) ReleaseObject(request *ReleaseObjectRequest) (err error) {
	err = obj.conn.Send(ReleaseObject, request, nil)
	return
}

type ReleaseObjectGroupRequest struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

// Releases all remote objects that belong to a given group.
func (obj *Runtime) ReleaseObjectGroup(request *ReleaseObjectGroupRequest) (err error) {
	err = obj.conn.Send(ReleaseObjectGroup, request, nil)
	return
}

// Tells inspected instance to run if it was waiting for debugger to attach.
func (obj *Runtime) RunIfWaitingForDebugger() (err error) {
	err = obj.conn.Send(RunIfWaitingForDebugger, nil, nil)
	return
}

type RunScriptRequest struct {
	// Id of the script to run.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	Silent *bool `json:"silent,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should `await` for resulting value and return once awaited promise is resolved.
	AwaitPromise *bool `json:"awaitPromise,omitempty"`
}

type RunScriptResponse struct {
	// Run result.
	Result types.Runtime_RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Runs script with given id in a given context.
func (obj *Runtime) RunScript(request *RunScriptRequest) (response RunScriptResponse, err error) {
	err = obj.conn.Send(RunScript, request, &response)
	return
}

type SetAsyncCallStackDepthRequest struct {
	// Maximum depth of async call stacks. Setting to `0` will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// Enables or disables async call stacks tracking.
func (obj *Runtime) SetAsyncCallStackDepth(request *SetAsyncCallStackDepthRequest) (err error) {
	err = obj.conn.Send(SetAsyncCallStackDepth, request, nil)
	return
}

type SetCustomObjectFormatterEnabledRequest struct {
	Enabled bool `json:"enabled"`
}

func (obj *Runtime) SetCustomObjectFormatterEnabled(request *SetCustomObjectFormatterEnabledRequest) (err error) {
	err = obj.conn.Send(SetCustomObjectFormatterEnabled, request, nil)
	return
}

type SetMaxCallStackSizeToCaptureRequest struct {
	Size int `json:"size"`
}

func (obj *Runtime) SetMaxCallStackSizeToCapture(request *SetMaxCallStackSizeToCaptureRequest) (err error) {
	err = obj.conn.Send(SetMaxCallStackSizeToCapture, request, nil)
	return
}

// Terminate current or next JavaScript execution. Will cancel the termination when the outer-most script execution ends.
func (obj *Runtime) TerminateExecution() (err error) {
	err = obj.conn.Send(TerminateExecution, nil, nil)
	return
}

type AddBindingRequest struct {
	Name               string                            `json:"name"`
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

// If executionContextId is empty, adds binding with the given name on the global objects of all inspected contexts, including those created later, bindings survive reloads. If executionContextId is specified, adds binding only on global object of given execution context. Binding function takes exactly one argument, this argument should be string, in case of any other input, function throws an exception. Each binding function call produces Runtime.bindingCalled notification.
func (obj *Runtime) AddBinding(request *AddBindingRequest) (err error) {
	err = obj.conn.Send(AddBinding, request, nil)
	return
}

type RemoveBindingRequest struct {
	Name string `json:"name"`
}

// This method does not remove binding function from global object but unsubscribes current runtime agent from Runtime.bindingCalled notifications.
func (obj *Runtime) RemoveBinding(request *RemoveBindingRequest) (err error) {
	err = obj.conn.Send(RemoveBinding, request, nil)
	return
}

type BindingCalledParams struct {
	Name    string `json:"name"`
	Payload string `json:"payload"`
	// Identifier of the context where the call was made.
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
}

// Notification is issued every time when binding is called.
// NOTE Experimental
func (obj *Runtime) BindingCalled(fn func(event string, params BindingCalledParams, err error) bool) {
	listen, closer := obj.conn.On(BindingCalled)
	go func() {
		defer closer()
		for {
			var params BindingCalledParams
			if !fn(BindingCalled, params, listen(&params)) {
				return
			}
		}
	}()
}

type ConsoleAPICalledParams struct {
	// Type of the call.
	Type string `json:"type"`
	// Call arguments.
	Args []types.Runtime_RemoteObject `json:"args"`
	// Identifier of the context where the call was made.
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
	// Call timestamp.
	Timestamp types.Runtime_Timestamp `json:"timestamp"`
	// Stack trace captured when the call was made.
	StackTrace *types.Runtime_StackTrace `json:"stackTrace,omitempty"`
	// Console context descriptor for calls on non-default console context (not console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call on named context.
	// NOTE Experimental
	Context *string `json:"context,omitempty"`
}

// Issued when console API was called.
func (obj *Runtime) ConsoleAPICalled(fn func(event string, params ConsoleAPICalledParams, err error) bool) {
	listen, closer := obj.conn.On(ConsoleAPICalled)
	go func() {
		defer closer()
		for {
			var params ConsoleAPICalledParams
			if !fn(ConsoleAPICalled, params, listen(&params)) {
				return
			}
		}
	}()
}

type ExceptionRevokedParams struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`
	// The id of revoked exception, as reported in `exceptionThrown`.
	ExceptionId int `json:"exceptionId"`
}

// Issued when unhandled exception was revoked.
func (obj *Runtime) ExceptionRevoked(fn func(event string, params ExceptionRevokedParams, err error) bool) {
	listen, closer := obj.conn.On(ExceptionRevoked)
	go func() {
		defer closer()
		for {
			var params ExceptionRevokedParams
			if !fn(ExceptionRevoked, params, listen(&params)) {
				return
			}
		}
	}()
}

type ExceptionThrownParams struct {
	// Timestamp of the exception.
	Timestamp        types.Runtime_Timestamp        `json:"timestamp"`
	ExceptionDetails types.Runtime_ExceptionDetails `json:"exceptionDetails"`
}

// Issued when exception was thrown and unhandled.
func (obj *Runtime) ExceptionThrown(fn func(event string, params ExceptionThrownParams, err error) bool) {
	listen, closer := obj.conn.On(ExceptionThrown)
	go func() {
		defer closer()
		for {
			var params ExceptionThrownParams
			if !fn(ExceptionThrown, params, listen(&params)) {
				return
			}
		}
	}()
}

type ExecutionContextCreatedParams struct {
	// A newly created execution context.
	Context types.Runtime_ExecutionContextDescription `json:"context"`
}

// Issued when new execution context is created.
func (obj *Runtime) ExecutionContextCreated(fn func(event string, params ExecutionContextCreatedParams, err error) bool) {
	listen, closer := obj.conn.On(ExecutionContextCreated)
	go func() {
		defer closer()
		for {
			var params ExecutionContextCreatedParams
			if !fn(ExecutionContextCreated, params, listen(&params)) {
				return
			}
		}
	}()
}

type ExecutionContextDestroyedParams struct {
	// Id of the destroyed context
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
}

// Issued when execution context is destroyed.
func (obj *Runtime) ExecutionContextDestroyed(fn func(event string, params ExecutionContextDestroyedParams, err error) bool) {
	listen, closer := obj.conn.On(ExecutionContextDestroyed)
	go func() {
		defer closer()
		for {
			var params ExecutionContextDestroyedParams
			if !fn(ExecutionContextDestroyed, params, listen(&params)) {
				return
			}
		}
	}()
}

// Issued when all executionContexts were cleared in browser
func (obj *Runtime) ExecutionContextsCleared(fn func(event string, err error) bool) {
	listen, closer := obj.conn.On(ExecutionContextsCleared)
	go func() {
		defer closer()
		for {
			if !fn(ExecutionContextsCleared, listen(nil)) {
				return
			}
		}
	}()
}

type InspectRequestedParams struct {
	Object types.Runtime_RemoteObject `json:"object"`
	Hints  map[string]interface{}     `json:"hints"`
}

// Issued when object should be inspected (for example, as a result of inspect() command line API call).
func (obj *Runtime) InspectRequested(fn func(event string, params InspectRequestedParams, err error) bool) {
	listen, closer := obj.conn.On(InspectRequested)
	go func() {
		defer closer()
		for {
			var params InspectRequestedParams
			if !fn(InspectRequested, params, listen(&params)) {
				return
			}
		}
	}()
}
