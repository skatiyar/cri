/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package emulation provides commands and events for Emulation domain.
package emulation

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Emulation domain
const (
	SetDeviceMetricsOverride          = "Emulation.setDeviceMetricsOverride"
	ClearDeviceMetricsOverride        = "Emulation.clearDeviceMetricsOverride"
	ResetPageScaleFactor              = "Emulation.resetPageScaleFactor"
	SetPageScaleFactor                = "Emulation.setPageScaleFactor"
	SetVisibleSize                    = "Emulation.setVisibleSize"
	SetScriptExecutionDisabled        = "Emulation.setScriptExecutionDisabled"
	SetGeolocationOverride            = "Emulation.setGeolocationOverride"
	ClearGeolocationOverride          = "Emulation.clearGeolocationOverride"
	SetTouchEmulationEnabled          = "Emulation.setTouchEmulationEnabled"
	SetEmitTouchEventsForMouse        = "Emulation.setEmitTouchEventsForMouse"
	SetEmulatedMedia                  = "Emulation.setEmulatedMedia"
	SetCPUThrottlingRate              = "Emulation.setCPUThrottlingRate"
	CanEmulate                        = "Emulation.canEmulate"
	SetVirtualTimePolicy              = "Emulation.setVirtualTimePolicy"
	SetNavigatorOverrides             = "Emulation.setNavigatorOverrides"
	SetDefaultBackgroundColorOverride = "Emulation.setDefaultBackgroundColorOverride"
)

// List of events in Emulation domain
const (
	VirtualTimeBudgetExpired = "Emulation.virtualTimeBudgetExpired"
	VirtualTimeAdvanced      = "Emulation.virtualTimeAdvanced"
	VirtualTimePaused        = "Emulation.virtualTimePaused"
)

// This domain emulates different environments for the page.
type Emulation struct {
	conn cri.Connector
}

// New creates a Emulation instance
func New(conn cri.Connector) *Emulation {
	return &Emulation{conn}
}

