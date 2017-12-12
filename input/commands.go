/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package input provides commands and events for Input domain.
package input

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Input domain
const (
	DispatchKeyEvent           = "Input.dispatchKeyEvent"
	DispatchMouseEvent         = "Input.dispatchMouseEvent"
	DispatchTouchEvent         = "Input.dispatchTouchEvent"
	EmulateTouchFromMouseEvent = "Input.emulateTouchFromMouseEvent"
	SetIgnoreInputEvents       = "Input.setIgnoreInputEvents"
	SynthesizePinchGesture     = "Input.synthesizePinchGesture"
	SynthesizeScrollGesture    = "Input.synthesizeScrollGesture"
	SynthesizeTapGesture       = "Input.synthesizeTapGesture"
)

type Input struct {
	conn cri.Connector
}

// New creates a Input instance
func New(conn cri.Connector) *Input {
	return &Input{conn}
}

type DispatchKeyEventRequest struct {
	// Type of the key event.
	Type string `json:"type"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers *int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
	// Text as generated by processing a virtual key code with a keyboard layout. Not needed for for `keyUp` and `rawKeyDown` events (default: "")
	Text *string `json:"text,omitempty"`
	// Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
	UnmodifiedText *string `json:"unmodifiedText,omitempty"`
	// Unique key identifier (e.g., 'U+0041') (default: "").
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`
	// Unique DOM defined string value for each physical key (e.g., 'KeyA') (default: "").
	Code *string `json:"code,omitempty"`
	// Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc (e.g., 'AltGr') (default: "").
	Key *string `json:"key,omitempty"`
	// Windows virtual key code (default: 0).
	WindowsVirtualKeyCode *int `json:"windowsVirtualKeyCode,omitempty"`
	// Native virtual key code (default: 0).
	NativeVirtualKeyCode *int `json:"nativeVirtualKeyCode,omitempty"`
	// Whether the event was generated from auto repeat (default: false).
	AutoRepeat *bool `json:"autoRepeat,omitempty"`
	// Whether the event was generated from the keypad (default: false).
	IsKeypad *bool `json:"isKeypad,omitempty"`
	// Whether the event was a system key event (default: false).
	IsSystemKey *bool `json:"isSystemKey,omitempty"`
	// Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right (default: 0).
	Location *int `json:"location,omitempty"`
}

// Dispatches a key event to the page.
func (obj *Input) DispatchKeyEvent(request *DispatchKeyEventRequest) (err error) {
	err = obj.conn.Send(DispatchKeyEvent, request, nil)
	return
}

type DispatchMouseEventRequest struct {
	// Type of the mouse event.
	Type string `json:"type"`
	// X coordinate of the event relative to the main frame's viewport in CSS pixels.
	X float32 `json:"x"`
	// Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	Y float32 `json:"y"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers *int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
	// Mouse button (default: "none").
	Button *string `json:"button,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount *int `json:"clickCount,omitempty"`
	// X delta in CSS pixels for mouse wheel event (default: 0).
	DeltaX *float32 `json:"deltaX,omitempty"`
	// Y delta in CSS pixels for mouse wheel event (default: 0).
	DeltaY *float32 `json:"deltaY,omitempty"`
}

// Dispatches a mouse event to the page.
func (obj *Input) DispatchMouseEvent(request *DispatchMouseEventRequest) (err error) {
	err = obj.conn.Send(DispatchMouseEvent, request, nil)
	return
}

type DispatchTouchEventRequest struct {
	// Type of the touch event. TouchEnd and TouchCancel must not contain any touch points, while TouchStart and TouchMove must contains at least one.
	Type string `json:"type"`
	// Active touch points on the touch device. One event per any changed point (compared to previous touch event in a sequence) is generated, emulating pressing/moving/releasing points one by one.
	TouchPoints []types.Input_TouchPoint `json:"touchPoints"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers *int `json:"modifiers,omitempty"`
	// Time at which the event occurred.
	Timestamp *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
}

// Dispatches a touch event to the page.
func (obj *Input) DispatchTouchEvent(request *DispatchTouchEventRequest) (err error) {
	err = obj.conn.Send(DispatchTouchEvent, request, nil)
	return
}

