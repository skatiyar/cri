/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package memory provides commands and events for Memory domain.
package memory

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Memory domain
const (
	GetDOMCounters                     = "Memory.getDOMCounters"
	PrepareForLeakDetection            = "Memory.prepareForLeakDetection"
	SetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"
	SimulatePressureNotification       = "Memory.simulatePressureNotification"
	StartSampling                      = "Memory.startSampling"
	StopSampling                       = "Memory.stopSampling"
	GetAllTimeSamplingProfile          = "Memory.getAllTimeSamplingProfile"
	GetBrowserSamplingProfile          = "Memory.getBrowserSamplingProfile"
	GetSamplingProfile                 = "Memory.getSamplingProfile"
)

type Memory struct {
	conn cri.Connector
}

// New creates a Memory instance
func New(conn cri.Connector) *Memory {
	return &Memory{conn}
}

type GetDOMCountersResponse struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}

func (obj *Memory) GetDOMCounters() (response GetDOMCountersResponse, err error) {
	err = obj.conn.Send(GetDOMCounters, nil, &response)
	return
}

func (obj *Memory) PrepareForLeakDetection() (err error) {
	err = obj.conn.Send(PrepareForLeakDetection, nil, nil)
	return
}

type SetPressureNotificationsSuppressedRequest struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// Enable/disable suppressing memory pressure notifications in all processes.
func (obj *Memory) SetPressureNotificationsSuppressed(request *SetPressureNotificationsSuppressedRequest) (err error) {
	err = obj.conn.Send(SetPressureNotificationsSuppressed, request, nil)
	return
}

type SimulatePressureNotificationRequest struct {
	// Memory pressure level of the notification.
	Level types.Memory_PressureLevel `json:"level"`
}

// Simulate a memory pressure notification in all processes.
func (obj *Memory) SimulatePressureNotification(request *SimulatePressureNotificationRequest) (err error) {
	err = obj.conn.Send(SimulatePressureNotification, request, nil)
	return
}

type StartSamplingRequest struct {
	// Average number of bytes between samples.
	SamplingInterval *int `json:"samplingInterval,omitempty"`
	// Do not randomize intervals between samples.
	SuppressRandomness *bool `json:"suppressRandomness,omitempty"`
}

// Start collecting native memory profile.
func (obj *Memory) StartSampling(request *StartSamplingRequest) (err error) {
	err = obj.conn.Send(StartSampling, request, nil)
	return
}

// Stop collecting native memory profile.
func (obj *Memory) StopSampling() (err error) {
	err = obj.conn.Send(StopSampling, nil, nil)
	return
}

type GetAllTimeSamplingProfileResponse struct {
	Profile types.Memory_SamplingProfile `json:"profile"`
}

// Retrieve native memory allocations profile collected since renderer process startup.
func (obj *Memory) GetAllTimeSamplingProfile() (response GetAllTimeSamplingProfileResponse, err error) {
	err = obj.conn.Send(GetAllTimeSamplingProfile, nil, &response)
	return
}

type GetBrowserSamplingProfileResponse struct {
	Profile types.Memory_SamplingProfile `json:"profile"`
}

// Retrieve native memory allocations profile collected since browser process startup.
func (obj *Memory) GetBrowserSamplingProfile() (response GetBrowserSamplingProfileResponse, err error) {
	err = obj.conn.Send(GetBrowserSamplingProfile, nil, &response)
	return
}

type GetSamplingProfileResponse struct {
	Profile types.Memory_SamplingProfile `json:"profile"`
}

// Retrieve native memory allocations profile collected since last `startSampling` call.
func (obj *Memory) GetSamplingProfile() (response GetSamplingProfileResponse, err error) {
	err = obj.conn.Send(GetSamplingProfile, nil, &response)
	return
}
