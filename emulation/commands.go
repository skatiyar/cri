/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// This domain emulates different environments for the page.
package emulation

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

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
	err = obj.conn.Send("Emulation.setDeviceMetricsOverride", request, nil)
	return
}

// Clears the overriden device metrics.
func (obj *Emulation) ClearDeviceMetricsOverride() (err error) {
	err = obj.conn.Send("Emulation.clearDeviceMetricsOverride", nil, nil)
	return
}

// Requests that page scale factor is reset to initial values.
func (obj *Emulation) ResetPageScaleFactor() (err error) {
	err = obj.conn.Send("Emulation.resetPageScaleFactor", nil, nil)
	return
}

type SetPageScaleFactorRequest struct {
	// Page scale factor.
	PageScaleFactor float32 `json:"pageScaleFactor"`
}

// Sets a specified page scale factor.
func (obj *Emulation) SetPageScaleFactor(request *SetPageScaleFactorRequest) (err error) {
	err = obj.conn.Send("Emulation.setPageScaleFactor", request, nil)
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
	err = obj.conn.Send("Emulation.setVisibleSize", request, nil)
	return
}

type SetScriptExecutionDisabledRequest struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// Switches script execution in the page.
func (obj *Emulation) SetScriptExecutionDisabled(request *SetScriptExecutionDisabledRequest) (err error) {
	err = obj.conn.Send("Emulation.setScriptExecutionDisabled", request, nil)
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
	err = obj.conn.Send("Emulation.setGeolocationOverride", request, nil)
	return
}

// Clears the overriden Geolocation Position and Error.
func (obj *Emulation) ClearGeolocationOverride() (err error) {
	err = obj.conn.Send("Emulation.clearGeolocationOverride", nil, nil)
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
	err = obj.conn.Send("Emulation.setTouchEmulationEnabled", request, nil)
	return
}

type SetEmitTouchEventsForMouseRequest struct {
	// Whether touch emulation based on mouse input should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration *string `json:"configuration,omitempty"`
}

func (obj *Emulation) SetEmitTouchEventsForMouse(request *SetEmitTouchEventsForMouseRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmitTouchEventsForMouse", request, nil)
	return
}

type SetEmulatedMediaRequest struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

// Emulates the given media for CSS media queries.
func (obj *Emulation) SetEmulatedMedia(request *SetEmulatedMediaRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmulatedMedia", request, nil)
	return
}

type SetCPUThrottlingRateRequest struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float32 `json:"rate"`
}

// Enables CPU throttling to emulate slow CPUs.
func (obj *Emulation) SetCPUThrottlingRate(request *SetCPUThrottlingRateRequest) (err error) {
	err = obj.conn.Send("Emulation.setCPUThrottlingRate", request, nil)
	return
}

type CanEmulateResponse struct {
	// True if emulation is supported.
	Result bool `json:"result"`
}

// Tells whether emulation is supported.
func (obj *Emulation) CanEmulate() (response CanEmulateResponse, err error) {
	err = obj.conn.Send("Emulation.canEmulate", nil, &response)
	return
}

type SetVirtualTimePolicyRequest struct {
	Policy types.Emulation_VirtualTimePolicy `json:"policy"`
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	Budget *int `json:"budget,omitempty"`
	// If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
	MaxVirtualTimeTaskStarvationCount *int `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
}

// Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
func (obj *Emulation) SetVirtualTimePolicy(request *SetVirtualTimePolicyRequest) (err error) {
	err = obj.conn.Send("Emulation.setVirtualTimePolicy", request, nil)
	return
}

type SetNavigatorOverridesRequest struct {
	// The platform navigator.platform should return.
	Platform string `json:"platform"`
}

// Overrides value returned by the javascript navigator object.
func (obj *Emulation) SetNavigatorOverrides(request *SetNavigatorOverridesRequest) (err error) {
	err = obj.conn.Send("Emulation.setNavigatorOverrides", request, nil)
	return
}

type SetDefaultBackgroundColorOverrideRequest struct {
	// RGBA of the default background color. If not specified, any existing override will be cleared.
	Color *types.DOM_RGBA `json:"color,omitempty"`
}

// Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
func (obj *Emulation) SetDefaultBackgroundColorOverride(request *SetDefaultBackgroundColorOverrideRequest) (err error) {
	err = obj.conn.Send("Emulation.setDefaultBackgroundColorOverride", request, nil)
	return
}

// Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
func (obj *Emulation) VirtualTimeBudgetExpired(fn func() bool) {

	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Emulation.virtualTimeBudgetExpired", closeChn, nil)
			if !fn() {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}

type VirtualTimeAdvancedParams struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

// Notification sent after the virtual time has advanced.
func (obj *Emulation) VirtualTimeAdvanced(fn func(params *VirtualTimeAdvancedParams) bool) {
	params := VirtualTimeAdvancedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Emulation.virtualTimeAdvanced", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}

type VirtualTimePausedParams struct {
	// The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
	VirtualTimeElapsed int `json:"virtualTimeElapsed"`
}

// Notification sent after the virtual time has paused.
func (obj *Emulation) VirtualTimePaused(fn func(params *VirtualTimePausedParams) bool) {
	params := VirtualTimePausedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Emulation.virtualTimePaused", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}
