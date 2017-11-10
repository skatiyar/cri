package types

type HeadlessExperimental_ScreenshotParams struct {
	Format	*string	`json:"format,omitempty"`// Image compression format (defaults to png).
	Quality	*int	`json:"quality,omitempty"`// Compression quality from range [0..100] (jpeg only).
}
