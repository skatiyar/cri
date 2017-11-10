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
	Depth	*int	`json:"depth,omitempty"`// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce	*bool	`json:"pierce,omitempty"`// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}

func (obj *DOM) GetDocument(request *GetDocumentRequest) (response struct {
	Root types.DOM_Node `json:"root"`// Resulting node.
}, err error) {
	err = obj.conn.Send("DOM.getDocument", request, &response)
	return
}

type GetFlattenedDocumentRequest struct {
	Depth	*int	`json:"depth,omitempty"`// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce	*bool	`json:"pierce,omitempty"`// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}

func (obj *DOM) GetFlattenedDocument(request *GetFlattenedDocumentRequest) (response struct {
	Nodes []types.DOM_Node `json:"nodes"`// Resulting node.
}, err error) {
	err = obj.conn.Send("DOM.getFlattenedDocument", request, &response)
	return
}

type CollectClassNamesFromSubtreeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node to collect class names.
}

func (obj *DOM) CollectClassNamesFromSubtree(request *CollectClassNamesFromSubtreeRequest) (response struct {
	ClassNames []string `json:"classNames"`// Class name list.
}, err error) {
	err = obj.conn.Send("DOM.collectClassNamesFromSubtree", request, &response)
	return
}

type RequestChildNodesRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the node to get children for.
	Depth	*int			`json:"depth,omitempty"`// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce	*bool			`json:"pierce,omitempty"`// Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
}

func (obj *DOM) RequestChildNodes(request *RequestChildNodesRequest) (err error) {
	err = obj.conn.Send("DOM.requestChildNodes", request, nil)
	return
}

type QuerySelectorRequest struct {
	NodeId		types.DOM_NodeId	`json:"nodeId"`// Id of the node to query upon.
	Selector	string			`json:"selector"`// Selector string.
}

func (obj *DOM) QuerySelector(request *QuerySelectorRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Query selector result.
}, err error) {
	err = obj.conn.Send("DOM.querySelector", request, &response)
	return
}

type QuerySelectorAllRequest struct {
	NodeId		types.DOM_NodeId	`json:"nodeId"`// Id of the node to query upon.
	Selector	string			`json:"selector"`// Selector string.
}

