/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Memory pressure level.
type Memory_PressureLevel string

// Heap profile sample.
type Memory_SamplingProfileNode struct {
	// Size of the sampled allocation.
	Size float32 `json:"size"`
	// Total bytes attributed to this sample.
	Total float32 `json:"total"`
	// Execution stack at the point of allocation.
	Stack []string `json:"stack"`
}

// Array of heap profile samples.
type Memory_SamplingProfile struct {
	Samples []Memory_SamplingProfileNode `json:"samples"`
}
