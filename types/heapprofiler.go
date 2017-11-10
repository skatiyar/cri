package types

type HeapProfiler_HeapSnapshotObjectId string
type HeapProfiler_SamplingHeapProfileNode struct {
	CallFrame	Runtime_CallFrame			`json:"callFrame"`// Function location.
	SelfSize	float32					`json:"selfSize"`// Allocations size in bytes for the node excluding children.
	Children	[]*HeapProfiler_SamplingHeapProfileNode	`json:"children"`// Child nodes.
}
type HeapProfiler_SamplingHeapProfile struct {
	Head HeapProfiler_SamplingHeapProfileNode `json:"head"`
}
