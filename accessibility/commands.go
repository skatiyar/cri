/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package accessibility provides commands and events for Accessibility domain.
package accessibility

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Accessibility domain
const (
	GetPartialAXTree = "Accessibility.getPartialAXTree"
)

type Accessibility struct {
	conn cri.Connector
}

// New creates a Accessibility instance
func New(conn cri.Connector) *Accessibility {
	return &Accessibility{conn}
}

type GetPartialAXTreeRequest struct {
	// Identifier of the node to get the partial accessibility tree for.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node to get the partial accessibility tree for.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper to get the partial accessibility tree for.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives *bool `json:"fetchRelatives,omitempty"`
}

type GetPartialAXTreeResponse struct {
	// The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
	Nodes []types.Accessibility_AXNode `json:"nodes"`
}

// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
func (obj *Accessibility) GetPartialAXTree(request *GetPartialAXTreeRequest) (response GetPartialAXTreeResponse, err error) {
	err = obj.conn.Send(GetPartialAXTree, request, &response)
	return
}
