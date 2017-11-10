package dom

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type DOM struct {
	conn cri.Connector
}

func New(conn cri.Connector) *DOM {
	return &DOM{conn}
}
func (obj *DOM) Enable() (err error) {
	err = obj.conn.Send("DOM.enable", nil, nil)
	return
}
func (obj *DOM) Disable() (err error) {
	err = obj.conn.Send("DOM.disable", nil, nil)
	return
}

type GetDocumentRequest struct {
	Depth  *int  `json:"depth,omitempty"`
	Pierce *bool `json:"pierce,omitempty"`
}

func (obj *DOM) GetDocument(request *GetDocumentRequest) (response struct {
	Root types.DOM_Node `json:"root"`
}, err error) {
	err = obj.conn.Send("DOM.getDocument", request, &response)
	return
}

type GetFlattenedDocumentRequest struct {
	Depth  *int  `json:"depth,omitempty"`
	Pierce *bool `json:"pierce,omitempty"`
}

func (obj *DOM) GetFlattenedDocument(request *GetFlattenedDocumentRequest) (response struct {
	Nodes []types.DOM_Node `json:"nodes"`
}, err error) {
	err = obj.conn.Send("DOM.getFlattenedDocument", request, &response)
	return
}

type CollectClassNamesFromSubtreeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *DOM) CollectClassNamesFromSubtree(request *CollectClassNamesFromSubtreeRequest) (response struct {
	ClassNames []string `json:"classNames"`
}, err error) {
	err = obj.conn.Send("DOM.collectClassNamesFromSubtree", request, &response)
	return
}

type RequestChildNodesRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Depth  *int             `json:"depth,omitempty"`
	Pierce *bool            `json:"pierce,omitempty"`
}

func (obj *DOM) RequestChildNodes(request *RequestChildNodesRequest) (err error) {
	err = obj.conn.Send("DOM.requestChildNodes", request, nil)
	return
}

type QuerySelectorRequest struct {
	NodeId   types.DOM_NodeId `json:"nodeId"`
	Selector string           `json:"selector"`
}

func (obj *DOM) QuerySelector(request *QuerySelectorRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.querySelector", request, &response)
	return
}

type QuerySelectorAllRequest struct {
	NodeId   types.DOM_NodeId `json:"nodeId"`
	Selector string           `json:"selector"`
}

func (obj *DOM) QuerySelectorAll(request *QuerySelectorAllRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}, err error) {
	err = obj.conn.Send("DOM.querySelectorAll", request, &response)
	return
}

type SetNodeNameRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Name   string           `json:"name"`
}

func (obj *DOM) SetNodeName(request *SetNodeNameRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.setNodeName", request, &response)
	return
}

type SetNodeValueRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Value  string           `json:"value"`
}

func (obj *DOM) SetNodeValue(request *SetNodeValueRequest) (err error) {
	err = obj.conn.Send("DOM.setNodeValue", request, nil)
	return
}

type RemoveNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *DOM) RemoveNode(request *RemoveNodeRequest) (err error) {
	err = obj.conn.Send("DOM.removeNode", request, nil)
	return
}

type SetAttributeValueRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Name   string           `json:"name"`
	Value  string           `json:"value"`
}

func (obj *DOM) SetAttributeValue(request *SetAttributeValueRequest) (err error) {
	err = obj.conn.Send("DOM.setAttributeValue", request, nil)
	return
}

type SetAttributesAsTextRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Text   string           `json:"text"`
	Name   *string          `json:"name,omitempty"`
}

func (obj *DOM) SetAttributesAsText(request *SetAttributesAsTextRequest) (err error) {
	err = obj.conn.Send("DOM.setAttributesAsText", request, nil)
	return
}

type RemoveAttributeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
	Name   string           `json:"name"`
}

func (obj *DOM) RemoveAttribute(request *RemoveAttributeRequest) (err error) {
	err = obj.conn.Send("DOM.removeAttribute", request, nil)
	return
}

type GetOuterHTMLRequest struct {
	NodeId        *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId      *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

func (obj *DOM) GetOuterHTML(request *GetOuterHTMLRequest) (response struct {
	OuterHTML string `json:"outerHTML"`
}, err error) {
	err = obj.conn.Send("DOM.getOuterHTML", request, &response)
	return
}

type SetOuterHTMLRequest struct {
	NodeId    types.DOM_NodeId `json:"nodeId"`
	OuterHTML string           `json:"outerHTML"`
}

func (obj *DOM) SetOuterHTML(request *SetOuterHTMLRequest) (err error) {
	err = obj.conn.Send("DOM.setOuterHTML", request, nil)
	return
}

type PerformSearchRequest struct {
	Query                     string `json:"query"`
	IncludeUserAgentShadowDOM *bool  `json:"includeUserAgentShadowDOM,omitempty"`
}

func (obj *DOM) PerformSearch(request *PerformSearchRequest) (response struct {
	SearchId    string `json:"searchId"`
	ResultCount int    `json:"resultCount"`
}, err error) {
	err = obj.conn.Send("DOM.performSearch", request, &response)
	return
}

type GetSearchResultsRequest struct {
	SearchId  string `json:"searchId"`
	FromIndex int    `json:"fromIndex"`
	ToIndex   int    `json:"toIndex"`
}

func (obj *DOM) GetSearchResults(request *GetSearchResultsRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}, err error) {
	err = obj.conn.Send("DOM.getSearchResults", request, &response)
	return
}

