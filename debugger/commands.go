/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package debugger provides commands and events for Debugger domain.
package debugger

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Debugger domain
const (
	Enable                 = "Debugger.enable"
	Disable                = "Debugger.disable"
	SetBreakpointsActive   = "Debugger.setBreakpointsActive"
	SetSkipAllPauses       = "Debugger.setSkipAllPauses"
	SetBreakpointByUrl     = "Debugger.setBreakpointByUrl"
	SetBreakpoint          = "Debugger.setBreakpoint"
	RemoveBreakpoint       = "Debugger.removeBreakpoint"
	GetPossibleBreakpoints = "Debugger.getPossibleBreakpoints"
	ContinueToLocation     = "Debugger.continueToLocation"
	PauseOnAsyncCall       = "Debugger.pauseOnAsyncCall"
	StepOver               = "Debugger.stepOver"
	StepInto               = "Debugger.stepInto"
	StepOut                = "Debugger.stepOut"
	Pause                  = "Debugger.pause"
	ScheduleStepIntoAsync  = "Debugger.scheduleStepIntoAsync"
	Resume                 = "Debugger.resume"
	GetStackTrace          = "Debugger.getStackTrace"
	SearchInContent        = "Debugger.searchInContent"
	SetScriptSource        = "Debugger.setScriptSource"
	RestartFrame           = "Debugger.restartFrame"
	GetScriptSource        = "Debugger.getScriptSource"
	SetPauseOnExceptions   = "Debugger.setPauseOnExceptions"
	EvaluateOnCallFrame    = "Debugger.evaluateOnCallFrame"
	SetVariableValue       = "Debugger.setVariableValue"
	SetReturnValue         = "Debugger.setReturnValue"
	SetAsyncCallStackDepth = "Debugger.setAsyncCallStackDepth"
	SetBlackboxPatterns    = "Debugger.setBlackboxPatterns"
	SetBlackboxedRanges    = "Debugger.setBlackboxedRanges"
)

// List of events in Debugger domain
const (
	ScriptParsed        = "Debugger.scriptParsed"
	ScriptFailedToParse = "Debugger.scriptFailedToParse"
	BreakpointResolved  = "Debugger.breakpointResolved"
	Paused              = "Debugger.paused"
	Resumed             = "Debugger.resumed"
)

// Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
type Debugger struct {
	conn cri.Connector
}

// New creates a Debugger instance
func New(conn cri.Connector) *Debugger {
	return &Debugger{conn}
}

type EnableResponse struct {
	// Unique identifier of the debugger.
	// NOTE Experimental
	DebuggerId types.Runtime_UniqueDebuggerId `json:"debuggerId"`
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (obj *Debugger) Enable() (response EnableResponse, err error) {
	err = obj.conn.Send(Enable, nil, &response)
	return
}

// Disables debugger for given page.
func (obj *Debugger) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type SetBreakpointsActiveRequest struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

// Activates / deactivates all breakpoints on the page.
func (obj *Debugger) SetBreakpointsActive(request *SetBreakpointsActiveRequest) (err error) {
	err = obj.conn.Send(SetBreakpointsActive, request, nil)
	return
}

type SetSkipAllPausesRequest struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

// Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (obj *Debugger) SetSkipAllPauses(request *SetSkipAllPausesRequest) (err error) {
	err = obj.conn.Send(SetSkipAllPauses, request, nil)
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
	err = obj.conn.Send(SetBreakpointByUrl, request, &response)
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
	err = obj.conn.Send(SetBreakpoint, request, &response)
	return
}

type RemoveBreakpointRequest struct {
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
}

// Removes JavaScript breakpoint.
func (obj *Debugger) RemoveBreakpoint(request *RemoveBreakpointRequest) (err error) {
	err = obj.conn.Send(RemoveBreakpoint, request, nil)
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
	err = obj.conn.Send(GetPossibleBreakpoints, request, &response)
	return
}

type ContinueToLocationRequest struct {
	// Location to continue to.
	Location         types.Debugger_Location `json:"location"`
	TargetCallFrames *string                 `json:"targetCallFrames,omitempty"`
}

// Continues execution until specific location is reached.
func (obj *Debugger) ContinueToLocation(request *ContinueToLocationRequest) (err error) {
	err = obj.conn.Send(ContinueToLocation, request, nil)
	return
}

type PauseOnAsyncCallRequest struct {
	// Debugger will pause when async call with given stack trace is started.
	ParentStackTraceId types.Runtime_StackTraceId `json:"parentStackTraceId"`
}

func (obj *Debugger) PauseOnAsyncCall(request *PauseOnAsyncCallRequest) (err error) {
	err = obj.conn.Send(PauseOnAsyncCall, request, nil)
	return
}

