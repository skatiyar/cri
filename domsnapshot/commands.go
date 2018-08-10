/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package domsnapshot provides commands and events for DOMSnapshot domain.
package domsnapshot

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in DOMSnapshot domain
const (
	Disable         = "DOMSnapshot.disable"
	Enable          = "DOMSnapshot.enable"
	GetSnapshot     = "DOMSnapshot.getSnapshot"
	CaptureSnapshot = "DOMSnapshot.captureSnapshot"
)

// This domain facilitates obtaining document snapshots with DOM, layout, and style information.
type DOMSnapshot struct {
	conn cri.Connector
}

// New creates a DOMSnapshot instance
func New(conn cri.Connector) *DOMSnapshot {
	return &DOMSnapshot{conn}
}

// Disables DOM snapshot agent for the given page.
func (obj *DOMSnapshot) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables DOM snapshot agent for the given page.
func (obj *DOMSnapshot) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type GetSnapshotRequest struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
	// Whether or not to retrieve details of DOM listeners (default false).
	IncludeEventListeners *bool `json:"includeEventListeners,omitempty"`
	// Whether to determine and include the paint order index of LayoutTreeNodes (default false).
	IncludePaintOrder *bool `json:"includePaintOrder,omitempty"`
	// Whether to include UA shadow tree in the snapshot (default false).
	IncludeUserAgentShadowTree *bool `json:"includeUserAgentShadowTree,omitempty"`
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
	err = obj.conn.Send(GetSnapshot, request, &response)
	return
}

type CaptureSnapshotRequest struct {
	// Whitelist of computed styles to return.
	ComputedStyles []string `json:"computedStyles"`
}

type CaptureSnapshotResponse struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	Documents []types.DOMSnapshot_DocumentSnapshot `json:"documents"`
	// Shared string table that all string properties refer to with indexes.
	Strings []string `json:"strings"`
}

// Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
func (obj *DOMSnapshot) CaptureSnapshot(request *CaptureSnapshotRequest) (response CaptureSnapshotResponse, err error) {
	err = obj.conn.Send(CaptureSnapshot, request, &response)
	return
}
