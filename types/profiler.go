package types

type Profiler_ProfileNode struct {
	Id            int                         `json:"id"`
	CallFrame     Runtime_CallFrame           `json:"callFrame"`
	HitCount      *int                        `json:"hitCount,omitempty"`
	Children      []int                       `json:"children,omitempty"`
	DeoptReason   *string                     `json:"deoptReason,omitempty"`
	PositionTicks []Profiler_PositionTickInfo `json:"positionTicks,omitempty"`
}
type Profiler_Profile struct {
	Nodes      []Profiler_ProfileNode `json:"nodes"`
	StartTime  float32                `json:"startTime"`
	EndTime    float32                `json:"endTime"`
	Samples    []int                  `json:"samples,omitempty"`
	TimeDeltas []int                  `json:"timeDeltas,omitempty"`
}
type Profiler_PositionTickInfo struct {
	Line  int `json:"line"`
	Ticks int `json:"ticks"`
}
type Profiler_CoverageRange struct {
	StartOffset int `json:"startOffset"`
	EndOffset   int `json:"endOffset"`
	Count       int `json:"count"`
}
type Profiler_FunctionCoverage struct {
	FunctionName    string                   `json:"functionName"`
	Ranges          []Profiler_CoverageRange `json:"ranges"`
	IsBlockCoverage bool                     `json:"isBlockCoverage"`
}
type Profiler_ScriptCoverage struct {
	ScriptId  Runtime_ScriptId            `json:"scriptId"`
	Url       string                      `json:"url"`
	Functions []Profiler_FunctionCoverage `json:"functions"`
}
type Profiler_TypeObject struct {
	Name string `json:"name"`
}
type Profiler_TypeProfileEntry struct {
	Offset int                   `json:"offset"`
	Types  []Profiler_TypeObject `json:"types"`
}
type Profiler_ScriptTypeProfile struct {
	ScriptId Runtime_ScriptId            `json:"scriptId"`
	Url      string                      `json:"url"`
	Entries  []Profiler_TypeProfileEntry `json:"entries"`
}
