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
	Width              int                                `json:"width"`
	Height             int                                `json:"height"`
	DeviceScaleFactor  float32                            `json:"deviceScaleFactor"`
	Mobile             bool                               `json:"mobile"`
	Scale              *float32                           `json:"scale,omitempty"`
	ScreenWidth        *int                               `json:"screenWidth,omitempty"`
	ScreenHeight       *int                               `json:"screenHeight,omitempty"`
	PositionX          *int                               `json:"positionX,omitempty"`
	PositionY          *int                               `json:"positionY,omitempty"`
	DontSetVisibleSize *bool                              `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation  *types.Emulation_ScreenOrientation `json:"screenOrientation,omitempty"`
	Viewport           *types.Page_Viewport               `json:"viewport,omitempty"`
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
	PageScaleFactor float32 `json:"pageScaleFactor"`
}

func (obj *Emulation) SetPageScaleFactor(request *SetPageScaleFactorRequest) (err error) {
	err = obj.conn.Send("Emulation.setPageScaleFactor", request, nil)
	return
}

type SetVisibleSizeRequest struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (obj *Emulation) SetVisibleSize(request *SetVisibleSizeRequest) (err error) {
	err = obj.conn.Send("Emulation.setVisibleSize", request, nil)
	return
}

type SetScriptExecutionDisabledRequest struct {
	Value bool `json:"value"`
}

func (obj *Emulation) SetScriptExecutionDisabled(request *SetScriptExecutionDisabledRequest) (err error) {
	err = obj.conn.Send("Emulation.setScriptExecutionDisabled", request, nil)
	return
}

type SetGeolocationOverrideRequest struct {
	Latitude  *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Accuracy  *float32 `json:"accuracy,omitempty"`
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
	Enabled        bool `json:"enabled"`
	MaxTouchPoints *int `json:"maxTouchPoints,omitempty"`
}

func (obj *Emulation) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send("Emulation.setTouchEmulationEnabled", request, nil)
	return
}

type SetEmitTouchEventsForMouseRequest struct {
	Enabled       bool    `json:"enabled"`
	Configuration *string `json:"configuration,omitempty"`
}

func (obj *Emulation) SetEmitTouchEventsForMouse(request *SetEmitTouchEventsForMouseRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmitTouchEventsForMouse", request, nil)
	return
}

type SetEmulatedMediaRequest struct {
	Media string `json:"media"`
}

func (obj *Emulation) SetEmulatedMedia(request *SetEmulatedMediaRequest) (err error) {
	err = obj.conn.Send("Emulation.setEmulatedMedia", request, nil)
	return
}

type SetCPUThrottlingRateRequest struct {
	Rate float32 `json:"rate"`
}

func (obj *Emulation) SetCPUThrottlingRate(request *SetCPUThrottlingRateRequest) (err error) {
	err = obj.conn.Send("Emulation.setCPUThrottlingRate", request, nil)
	return
}
func (obj *Emulation) CanEmulate() (response struct {
	Result bool `json:"result"`
}, err error) {
	err = obj.conn.Send("Emulation.canEmulate", nil, &response)
	return
}

type SetVirtualTimePolicyRequest struct {
	Policy                            types.Emulation_VirtualTimePolicy `json:"policy"`
	Budget                            *int                              `json:"budget,omitempty"`
	MaxVirtualTimeTaskStarvationCount *int                              `json:"maxVirtualTimeTaskStarvationCount,omitempty"`
}

func (obj *Emulation) SetVirtualTimePolicy(request *SetVirtualTimePolicyRequest) (err error) {
	err = obj.conn.Send("Emulation.setVirtualTimePolicy", request, nil)
	return
}

type SetNavigatorOverridesRequest struct {
	Platform string `json:"platform"`
}

func (obj *Emulation) SetNavigatorOverrides(request *SetNavigatorOverridesRequest) (err error) {
	err = obj.conn.Send("Emulation.setNavigatorOverrides", request, nil)
	return
}

type SetDefaultBackgroundColorOverrideRequest struct {
	Color *types.DOM_RGBA `json:"color,omitempty"`
}

func (obj *Emulation) SetDefaultBackgroundColorOverride(request *SetDefaultBackgroundColorOverrideRequest) (err error) {
	err = obj.conn.Send("Emulation.setDefaultBackgroundColorOverride", request, nil)
	return
}
