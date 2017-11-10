package emulation

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Emulation struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Emulation {
	return &Emulation{conn}
}

type SetDeviceMetricsOverrideRequest struct {
	Width			int					`json:"width"`// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height			int					`json:"height"`// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor	float32					`json:"deviceScaleFactor"`// Overriding device scale factor value. 0 disables the override.
	Mobile			bool					`json:"mobile"`// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Scale			*float32				`json:"scale,omitempty"`// Scale to apply to resulting view image.
	ScreenWidth		*int					`json:"screenWidth,omitempty"`// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenHeight		*int					`json:"screenHeight,omitempty"`// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	PositionX		*int					`json:"positionX,omitempty"`// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionY		*int					`json:"positionY,omitempty"`// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	DontSetVisibleSize	*bool					`json:"dontSetVisibleSize,omitempty"`// Do not set visible view size, rely upon explicit setVisibleSize call.
	ScreenOrientation	*types.Emulation_ScreenOrientation	`json:"screenOrientation,omitempty"`// Screen orientation override.
	Viewport		*types.Page_Viewport			`json:"viewport,omitempty"`// If set, the visible area of the page will be overridden to this viewport. This viewport change is not observed by the page, e.g. viewport-relative elements do not change positions.
}

func (obj *Emulation) SetDeviceMetricsOverride(request *SetDeviceMetricsOverrideRequest) (err error) {
	err = obj.conn.Send("Emulation.setDeviceMetricsOverride", request, nil)
	return
}
func (obj *Emulation) ClearDeviceMetricsOverride() (err error) {
	err = obj.conn.Send("Emulation.clearDeviceMetricsOverride", nil, nil)
	return
}
func (obj *Emulation) ResetPageScaleFactor() (err error) {
	err = obj.conn.Send("Emulation.resetPageScaleFactor", nil, nil)
	return
}

type SetPageScaleFactorRequest struct {
	PageScaleFactor float32 `json:"pageScaleFactor"`// Page scale factor.
}

func (obj *Emulation) SetPageScaleFactor(request *SetPageScaleFactorRequest) (err error) {
	err = obj.conn.Send("Emulation.setPageScaleFactor", request, nil)
	return
}

type SetVisibleSizeRequest struct {
	Width	int	`json:"width"`// Frame width (DIP).
	Height	int	`json:"height"`// Frame height (DIP).
}

func (obj *Emulation) SetVisibleSize(request *SetVisibleSizeRequest) (err error) {
	err = obj.conn.Send("Emulation.setVisibleSize", request, nil)
	return
}

type SetScriptExecutionDisabledRequest struct {
	Value bool `json:"value"`// Whether script execution should be disabled in the page.
}

func (obj *Emulation) SetScriptExecutionDisabled(request *SetScriptExecutionDisabledRequest) (err error) {
	err = obj.conn.Send("Emulation.setScriptExecutionDisabled", request, nil)
	return
}

type SetGeolocationOverrideRequest struct {
	Latitude	*float32	`json:"latitude,omitempty"`// Mock latitude
	Longitude	*float32	`json:"longitude,omitempty"`// Mock longitude
	Accuracy	*float32	`json:"accuracy,omitempty"`// Mock accuracy
}

func (obj *Emulation) SetGeolocationOverride(request *SetGeolocationOverrideRequest) (err error) {
	err = obj.conn.Send("Emulation.setGeolocationOverride", request, nil)
	return
}
func (obj *Emulation) ClearGeolocationOverride() (err error) {
	err = obj.conn.Send("Emulation.clearGeolocationOverride", nil, nil)
	return
}

type SetTouchEmulationEnabledRequest struct {
	Enabled		bool	`json:"enabled"`// Whether the touch event emulation should be enabled.
	MaxTouchPoints	*int	`json:"maxTouchPoints,omitempty"`// Maximum touch points supported. Defaults to one.
}

func (obj *Emulation) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send("Emulation.setTouchEmulationEnabled", request, nil)
	return
}

type SetEmitTouchEventsForMouseRequest struct {
	Enabled		bool	`json:"enabled"`// Whether touch emulation based on mouse input should be enabled.
	Configuration	*string	`json:"configuration,omitempty"`// Touch/gesture events configuration. Default: current platform.
}

func (obj *Emulation) SetEmitTouchEventsForMouse(request *SetEmitTouchEventsForMouseRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmitTouchEventsForMouse", request, nil)
	return
}

type SetEmulatedMediaRequest struct {
	Media string `json:"media"`// Media type to emulate. Empty string disables the override.
}

func (obj *Emulation) SetEmulatedMedia(request *SetEmulatedMediaRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmulatedMedia", request, nil)
	return
}

type SetCPUThrottlingRateRequest struct {
	Rate float32 `json:"rate"`// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
}

func (obj *Emulation) SetCPUThrottlingRate(request *SetCPUThrottlingRateRequest) (err error) {
	err = obj.conn.Send("Emulation.setCPUThrottlingRate", request, nil)
	return
}
func (obj *Emulation) CanEmulate() (response struct {
	Result bool `json:"result"`// True if emulation is supported.
}, err error) {
	err = obj.conn.Send("Emulation.canEmulate", nil, &response)
	return
}

type SetVirtualTimePolicyRequest struct {
	Policy					types.Emulation_VirtualTimePolicy	`json:"policy"`
	Budget					*int					`json:"budget,omitempty"`// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	MaxVirtualTimeTaskStarvationCount	*int					`json:"maxVirtualTimeTaskStarvationCount,omitempty"`// If set this specifies the maximum number of tasks that can be run before virtual is forced forwards to prevent deadlock.
}

func (obj *Emulation) SetVirtualTimePolicy(request *SetVirtualTimePolicyRequest) (err error) {
	err = obj.conn.Send("Emulation.setVirtualTimePolicy", request, nil)
	return
}

type SetNavigatorOverridesRequest struct {
	Platform string `json:"platform"`// The platform navigator.platform should return.
}

func (obj *Emulation) SetNavigatorOverrides(request *SetNavigatorOverridesRequest) (err error) {
	err = obj.conn.Send("Emulation.setNavigatorOverrides", request, nil)
	return
}

type SetDefaultBackgroundColorOverrideRequest struct {
	Color *types.DOM_RGBA `json:"color,omitempty"`// RGBA of the default background color. If not specified, any existing override will be cleared.
}

func (obj *Emulation) SetDefaultBackgroundColorOverride(request *SetDefaultBackgroundColorOverrideRequest) (err error) {
	err = obj.conn.Send("Emulation.setDefaultBackgroundColorOverride", request, nil)
	return
}
