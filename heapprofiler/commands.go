/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package heapprofiler provides commands and events for HeapProfiler domain.
package heapprofiler

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in HeapProfiler domain
const (
	Enable                   = "HeapProfiler.enable"
	Disable                  = "HeapProfiler.disable"
	StartTrackingHeapObjects = "HeapProfiler.startTrackingHeapObjects"
	StopTrackingHeapObjects  = "HeapProfiler.stopTrackingHeapObjects"
	TakeHeapSnapshot         = "HeapProfiler.takeHeapSnapshot"
	CollectGarbage           = "HeapProfiler.collectGarbage"
	GetObjectByHeapObjectId  = "HeapProfiler.getObjectByHeapObjectId"
	AddInspectedHeapObject   = "HeapProfiler.addInspectedHeapObject"
	GetHeapObjectId          = "HeapProfiler.getHeapObjectId"
	StartSampling            = "HeapProfiler.startSampling"
	StopSampling             = "HeapProfiler.stopSampling"
	GetSamplingProfile       = "HeapProfiler.getSamplingProfile"
)

// List of events in HeapProfiler domain
const (
	AddHeapSnapshotChunk       = "HeapProfiler.addHeapSnapshotChunk"
	ResetProfiles              = "HeapProfiler.resetProfiles"
	ReportHeapSnapshotProgress = "HeapProfiler.reportHeapSnapshotProgress"
	LastSeenObjectId           = "HeapProfiler.lastSeenObjectId"
	HeapStatsUpdate            = "HeapProfiler.heapStatsUpdate"
)

type HeapProfiler struct {
	conn cri.Connector
}

// New creates a HeapProfiler instance
func New(conn cri.Connector) *HeapProfiler {
	return &HeapProfiler{conn}
}

func (obj *HeapProfiler) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

func (obj *HeapProfiler) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type StartTrackingHeapObjectsRequest struct {
	TrackAllocations *bool `json:"trackAllocations,omitempty"`
}

func (obj *HeapProfiler) StartTrackingHeapObjects(request *StartTrackingHeapObjectsRequest) (err error) {
	err = obj.conn.Send(StartTrackingHeapObjects, request, nil)
	return
}

type StopTrackingHeapObjectsRequest struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
	ReportProgress *bool `json:"reportProgress,omitempty"`
}

func (obj *HeapProfiler) StopTrackingHeapObjects(request *StopTrackingHeapObjectsRequest) (err error) {
	err = obj.conn.Send(StopTrackingHeapObjects, request, nil)
	return
}

type TakeHeapSnapshotRequest struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress *bool `json:"reportProgress,omitempty"`
}

func (obj *HeapProfiler) TakeHeapSnapshot(request *TakeHeapSnapshotRequest) (err error) {
	err = obj.conn.Send(TakeHeapSnapshot, request, nil)
	return
}

func (obj *HeapProfiler) CollectGarbage() (err error) {
	err = obj.conn.Send(CollectGarbage, nil, nil)
	return
}

type GetObjectByHeapObjectIdRequest struct {
	ObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"objectId"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
}

type GetObjectByHeapObjectIdResponse struct {
	// Evaluation result.
	Result types.Runtime_RemoteObject `json:"result"`
}

func (obj *HeapProfiler) GetObjectByHeapObjectId(request *GetObjectByHeapObjectIdRequest) (response GetObjectByHeapObjectIdResponse, err error) {
	err = obj.conn.Send(GetObjectByHeapObjectId, request, &response)
	return
}

type AddInspectedHeapObjectRequest struct {
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"heapObjectId"`
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (obj *HeapProfiler) AddInspectedHeapObject(request *AddInspectedHeapObjectRequest) (err error) {
	err = obj.conn.Send(AddInspectedHeapObject, request, nil)
	return
}

type GetHeapObjectIdRequest struct {
	// Identifier of the object to get heap object id for.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

type GetHeapObjectIdResponse struct {
	// Id of the heap snapshot object corresponding to the passed remote object id.
	HeapSnapshotObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"heapSnapshotObjectId"`
}

func (obj *HeapProfiler) GetHeapObjectId(request *GetHeapObjectIdRequest) (response GetHeapObjectIdResponse, err error) {
	err = obj.conn.Send(GetHeapObjectId, request, &response)
	return
}

type StartSamplingRequest struct {
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
	SamplingInterval *float32 `json:"samplingInterval,omitempty"`
}

func (obj *HeapProfiler) StartSampling(request *StartSamplingRequest) (err error) {
	err = obj.conn.Send(StartSampling, request, nil)
	return
}

type StopSamplingResponse struct {
	// Recorded sampling heap profile.
	Profile types.HeapProfiler_SamplingHeapProfile `json:"profile"`
}

func (obj *HeapProfiler) StopSampling() (response StopSamplingResponse, err error) {
	err = obj.conn.Send(StopSampling, nil, &response)
	return
}

type GetSamplingProfileResponse struct {
	// Return the sampling profile being collected.
	Profile types.HeapProfiler_SamplingHeapProfile `json:"profile"`
}

func (obj *HeapProfiler) GetSamplingProfile() (response GetSamplingProfileResponse, err error) {
	err = obj.conn.Send(GetSamplingProfile, nil, &response)
	return
}

type AddHeapSnapshotChunkParams struct {
	Chunk string `json:"chunk"`
}

func (obj *HeapProfiler) AddHeapSnapshotChunk(fn func(params *AddHeapSnapshotChunkParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AddHeapSnapshotChunk, closeChn)
	go func() {
		for {
			params := AddHeapSnapshotChunkParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

func (obj *HeapProfiler) ResetProfiles(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ResetProfiles, closeChn)
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

type ReportHeapSnapshotProgressParams struct {
	Done     int   `json:"done"`
	Total    int   `json:"total"`
	Finished *bool `json:"finished,omitempty"`
}

func (obj *HeapProfiler) ReportHeapSnapshotProgress(fn func(params *ReportHeapSnapshotProgressParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ReportHeapSnapshotProgress, closeChn)
	go func() {
		for {
			params := ReportHeapSnapshotProgressParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type LastSeenObjectIdParams struct {
	LastSeenObjectId int     `json:"lastSeenObjectId"`
	Timestamp        float32 `json:"timestamp"`
}

// If heap objects tracking has been started then backend regularly sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
func (obj *HeapProfiler) LastSeenObjectId(fn func(params *LastSeenObjectIdParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(LastSeenObjectId, closeChn)
	go func() {
		for {
			params := LastSeenObjectIdParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type HeapStatsUpdateParams struct {
	// An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
	StatsUpdate []int `json:"statsUpdate"`
}

// If heap objects tracking has been started then backend may send update for one or more fragments
func (obj *HeapProfiler) HeapStatsUpdate(fn func(params *HeapStatsUpdateParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(HeapStatsUpdate, closeChn)
	go func() {
		for {
			params := HeapStatsUpdateParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
