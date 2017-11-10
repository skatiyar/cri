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
	Expression            string                            `json:"expression"`
	ObjectGroup           *string                           `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI *bool                             `json:"includeCommandLineAPI,omitempty"`
	Silent                *bool                             `json:"silent,omitempty"`
	ContextId             *types.Runtime_ExecutionContextId `json:"contextId,omitempty"`
	ReturnByValue         *bool                             `json:"returnByValue,omitempty"`
	GeneratePreview       *bool                             `json:"generatePreview,omitempty"`
	UserGesture           *bool                             `json:"userGesture,omitempty"`
	AwaitPromise          *bool                             `json:"awaitPromise,omitempty"`
}

func (obj *Runtime) Evaluate(request *EvaluateRequest) (response struct {
	Result           types.Runtime_RemoteObject      `json:"result"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.evaluate", request, &response)
	return
}

type AwaitPromiseRequest struct {
	PromiseObjectId types.Runtime_RemoteObjectId `json:"promiseObjectId"`
	ReturnByValue   *bool                        `json:"returnByValue,omitempty"`
	GeneratePreview *bool                        `json:"generatePreview,omitempty"`
}

func (obj *Runtime) AwaitPromise(request *AwaitPromiseRequest) (response struct {
	Result           types.Runtime_RemoteObject      `json:"result"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.awaitPromise", request, &response)
	return
}

type CallFunctionOnRequest struct {
	FunctionDeclaration string                            `json:"functionDeclaration"`
	ObjectId            *types.Runtime_RemoteObjectId     `json:"objectId,omitempty"`
	Arguments           []types.Runtime_CallArgument      `json:"arguments,omitempty"`
	Silent              *bool                             `json:"silent,omitempty"`
	ReturnByValue       *bool                             `json:"returnByValue,omitempty"`
	GeneratePreview     *bool                             `json:"generatePreview,omitempty"`
	UserGesture         *bool                             `json:"userGesture,omitempty"`
	AwaitPromise        *bool                             `json:"awaitPromise,omitempty"`
	ExecutionContextId  *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
	ObjectGroup         *string                           `json:"objectGroup,omitempty"`
}

func (obj *Runtime) CallFunctionOn(request *CallFunctionOnRequest) (response struct {
	Result           types.Runtime_RemoteObject      `json:"result"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.callFunctionOn", request, &response)
	return
}

type GetPropertiesRequest struct {
	ObjectId               types.Runtime_RemoteObjectId `json:"objectId"`
	OwnProperties          *bool                        `json:"ownProperties,omitempty"`
	AccessorPropertiesOnly *bool                        `json:"accessorPropertiesOnly,omitempty"`
	GeneratePreview        *bool                        `json:"generatePreview,omitempty"`
}

func (obj *Runtime) GetProperties(request *GetPropertiesRequest) (response struct {
	Result             []types.Runtime_PropertyDescriptor         `json:"result"`
	InternalProperties []types.Runtime_InternalPropertyDescriptor `json:"internalProperties,omitempty"`
	ExceptionDetails   *types.Runtime_ExceptionDetails            `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.getProperties", request, &response)
	return
}

type ReleaseObjectRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

func (obj *Runtime) ReleaseObject(request *ReleaseObjectRequest) (err error) {
	err = obj.conn.Send("Runtime.releaseObject", request, nil)
	return
}

type ReleaseObjectGroupRequest struct {
	ObjectGroup string `json:"objectGroup"`
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
	Expression         string                            `json:"expression"`
	SourceURL          string                            `json:"sourceURL"`
	PersistScript      bool                              `json:"persistScript"`
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

func (obj *Runtime) CompileScript(request *CompileScriptRequest) (response struct {
	ScriptId         *types.Runtime_ScriptId         `json:"scriptId,omitempty"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.compileScript", request, &response)
	return
}

type RunScriptRequest struct {
	ScriptId              types.Runtime_ScriptId            `json:"scriptId"`
	ExecutionContextId    *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
	ObjectGroup           *string                           `json:"objectGroup,omitempty"`
	Silent                *bool                             `json:"silent,omitempty"`
	IncludeCommandLineAPI *bool                             `json:"includeCommandLineAPI,omitempty"`
	ReturnByValue         *bool                             `json:"returnByValue,omitempty"`
	GeneratePreview       *bool                             `json:"generatePreview,omitempty"`
	AwaitPromise          *bool                             `json:"awaitPromise,omitempty"`
}

func (obj *Runtime) RunScript(request *RunScriptRequest) (response struct {
	Result           types.Runtime_RemoteObject      `json:"result"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Runtime.runScript", request, &response)
	return
}

type QueryObjectsRequest struct {
	PrototypeObjectId types.Runtime_RemoteObjectId `json:"prototypeObjectId"`
}

func (obj *Runtime) QueryObjects(request *QueryObjectsRequest) (response struct {
	Objects types.Runtime_RemoteObject `json:"objects"`
}, err error) {
	err = obj.conn.Send("Runtime.queryObjects", request, &response)
	return
}

type GlobalLexicalScopeNamesRequest struct {
	ExecutionContextId *types.Runtime_ExecutionContextId `json:"executionContextId,omitempty"`
}

func (obj *Runtime) GlobalLexicalScopeNames(request *GlobalLexicalScopeNamesRequest) (response struct {
	Names []string `json:"names"`
}, err error) {
	err = obj.conn.Send("Runtime.globalLexicalScopeNames", request, &response)
	return
}
