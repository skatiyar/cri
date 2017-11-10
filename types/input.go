package types

type Input_TouchPoint struct {
	X		float32		`json:"x"`// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y		float32		`json:"y"`// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX		*float32	`json:"radiusX,omitempty"`// X radius of the touch area (default: 1.0).
	RadiusY		*float32	`json:"radiusY,omitempty"`// Y radius of the touch area (default: 1.0).
	RotationAngle	*float32	`json:"rotationAngle,omitempty"`// Rotation angle (default: 0.0).
	Force		*float32	`json:"force,omitempty"`// Force (default: 1.0).
	Id		*float32	`json:"id,omitempty"`// Identifier used to track touch sources between events, must be unique within an event.
}
type Input_GestureSourceType string
type Input_TimeSinceEpoch float32