// Steps over the statement.
func (obj *Debugger) StepOver() (err error) {
	err = obj.conn.Send(StepOver, nil, nil)
	return
}

type StepIntoRequest struct {
	// Debugger will issue additional Debugger.paused notification if any async task is scheduled before next pause.
	// NOTE Experimental
	BreakOnAsyncCall *bool `json:"breakOnAsyncCall,omitempty"`
}

// Steps into the function call.
func (obj *Debugger) StepInto(request *StepIntoRequest) (err error) {
	err = obj.conn.Send(StepInto, request, nil)
	return
}

// Steps out of the function call.
func (obj *Debugger) StepOut() (err error) {
	err = obj.conn.Send(StepOut, nil, nil)
	return
}

// Stops on the next JavaScript statement.
func (obj *Debugger) Pause() (err error) {
	err = obj.conn.Send(Pause, nil, nil)
	return
}

// This method is deprecated - use Debugger.stepInto with breakOnAsyncCall and Debugger.pauseOnAsyncTask instead. Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called.
func (obj *Debugger) ScheduleStepIntoAsync() (err error) {
	err = obj.conn.Send(ScheduleStepIntoAsync, nil, nil)
	return
}

// Resumes JavaScript execution.
func (obj *Debugger) Resume() (err error) {
	err = obj.conn.Send(Resume, nil, nil)
	return
}

type GetStackTraceRequest struct {
	StackTraceId types.Runtime_StackTraceId `json:"stackTraceId"`
}

type GetStackTraceResponse struct {
	StackTrace types.Runtime_StackTrace `json:"stackTrace"`
}

// Returns stack trace with given <code>stackTraceId</code>.
func (obj *Debugger) GetStackTrace(request *GetStackTraceRequest) (response GetStackTraceResponse, err error) {
	err = obj.conn.Send(GetStackTrace, request, &response)
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
	err = obj.conn.Send(SearchInContent, request, &response)
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
	// Async stack trace, if any.
	// NOTE Experimental
	AsyncStackTraceId *types.Runtime_StackTraceId `json:"asyncStackTraceId,omitempty"`
	// Exception details if any.
	ExceptionDetails *types.Runtime_ExceptionDetails `json:"exceptionDetails,omitempty"`
}

// Edits JavaScript source live.
func (obj *Debugger) SetScriptSource(request *SetScriptSourceRequest) (response SetScriptSourceResponse, err error) {
	err = obj.conn.Send(SetScriptSource, request, &response)
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
	// Async stack trace, if any.
	// NOTE Experimental
	AsyncStackTraceId *types.Runtime_StackTraceId `json:"asyncStackTraceId,omitempty"`
}

// Restarts particular call frame from the beginning.
func (obj *Debugger) RestartFrame(request *RestartFrameRequest) (response RestartFrameResponse, err error) {
	err = obj.conn.Send(RestartFrame, request, &response)
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
	err = obj.conn.Send(GetScriptSource, request, &response)
	return
}

type SetPauseOnExceptionsRequest struct {
	// Pause on exceptions mode.
	State string `json:"state"`
}

// Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
func (obj *Debugger) SetPauseOnExceptions(request *SetPauseOnExceptionsRequest) (err error) {
	err = obj.conn.Send(SetPauseOnExceptions, request, nil)
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
	err = obj.conn.Send(EvaluateOnCallFrame, request, &response)
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
	err = obj.conn.Send(SetVariableValue, request, nil)
	return
}

type SetReturnValueRequest struct {
	// New return value.
	NewValue types.Runtime_CallArgument `json:"newValue"`
}

// Changes return value in top frame. Available only at return break position.
func (obj *Debugger) SetReturnValue(request *SetReturnValueRequest) (err error) {
	err = obj.conn.Send(SetReturnValue, request, nil)
	return
}