func (obj *DOM) QuerySelectorAll(request *QuerySelectorAllRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`// Query selector result.
}, err error) {
	err = obj.conn.Send("DOM.querySelectorAll", request, &response)
	return
}

type SetNodeNameRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the node to set name for.
	Name	string			`json:"name"`// New node's name.
}

func (obj *DOM) SetNodeName(request *SetNodeNameRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// New node's id.
}, err error) {
	err = obj.conn.Send("DOM.setNodeName", request, &response)
	return
}

type SetNodeValueRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the node to set value for.
	Value	string			`json:"value"`// New node's value.
}

func (obj *DOM) SetNodeValue(request *SetNodeValueRequest) (err error) {
	err = obj.conn.Send("DOM.setNodeValue", request, nil)
	return
}

type RemoveNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node to remove.
}

func (obj *DOM) RemoveNode(request *RemoveNodeRequest) (err error) {
	err = obj.conn.Send("DOM.removeNode", request, nil)
	return
}

type SetAttributeValueRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the element to set attribute for.
	Name	string			`json:"name"`// Attribute name.
	Value	string			`json:"value"`// Attribute value.
}

func (obj *DOM) SetAttributeValue(request *SetAttributeValueRequest) (err error) {
	err = obj.conn.Send("DOM.setAttributeValue", request, nil)
	return
}

type SetAttributesAsTextRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the element to set attributes for.
	Text	string			`json:"text"`// Text with a number of attributes. Will parse this text using HTML parser.
	Name	*string			`json:"name,omitempty"`// Attribute name to replace with new attributes derived from text in case text parsed successfully.
}

func (obj *DOM) SetAttributesAsText(request *SetAttributesAsTextRequest) (err error) {
	err = obj.conn.Send("DOM.setAttributesAsText", request, nil)
	return
}

type RemoveAttributeRequest struct {
	NodeId	types.DOM_NodeId	`json:"nodeId"`// Id of the element to remove attribute from.
	Name	string			`json:"name"`// Name of the attribute to remove.
}

func (obj *DOM) RemoveAttribute(request *RemoveAttributeRequest) (err error) {
	err = obj.conn.Send("DOM.removeAttribute", request, nil)
	return
}

type GetOuterHTMLRequest struct {
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node wrapper.
}

func (obj *DOM) GetOuterHTML(request *GetOuterHTMLRequest) (response struct {
	OuterHTML string `json:"outerHTML"`// Outer HTML markup.
}, err error) {
	err = obj.conn.Send("DOM.getOuterHTML", request, &response)
	return
}

type SetOuterHTMLRequest struct {
	NodeId		types.DOM_NodeId	`json:"nodeId"`// Id of the node to set markup for.
	OuterHTML	string			`json:"outerHTML"`// Outer HTML markup to set.
}

func (obj *DOM) SetOuterHTML(request *SetOuterHTMLRequest) (err error) {
	err = obj.conn.Send("DOM.setOuterHTML", request, nil)
	return
}

type PerformSearchRequest struct {
	Query				string	`json:"query"`// Plain text or query selector or XPath search query.
	IncludeUserAgentShadowDOM	*bool	`json:"includeUserAgentShadowDOM,omitempty"`// True to search in user agent shadow DOM.
}

func (obj *DOM) PerformSearch(request *PerformSearchRequest) (response struct {
	SearchId	string	`json:"searchId"`// Unique search session identifier.
	ResultCount	int	`json:"resultCount"`// Number of search results.
}, err error) {
	err = obj.conn.Send("DOM.performSearch", request, &response)
	return
}

type GetSearchResultsRequest struct {
	SearchId	string	`json:"searchId"`// Unique search session identifier.
	FromIndex	int	`json:"fromIndex"`// Start index of the search result to be returned.
	ToIndex		int	`json:"toIndex"`// End index of the search result to be returned.
}

func (obj *DOM) GetSearchResults(request *GetSearchResultsRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`// Ids of the search result nodes.
}, err error) {
	err = obj.conn.Send("DOM.getSearchResults", request, &response)
	return
}

type DiscardSearchResultsRequest struct {
	SearchId string `json:"searchId"`// Unique search session identifier.
}

func (obj *DOM) DiscardSearchResults(request *DiscardSearchResultsRequest) (err error) {
	err = obj.conn.Send("DOM.discardSearchResults", request, nil)
	return
}

type RequestNodeRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`// JavaScript object id to convert into node.
}

func (obj *DOM) RequestNode(request *RequestNodeRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Node id for given object.
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
	Path string `json:"path"`// Path to node in the proprietary format.
}

func (obj *DOM) PushNodeByPathToFrontend(request *PushNodeByPathToFrontendRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node for given path.
}, err error) {
	err = obj.conn.Send("DOM.pushNodeByPathToFrontend", request, &response)
	return
}

type PushNodesByBackendIdsToFrontendRequest struct {
	BackendNodeIds []types.DOM_BackendNodeId `json:"backendNodeIds"`// The array of backend node ids.
}

func (obj *DOM) PushNodesByBackendIdsToFrontend(request *PushNodesByBackendIdsToFrontendRequest) (response struct {
	NodeIds []types.DOM_NodeId `json:"nodeIds"`// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
}, err error) {
	err = obj.conn.Send("DOM.pushNodesByBackendIdsToFrontend", request, &response)
	return
}

type SetInspectedNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// DOM node id to be accessible by means of $x command line API.
}

func (obj *DOM) SetInspectedNode(request *SetInspectedNodeRequest) (err error) {
	err = obj.conn.Send("DOM.setInspectedNode", request, nil)
	return
}

type ResolveNodeRequest struct {
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Id of the node to resolve.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Backend identifier of the node to resolve.
	ObjectGroup	*string				`json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
}

