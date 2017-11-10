package headlessexperimental

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type HeadlessExperimental struct {
	conn cri.Connector
}

func New(conn cri.Connector) *HeadlessExperimental {
	return &HeadlessExperimental{conn}
}
func (obj *HeadlessExperimental) Enable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.enable", nil, nil)
	return
}
func (obj *HeadlessExperimental) Disable() (err error) {
	err = obj.conn.Send("HeadlessExperimental.disable", nil, nil)
	return
}

type BeginFrameRequest struct {
	FrameTime	*types.Runtime_Timestamp			`json:"frameTime,omitempty"`// Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used.
	Deadline	*types.Runtime_Timestamp			`json:"deadline,omitempty"`// Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval.
	Interval	*float32					`json:"interval,omitempty"`// The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Screenshot	*types.HeadlessExperimental_ScreenshotParams	`json:"screenshot,omitempty"`// If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured.
}

func (obj *HeadlessExperimental) BeginFrame(request *BeginFrameRequest) (response struct {
	HasDamage		bool	`json:"hasDamage"`// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display.
	MainFrameContentUpdated	bool	`json:"mainFrameContentUpdated"`// Whether the main frame submitted a new display frame in response to this BeginFrame.
	ScreenshotData		*string	`json:"screenshotData,omitempty"`// Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}, err error) {
	err = obj.conn.Send("HeadlessExperimental.beginFrame", request, &response)
	return
}
