/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types


//Unique DOM node identifier.
type DOM_NodeId int

//Unique DOM node identifier used to reference a node that may not have been pushed to the front-end.
type DOM_BackendNodeId int

//Backend node with a friendly name.
type DOM_BackendNode struct {
	// <code>Node</code>'s nodeType.
	NodeType int `json:"nodeType"`
	// <code>Node</code>'s nodeName.
	NodeName string `json:"nodeName"`
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
}

//Pseudo element type.
type DOM_PseudoType string

//Shadow root type.
type DOM_ShadowRootType string

//DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type DOM_Node struct {
	// Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeId DOM_NodeId `json:"nodeId"`
	// The id of the parent node if any.
	// NOTE Experimental
	ParentId *DOM_NodeId `json:"parentId,omitempty"`
	// The BackendNodeId for this node.
	// NOTE Experimental
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
	// <code>Node</code>'s nodeType.
	NodeType int `json:"nodeType"`
	// <code>Node</code>'s nodeName.
	NodeName string `json:"nodeName"`
	// <code>Node</code>'s localName.
	LocalName string `json:"localName"`
	// <code>Node</code>'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Child count for <code>Container</code> nodes.
	ChildNodeCount *int `json:"childNodeCount,omitempty"`
	// Child nodes of this node when requested with children.
	Children []*DOM_Node `json:"children,omitempty"`
	// Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
	Attributes []string `json:"attributes,omitempty"`
	// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	DocumentURL *string `json:"documentURL,omitempty"`
	// Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	// NOTE Experimental
	BaseURL *string `json:"baseURL,omitempty"`
	// <code>DocumentType</code>'s publicId.
	PublicId *string `json:"publicId,omitempty"`
	// <code>DocumentType</code>'s systemId.
	SystemId *string `json:"systemId,omitempty"`
	// <code>DocumentType</code>'s internalSubset.
	InternalSubset *string `json:"internalSubset,omitempty"`
	// <code>Document</code>'s XML version in case of XML documents.
	XmlVersion *string `json:"xmlVersion,omitempty"`
	// <code>Attr</code>'s name.
	Name *string `json:"name,omitempty"`
	// <code>Attr</code>'s value.
	Value *string `json:"value,omitempty"`
	// Pseudo element type for this node.
	PseudoType *DOM_PseudoType `json:"pseudoType,omitempty"`
	// Shadow root type.
	ShadowRootType *DOM_ShadowRootType `json:"shadowRootType,omitempty"`
	// Frame ID for frame owner elements.
	// NOTE Experimental
	FrameId *Page_FrameId `json:"frameId,omitempty"`
	// Content document for frame owner elements.
	ContentDocument *DOM_Node `json:"contentDocument,omitempty"`
	// Shadow root list for given element host.
	// NOTE Experimental
	ShadowRoots []*DOM_Node `json:"shadowRoots,omitempty"`
	// Content document fragment for template elements.
	// NOTE Experimental
	TemplateContent *DOM_Node `json:"templateContent,omitempty"`
	// Pseudo elements associated with this node.
	// NOTE Experimental
	PseudoElements []*DOM_Node `json:"pseudoElements,omitempty"`
	// Import document for the HTMLImport links.
	ImportedDocument *DOM_Node `json:"importedDocument,omitempty"`
	// Distributed nodes for given insertion point.
	// NOTE Experimental
	DistributedNodes []DOM_BackendNode `json:"distributedNodes,omitempty"`
	// Whether the node is SVG.
	// NOTE Experimental
	IsSVG *bool `json:"isSVG,omitempty"`
}

//A structure holding an RGBA color.
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

//An array of quad vertices, x immediately followed by y for each point, points clock-wise.
type DOM_Quad []float32

//Box model.
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

//CSS Shape Outside details.
type DOM_ShapeOutsideInfo struct {
	// Shape bounds
	Bounds DOM_Quad `json:"bounds"`
	// Shape coordinate details
	Shape []interface{} `json:"shape"`
	// Margin shape bounds
	MarginShape []interface{} `json:"marginShape"`
}

//Rectangle.
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

