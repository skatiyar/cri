package input

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Input struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Input {
	return &Input{conn}
}

type SetIgnoreInputEventsRequest struct {
	Ignore bool `json:"ignore"`
}

func (obj *Input) SetIgnoreInputEvents(request *SetIgnoreInputEventsRequest) (err error) {
	err = obj.conn.Send("Input.setIgnoreInputEvents", request, nil)
	return
}

type DispatchKeyEventRequest struct {
	Type                  string                      `json:"type"`
	Modifiers             *int                        `json:"modifiers,omitempty"`
	Timestamp             *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
	Text                  *string                     `json:"text,omitempty"`
	UnmodifiedText        *string                     `json:"unmodifiedText,omitempty"`
	KeyIdentifier         *string                     `json:"keyIdentifier,omitempty"`
	Code                  *string                     `json:"code,omitempty"`
	Key                   *string                     `json:"key,omitempty"`
	WindowsVirtualKeyCode *int                        `json:"windowsVirtualKeyCode,omitempty"`
	NativeVirtualKeyCode  *int                        `json:"nativeVirtualKeyCode,omitempty"`
	AutoRepeat            *bool                       `json:"autoRepeat,omitempty"`
	IsKeypad              *bool                       `json:"isKeypad,omitempty"`
	IsSystemKey           *bool                       `json:"isSystemKey,omitempty"`
	Location              *int                        `json:"location,omitempty"`
}

func (obj *Input) DispatchKeyEvent(request *DispatchKeyEventRequest) (err error) {
	err = obj.conn.Send("Input.dispatchKeyEvent", request, nil)
	return
}

type DispatchMouseEventRequest struct {
	Type       string                      `json:"type"`
	X          float32                     `json:"x"`
	Y          float32                     `json:"y"`
	Modifiers  *int                        `json:"modifiers,omitempty"`
	Timestamp  *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
	Button     *string                     `json:"button,omitempty"`
	ClickCount *int                        `json:"clickCount,omitempty"`
	DeltaX     *float32                    `json:"deltaX,omitempty"`
	DeltaY     *float32                    `json:"deltaY,omitempty"`
}

func (obj *Input) DispatchMouseEvent(request *DispatchMouseEventRequest) (err error) {
	err = obj.conn.Send("Input.dispatchMouseEvent", request, nil)
	return
}

type DispatchTouchEventRequest struct {
	Type        string                      `json:"type"`
	TouchPoints []types.Input_TouchPoint    `json:"touchPoints"`
	Modifiers   *int                        `json:"modifiers,omitempty"`
	Timestamp   *types.Input_TimeSinceEpoch `json:"timestamp,omitempty"`
}

func (obj *Input) DispatchTouchEvent(request *DispatchTouchEventRequest) (err error) {
	err = obj.conn.Send("Input.dispatchTouchEvent", request, nil)
	return
}

type EmulateTouchFromMouseEventRequest struct {
	Type       string                     `json:"type"`
	X          int                        `json:"x"`
	Y          int                        `json:"y"`
	Timestamp  types.Input_TimeSinceEpoch `json:"timestamp"`
	Button     string                     `json:"button"`
	DeltaX     *float32                   `json:"deltaX,omitempty"`
	DeltaY     *float32                   `json:"deltaY,omitempty"`
	Modifiers  *int                       `json:"modifiers,omitempty"`
	ClickCount *int                       `json:"clickCount,omitempty"`
}

func (obj *Input) EmulateTouchFromMouseEvent(request *EmulateTouchFromMouseEventRequest) (err error) {
	err = obj.conn.Send("Input.emulateTouchFromMouseEvent", request, nil)
	return
}

type SynthesizePinchGestureRequest struct {
	X                 float32                        `json:"x"`
	Y                 float32                        `json:"y"`
	ScaleFactor       float32                        `json:"scaleFactor"`
	RelativeSpeed     *int                           `json:"relativeSpeed,omitempty"`
	GestureSourceType *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
}

func (obj *Input) SynthesizePinchGesture(request *SynthesizePinchGestureRequest) (err error) {
	err = obj.conn.Send("Input.synthesizePinchGesture", request, nil)
	return
}

type SynthesizeScrollGestureRequest struct {
	X                     float32                        `json:"x"`
	Y                     float32                        `json:"y"`
	XDistance             *float32                       `json:"xDistance,omitempty"`
	YDistance             *float32                       `json:"yDistance,omitempty"`
	XOverscroll           *float32                       `json:"xOverscroll,omitempty"`
	YOverscroll           *float32                       `json:"yOverscroll,omitempty"`
	PreventFling          *bool                          `json:"preventFling,omitempty"`
	Speed                 *int                           `json:"speed,omitempty"`
	GestureSourceType     *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
	RepeatCount           *int                           `json:"repeatCount,omitempty"`
	RepeatDelayMs         *int                           `json:"repeatDelayMs,omitempty"`
	InteractionMarkerName *string                        `json:"interactionMarkerName,omitempty"`
}

func (obj *Input) SynthesizeScrollGesture(request *SynthesizeScrollGestureRequest) (err error) {
	err = obj.conn.Send("Input.synthesizeScrollGesture", request, nil)
	return
}

type SynthesizeTapGestureRequest struct {
	X                 float32                        `json:"x"`
	Y                 float32                        `json:"y"`
	Duration          *int                           `json:"duration,omitempty"`
	TapCount          *int                           `json:"tapCount,omitempty"`
	GestureSourceType *types.Input_GestureSourceType `json:"gestureSourceType,omitempty"`
}

func (obj *Input) SynthesizeTapGesture(request *SynthesizeTapGestureRequest) (err error) {
	err = obj.conn.Send("Input.synthesizeTapGesture", request, nil)
	return
}
