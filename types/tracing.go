/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Configuration for memory dump. Used only when "memory-infra" category is enabled.
type Tracing_MemoryDumpConfig struct {
}

type Tracing_TraceConfig struct {
	// Controls how the trace buffer stores data.
	RecordMode *string `json:"recordMode,omitempty"`
	// Turns on JavaScript stack sampling.
	EnableSampling *bool `json:"enableSampling,omitempty"`
	// Turns on system tracing.
	EnableSystrace *bool `json:"enableSystrace,omitempty"`
	// Turns on argument filter.
	EnableArgumentFilter *bool `json:"enableArgumentFilter,omitempty"`
	// Included category filters.
	IncludedCategories []string `json:"includedCategories,omitempty"`
	// Excluded category filters.
	ExcludedCategories []string `json:"excludedCategories,omitempty"`
	// Configuration to synthesize the delays in tracing.
	SyntheticDelays []string `json:"syntheticDelays,omitempty"`
	// Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
	MemoryDumpConfig *Tracing_MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}
