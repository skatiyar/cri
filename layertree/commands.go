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
	LayerId types.LayerTree_LayerId `json:"layerId"`// The id of the layer for which we want to get the reasons it was composited.
}

func (obj *LayerTree) CompositingReasons(request *CompositingReasonsRequest) (response struct {
	CompositingReasons []string `json:"compositingReasons"`// A list of strings specifying reasons for the given layer to become composited.
}, err error) {
	err = obj.conn.Send("LayerTree.compositingReasons", request, &response)
	return
}

type MakeSnapshotRequest struct {
	LayerId types.LayerTree_LayerId `json:"layerId"`// The id of the layer.
}

func (obj *LayerTree) MakeSnapshot(request *MakeSnapshotRequest) (response struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}, err error) {
	err = obj.conn.Send("LayerTree.makeSnapshot", request, &response)
	return
}

type LoadSnapshotRequest struct {
	Tiles []types.LayerTree_PictureTile `json:"tiles"`// An array of tiles composing the snapshot.
}

func (obj *LayerTree) LoadSnapshot(request *LoadSnapshotRequest) (response struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`// The id of the snapshot.
}, err error) {
	err = obj.conn.Send("LayerTree.loadSnapshot", request, &response)
	return
}

type ReleaseSnapshotRequest struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}

func (obj *LayerTree) ReleaseSnapshot(request *ReleaseSnapshotRequest) (err error) {
	err = obj.conn.Send("LayerTree.releaseSnapshot", request, nil)
	return
}

type ProfileSnapshotRequest struct {
	SnapshotId	types.LayerTree_SnapshotId	`json:"snapshotId"`// The id of the layer snapshot.
	MinRepeatCount	*int				`json:"minRepeatCount,omitempty"`// The maximum number of times to replay the snapshot (1, if not specified).
	MinDuration	*float32			`json:"minDuration,omitempty"`// The minimum duration (in seconds) to replay the snapshot.
	ClipRect	*types.DOM_Rect			`json:"clipRect,omitempty"`// The clip rectangle to apply when replaying the snapshot.
}

func (obj *LayerTree) ProfileSnapshot(request *ProfileSnapshotRequest) (response struct {
	Timings []types.LayerTree_PaintProfile `json:"timings"`// The array of paint profiles, one per run.
}, err error) {
	err = obj.conn.Send("LayerTree.profileSnapshot", request, &response)
	return
}

type ReplaySnapshotRequest struct {
	SnapshotId	types.LayerTree_SnapshotId	`json:"snapshotId"`// The id of the layer snapshot.
	FromStep	*int				`json:"fromStep,omitempty"`// The first step to replay from (replay from the very start if not specified).
	ToStep		*int				`json:"toStep,omitempty"`// The last step to replay to (replay till the end if not specified).
	Scale		*float32			`json:"scale,omitempty"`// The scale to apply while replaying (defaults to 1).
}

func (obj *LayerTree) ReplaySnapshot(request *ReplaySnapshotRequest) (response struct {
	DataURL string `json:"dataURL"`// A data: URL for resulting image.
}, err error) {
	err = obj.conn.Send("LayerTree.replaySnapshot", request, &response)
	return
}

type SnapshotCommandLogRequest struct {
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}

func (obj *LayerTree) SnapshotCommandLog(request *SnapshotCommandLogRequest) (response struct {
	CommandLog []map[string]interface{} `json:"commandLog"`// The array of canvas function calls.
}, err error) {
	err = obj.conn.Send("LayerTree.snapshotCommandLog", request, &response)
	return
}
