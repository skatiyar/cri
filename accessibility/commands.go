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
	NodeId         types.DOM_NodeId `json:"nodeId"`
	FetchRelatives *bool            `json:"fetchRelatives,omitempty"`
}

func (obj *Accessibility) GetPartialAXTree(request *GetPartialAXTreeRequest) (response struct {
	Nodes []types.Accessibility_AXNode `json:"nodes"`
}, err error) {
	err = obj.conn.Send("Accessibility.getPartialAXTree", request, &response)
	return
}
