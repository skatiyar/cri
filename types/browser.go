/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

type Browser_WindowID int

// The state of the browser window.
type Browser_WindowState string

// Browser window bounds information
type Browser_Bounds struct {
	// The offset from the left edge of the screen to the window in pixels.
	Left *int `json:"left,omitempty"`
	// The offset from the top edge of the screen to the window in pixels.
	Top *int `json:"top,omitempty"`
	// The window width in pixels.
	Width *int `json:"width,omitempty"`
	// The window height in pixels.
	Height *int `json:"height,omitempty"`
	// The window state. Default to normal.
	WindowState *Browser_WindowState `json:"windowState,omitempty"`
}

// Chrome histogram bucket.
type Browser_Bucket struct {
	// Minimum value (inclusive).
	Low int `json:"low"`
	// Maximum value (exclusive).
	High int `json:"high"`
	// Number of samples.
	Count int `json:"count"`
}

// Chrome histogram.
type Browser_Histogram struct {
	// Name.
	Name string `json:"name"`
	// Sum of sample values.
	Sum int `json:"sum"`
	// Total number of samples.
	Count int `json:"count"`
	// Buckets.
	Buckets []Browser_Bucket `json:"buckets"`
}
