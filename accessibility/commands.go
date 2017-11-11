package accessibility

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Accessibility struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Accessibility {
	return &Accessibility{conn}
}

type GetPartialAXTreeRequest struct {
	// ID of node to get the partial accessibility tree for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
	FetchRelatives *bool `json:"fetchRelatives,omitempty"`
}
type GetPartialAXTreeResponse struct {
	// The <code>Accessibility.AXNode</code> for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
	Nodes []types.Accessibility_AXNode `json:"nodes"`
}

// Fetches the accessibility node and partial accessibility tree for this DOM node, if it exists.
func (obj *Accessibility) GetPartialAXTree(request *GetPartialAXTreeRequest) (response GetPartialAXTreeResponse, err error) {
	err = obj.conn.Send("Accessibility.getPartialAXTree", request, &response)
	return
}
