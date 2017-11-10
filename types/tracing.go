package types

type Tracing_MemoryDumpConfig struct {
}
type Tracing_TraceConfig struct {
	RecordMode           *string                   `json:"recordMode,omitempty"`
	EnableSampling       *bool                     `json:"enableSampling,omitempty"`
	EnableSystrace       *bool                     `json:"enableSystrace,omitempty"`
	EnableArgumentFilter *bool                     `json:"enableArgumentFilter,omitempty"`
	IncludedCategories   []string                  `json:"includedCategories,omitempty"`
	ExcludedCategories   []string                  `json:"excludedCategories,omitempty"`
	SyntheticDelays      []string                  `json:"syntheticDelays,omitempty"`
	MemoryDumpConfig     *Tracing_MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}
