package types

type DOM_NodeId int
type DOM_BackendNodeId int
type DOM_BackendNode struct {
	NodeType	int			`json:"nodeType"`// <code>Node</code>'s nodeType.
	NodeName	string			`json:"nodeName"`// <code>Node</code>'s nodeName.
	BackendNodeId	DOM_BackendNodeId	`json:"backendNodeId"`
}
type DOM_PseudoType string
type DOM_ShadowRootType string
type DOM_Node struct {
	NodeId			DOM_NodeId		`json:"nodeId"`// Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentId		*DOM_NodeId		`json:"parentId,omitempty"`// The id of the parent node if any.
	BackendNodeId		DOM_BackendNodeId	`json:"backendNodeId"`// The BackendNodeId for this node.
	NodeType		int			`json:"nodeType"`// <code>Node</code>'s nodeType.
	NodeName		string			`json:"nodeName"`// <code>Node</code>'s nodeName.
	LocalName		string			`json:"localName"`// <code>Node</code>'s localName.
	NodeValue		string			`json:"nodeValue"`// <code>Node</code>'s nodeValue.
	ChildNodeCount		*int			`json:"childNodeCount,omitempty"`// Child count for <code>Container</code> nodes.
	Children		[]*DOM_Node		`json:"children,omitempty"`// Child nodes of this node when requested with children.
	Attributes		[]string		`json:"attributes,omitempty"`// Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
	DocumentURL		*string			`json:"documentURL,omitempty"`// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL			*string			`json:"baseURL,omitempty"`// Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	PublicId		*string			`json:"publicId,omitempty"`// <code>DocumentType</code>'s publicId.
	SystemId		*string			`json:"systemId,omitempty"`// <code>DocumentType</code>'s systemId.
	InternalSubset		*string			`json:"internalSubset,omitempty"`// <code>DocumentType</code>'s internalSubset.
	XmlVersion		*string			`json:"xmlVersion,omitempty"`// <code>Document</code>'s XML version in case of XML documents.
	Name			*string			`json:"name,omitempty"`// <code>Attr</code>'s name.
	Value			*string			`json:"value,omitempty"`// <code>Attr</code>'s value.
	PseudoType		*DOM_PseudoType		`json:"pseudoType,omitempty"`// Pseudo element type for this node.
	ShadowRootType		*DOM_ShadowRootType	`json:"shadowRootType,omitempty"`// Shadow root type.
	FrameId			*Page_FrameId		`json:"frameId,omitempty"`// Frame ID for frame owner elements.
	ContentDocument		*DOM_Node		`json:"contentDocument,omitempty"`// Content document for frame owner elements.
	ShadowRoots		[]*DOM_Node		`json:"shadowRoots,omitempty"`// Shadow root list for given element host.
	TemplateContent		*DOM_Node		`json:"templateContent,omitempty"`// Content document fragment for template elements.
	PseudoElements		[]*DOM_Node		`json:"pseudoElements,omitempty"`// Pseudo elements associated with this node.
	ImportedDocument	*DOM_Node		`json:"importedDocument,omitempty"`// Import document for the HTMLImport links.
	DistributedNodes	[]DOM_BackendNode	`json:"distributedNodes,omitempty"`// Distributed nodes for given insertion point.
	IsSVG			*bool			`json:"isSVG,omitempty"`// Whether the node is SVG.
}
type DOM_RGBA struct {
	R	int		`json:"r"`// The red component, in the [0-255] range.
	G	int		`json:"g"`// The green component, in the [0-255] range.
	B	int		`json:"b"`// The blue component, in the [0-255] range.
	A	*float32	`json:"a,omitempty"`// The alpha component, in the [0-1] range (default: 1).
}
type DOM_Quad []float32
type DOM_BoxModel struct {
	Content		DOM_Quad		`json:"content"`// Content box
	Padding		DOM_Quad		`json:"padding"`// Padding box
	Border		DOM_Quad		`json:"border"`// Border box
	Margin		DOM_Quad		`json:"margin"`// Margin box
	Width		int			`json:"width"`// Node width
	Height		int			`json:"height"`// Node height
	ShapeOutside	*DOM_ShapeOutsideInfo	`json:"shapeOutside,omitempty"`// Shape outside coordinates
}
type DOM_ShapeOutsideInfo struct {
	Bounds		DOM_Quad	`json:"bounds"`// Shape bounds
	Shape		[]interface{}	`json:"shape"`// Shape coordinate details
	MarginShape	[]interface{}	`json:"marginShape"`// Margin shape bounds
}
type DOM_Rect struct {
	X	float32	`json:"x"`// X coordinate
	Y	float32	`json:"y"`// Y coordinate
	Width	float32	`json:"width"`// Rectangle width
	Height	float32	`json:"height"`// Rectangle height
}
