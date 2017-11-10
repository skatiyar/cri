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
	Active bool `json:"active"`// New value for breakpoints active state.
}

func (obj *Debugger) SetBreakpointsActive(request *SetBreakpointsActiveRequest) (err error) {
	err = obj.conn.Send("Debugger.setBreakpointsActive", request, nil)
	return
}

type SetSkipAllPausesRequest struct {
	Skip bool `json:"skip"`// New value for skip pauses state.
}

func (obj *Debugger) SetSkipAllPauses(request *SetSkipAllPausesRequest) (err error) {
	err = obj.conn.Send("Debugger.setSkipAllPauses", request, nil)
	return
}

type SetBreakpointByUrlRequest struct {
	LineNumber	int	`json:"lineNumber"`// Line number to set breakpoint at.
	Url		*string	`json:"url,omitempty"`// URL of the resources to set breakpoint on.
	UrlRegex	*string	`json:"urlRegex,omitempty"`// Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
	ScriptHash	*string	`json:"scriptHash,omitempty"`// Script hash of the resources to set breakpoint on.
	ColumnNumber	*int	`json:"columnNumber,omitempty"`// Offset in the line to set breakpoint at.
	Condition	*string	`json:"condition,omitempty"`// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

func (obj *Debugger) SetBreakpointByUrl(request *SetBreakpointByUrlRequest) (response struct {
	BreakpointId	types.Debugger_BreakpointId	`json:"breakpointId"`// Id of the created breakpoint for further reference.
	Locations	[]types.Debugger_Location	`json:"locations"`// List of the locations this breakpoint resolved into upon addition.
}, err error) {
	err = obj.conn.Send("Debugger.setBreakpointByUrl", request, &response)
	return
}

type SetBreakpointRequest struct {
	Location	types.Debugger_Location	`json:"location"`// Location to set breakpoint in.
	Condition	*string			`json:"condition,omitempty"`// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}

func (obj *Debugger) SetBreakpoint(request *SetBreakpointRequest) (response struct {
	BreakpointId	types.Debugger_BreakpointId	`json:"breakpointId"`// Id of the created breakpoint for further reference.
	ActualLocation	types.Debugger_Location		`json:"actualLocation"`// Location this breakpoint resolved into.
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
	Start			types.Debugger_Location		`json:"start"`// Start of range to search possible breakpoint locations in.
	End			*types.Debugger_Location	`json:"end,omitempty"`// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	RestrictToFunction	*bool				`json:"restrictToFunction,omitempty"`// Only consider locations which are in the same (non-nested) function as start.
}

func (obj *Debugger) GetPossibleBreakpoints(request *GetPossibleBreakpointsRequest) (response struct {
	Locations []types.Debugger_BreakLocation `json:"locations"`// List of the possible breakpoint locations.
}, err error) {
	err = obj.conn.Send("Debugger.getPossibleBreakpoints", request, &response)
	return
}

type ContinueToLocationRequest struct {
	Location		types.Debugger_Location	`json:"location"`// Location to continue to.
	TargetCallFrames	*string			`json:"targetCallFrames,omitempty"`
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
	ScriptId	types.Runtime_ScriptId	`json:"scriptId"`// Id of the script to search in.
	Query		string			`json:"query"`// String to search for.
	CaseSensitive	*bool			`json:"caseSensitive,omitempty"`// If true, search is case sensitive.
	IsRegex		*bool			`json:"isRegex,omitempty"`// If true, treats string parameter as regex.
}

func (obj *Debugger) SearchInContent(request *SearchInContentRequest) (response struct {
	Result []types.Debugger_SearchMatch `json:"result"`// List of search matches.
}, err error) {
	err = obj.conn.Send("Debugger.searchInContent", request, &response)
	return
}

type SetScriptSourceRequest struct {
	ScriptId	types.Runtime_ScriptId	`json:"scriptId"`// Id of the script to edit.
	ScriptSource	string			`json:"scriptSource"`// New content of the script.
	DryRun		*bool			`json:"dryRun,omitempty"`//  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
}

func (obj *Debugger) SetScriptSource(request *SetScriptSourceRequest) (response struct {
	CallFrames		[]types.Debugger_CallFrame	`json:"callFrames,omitempty"`// New stack trace in case editing has happened while VM was stopped.
	StackChanged		*bool				`json:"stackChanged,omitempty"`// Whether current call stack  was modified after applying the changes.
	AsyncStackTrace		*types.Runtime_StackTrace	`json:"asyncStackTrace,omitempty"`// Async stack trace, if any.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details if any.
}, err error) {
	err = obj.conn.Send("Debugger.setScriptSource", request, &response)
	return
}

type RestartFrameRequest struct {
	CallFrameId types.Debugger_CallFrameId `json:"callFrameId"`// Call frame identifier to evaluate on.
}

func (obj *Debugger) RestartFrame(request *RestartFrameRequest) (response struct {
	CallFrames	[]types.Debugger_CallFrame	`json:"callFrames"`// New stack trace.
	AsyncStackTrace	*types.Runtime_StackTrace	`json:"asyncStackTrace,omitempty"`// Async stack trace, if any.
}, err error) {
	err = obj.conn.Send("Debugger.restartFrame", request, &response)
	return
}

type GetScriptSourceRequest struct {
	ScriptId types.Runtime_ScriptId `json:"scriptId"`// Id of the script to get source for.
}

func (obj *Debugger) GetScriptSource(request *GetScriptSourceRequest) (response struct {
	ScriptSource string `json:"scriptSource"`// Script source.
}, err error) {
	err = obj.conn.Send("Debugger.getScriptSource", request, &response)
	return
}

type SetPauseOnExceptionsRequest struct {
	State string `json:"state"`// Pause on exceptions mode.
}

func (obj *Debugger) SetPauseOnExceptions(request *SetPauseOnExceptionsRequest) (err error) {
	err = obj.conn.Send("Debugger.setPauseOnExceptions", request, nil)
	return
}

type EvaluateOnCallFrameRequest struct {
	CallFrameId		types.Debugger_CallFrameId	`json:"callFrameId"`// Call frame identifier to evaluate on.
	Expression		string				`json:"expression"`// Expression to evaluate.
	ObjectGroup		*string				`json:"objectGroup,omitempty"`// String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
	IncludeCommandLineAPI	*bool				`json:"includeCommandLineAPI,omitempty"`// Specifies whether command line API should be available to the evaluated expression, defaults to false.
	Silent			*bool				`json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	ReturnByValue		*bool				`json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview		*bool				`json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
	ThrowOnSideEffect	*bool				`json:"throwOnSideEffect,omitempty"`// Whether to throw an exception if side effect cannot be ruled out during evaluation.
}

func (obj *Debugger) EvaluateOnCallFrame(request *EvaluateOnCallFrameRequest) (response struct {
	Result			types.Runtime_RemoteObject	`json:"result"`// Object wrapper for the evaluation result.
	ExceptionDetails	*types.Runtime_ExceptionDetails	`json:"exceptionDetails,omitempty"`// Exception details.
}, err error) {
	err = obj.conn.Send("Debugger.evaluateOnCallFrame", request, &response)
	return
}

type SetVariableValueRequest struct {
	ScopeNumber	int				`json:"scopeNumber"`// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	VariableName	string				`json:"variableName"`// Variable name.
	NewValue	types.Runtime_CallArgument	`json:"newValue"`// New variable value.
	CallFrameId	types.Debugger_CallFrameId	`json:"callFrameId"`// Id of callframe that holds variable.
}

func (obj *Debugger) SetVariableValue(request *SetVariableValueRequest) (err error) {
	err = obj.conn.Send("Debugger.setVariableValue", request, nil)
	return
}

type SetAsyncCallStackDepthRequest struct {
	MaxDepth int `json:"maxDepth"`// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
}

func (obj *Debugger) SetAsyncCallStackDepth(request *SetAsyncCallStackDepthRequest) (err error) {
	err = obj.conn.Send("Debugger.setAsyncCallStackDepth", request, nil)
	return
}

type SetBlackboxPatternsRequest struct {
	Patterns []string `json:"patterns"`// Array of regexps that will be used to check script url for blackbox state.
}

func (obj *Debugger) SetBlackboxPatterns(request *SetBlackboxPatternsRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxPatterns", request, nil)
	return
}

type SetBlackboxedRangesRequest struct {
	ScriptId	types.Runtime_ScriptId		`json:"scriptId"`// Id of the script.
	Positions	[]types.Debugger_ScriptPosition	`json:"positions"`
}

func (obj *Debugger) SetBlackboxedRanges(request *SetBlackboxedRangesRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxedRanges", request, nil)
	return
}
