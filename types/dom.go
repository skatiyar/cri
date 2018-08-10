/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Unique DOM node identifier.
type DOM_NodeId int

// Unique DOM node identifier used to reference a node that may not have been pushed to the front-end.
type DOM_BackendNodeId int

// Backend node with a friendly name.
type DOM_BackendNode struct {
	// `Node`'s nodeType.
	NodeType int `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName      string            `json:"nodeName"`
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
}

// Pseudo element type.
type DOM_PseudoType string

// Shadow root type.
type DOM_ShadowRootType string

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type DOM_Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend will only push node with given `id` once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeId DOM_NodeId `json:"nodeId"`
	// The id of the parent node if any.
	ParentId *DOM_NodeId `json:"parentId,omitempty"`
	// The BackendNodeId for this node.
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
	// `Node`'s nodeType.
	NodeType int `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName string `json:"nodeName"`
	// `Node`'s localName.
	LocalName string `json:"localName"`
	// `Node`'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Child count for `Container` nodes.
	ChildNodeCount *int `json:"childNodeCount,omitempty"`
	// Child nodes of this node when requested with children.
	Children []*DOM_Node `json:"children,omitempty"`
	// Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	Attributes []string `json:"attributes,omitempty"`
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL *string `json:"documentURL,omitempty"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL *string `json:"baseURL,omitempty"`
	// `DocumentType`'s publicId.
	PublicId *string `json:"publicId,omitempty"`
	// `DocumentType`'s systemId.
	SystemId *string `json:"systemId,omitempty"`
	// `DocumentType`'s internalSubset.
	InternalSubset *string `json:"internalSubset,omitempty"`
	// `Document`'s XML version in case of XML documents.
	XmlVersion *string `json:"xmlVersion,omitempty"`
	// `Attr`'s name.
	Name *string `json:"name,omitempty"`
	// `Attr`'s value.
	Value *string `json:"value,omitempty"`
	// Pseudo element type for this node.
	PseudoType *DOM_PseudoType `json:"pseudoType,omitempty"`
	// Shadow root type.
	ShadowRootType *DOM_ShadowRootType `json:"shadowRootType,omitempty"`
	// Frame ID for frame owner elements.
	FrameId *Page_FrameId `json:"frameId,omitempty"`
	// Content document for frame owner elements.
	ContentDocument *DOM_Node `json:"contentDocument,omitempty"`
	// Shadow root list for given element host.
	ShadowRoots []*DOM_Node `json:"shadowRoots,omitempty"`
	// Content document fragment for template elements.
	TemplateContent *DOM_Node `json:"templateContent,omitempty"`
	// Pseudo elements associated with this node.
	PseudoElements []*DOM_Node `json:"pseudoElements,omitempty"`
	// Import document for the HTMLImport links.
	ImportedDocument *DOM_Node `json:"importedDocument,omitempty"`
	// Distributed nodes for given insertion point.
	DistributedNodes []DOM_BackendNode `json:"distributedNodes,omitempty"`
	// Whether the node is SVG.
	IsSVG *bool `json:"isSVG,omitempty"`
}

// A structure holding an RGBA color.
type DOM_RGBA struct {
	// The red component, in the [0-255] range.
	R int `json:"r"`
	// The green component, in the [0-255] range.
	G int `json:"g"`
	// The blue component, in the [0-255] range.
	B int `json:"b"`
	// The alpha component, in the [0-1] range (default: 1).
	A *float32 `json:"a,omitempty"`
}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise.
type DOM_Quad []float32

// Box model.
type DOM_BoxModel struct {
	// Content box
	Content DOM_Quad `json:"content"`
	// Padding box
	Padding DOM_Quad `json:"padding"`
	// Border box
	Border DOM_Quad `json:"border"`
	// Margin box
	Margin DOM_Quad `json:"margin"`
	// Node width
	Width int `json:"width"`
	// Node height
	Height int `json:"height"`
	// Shape outside coordinates
	ShapeOutside *DOM_ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}

// CSS Shape Outside details.
type DOM_ShapeOutsideInfo struct {
	// Shape bounds
	Bounds DOM_Quad `json:"bounds"`
	// Shape coordinate details
	Shape []interface{} `json:"shape"`
	// Margin shape bounds
	MarginShape []interface{} `json:"marginShape"`
}

// Rectangle.
type DOM_Rect struct {
	// X coordinate
	X float32 `json:"x"`
	// Y coordinate
	Y float32 `json:"y"`
	// Rectangle width
	Width float32 `json:"width"`
	// Rectangle height
	Height float32 `json:"height"`
}
