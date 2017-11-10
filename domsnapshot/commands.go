package domsnapshot

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type DOMSnapshot struct {
	conn cri.Connector
}

func New(conn cri.Connector) *DOMSnapshot {
	return &DOMSnapshot{conn}
}

type GetSnapshotRequest struct {
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

func (obj *DOMSnapshot) GetSnapshot(request *GetSnapshotRequest) (response struct {
	DomNodes        []types.DOMSnapshot_DOMNode        `json:"domNodes"`
	LayoutTreeNodes []types.DOMSnapshot_LayoutTreeNode `json:"layoutTreeNodes"`
	ComputedStyles  []types.DOMSnapshot_ComputedStyle  `json:"computedStyles"`
}, err error) {
	err = obj.conn.Send("DOMSnapshot.getSnapshot", request, &response)
	return
}
