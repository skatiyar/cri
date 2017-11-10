package types

type Profiler_ProfileNode struct {
	Id		int				`json:"id"`// Unique id of the node.
	CallFrame	Runtime_CallFrame		`json:"callFrame"`// Function location.
	HitCount	*int				`json:"hitCount,omitempty"`// Number of samples where this node was on top of the call stack.
	Children	[]int				`json:"children,omitempty"`// Child node ids.
	DeoptReason	*string				`json:"deoptReason,omitempty"`// The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	PositionTicks	[]Profiler_PositionTickInfo	`json:"positionTicks,omitempty"`// An array of source position ticks.
}
type Profiler_Profile struct {
	Nodes		[]Profiler_ProfileNode	`json:"nodes"`// The list of profile nodes. First item is the root node.
	StartTime	float32			`json:"startTime"`// Profiling start timestamp in microseconds.
	EndTime		float32			`json:"endTime"`// Profiling end timestamp in microseconds.
	Samples		[]int			`json:"samples,omitempty"`// Ids of samples top nodes.
	TimeDeltas	[]int			`json:"timeDeltas,omitempty"`// Time intervals between adjacent samples in microseconds. The first delta is relative to the profile startTime.
}
type Profiler_PositionTickInfo struct {
	Line	int	`json:"line"`// Source line number (1-based).
	Ticks	int	`json:"ticks"`// Number of samples attributed to the source line.
}
type Profiler_CoverageRange struct {
	StartOffset	int	`json:"startOffset"`// JavaScript script source offset for the range start.
	EndOffset	int	`json:"endOffset"`// JavaScript script source offset for the range end.
	Count		int	`json:"count"`// Collected execution count of the source range.
}
type Profiler_FunctionCoverage struct {
	FunctionName	string				`json:"functionName"`// JavaScript function name.
	Ranges		[]Profiler_CoverageRange	`json:"ranges"`// Source ranges inside the function with coverage data.
	IsBlockCoverage	bool				`json:"isBlockCoverage"`// Whether coverage data for this function has block granularity.
}
type Profiler_ScriptCoverage struct {
	ScriptId	Runtime_ScriptId		`json:"scriptId"`// JavaScript script id.
	Url		string				`json:"url"`// JavaScript script name or url.
	Functions	[]Profiler_FunctionCoverage	`json:"functions"`// Functions contained in the script that has coverage data.
}
type Profiler_TypeObject struct {
	Name string `json:"name"`// Name of a type collected with type profiling.
}
type Profiler_TypeProfileEntry struct {
	Offset	int			`json:"offset"`// Source offset of the parameter or end of function for return values.
	Types	[]Profiler_TypeObject	`json:"types"`// The types for this parameter or return value.
}
type Profiler_ScriptTypeProfile struct {
	ScriptId	Runtime_ScriptId		`json:"scriptId"`// JavaScript script id.
	Url		string				`json:"url"`// JavaScript script name or url.
	Entries		[]Profiler_TypeProfileEntry	`json:"entries"`// Type profile entries for parameters and return values of the functions in the script.
}
