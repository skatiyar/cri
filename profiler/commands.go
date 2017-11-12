/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package profiler

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Profiler struct {
	conn cri.Connector
}

// New creates a Profiler instance
func New(conn cri.Connector) *Profiler {
	return &Profiler{conn}
}

func (obj *Profiler) Enable() (err error) {
	err = obj.conn.Send("Profiler.enable", nil, nil)
	return
}

func (obj *Profiler) Disable() (err error) {
	err = obj.conn.Send("Profiler.disable", nil, nil)
	return
}

type SetSamplingIntervalRequest struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (obj *Profiler) SetSamplingInterval(request *SetSamplingIntervalRequest) (err error) {
	err = obj.conn.Send("Profiler.setSamplingInterval", request, nil)
	return
}

func (obj *Profiler) Start() (err error) {
	err = obj.conn.Send("Profiler.start", nil, nil)
	return
}

type StopResponse struct {
	// Recorded profile.
	Profile types.Profiler_Profile `json:"profile"`
}

func (obj *Profiler) Stop() (response StopResponse, err error) {
	err = obj.conn.Send("Profiler.stop", nil, &response)
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
	err = obj.conn.Send("Profiler.startPreciseCoverage", request, nil)
	return
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (obj *Profiler) StopPreciseCoverage() (err error) {
	err = obj.conn.Send("Profiler.stopPreciseCoverage", nil, nil)
	return
}

type TakePreciseCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
func (obj *Profiler) TakePreciseCoverage() (response TakePreciseCoverageResponse, err error) {
	err = obj.conn.Send("Profiler.takePreciseCoverage", nil, &response)
	return
}

type GetBestEffortCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
func (obj *Profiler) GetBestEffortCoverage() (response GetBestEffortCoverageResponse, err error) {
	err = obj.conn.Send("Profiler.getBestEffortCoverage", nil, &response)
	return
}

// Enable type profile.
func (obj *Profiler) StartTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.startTypeProfile", nil, nil)
	return
}

// Disable type profile. Disabling releases type profile data collected so far.
func (obj *Profiler) StopTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.stopTypeProfile", nil, nil)
	return
}

type TakeTypeProfileResponse struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []types.Profiler_ScriptTypeProfile `json:"result"`
}

// Collect type profile.
func (obj *Profiler) TakeTypeProfile() (response TakeTypeProfileResponse, err error) {
	err = obj.conn.Send("Profiler.takeTypeProfile", nil, &response)
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
func (obj *Profiler) ConsoleProfileStarted(fn func(params *ConsoleProfileStartedParams) bool) {
	params := ConsoleProfileStartedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Profiler.consoleProfileStarted", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
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

func (obj *Profiler) ConsoleProfileFinished(fn func(params *ConsoleProfileFinishedParams) bool) {
	params := ConsoleProfileFinishedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Profiler.consoleProfileFinished", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}