type DiscardSearchResultsRequest struct {
	SearchId string `json:"searchId"`
}

func (obj *DOM) DiscardSearchResults(request *DiscardSearchResultsRequest) (err error) {
	err = obj.conn.Send("DOM.discardSearchResults", request, nil)
	return
}

type RequestNodeRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

func (obj *DOM) RequestNode(request *RequestNodeRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.requestNode", request, &response)
	return
}
func (obj *DOM) HighlightRect() (err error) {
	err = obj.conn.Send("DOM.highlightRect", nil, nil)
	return
}
func (obj *DOM) HighlightNode() (err error) {
	err = obj.conn.Send("DOM.highlightNode", nil, nil)
	return
}
func (obj *DOM) HideHighlight() (err error) {
	err = obj.conn.Send("DOM.hideHighlight", nil, nil)
	return
}

type PushNodeByPathToFrontendRequest struct {
	Path string `json:"path"`
}

func (obj *DOM) PushNodeByPathToFrontend(request *PushNodeByPathToFrontendRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.pushNodeByPathToFrontend", request, &response)
	return
}

type PushNodesByBackendIdsToFrontendRequest struct {
	BackendNodeIds []types.DOM_BackendNodeId `json:"backendNodeIds"`
}

func (obj *DOM) PushNodesByBackendIdsToFrontend(request *PushNodesByBackendIdsToFrontendRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}, err error) {
	err = obj.conn.Send("DOM.pushNodesByBackendIdsToFrontend", request, &response)
	return
}

type SetInspectedNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *DOM) SetInspectedNode(request *SetInspectedNodeRequest) (err error) {
	err = obj.conn.Send("DOM.setInspectedNode", request, nil)
	return
}

type ResolveNodeRequest struct {
	NodeId        *types.DOM_NodeId        `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	ObjectGroup   *string                  `json:"objectGroup,omitempty"`
}

func (obj *DOM) ResolveNode(request *ResolveNodeRequest) (response struct {
	Object types.Runtime_RemoteObject `json:"object"`
}, err error) {
	err = obj.conn.Send("DOM.resolveNode", request, &response)
	return
}

type GetAttributesRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *DOM) GetAttributes(request *GetAttributesRequest) (response struct {
	Attributes []string `json:"attributes"`
}, err error) {
	err = obj.conn.Send("DOM.getAttributes", request, &response)
	return
}

type CopyToRequest struct {
	NodeId             types.DOM_NodeId  `json:"nodeId"`
	TargetNodeId       types.DOM_NodeId  `json:"targetNodeId"`
	InsertBeforeNodeId *types.DOM_NodeId `json:"insertBeforeNodeId,omitempty"`
}

func (obj *DOM) CopyTo(request *CopyToRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.copyTo", request, &response)
	return
}

type MoveToRequest struct {
	NodeId             types.DOM_NodeId  `json:"nodeId"`
	TargetNodeId       types.DOM_NodeId  `json:"targetNodeId"`
	InsertBeforeNodeId *types.DOM_NodeId `json:"insertBeforeNodeId,omitempty"`
}

func (obj *DOM) MoveTo(request *MoveToRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.moveTo", request, &response)
	return
}
func (obj *DOM) Undo() (err error) {
	err = obj.conn.Send("DOM.undo", nil, nil)
	return
}
func (obj *DOM) Redo() (err error) {
	err = obj.conn.Send("DOM.redo", nil, nil)
	return
}
func (obj *DOM) MarkUndoableState() (err error) {
	err = obj.conn.Send("DOM.markUndoableState", nil, nil)
	return
}

type FocusRequest struct {
	NodeId        *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId      *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

func (obj *DOM) Focus(request *FocusRequest) (err error) {
	err = obj.conn.Send("DOM.focus", request, nil)
	return
}

type SetFileInputFilesRequest struct {
	Files         []string                      `json:"files"`
	NodeId        *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId      *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

func (obj *DOM) SetFileInputFiles(request *SetFileInputFilesRequest) (err error) {
	err = obj.conn.Send("DOM.setFileInputFiles", request, nil)
	return
}

type GetBoxModelRequest struct {
	NodeId        *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId      *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

func (obj *DOM) GetBoxModel(request *GetBoxModelRequest) (response struct {
	Model types.DOM_BoxModel `json:"model"`
}, err error) {
	err = obj.conn.Send("DOM.getBoxModel", request, &response)
	return
}

type GetNodeForLocationRequest struct {
	X                         int   `json:"x"`
	Y                         int   `json:"y"`
	IncludeUserAgentShadowDOM *bool `json:"includeUserAgentShadowDOM,omitempty"`
}

func (obj *DOM) GetNodeForLocation(request *GetNodeForLocationRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.getNodeForLocation", request, &response)
	return
}

type GetRelayoutBoundaryRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *DOM) GetRelayoutBoundary(request *GetRelayoutBoundaryRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}, err error) {
	err = obj.conn.Send("DOM.getRelayoutBoundary", request, &response)
	return
}

type DescribeNodeRequest struct {
	NodeId        *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId      *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
	Depth         *int                          `json:"depth,omitempty"`
	Pierce        *bool                         `json:"pierce,omitempty"`
}

func (obj *DOM) DescribeNode(request *DescribeNodeRequest) (response struct {
	Node types.DOM_Node `json:"node"`
}, err error) {
	err = obj.conn.Send("DOM.describeNode", request, &response)
	return
}
