/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

// Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
package debugger

import (
    "github.com/SKatiyar/cri"
    types "github.com/SKatiyar/cri/types"
)

type Debugger struct {
	conn cri.Connector
}

// New creates a Debugger instance
func New(conn cri.Connector) *Debugger {
	return &Debugger{conn}
}
// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (obj *Debugger) Enable() (err error) {
	err = obj.conn.Send("Debugger.enable", nil, nil)
	return
}

// Disables debugger for given page.
func (obj *Debugger) Disable() (err error) {
	err = obj.conn.Send("Debugger.disable", nil, nil)
	return
}


type SetBreakpointsActiveRequest struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

// Activates / deactivates all breakpoints on the page.
func (obj *Debugger) SetBreakpointsActive(request *SetBreakpointsActiveRequest) (err error) {
	err = obj.conn.Send("Debugger.setBreakpointsActive", request, nil)
	return
}


type SetSkipAllPausesRequest struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (obj *Debugger) SetSkipAllPauses(request *SetSkipAllPausesRequest) (err error) {
	err = obj.conn.Send("Debugger.setSkipAllPauses", request, nil)
	return
}


type SetBreakpointByUrlRequest struct {
	// Line number to set breakpoint at.
	LineNumber int `json:"lineNumber"`
	// URL of the resources to set breakpoint on.
	Url *string `json:"url,omitempty"`
	// Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
	UrlRegex *string `json:"urlRegex,omitempty"`
	// Script hash of the resources to set breakpoint on.
	// NOTE Experimental
	ScriptHash *string `json:"scriptHash,omitempty"`
	// Offset in the line to set breakpoint at.
	ColumnNumber *int `json:"columnNumber,omitempty"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition *string `json:"condition,omitempty"`
}


type SetBreakpointByUrlResponse struct {
	// Id of the created breakpoint for further reference.
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
	// List of the locations this breakpoint resolved into upon addition.
	Locations []types.Debugger_Location `json:"locations"`
}

// Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
func (obj *Debugger) SetBreakpointByUrl(request *SetBreakpointByUrlRequest) (response SetBreakpointByUrlResponse, err error) {
	err = obj.conn.Send("Debugger.setBreakpointByUrl", request, &response)
	return
}


type SetBreakpointRequest struct {
	// Location to set breakpoint in.
	Location types.Debugger_Location `json:"location"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition *string `json:"condition,omitempty"`
}


type SetBreakpointResponse struct {
	// Id of the created breakpoint for further reference.
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
	// Location this breakpoint resolved into.
	ActualLocation types.Debugger_Location `json:"actualLocation"`
}

// Sets JavaScript breakpoint at a given location.
func (obj *Debugger) SetBreakpoint(request *SetBreakpointRequest) (response SetBreakpointResponse, err error) {
	err = obj.conn.Send("Debugger.setBreakpoint", request, &response)
	return
}


type RemoveBreakpointRequest struct {
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
}

// Removes JavaScript breakpoint.
func (obj *Debugger) RemoveBreakpoint(request *RemoveBreakpointRequest) (err error) {
	err = obj.conn.Send("Debugger.removeBreakpoint", request, nil)
	return
}


type GetPossibleBreakpointsRequest struct {
	// Start of range to search possible breakpoint locations in.
	Start types.Debugger_Location `json:"start"`
	// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	End *types.Debugger_Location `json:"end,omitempty"`
	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction *bool `json:"restrictToFunction,omitempty"`
}


type GetPossibleBreakpointsResponse struct {
	// List of the possible breakpoint locations.
	Locations []types.Debugger_BreakLocation `json:"locations"`
}

// Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
func (obj *Debugger) GetPossibleBreakpoints(request *GetPossibleBreakpointsRequest) (response GetPossibleBreakpointsResponse, err error) {
	err = obj.conn.Send("Debugger.getPossibleBreakpoints", request, &response)
	return
}


type ContinueToLocationRequest struct {
	// Location to continue to.
	Location types.Debugger_Location `json:"location"`
	// NOTE Experimental
	TargetCallFrames *string `json:"targetCallFrames,omitempty"`
}

// Continues execution until specific location is reached.
func (obj *Debugger) ContinueToLocation(request *ContinueToLocationRequest) (err error) {
	err = obj.conn.Send("Debugger.continueToLocation", request, nil)
	return
}

// Steps over the statement.
func (obj *Debugger) StepOver() (err error) {
	err = obj.conn.Send("Debugger.stepOver", nil, nil)
	return
}

// Steps into the function call.
func (obj *Debugger) StepInto() (err error) {
	err = obj.conn.Send("Debugger.stepInto", nil, nil)
	return
}

// Steps out of the function call.
func (obj *Debugger) StepOut() (err error) {
	err = obj.conn.Send("Debugger.stepOut", nil, nil)
	return
}

// Stops on the next JavaScript statement.
func (obj *Debugger) Pause() (err error) {
	err = obj.conn.Send("Debugger.pause", nil, nil)
	return
}

// Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called.
func (obj *Debugger) ScheduleStepIntoAsync() (err error) {
	err = obj.conn.Send("Debugger.scheduleStepIntoAsync", nil, nil)
	return
}

// Resumes JavaScript execution.
func (obj *Debugger) Resume() (err error) {
	err = obj.conn.Send("Debugger.resume", nil, nil)
	return
}


