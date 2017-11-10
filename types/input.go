package types

type Input_TouchPoint struct {
	X             float32  `json:"x"`
	Y             float32  `json:"y"`
	RadiusX       *float32 `json:"radiusX,omitempty"`
	RadiusY       *float32 `json:"radiusY,omitempty"`
	RotationAngle *float32 `json:"rotationAngle,omitempty"`
	Force         *float32 `json:"force,omitempty"`
	Id            *float32 `json:"id,omitempty"`
}
type Input_GestureSourceType string
type Input_TimeSinceEpoch float32
