/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types



type Input_TouchPoint struct {
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float32 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float32 `json:"y"`
	// X radius of the touch area (default: 1.0).
	RadiusX *float32 `json:"radiusX,omitempty"`
	// Y radius of the touch area (default: 1.0).
	RadiusY *float32 `json:"radiusY,omitempty"`
	// Rotation angle (default: 0.0).
	RotationAngle *float32 `json:"rotationAngle,omitempty"`
	// Force (default: 1.0).
	Force *float32 `json:"force,omitempty"`
	// Identifier used to track touch sources between events, must be unique within an event.
	Id *float32 `json:"id,omitempty"`
}


type Input_GestureSourceType string

//UTC time in seconds, counted from January 1, 1970.
type Input_TimeSinceEpoch float32

