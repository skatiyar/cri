/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/


package heapprofiler

import (
    "github.com/SKatiyar/cri"
    types "github.com/SKatiyar/cri/types"
)

type HeapProfiler struct {
	conn cri.Connector
}

// New creates a HeapProfiler instance
func New(conn cri.Connector) *HeapProfiler {
	return &HeapProfiler{conn}
}

func (obj *HeapProfiler) Enable() (err error) {
	err = obj.conn.Send("HeapProfiler.enable", nil, nil)
	return
}


func (obj *HeapProfiler) Disable() (err error) {
	err = obj.conn.Send("HeapProfiler.disable", nil, nil)
	return
}


type StartTrackingHeapObjectsRequest struct {
	TrackAllocations *bool `json:"trackAllocations,omitempty"`
}


func (obj *HeapProfiler) StartTrackingHeapObjects(request *StartTrackingHeapObjectsRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.startTrackingHeapObjects", request, nil)
	return
}


type StopTrackingHeapObjectsRequest struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
	ReportProgress *bool `json:"reportProgress,omitempty"`
}


func (obj *HeapProfiler) StopTrackingHeapObjects(request *StopTrackingHeapObjectsRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.stopTrackingHeapObjects", request, nil)
	return
}


type TakeHeapSnapshotRequest struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress *bool `json:"reportProgress,omitempty"`
}


func (obj *HeapProfiler) TakeHeapSnapshot(request *TakeHeapSnapshotRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.takeHeapSnapshot", request, nil)
	return
}


func (obj *HeapProfiler) CollectGarbage() (err error) {
	err = obj.conn.Send("HeapProfiler.collectGarbage", nil, nil)
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
	err = obj.conn.Send("HeapProfiler.getObjectByHeapObjectId", request, &response)
	return
}


type AddInspectedHeapObjectRequest struct {
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"heapObjectId"`
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (obj *HeapProfiler) AddInspectedHeapObject(request *AddInspectedHeapObjectRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.addInspectedHeapObject", request, nil)
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
	err = obj.conn.Send("HeapProfiler.getHeapObjectId", request, &response)
	return
}


type StartSamplingRequest struct {
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
	SamplingInterval *float32 `json:"samplingInterval,omitempty"`
}


func (obj *HeapProfiler) StartSampling(request *StartSamplingRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.startSampling", request, nil)
	return
}


type StopSamplingResponse struct {
	// Recorded sampling heap profile.
	Profile types.HeapProfiler_SamplingHeapProfile `json:"profile"`
}


func (obj *HeapProfiler) StopSampling() (response StopSamplingResponse, err error) {
	err = obj.conn.Send("HeapProfiler.stopSampling", nil, &response)
	return
}
