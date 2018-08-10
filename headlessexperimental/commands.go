/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package headlessexperimental provides commands and events for HeadlessExperimental domain.
package headlessexperimental

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in HeadlessExperimental domain
const (
	BeginFrame = "HeadlessExperimental.beginFrame"
	Disable    = "HeadlessExperimental.disable"
	Enable     = "HeadlessExperimental.enable"
)

// List of events in HeadlessExperimental domain
const (
	NeedsBeginFramesChanged = "HeadlessExperimental.needsBeginFramesChanged"
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
	// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used.
	FrameTimeTicks *float32 `json:"frameTimeTicks,omitempty"`
	// The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval *float32 `json:"interval,omitempty"`
	// Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	NoDisplayUpdates *bool `json:"noDisplayUpdates,omitempty"`
	// If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
	Screenshot *types.HeadlessExperimental_ScreenshotParams `json:"screenshot,omitempty"`
}

type BeginFrameResponse struct {
	// Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
	HasDamage bool `json:"hasDamage"`
	// Base64-encoded image data of the screenshot, if one was requested and successfully taken.
	ScreenshotData *string `json:"screenshotData,omitempty"`
}

// Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a screenshot from the resulting frame. Requires that the target was created with enabled BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also https://goo.gl/3zHXhB for more background.
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

type NeedsBeginFramesChangedParams struct {
	// True if BeginFrames are needed, false otherwise.
	NeedsBeginFrames bool `json:"needsBeginFrames"`
}

// Issued when the target starts or stops needing BeginFrames.
func (obj *HeadlessExperimental) NeedsBeginFramesChanged(fn func(event string, params NeedsBeginFramesChangedParams, err error) bool) {
	listen, closer := obj.conn.On(NeedsBeginFramesChanged)
	go func() {
		defer closer()
		for {
			var params NeedsBeginFramesChangedParams
			if !fn(NeedsBeginFramesChanged, params, listen(&params)) {
				return
			}
		}
	}()
}
