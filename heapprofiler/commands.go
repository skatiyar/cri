package heapprofiler

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type HeapProfiler struct {
	conn cri.Connector
}

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
	ReportProgress *bool `json:"reportProgress,omitempty"`
}

func (obj *HeapProfiler) StopTrackingHeapObjects(request *StopTrackingHeapObjectsRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.stopTrackingHeapObjects", request, nil)
	return
}

type TakeHeapSnapshotRequest struct {
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
	ObjectId    types.HeapProfiler_HeapSnapshotObjectId `json:"objectId"`
	ObjectGroup *string                                 `json:"objectGroup,omitempty"`
}

func (obj *HeapProfiler) GetObjectByHeapObjectId(request *GetObjectByHeapObjectIdRequest) (response struct {
	Result types.Runtime_RemoteObject `json:"result"`
}, err error) {
	err = obj.conn.Send("HeapProfiler.getObjectByHeapObjectId", request, &response)
	return
}

type AddInspectedHeapObjectRequest struct {
	HeapObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"heapObjectId"`
}

func (obj *HeapProfiler) AddInspectedHeapObject(request *AddInspectedHeapObjectRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.addInspectedHeapObject", request, nil)
	return
}

type GetHeapObjectIdRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

func (obj *HeapProfiler) GetHeapObjectId(request *GetHeapObjectIdRequest) (response struct {
	HeapSnapshotObjectId types.HeapProfiler_HeapSnapshotObjectId `json:"heapSnapshotObjectId"`
}, err error) {
	err = obj.conn.Send("HeapProfiler.getHeapObjectId", request, &response)
	return
}

type StartSamplingRequest struct {
	SamplingInterval *float32 `json:"samplingInterval,omitempty"`
}

func (obj *HeapProfiler) StartSampling(request *StartSamplingRequest) (err error) {
	err = obj.conn.Send("HeapProfiler.startSampling", request, nil)
	return
}
func (obj *HeapProfiler) StopSampling() (response struct {
	Profile types.HeapProfiler_SamplingHeapProfile `json:"profile"`
}, err error) {
	err = obj.conn.Send("HeapProfiler.stopSampling", nil, &response)
	return
}
