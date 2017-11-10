package types

type HeapProfiler_HeapSnapshotObjectId string
type HeapProfiler_SamplingHeapProfileNode struct {
	// Function location.
	CallFrame Runtime_CallFrame `json:"callFrame"`
	// Allocations size in bytes for the node excluding children.
	SelfSize float32 `json:"selfSize"`
	// Child nodes.
	Children []*HeapProfiler_SamplingHeapProfileNode `json:"children"`
}
type HeapProfiler_SamplingHeapProfile struct {
	Head HeapProfiler_SamplingHeapProfileNode `json:"head"`
}
