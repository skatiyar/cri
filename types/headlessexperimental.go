/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types


//Encoding options for a screenshot.
type HeadlessExperimental_ScreenshotParams struct {
	// Image compression format (defaults to png).
	Format *string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality *int `json:"quality,omitempty"`
}