type EmulateTouchFromMouseEventRequest struct {
	// Type of the mouse event.
	Type string `json:"type"`
	// X coordinate of the mouse pointer in DIP.
	X int `json:"x"`
	// Y coordinate of the mouse pointer in DIP.
	Y int `json:"y"`
	// Time at which the event occurred.
	Timestamp types.Input_TimeSinceEpoch `json:"timestamp"`
	// Mouse button.
	Button string `json:"button"`
	// X delta in DIP for mouse wheel event (default: 0).
	DeltaX *float32 `json:"deltaX,omitempty"`
	// Y delta in DIP for mouse wheel event (default: 0).
	DeltaY *float32 `json:"deltaY,omitempty"`
	// Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
	Modifiers *int `json:"modifiers,omitempty"`
	// Number of times the mouse button was clicked (default: 0).
	ClickCount *int `json:"clickCount,omitempty"`
}

// Emulates touch event from the mouse event parameters.
func (obj *Input) EmulateTouchFromMouseEvent(request *EmulateTouchFromMouseEventRequest) (err error) {
	err = obj.conn.Send(EmulateTouchFromMouseEvent, request, nil)
	return
}

type SetIgnoreInputEventsRequest struct {
	// Ignores input events processing when set to true.
	Ignore bool `json:"ignore"`
}

// Ignores input events (useful while auditing page).
func (obj *Input) SetIgnoreInputEvents(request *SetIgnoreInputEventsRequest) (err error) {
	err = obj.conn.Send(SetIgnoreInputEvents, request, nil)
	return
}

type SynthesizePinchGestureRequest struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float32 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float32 `json:"y"`
	// Relative scale factor after zooming (>1.0 zooms in, <1.0 zooms out).
	ScaleFactor float32 `json:"scaleFactor"`
	// Relative pointer speed in pixels per second (default: 800).
	RelativeSpeed *int `json:"relativeSpeed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	GestureSourceType *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
}

// Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
func (obj *Input) SynthesizePinchGesture(request *SynthesizePinchGestureRequest) (err error) {
	err = obj.conn.Send(SynthesizePinchGesture, request, nil)
	return
}

type SynthesizeScrollGestureRequest struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float32 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float32 `json:"y"`
	// The distance to scroll along the X axis (positive to scroll left).
	XDistance *float32 `json:"xDistance,omitempty"`
	// The distance to scroll along the Y axis (positive to scroll up).
	YDistance *float32 `json:"yDistance,omitempty"`
	// The number of additional pixels to scroll back along the X axis, in addition to the given distance.
	XOverscroll *float32 `json:"xOverscroll,omitempty"`
	// The number of additional pixels to scroll back along the Y axis, in addition to the given distance.
	YOverscroll *float32 `json:"yOverscroll,omitempty"`
	// Prevent fling (default: true).
	PreventFling *bool `json:"preventFling,omitempty"`
	// Swipe speed in pixels per second (default: 800).
	Speed *int `json:"speed,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	GestureSourceType *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
	// The number of times to repeat the gesture (default: 0).
	RepeatCount *int `json:"repeatCount,omitempty"`
	// The number of milliseconds delay between each repeat. (default: 250).
	RepeatDelayMs *int `json:"repeatDelayMs,omitempty"`
	// The name of the interaction markers to generate, if not empty (default: "").
	InteractionMarkerName *string `json:"interactionMarkerName,omitempty"`
}

// Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
func (obj *Input) SynthesizeScrollGesture(request *SynthesizeScrollGestureRequest) (err error) {
	err = obj.conn.Send(SynthesizeScrollGesture, request, nil)
	return
}

type SynthesizeTapGestureRequest struct {
	// X coordinate of the start of the gesture in CSS pixels.
	X float32 `json:"x"`
	// Y coordinate of the start of the gesture in CSS pixels.
	Y float32 `json:"y"`
	// Duration between touchdown and touchup events in ms (default: 50).
	Duration *int `json:"duration,omitempty"`
	// Number of times to perform the tap (e.g. 2 for double tap, default: 1).
	TapCount *int `json:"tapCount,omitempty"`
	// Which type of input events to be generated (default: 'default', which queries the platform for the preferred input type).
	GestureSourceType *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
}

// Synthesizes a tap gesture over a time period by issuing appropriate touch events.
func (obj *Input) SynthesizeTapGesture(request *SynthesizeTapGestureRequest) (err error) {
	err = obj.conn.Send(SynthesizeTapGesture, request, nil)
	return
}
