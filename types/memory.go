/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
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
	Modules []Memory_Module              `json:"modules"`
}

// Executable module information
type Memory_Module struct {
	// Name of the module.
	Name string `json:"name"`
	// UUID of the module.
	Uuid string `json:"uuid"`
	// Base address where the module is loaded into memory. Encoded as a decimal or hexadecimal (0x prefixed) string.
	BaseAddress string `json:"baseAddress"`
	// Size of the module in bytes.
	Size float32 `json:"size"`
}