type SetAsyncCallStackDepthRequest struct {
	// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// Enables or disables async call stacks tracking.
func (obj *Debugger) SetAsyncCallStackDepth(request *SetAsyncCallStackDepthRequest) (err error) {
	err = obj.conn.Send(SetAsyncCallStackDepth, request, nil)
	return
}

type SetBlackboxPatternsRequest struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

// Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
func (obj *Debugger) SetBlackboxPatterns(request *SetBlackboxPatternsRequest) (err error) {
	err = obj.conn.Send(SetBlackboxPatterns, request, nil)
	return
}

type SetBlackboxedRangesRequest struct {
	// Id of the script.
	ScriptId  types.Runtime_ScriptId          `json:"scriptId"`
	Positions []types.Debugger_ScriptPosition `json:"positions"`
}

// Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
func (obj *Debugger) SetBlackboxedRanges(request *SetBlackboxedRangesRequest) (err error) {
	err = obj.conn.Send(SetBlackboxedRanges, request, nil)
	return
}

type ScriptParsedParams struct {
	// Identifier of the script parsed.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// URL or name of the script parsed (if any).
	Url string `json:"url"`
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`
	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`
	// Last line of the script.
	EndLine int `json:"endLine"`
	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`
	// Specifies script creation context.
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
	// Content hash of the script.
	Hash string `json:"hash"`
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"`
	// True, if this script is generated as a result of the live edit operation.
	// NOTE Experimental
	IsLiveEdit *bool `json:"isLiveEdit,omitempty"`
	// URL of source map associated with script (if any).
	SourceMapURL *string `json:"sourceMapURL,omitempty"`
	// True, if this script has sourceURL.
	HasSourceURL *bool `json:"hasSourceURL,omitempty"`
	// True, if this script is ES6 module.
	IsModule *bool `json:"isModule,omitempty"`
	// This script length.
	Length *int `json:"length,omitempty"`
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	// NOTE Experimental
	StackTrace *types.Runtime_StackTrace `json:"stackTrace,omitempty"`
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
func (obj *Debugger) ScriptParsed(fn func(params *ScriptParsedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ScriptParsed, closeChn)
	go func() {
		for {
			params := ScriptParsedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ScriptFailedToParseParams struct {
	// Identifier of the script parsed.
	ScriptId types.Runtime_ScriptId `json:"scriptId"`
	// URL or name of the script parsed (if any).
	Url string `json:"url"`
	// Line offset of the script within the resource with given URL (for script tags).
	StartLine int `json:"startLine"`
	// Column offset of the script within the resource with given URL.
	StartColumn int `json:"startColumn"`
	// Last line of the script.
	EndLine int `json:"endLine"`
	// Length of the last line of the script.
	EndColumn int `json:"endColumn"`
	// Specifies script creation context.
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
	// Content hash of the script.
	Hash string `json:"hash"`
	// Embedder-specific auxiliary data.
	ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"`
	// URL of source map associated with script (if any).
	SourceMapURL *string `json:"sourceMapURL,omitempty"`
	// True, if this script has sourceURL.
	HasSourceURL *bool `json:"hasSourceURL,omitempty"`
	// True, if this script is ES6 module.
	IsModule *bool `json:"isModule,omitempty"`
	// This script length.
	Length *int `json:"length,omitempty"`
	// JavaScript top stack frame of where the script parsed event was triggered if available.
	// NOTE Experimental
	StackTrace *types.Runtime_StackTrace `json:"stackTrace,omitempty"`
}

// Fired when virtual machine fails to parse the script.
func (obj *Debugger) ScriptFailedToParse(fn func(params *ScriptFailedToParseParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ScriptFailedToParse, closeChn)
	go func() {
		for {
			params := ScriptFailedToParseParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type BreakpointResolvedParams struct {
	// Breakpoint unique identifier.
	BreakpointId types.Debugger_BreakpointId `json:"breakpointId"`
	// Actual breakpoint location.
	Location types.Debugger_Location `json:"location"`
}

// Fired when breakpoint is resolved to an actual script and location.
func (obj *Debugger) BreakpointResolved(fn func(params *BreakpointResolvedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(BreakpointResolved, closeChn)
	go func() {
		for {
			params := BreakpointResolvedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type PausedParams struct {
	// Call stack the virtual machine stopped on.
	CallFrames []types.Debugger_CallFrame `json:"callFrames"`
	// Pause reason.
	Reason string `json:"reason"`
	// Object containing break-specific auxiliary properties.
	Data map[string]interface{} `json:"data,omitempty"`
	// Hit breakpoints IDs
	HitBreakpoints []string `json:"hitBreakpoints,omitempty"`
	// Async stack trace, if any.
	AsyncStackTrace *types.Runtime_StackTrace `json:"asyncStackTrace,omitempty"`
	// Async stack trace, if any.
	// NOTE Experimental
	AsyncStackTraceId *types.Runtime_StackTraceId `json:"asyncStackTraceId,omitempty"`
	// Just scheduled async call will have this stack trace as parent stack during async execution. This field is available only after <code>Debugger.stepInto</code> call with <code>breakOnAsynCall</code> flag.
	// NOTE Experimental
	AsyncCallStackTraceId *types.Runtime_StackTraceId `json:"asyncCallStackTraceId,omitempty"`
}

// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
func (obj *Debugger) Paused(fn func(params *PausedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(Paused, closeChn)
	go func() {
		for {
			params := PausedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

// Fired when the virtual machine resumed execution.
func (obj *Debugger) Resumed(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(Resumed, closeChn)
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