func (obj *DOM) ResolveNode(request *ResolveNodeRequest) (response struct {
	Object types.Runtime_RemoteObject `json:"object"`// JavaScript object wrapper for given node.
}, err error) {
	err = obj.conn.Send("DOM.resolveNode", request, &response)
	return
}

type GetAttributesRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node to retrieve attibutes for.
}

func (obj *DOM) GetAttributes(request *GetAttributesRequest) (response struct {
	Attributes []string `json:"attributes"`// An interleaved array of node attribute names and values.
}, err error) {
	err = obj.conn.Send("DOM.getAttributes", request, &response)
	return
}

type CopyToRequest struct {
	NodeId			types.DOM_NodeId	`json:"nodeId"`// Id of the node to copy.
	TargetNodeId		types.DOM_NodeId	`json:"targetNodeId"`// Id of the element to drop the copy into.
	InsertBeforeNodeId	*types.DOM_NodeId	`json:"insertBeforeNodeId,omitempty"`// Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>).
}

func (obj *DOM) CopyTo(request *CopyToRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node clone.
}, err error) {
	err = obj.conn.Send("DOM.copyTo", request, &response)
	return
}

type MoveToRequest struct {
	NodeId			types.DOM_NodeId	`json:"nodeId"`// Id of the node to move.
	TargetNodeId		types.DOM_NodeId	`json:"targetNodeId"`// Id of the element to drop the moved node into.
	InsertBeforeNodeId	*types.DOM_NodeId	`json:"insertBeforeNodeId,omitempty"`// Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>).
}

func (obj *DOM) MoveTo(request *MoveToRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// New id of the moved node.
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
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node wrapper.
}

func (obj *DOM) Focus(request *FocusRequest) (err error) {
	err = obj.conn.Send("DOM.focus", request, nil)
	return
}

type SetFileInputFilesRequest struct {
	Files		[]string			`json:"files"`// Array of file paths to set.
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node wrapper.
}

func (obj *DOM) SetFileInputFiles(request *SetFileInputFilesRequest) (err error) {
	err = obj.conn.Send("DOM.setFileInputFiles", request, nil)
	return
}

type GetBoxModelRequest struct {
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node wrapper.
}

func (obj *DOM) GetBoxModel(request *GetBoxModelRequest) (response struct {
	Model types.DOM_BoxModel `json:"model"`// Box model for the node.
}, err error) {
	err = obj.conn.Send("DOM.getBoxModel", request, &response)
	return
}

type GetNodeForLocationRequest struct {
	X				int	`json:"x"`// X coordinate.
	Y				int	`json:"y"`// Y coordinate.
	IncludeUserAgentShadowDOM	*bool	`json:"includeUserAgentShadowDOM,omitempty"`// False to skip to the nearest non-UA shadow root ancestor (default: false).
}

func (obj *DOM) GetNodeForLocation(request *GetNodeForLocationRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node at given coordinates.
}, err error) {
	err = obj.conn.Send("DOM.getNodeForLocation", request, &response)
	return
}

type GetRelayoutBoundaryRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node.
}

func (obj *DOM) GetRelayoutBoundary(request *GetRelayoutBoundaryRequest) (response struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Relayout boundary node id for the given node.
}, err error) {
	err = obj.conn.Send("DOM.getRelayoutBoundary", request, &response)
	return
}

type DescribeNodeRequest struct {
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node wrapper.
	Depth		*int				`json:"depth,omitempty"`// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce		*bool				`json:"pierce,omitempty"`// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}

func (obj *DOM) DescribeNode(request *DescribeNodeRequest) (response struct {
	Node types.DOM_Node `json:"node"`// Node description.
}, err error) {
	err = obj.conn.Send("DOM.describeNode", request, &response)
	return
}
