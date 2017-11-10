package types

type HeapProfiler_HeapSnapshotObjectId string
type HeapProfiler_SamplingHeapProfileNode struct {
	CallFrame Runtime_CallFrame                       `json:"callFrame"`
	SelfSize  float32                                 `json:"selfSize"`
	Children  []*HeapProfiler_SamplingHeapProfileNode `json:"children"`
}
type HeapProfiler_SamplingHeapProfile struct {
	Head HeapProfiler_SamplingHeapProfileNode `json:"head"`
}
