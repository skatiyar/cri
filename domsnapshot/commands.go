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
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`// Whitelist of computed styles to return.
}

func (obj *DOMSnapshot) GetSnapshot(request *GetSnapshotRequest) (response struct {
	DomNodes	[]types.DOMSnapshot_DOMNode		`json:"domNodes"`// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	LayoutTreeNodes	[]types.DOMSnapshot_LayoutTreeNode	`json:"layoutTreeNodes"`// The nodes in the layout tree.
	ComputedStyles	[]types.DOMSnapshot_ComputedStyle	`json:"computedStyles"`// Whitelisted ComputedStyle properties for each node in the layout tree.
}, err error) {
	err = obj.conn.Send("DOMSnapshot.getSnapshot", request, &response)
	return
}
