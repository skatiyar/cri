/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package runtime provides commands and events for Runtime domain.
package runtime

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Runtime domain
const (
	Evaluate                        = "Runtime.evaluate"
	AwaitPromise                    = "Runtime.awaitPromise"
	CallFunctionOn                  = "Runtime.callFunctionOn"
	GetProperties                   = "Runtime.getProperties"
	ReleaseObject                   = "Runtime.releaseObject"
	ReleaseObjectGroup              = "Runtime.releaseObjectGroup"
	RunIfWaitingForDebugger         = "Runtime.runIfWaitingForDebugger"
	Enable                          = "Runtime.enable"
	Disable                         = "Runtime.disable"
	DiscardConsoleEntries           = "Runtime.discardConsoleEntries"
	SetCustomObjectFormatterEnabled = "Runtime.setCustomObjectFormatterEnabled"
	CompileScript                   = "Runtime.compileScript"
	RunScript                       = "Runtime.runScript"
	QueryObjects                    = "Runtime.queryObjects"
	GlobalLexicalScopeNames         = "Runtime.globalLexicalScopeNames"
)

// List of events in Runtime domain
const (
	ExecutionContextCreated   = "Runtime.executionContextCreated"
	ExecutionContextDestroyed = "Runtime.executionContextDestroyed"
	ExecutionContextsCleared  = "Runtime.executionContextsCleared"
	ExceptionThrown           = "Runtime.exceptionThrown"
	ExceptionRevoked          = "Runtime.exceptionRevoked"
	ConsoleAPICalled          = "Runtime.consoleAPICalled"
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

type EvaluateRequest struct {
	// Expression to evaluate.
	Expression string `json:"expression"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent *bool `json:"silent,omitempty"`
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ContextId *types.Runtime_ExecutionContextId `json:"contextId,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	// NOTE Experimental
	UserGesture *bool `json:"userGesture,omitempty"`
	// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
	AwaitPromise *bool `json:"awaitPromise,omitempty"`
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
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent *bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	// NOTE Experimental
	UserGesture *bool `json:"userGesture,omitempty"`
	// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
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

// Enables reporting of execution contexts creation by means of <code>executionContextCreated</code> event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (obj *Runtime) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
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

type SetCustomObjectFormatterEnabledRequest struct {
	Enabled bool `json:"enabled"`
}

func (obj *Runtime) SetCustomObjectFormatterEnabled(request *SetCustomObjectFormatterEnabledRequest) (err error) {
	err = obj.conn.Send(SetCustomObjectFormatterEnabled, request, nil)
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

type RunScriptRequest struct {
	// Id of the script to run.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent *bool `json:"silent,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether execution should <code>await</code> for resulting value and return once awaited promise is resolved.
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

type QueryObjectsRequest struct {
	// Identifier of the prototype to return objects for.
	PrototypeObjectId types.Runtime_RemoteObjectId `json:"prototypeObjectId"`
}

type QueryObjectsResponse struct {
	// Array with objects.
	Objects types.Runtime_RemoteObject `json:"objects"`
}

func (obj *Runtime) QueryObjects(request *QueryObjectsRequest) (response QueryObjectsResponse, err error) {
	err = obj.conn.Send(QueryObjects, request, &response)
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

type ExecutionContextCreatedParams struct {
	// A newly created execution context.
	Context types.Runtime_ExecutionContextDescription `json:"context"`
}

// Issued when new execution context is created.
func (obj *Runtime) ExecutionContextCreated(fn func(params *ExecutionContextCreatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ExecutionContextCreated, closeChn)
	go func() {
		for {
			params := ExecutionContextCreatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ExecutionContextDestroyedParams struct {
	// Id of the destroyed context
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
}

// Issued when execution context is destroyed.
func (obj *Runtime) ExecutionContextDestroyed(fn func(params *ExecutionContextDestroyedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ExecutionContextDestroyed, closeChn)
	go func() {
		for {
			params := ExecutionContextDestroyedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

// Issued when all executionContexts were cleared in browser
func (obj *Runtime) ExecutionContextsCleared(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ExecutionContextsCleared, closeChn)
	go func() {
		for {

			readErr := decoder(nil)
			if !fn(readErr) {
				close(closeChn)
				break
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
func (obj *Runtime) ExceptionThrown(fn func(params *ExceptionThrownParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ExceptionThrown, closeChn)
	go func() {
		for {
			params := ExceptionThrownParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ExceptionRevokedParams struct {
	// Reason describing why exception was revoked.
	Reason string `json:"reason"`
	// The id of revoked exception, as reported in <code>exceptionThrown</code>.
	ExceptionId int `json:"exceptionId"`
}

// Issued when unhandled exception was revoked.
func (obj *Runtime) ExceptionRevoked(fn func(params *ExceptionRevokedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ExceptionRevoked, closeChn)
	go func() {
		for {
			params := ExceptionRevokedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
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
func (obj *Runtime) ConsoleAPICalled(fn func(params *ConsoleAPICalledParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ConsoleAPICalled, closeChn)
	go func() {
		for {
			params := ConsoleAPICalledParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type InspectRequestedParams struct {
	Object types.Runtime_RemoteObject `json:"object"`
	Hints  map[string]interface{}     `json:"hints"`
}

// Issued when object should be inspected (for example, as a result of inspect() command line API call).
func (obj *Runtime) InspectRequested(fn func(params *InspectRequestedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(InspectRequested, closeChn)
	go func() {
		for {
			params := InspectRequestedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
