package debugger

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Debugger struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Debugger {
	return &Debugger{conn}
}
func (obj *Debugger) Enable() (err error) {
	err = obj.conn.Send("Debugger.enable", nil, nil)
	return
}
func (obj *Debugger) Disable() (err error) {
	err = obj.conn.Send("Debugger.disable", nil, nil)
	return
}

type SetBreakpointsActiveRequest struct {
	Active bool `json:"active"`
}

func (obj *Debugger) SetBreakpointsActive(request *SetBreakpointsActiveRequest) (err error) {
	err = obj.conn.Send("Debugger.setBreakpointsActive", request, nil)
	return
}

type SetSkipAllPausesRequest struct {
	Skip bool `json:"skip"`
}

func (obj *Debugger) SetSkipAllPauses(request *SetSkipAllPausesRequest) (err error) {
	err = obj.conn.Send("Debugger.setSkipAllPauses", request, nil)
	return
}

type SetBreakpointByUrlRequest struct {
	LineNumber   int     `json:"lineNumber"`
	Url          *string `json:"url,omitempty"`
	UrlRegex     *string `json:"urlRegex,omitempty"`
	ScriptHash   *string `json:"scriptHash,omitempty"`
	ColumnNumber *int    `json:"columnNumber,omitempty"`
	Condition    *string `json:"condition,omitempty"`
}

func (obj *Debugger) SetBreakpointByUrl(request *SetBreakpointByUrlRequest) (response struct {
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
	Locations    []types.Debugger_Location   `json:"locations"`
}, err error) {
	err = obj.conn.Send("Debugger.setBreakpointByUrl", request, &response)
	return
}

type SetBreakpointRequest struct {
	Location  types.Debugger_Location `json:"location"`
	Condition *string                 `json:"condition,omitempty"`
}

func (obj *Debugger) SetBreakpoint(request *SetBreakpointRequest) (response struct {
	BreakpointId   types.Debugger_BreakpointId `json:"breakpointId"`
	ActualLocation types.Debugger_Location     `json:"actualLocation"`
}, err error) {
	err = obj.conn.Send("Debugger.setBreakpoint", request, &response)
	return
}

type RemoveBreakpointRequest struct {
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
}

func (obj *Debugger) RemoveBreakpoint(request *RemoveBreakpointRequest) (err error) {
	err = obj.conn.Send("Debugger.removeBreakpoint", request, nil)
	return
}

type GetPossibleBreakpointsRequest struct {
	Start              types.Debugger_Location  `json:"start"`
	End                *types.Debugger_Location `json:"end,omitempty"`
	RestrictToFunction *bool                    `json:"restrictToFunction,omitempty"`
}

func (obj *Debugger) GetPossibleBreakpoints(request *GetPossibleBreakpointsRequest) (response struct {
	Locations []types.Debugger_BreakLocation `json:"locations"`
}, err error) {
	err = obj.conn.Send("Debugger.getPossibleBreakpoints", request, &response)
	return
}

type ContinueToLocationRequest struct {
	Location         types.Debugger_Location `json:"location"`
	TargetCallFrames *string                 `json:"targetCallFrames,omitempty"`
}

func (obj *Debugger) ContinueToLocation(request *ContinueToLocationRequest) (err error) {
	err = obj.conn.Send("Debugger.continueToLocation", request, nil)
	return
}
func (obj *Debugger) StepOver() (err error) {
	err = obj.conn.Send("Debugger.stepOver", nil, nil)
	return
}
func (obj *Debugger) StepInto() (err error) {
	err = obj.conn.Send("Debugger.stepInto", nil, nil)
	return
}
func (obj *Debugger) StepOut() (err error) {
	err = obj.conn.Send("Debugger.stepOut", nil, nil)
	return
}
func (obj *Debugger) Pause() (err error) {
	err = obj.conn.Send("Debugger.pause", nil, nil)
	return
}
func (obj *Debugger) ScheduleStepIntoAsync() (err error) {
	err = obj.conn.Send("Debugger.scheduleStepIntoAsync", nil, nil)
	return
}
func (obj *Debugger) Resume() (err error) {
	err = obj.conn.Send("Debugger.resume", nil, nil)
	return
}

type SearchInContentRequest struct {
	ScriptId      types.Runtime_ScriptId `json:"scriptId"`
	Query         string                 `json:"query"`
	CaseSensitive *bool                  `json:"caseSensitive,omitempty"`
	IsRegex       *bool                  `json:"isRegex,omitempty"`
}

