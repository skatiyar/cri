/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Heap snapshot object id.
type HeapProfiler_HeapSnapshotObjectId string

//Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
type HeapProfiler_SamplingHeapProfileNode struct {
	// Function location.
	CallFrame Runtime_CallFrame `json:"callFrame"`
	// Allocations size in bytes for the node excluding children.
	SelfSize float32 `json:"selfSize"`
	// Child nodes.
	Children []*HeapProfiler_SamplingHeapProfileNode `json:"children"`
}

//Profile.
type HeapProfiler_SamplingHeapProfile struct {
	Head HeapProfiler_SamplingHeapProfileNode `json:"head"`
}
