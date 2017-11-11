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
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}
type GetSnapshotResponse struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	DomNodes []types.DOMSnapshot_DOMNode `json:"domNodes"`
	// The nodes in the layout tree.
	LayoutTreeNodes []types.DOMSnapshot_LayoutTreeNode `json:"layoutTreeNodes"`
	// Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles []types.DOMSnapshot_ComputedStyle `json:"computedStyles"`
}

// Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
func (obj *DOMSnapshot) GetSnapshot(request *GetSnapshotRequest) (response GetSnapshotResponse, err error) {
	err = obj.conn.Send("DOMSnapshot.getSnapshot", request, &response)
	return
}
