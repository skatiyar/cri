package types

type Tracing_MemoryDumpConfig struct {
}
type Tracing_TraceConfig struct {
	RecordMode		*string				`json:"recordMode,omitempty"`// Controls how the trace buffer stores data.
	EnableSampling		*bool				`json:"enableSampling,omitempty"`// Turns on JavaScript stack sampling.
	EnableSystrace		*bool				`json:"enableSystrace,omitempty"`// Turns on system tracing.
	EnableArgumentFilter	*bool				`json:"enableArgumentFilter,omitempty"`// Turns on argument filter.
	IncludedCategories	[]string			`json:"includedCategories,omitempty"`// Included category filters.
	ExcludedCategories	[]string			`json:"excludedCategories,omitempty"`// Excluded category filters.
	SyntheticDelays		[]string			`json:"syntheticDelays,omitempty"`// Configuration to synthesize the delays in tracing.
	MemoryDumpConfig	*Tracing_MemoryDumpConfig	`json:"memoryDumpConfig,omitempty"`// Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}
