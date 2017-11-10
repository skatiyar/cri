package headlessexperimental

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type HeadlessExperimental struct {
	conn cri.Connector
}

func New(conn cri.Connector) *HeadlessExperimental {
	return &HeadlessExperimental{conn}
}

// Enables headless events for the target.
func (obj *HeadlessExperimental) Enable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.enable", nil, nil)
	return
}

// Disables headless events for the target.
func (obj *HeadlessExperimental) Disable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.disable", nil, nil)
	return
}

type BeginFrameRequest struct {
	// Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used.
	FrameTime *types.Runtime_Timestamp `json:"frameTime,omitempty"`
	// Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval.
	Deadline *types.Runtime_Timestamp `json:"deadline,omitempty"`
	// The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval *float32 `json:"interval,omitempty"`
	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured.
	Screenshot *types.HeadlessExperimental_ScreenshotParams `json:"screenshot,omitempty"`
}
type BeginFrameResponse struct {
	// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display.
	HasDamage bool `json:"hasDamage"`
	// Whether the main frame submitted a new display frame in response to this BeginFrame.
	MainFrameContentUpdated bool `json:"mainFrameContentUpdated"`
	// Base64-encoded image data of the screenshot, if one was requested and successfully taken.
	ScreenshotData *string `json:"screenshotData,omitempty"`
}

// Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl.
func (obj *HeadlessExperimental) BeginFrame(request *BeginFrameRequest) (response BeginFrameResponse, err error) {
	err = obj.conn.Send("HeadlessExperimental.beginFrame", request, &response)
	return
}