type SetDeviceMetricsOverrideRequest struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float32 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`
	// Scale to apply to resulting view image.
	// NOTE Experimental
	Scale *float32 `json:"scale,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	// NOTE Experimental
	ScreenWidth *int `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	// NOTE Experimental
	ScreenHeight *int `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	// NOTE Experimental
	PositionX *int `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	// NOTE Experimental
	PositionY *int `json:"positionY,omitempty"`
	// Do not set visible view size, rely upon explicit setVisibleSize call.
	// NOTE Experimental
	DontSetVisibleSize *bool `json:"dontSetVisibleSize,omitempty"`
	// Screen orientation override.
	ScreenOrientation *types.Emulation_ScreenOrientation `json:"screenOrientation,omitempty"`
	// If set, the visible area of the page will be overridden to this viewport. This viewport change is not observed by the page, e.g. viewport-relative elements do not change positions.
	// NOTE Experimental
	Viewport *types.Page_Viewport `json:"viewport,omitempty"`
}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (obj *Emulation) SetDeviceMetricsOverride(request *SetDeviceMetricsOverrideRequest) (err error) {
	err = obj.conn.Send(SetDeviceMetricsOverride, request, nil)
	return
}

// Clears the overriden device metrics.
func (obj *Emulation) ClearDeviceMetricsOverride() (err error) {
	err = obj.conn.Send(ClearDeviceMetricsOverride, nil, nil)
	return
}

// Requests that page scale factor is reset to initial values.
func (obj *Emulation) ResetPageScaleFactor() (err error) {
	err = obj.conn.Send(ResetPageScaleFactor, nil, nil)
	return
}

type SetPageScaleFactorRequest struct {
	// Page scale factor.
	PageScaleFactor float32 `json:"pageScaleFactor"`
}

// Sets a specified page scale factor.
func (obj *Emulation) SetPageScaleFactor(request *SetPageScaleFactorRequest) (err error) {
	err = obj.conn.Send(SetPageScaleFactor, request, nil)
	return
}

type SetVisibleSizeRequest struct {
	// Frame width (DIP).
	Width int `json:"width"`
	// Frame height (DIP).
	Height int `json:"height"`
}

// Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
func (obj *Emulation) SetVisibleSize(request *SetVisibleSizeRequest) (err error) {
	err = obj.conn.Send(SetVisibleSize, request, nil)
	return
}

type SetScriptExecutionDisabledRequest struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// Switches script execution in the page.
func (obj *Emulation) SetScriptExecutionDisabled(request *SetScriptExecutionDisabledRequest) (err error) {
	err = obj.conn.Send(SetScriptExecutionDisabled, request, nil)
	return
}

type SetGeolocationOverrideRequest struct {
	// Mock latitude
	Latitude *float32 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude *float32 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy *float32 `json:"accuracy,omitempty"`
}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (obj *Emulation) SetGeolocationOverride(request *SetGeolocationOverrideRequest) (err error) {
	err = obj.conn.Send(SetGeolocationOverride, request, nil)
	return
}

// Clears the overriden Geolocation Position and Error.
func (obj *Emulation) ClearGeolocationOverride() (err error) {
	err = obj.conn.Send(ClearGeolocationOverride, nil, nil)
	return
}

type SetTouchEmulationEnabledRequest struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Maximum touch points supported. Defaults to one.
	MaxTouchPoints *int `json:"maxTouchPoints,omitempty"`
}

// Enables touch on platforms which do not support them.
func (obj *Emulation) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send(SetTouchEmulationEnabled, request, nil)
	return
}

type SetEmitTouchEventsForMouseRequest struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration *string `json:"configuration,omitempty"`
}

func (obj *Emulation) SetEmitTouchEventsForMouse(request *SetEmitTouchEventsForMouseRequest) (err error) {
	err = obj.conn.Send(SetEmitTouchEventsForMouse, request, nil)
	return
}

type SetEmulatedMediaRequest struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

// Emulates the given media for CSS media queries.
func (obj *Emulation) SetEmulatedMedia(request *SetEmulatedMediaRequest) (err error) {
	err = obj.conn.Send(SetEmulatedMedia, request, nil)
	return
}

type SetCPUThrottlingRateRequest struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float32 `json:"rate"`
}

// Enables CPU throttling to emulate slow CPUs.
func (obj *Emulation) SetCPUThrottlingRate(request *SetCPUThrottlingRateRequest) (err error) {
	err = obj.conn.Send(SetCPUThrottlingRate, request, nil)
	return
}

type CanEmulateResponse struct {
	// True if emulation is supported.
	Result bool `json:"result"`
}

// Tells whether emulation is supported.
func (obj *Emulation) CanEmulate() (response CanEmulateResponse, err error) {
	err = obj.conn.Send(CanEmulate, nil, &response)
	return
}

type SetVirtualTimePolicyRequest struct {
	Policy types.Emulation_VirtualTimePolicy `json:"policy"`
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	Budget *float32 `json:"budget,omitempty"`
	// If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount *int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
}

type SetVirtualTimePolicyResponse struct {
	// Absolute timestamp at which virtual time was first enabled (milliseconds since epoch).
	VirtualTimeBase types.Runtime_Timestamp `json:"virtualTimeBase"`
}

// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
func (obj *Emulation) SetVirtualTimePolicy(request *SetVirtualTimePolicyRequest) (response SetVirtualTimePolicyResponse, err error) {
	err = obj.conn.Send(SetVirtualTimePolicy, request, &response)
	return
}

type SetNavigatorOverridesRequest struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

// Overrides value returned by the javascript navigator object.
func (obj *Emulation) SetNavigatorOverrides(request *SetNavigatorOverridesRequest) (err error) {
	err = obj.conn.Send(SetNavigatorOverrides, request, nil)
	return
}

type SetDefaultBackgroundColorOverrideRequest struct {
	// RGBA of the default background color. If not specified, any existing override will be cleared.
	Color *types.DOM_RGBA `json:"color,omitempty"`
}

// Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
func (obj *Emulation) SetDefaultBackgroundColorOverride(request *SetDefaultBackgroundColorOverrideRequest) (err error) {
	err = obj.conn.Send(SetDefaultBackgroundColorOverride, request, nil)
	return
}

// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
// NOTE Experimental
func (obj *Emulation) VirtualTimeBudgetExpired(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(VirtualTimeBudgetExpired, closeChn)
	go func() {
		for {

			readErr := decoder(nil)
			if !fn(readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type VirtualTimeAdvancedParams struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed float32 `json:"virtualTimeElapsed"`
}

// Notification sent after the virtual time has advanced.
// NOTE Experimental
func (obj *Emulation) VirtualTimeAdvanced(fn func(params *VirtualTimeAdvancedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(VirtualTimeAdvanced, closeChn)
	go func() {
		for {
			params := VirtualTimeAdvancedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type VirtualTimePausedParams struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed float32 `json:"virtualTimeElapsed"`
}

// Notification sent after the virtual time has paused.
// NOTE Experimental
func (obj *Emulation) VirtualTimePaused(fn func(params *VirtualTimePausedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(VirtualTimePaused, closeChn)
	go func() {
		for {
			params := VirtualTimePausedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
