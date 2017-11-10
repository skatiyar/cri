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
	NodeId		types.DOM_NodeId	`json:"nodeId"`// ID of node to get the partial accessibility tree for.
	FetchRelatives	*bool			`json:"fetchRelatives,omitempty"`// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}

func (obj *Accessibility) GetPartialAXTree(request *GetPartialAXTreeRequest) (response struct {
	Nodes []types.Accessibility_AXNode `json:"nodes"`// The <code>Accessibility.AXNode</code> for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}, err error) {
	err = obj.conn.Send("Accessibility.getPartialAXTree", request, &response)
	return
}
