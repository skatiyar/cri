/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package headlessexperimental provides commands and events for HeadlessExperimental domain.
package headlessexperimental

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in HeadlessExperimental domain
const (
	BeginFrame = "HeadlessExperimental.beginFrame"
	Disable    = "HeadlessExperimental.disable"
	Enable     = "HeadlessExperimental.enable"
)

// List of events in HeadlessExperimental domain
const (
	MainFrameReadyForScreenshots = "HeadlessExperimental.mainFrameReadyForScreenshots"
	NeedsBeginFramesChanged      = "HeadlessExperimental.needsBeginFramesChanged"
)

// This domain provides experimental commands only supported in headless mode.
type HeadlessExperimental struct {
	conn cri.Connector
}

// New creates a HeadlessExperimental instance
func New(conn cri.Connector) *HeadlessExperimental {
	return &HeadlessExperimental{conn}
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
	err = obj.conn.Send(BeginFrame, request, &response)
	return
}

// Disables headless events for the target.
func (obj *HeadlessExperimental) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables headless events for the target.
func (obj *HeadlessExperimental) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Issued when the main frame has first submitted a frame to the browser. May only be fired while a BeginFrame is in flight. Before this event, screenshotting requests may fail.
func (obj *HeadlessExperimental) MainFrameReadyForScreenshots() (err error) {
	err = obj.conn.On(MainFrameReadyForScreenshots, nil)
	return
}

type NeedsBeginFramesChangedParams struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool `json:"needsBeginFrames"`
}

// Issued when the target starts or stops needing BeginFrames.
func (obj *HeadlessExperimental) NeedsBeginFramesChanged() (params NeedsBeginFramesChangedParams, err error) {
	err = obj.conn.On(NeedsBeginFramesChanged, &params)
	return
}
