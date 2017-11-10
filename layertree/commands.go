package layertree

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type LayerTree struct {
	conn cri.Connector
}

func New(conn cri.Connector) *LayerTree {
	return &LayerTree{conn}
}
func (obj *LayerTree) Enable() (err error) {
	err = obj.conn.Send("LayerTree.enable", nil, nil)
	return
}
func (obj *LayerTree) Disable() (err error) {
	err = obj.conn.Send("LayerTree.disable", nil, nil)
	return
}

type CompositingReasonsRequest struct {
	LayerId types.LayerTree_LayerId `json:"layerId"`
}

func (obj *LayerTree) CompositingReasons(request *CompositingReasonsRequest) (response struct {
	CompositingReasons []string `json:"compositingReasons"`
}, err error) {
	err = obj.conn.Send("LayerTree.compositingReasons", request, &response)
	return
}

type MakeSnapshotRequest struct {
	LayerId types.LayerTree_LayerId `json:"layerId"`
}

func (obj *LayerTree) MakeSnapshot(request *MakeSnapshotRequest) (response struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}, err error) {
	err = obj.conn.Send("LayerTree.makeSnapshot", request, &response)
	return
}

type LoadSnapshotRequest struct {
	Tiles []types.LayerTree_PictureTile `json:"tiles"`
}

func (obj *LayerTree) LoadSnapshot(request *LoadSnapshotRequest) (response struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}, err error) {
	err = obj.conn.Send("LayerTree.loadSnapshot", request, &response)
	return
}

type ReleaseSnapshotRequest struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

func (obj *LayerTree) ReleaseSnapshot(request *ReleaseSnapshotRequest) (err error) {
	err = obj.conn.Send("LayerTree.releaseSnapshot", request, nil)
	return
}

type ProfileSnapshotRequest struct {
	SnapshotId     types.LayerTree_SnapshotId `json:"snapshotId"`
	MinRepeatCount *int                       `json:"minRepeatCount,omitempty"`
	MinDuration    *float32                   `json:"minDuration,omitempty"`
	ClipRect       *types.DOM_Rect            `json:"clipRect,omitempty"`
}

func (obj *LayerTree) ProfileSnapshot(request *ProfileSnapshotRequest) (response struct {
	Timings []types.LayerTree_PaintProfile `json:"timings"`
}, err error) {
	err = obj.conn.Send("LayerTree.profileSnapshot", request, &response)
	return
}

type ReplaySnapshotRequest struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
	FromStep   *int                       `json:"fromStep,omitempty"`
	ToStep     *int                       `json:"toStep,omitempty"`
	Scale      *float32                   `json:"scale,omitempty"`
}

func (obj *LayerTree) ReplaySnapshot(request *ReplaySnapshotRequest) (response struct {
	DataURL string `json:"dataURL"`
}, err error) {
	err = obj.conn.Send("LayerTree.replaySnapshot", request, &response)
	return
}

type SnapshotCommandLogRequest struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

func (obj *LayerTree) SnapshotCommandLog(request *SnapshotCommandLogRequest) (response struct {
	CommandLog []map[string]interface{} `json:"commandLog"`
}, err error) {
	err = obj.conn.Send("LayerTree.snapshotCommandLog", request, &response)
	return
}
