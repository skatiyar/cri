package profiler

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Profiler struct {
	conn cri.Connector
}

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
	Interval int `json:"interval"`
}

func (obj *Profiler) SetSamplingInterval(request *SetSamplingIntervalRequest) (err error) {
	err = obj.conn.Send("Profiler.setSamplingInterval", request, nil)
	return
}
func (obj *Profiler) Start() (err error) {
	err = obj.conn.Send("Profiler.start", nil, nil)
	return
}
func (obj *Profiler) Stop() (response struct {
	Profile types.Profiler_Profile `json:"profile"`
}, err error) {
	err = obj.conn.Send("Profiler.stop", nil, &response)
	return
}

type StartPreciseCoverageRequest struct {
	CallCount *bool `json:"callCount,omitempty"`
	Detailed  *bool `json:"detailed,omitempty"`
}

func (obj *Profiler) StartPreciseCoverage(request *StartPreciseCoverageRequest) (err error) {
	err = obj.conn.Send("Profiler.startPreciseCoverage", request, nil)
	return
}
func (obj *Profiler) StopPreciseCoverage() (err error) {
	err = obj.conn.Send("Profiler.stopPreciseCoverage", nil, nil)
	return
}
func (obj *Profiler) TakePreciseCoverage() (response struct {
	Result []types.Profiler_ScriptCoverage `json:"result"`
}, err error) {
	err = obj.conn.Send("Profiler.takePreciseCoverage", nil, &response)
	return
}
func (obj *Profiler) GetBestEffortCoverage() (response struct {
	Result []types.Profiler_ScriptCoverage `json:"result"`
}, err error) {
	err = obj.conn.Send("Profiler.getBestEffortCoverage", nil, &response)
	return
}
func (obj *Profiler) StartTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.startTypeProfile", nil, nil)
	return
}
func (obj *Profiler) StopTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.stopTypeProfile", nil, nil)
	return
}
func (obj *Profiler) TakeTypeProfile() (response struct {
	Result []types.Profiler_ScriptTypeProfile `json:"result"`
}, err error) {
	err = obj.conn.Send("Profiler.takeTypeProfile", nil, &response)
	return
}