func (obj *Debugger) SearchInContent(request *SearchInContentRequest) (response struct {
	Result []types.Debugger_SearchMatch `json:"result"`
}, err error) {
	err = obj.conn.Send("Debugger.searchInContent", request, &response)
	return
}

type SetScriptSourceRequest struct {
	ScriptId     types.Runtime_ScriptId `json:"scriptId"`
	ScriptSource string                 `json:"scriptSource"`
	DryRun       *bool                  `json:"dryRun,omitempty"`
}

func (obj *Debugger) SetScriptSource(request *SetScriptSourceRequest) (response struct {
	CallFrames       []types.Debugger_CallFrame      `json:"callFrames,omitempty"`
	StackChanged     *bool                           `json:"stackChanged,omitempty"`
	AsyncStackTrace  *types.Runtime_StackTrace       `json:"asyncStackTrace,omitempty"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Debugger.setScriptSource", request, &response)
	return
}

type RestartFrameRequest struct {
	CallFrameId types.Debugger_CallFrameId `json:"callFrameId"`
}

func (obj *Debugger) RestartFrame(request *RestartFrameRequest) (response struct {
	CallFrames      []types.Debugger_CallFrame `json:"callFrames"`
	AsyncStackTrace *types.Runtime_StackTrace  `json:"asyncStackTrace,omitempty"`
}, err error) {
	err = obj.conn.Send("Debugger.restartFrame", request, &response)
	return
}

type GetScriptSourceRequest struct {
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
}

func (obj *Debugger) GetScriptSource(request *GetScriptSourceRequest) (response struct {
	ScriptSource string `json:"scriptSource"`
}, err error) {
	err = obj.conn.Send("Debugger.getScriptSource", request, &response)
	return
}

type SetPauseOnExceptionsRequest struct {
	State string `json:"state"`
}

func (obj *Debugger) SetPauseOnExceptions(request *SetPauseOnExceptionsRequest) (err error) {
	err = obj.conn.Send("Debugger.setPauseOnExceptions", request, nil)
	return
}

type EvaluateOnCallFrameRequest struct {
	CallFrameId           types.Debugger_CallFrameId `json:"callFrameId"`
	Expression            string                     `json:"expression"`
	ObjectGroup           *string                    `json:"objectGroup,omitempty"`
	IncludeCommandLineAPI *bool                      `json:"includeCommandLineAPI,omitempty"`
	Silent                *bool                      `json:"silent,omitempty"`
	ReturnByValue         *bool                      `json:"returnByValue,omitempty"`
	GeneratePreview       *bool                      `json:"generatePreview,omitempty"`
	ThrowOnSideEffect     *bool                      `json:"throwOnSideEffect,omitempty"`
}

func (obj *Debugger) EvaluateOnCallFrame(request *EvaluateOnCallFrameRequest) (response struct {
	Result           types.Runtime_RemoteObject      `json:"result"`
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}, err error) {
	err = obj.conn.Send("Debugger.evaluateOnCallFrame", request, &response)
	return
}

type SetVariableValueRequest struct {
	ScopeNumber  int                        `json:"scopeNumber"`
	VariableName string                     `json:"variableName"`
	NewValue     types.Runtime_CallArgument `json:"newValue"`
	CallFrameId  types.Debugger_CallFrameId `json:"callFrameId"`
}

func (obj *Debugger) SetVariableValue(request *SetVariableValueRequest) (err error) {
	err = obj.conn.Send("Debugger.setVariableValue", request, nil)
	return
}

type SetAsyncCallStackDepthRequest struct {
	MaxDepth int `json:"maxDepth"`
}

func (obj *Debugger) SetAsyncCallStackDepth(request *SetAsyncCallStackDepthRequest) (err error) {
	err = obj.conn.Send("Debugger.setAsyncCallStackDepth", request, nil)
	return
}

type SetBlackboxPatternsRequest struct {
	Patterns []string `json:"patterns"`
}

func (obj *Debugger) SetBlackboxPatterns(request *SetBlackboxPatternsRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxPatterns", request, nil)
	return
}

type SetBlackboxedRangesRequest struct {
	ScriptId  types.Runtime_ScriptId          `json:"scriptId"`
	Positions []types.Debugger_ScriptPosition `json:"positions"`
}

func (obj *Debugger) SetBlackboxedRanges(request *SetBlackboxedRangesRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxedRanges", request, nil)
	return
}