type SearchInContentRequest struct {
	// Id of the script to search in.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex *bool `json:"isRegex,omitempty"`
}


type SearchInContentResponse struct {
	// List of search matches.
	Result []types.Debugger_SearchMatch `json:"result"`
}

// Searches for given string in script content.
func (obj *Debugger) SearchInContent(request *SearchInContentRequest) (response SearchInContentResponse, err error) {
	err = obj.conn.Send("Debugger.searchInContent", request, &response)
	return
}


type SetScriptSourceRequest struct {
	// Id of the script to edit.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// New content of the script.
	ScriptSource string `json:"scriptSource"`
	//  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
	DryRun *bool `json:"dryRun,omitempty"`
}


type SetScriptSourceResponse struct {
	// New stack trace in case editing has happened while VM was stopped.
	CallFrames []types.Debugger_CallFrame `json:"callFrames,omitempty"`
	// Whether current call stack  was modified after applying the changes.
	StackChanged *bool `json:"stackChanged,omitempty"`
	// Async stack trace, if any.
	AsyncStackTrace *types.Runtime_StackTrace `json:"asyncStackTrace,omitempty"`
	// Exception details if any.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Edits JavaScript source live.
func (obj *Debugger) SetScriptSource(request *SetScriptSourceRequest) (response SetScriptSourceResponse, err error) {
	err = obj.conn.Send("Debugger.setScriptSource", request, &response)
	return
}


type RestartFrameRequest struct {
	// Call frame identifier to evaluate on.
	CallFrameId types.Debugger_CallFrameId `json:"callFrameId"`
}


type RestartFrameResponse struct {
	// New stack trace.
	CallFrames []types.Debugger_CallFrame `json:"callFrames"`
	// Async stack trace, if any.
	AsyncStackTrace *types.Runtime_StackTrace `json:"asyncStackTrace,omitempty"`
}

// Restarts particular call frame from the beginning.
func (obj *Debugger) RestartFrame(request *RestartFrameRequest) (response RestartFrameResponse, err error) {
	err = obj.conn.Send("Debugger.restartFrame", request, &response)
	return
}


type GetScriptSourceRequest struct {
	// Id of the script to get source for.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
}


type GetScriptSourceResponse struct {
	// Script source.
	ScriptSource string `json:"scriptSource"`
}

// Returns source for the script with given id.
func (obj *Debugger) GetScriptSource(request *GetScriptSourceRequest) (response GetScriptSourceResponse, err error) {
	err = obj.conn.Send("Debugger.getScriptSource", request, &response)
	return
}


type SetPauseOnExceptionsRequest struct {
	// Pause on exceptions mode.
	State string `json:"state"`
}

// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
func (obj *Debugger) SetPauseOnExceptions(request *SetPauseOnExceptionsRequest) (err error) {
	err = obj.conn.Send("Debugger.setPauseOnExceptions", request, nil)
	return
}


type EvaluateOnCallFrameRequest struct {
	// Call frame identifier to evaluate on.
	CallFrameId types.Debugger_CallFrameId `json:"callFrameId"`
	// Expression to evaluate.
	Expression string `json:"expression"`
	// String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
	ObjectGroup *string `json:"objectGroup,omitempty"`
	// Specifies whether command line API should be available to the evaluated expression, defaults to false.
	IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent *bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue *bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	// NOTE Experimental
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	// NOTE Experimental
	ThrowOnSideEffect *bool `json:"throwOnSideEffect,omitempty"`
}


type EvaluateOnCallFrameResponse struct {
	// Object wrapper for the evaluation result.
	Result types.Runtime_RemoteObject `json:"result"`
	// Exception details.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Evaluates expression on a given call frame.
func (obj *Debugger) EvaluateOnCallFrame(request *EvaluateOnCallFrameRequest) (response EvaluateOnCallFrameResponse, err error) {
	err = obj.conn.Send("Debugger.evaluateOnCallFrame", request, &response)
	return
}


type SetVariableValueRequest struct {
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber int `json:"scopeNumber"`
	// Variable name.
	VariableName string `json:"variableName"`
	// New variable value.
	NewValue types.Runtime_CallArgument `json:"newValue"`
	// Id of callframe that holds variable.
	CallFrameId types.Debugger_CallFrameId `json:"callFrameId"`
}

// Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
func (obj *Debugger) SetVariableValue(request *SetVariableValueRequest) (err error) {
	err = obj.conn.Send("Debugger.setVariableValue", request, nil)
	return
}


type SetAsyncCallStackDepthRequest struct {
	// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// Enables or disables async call stacks tracking.
func (obj *Debugger) SetAsyncCallStackDepth(request *SetAsyncCallStackDepthRequest) (err error) {
	err = obj.conn.Send("Debugger.setAsyncCallStackDepth", request, nil)
	return
}


type SetBlackboxPatternsRequest struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
func (obj *Debugger) SetBlackboxPatterns(request *SetBlackboxPatternsRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxPatterns", request, nil)
	return
}


type SetBlackboxedRangesRequest struct {
	// Id of the script.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	Positions []types.Debugger_ScriptPosition `json:"positions"`
}

// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
func (obj *Debugger) SetBlackboxedRanges(request *SetBlackboxedRangesRequest) (err error) {
	err = obj.conn.Send("Debugger.setBlackboxedRanges", request, nil)
	return
}
