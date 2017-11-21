/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package profiler provides commands and events for Profiler domain.
package profiler

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Profiler domain
const (
	Enable                = "Profiler.enable"
	Disable               = "Profiler.disable"
	SetSamplingInterval   = "Profiler.setSamplingInterval"
	Start                 = "Profiler.start"
	Stop                  = "Profiler.stop"
	StartPreciseCoverage  = "Profiler.startPreciseCoverage"
	StopPreciseCoverage   = "Profiler.stopPreciseCoverage"
	TakePreciseCoverage   = "Profiler.takePreciseCoverage"
	GetBestEffortCoverage = "Profiler.getBestEffortCoverage"
	StartTypeProfile      = "Profiler.startTypeProfile"
	StopTypeProfile       = "Profiler.stopTypeProfile"
	TakeTypeProfile       = "Profiler.takeTypeProfile"
)

// List of events in Profiler domain
const (
	ConsoleProfileStarted  = "Profiler.consoleProfileStarted"
	ConsoleProfileFinished = "Profiler.consoleProfileFinished"
)

type Profiler struct {
	conn cri.Connector
}

// New creates a Profiler instance
func New(conn cri.Connector) *Profiler {
	return &Profiler{conn}
}

func (obj *Profiler) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

func (obj *Profiler) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type SetSamplingIntervalRequest struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (obj *Profiler) SetSamplingInterval(request *SetSamplingIntervalRequest) (err error) {
	err = obj.conn.Send(SetSamplingInterval, request, nil)
	return
}

func (obj *Profiler) Start() (err error) {
	err = obj.conn.Send(Start, nil, nil)
	return
}

type StopResponse struct {
	// Recorded profile.
	Profile types.Profiler_Profile `json:"profile"`
}

func (obj *Profiler) Stop() (response StopResponse, err error) {
	err = obj.conn.Send(Stop, nil, &response)
	return
}

type StartPreciseCoverageRequest struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount *bool `json:"callCount,omitempty"`
	// Collect block-based coverage.
	Detailed *bool `json:"detailed,omitempty"`
}

// Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
func (obj *Profiler) StartPreciseCoverage(request *StartPreciseCoverageRequest) (err error) {
	err = obj.conn.Send(StartPreciseCoverage, request, nil)
	return
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (obj *Profiler) StopPreciseCoverage() (err error) {
	err = obj.conn.Send(StopPreciseCoverage, nil, nil)
	return
}

type TakePreciseCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
func (obj *Profiler) TakePreciseCoverage() (response TakePreciseCoverageResponse, err error) {
	err = obj.conn.Send(TakePreciseCoverage, nil, &response)
	return
}

type GetBestEffortCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
func (obj *Profiler) GetBestEffortCoverage() (response GetBestEffortCoverageResponse, err error) {
	err = obj.conn.Send(GetBestEffortCoverage, nil, &response)
	return
}

// Enable type profile.
func (obj *Profiler) StartTypeProfile() (err error) {
	err = obj.conn.Send(StartTypeProfile, nil, nil)
	return
}

// Disable type profile. Disabling releases type profile data collected so far.
func (obj *Profiler) StopTypeProfile() (err error) {
	err = obj.conn.Send(StopTypeProfile, nil, nil)
	return
}

type TakeTypeProfileResponse struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []types.Profiler_ScriptTypeProfile `json:"result"`
}

// Collect type profile.
func (obj *Profiler) TakeTypeProfile() (response TakeTypeProfileResponse, err error) {
	err = obj.conn.Send(TakeTypeProfile, nil, &response)
	return
}

type ConsoleProfileStartedParams struct {
	Id string `json:"id"`
	// Location of console.profile().
	Location types.Debugger_Location `json:"location"`
	// Profile title passed as an argument to console.profile().
	Title *string `json:"title,omitempty"`
}

// Sent when new profile recording is started using console.profile() call.
func (obj *Profiler) ConsoleProfileStarted(fn func(params *ConsoleProfileStartedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ConsoleProfileStarted, closeChn)
	go func() {
		for {
			params := ConsoleProfileStartedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ConsoleProfileFinishedParams struct {
	Id string `json:"id"`
	// Location of console.profileEnd().
	Location types.Debugger_Location `json:"location"`
	Profile  types.Profiler_Profile  `json:"profile"`
	// Profile title passed as an argument to console.profile().
	Title *string `json:"title,omitempty"`
}

func (obj *Profiler) ConsoleProfileFinished(fn func(params *ConsoleProfileFinishedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ConsoleProfileFinished, closeChn)
	go func() {
		for {
			params := ConsoleProfileFinishedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
