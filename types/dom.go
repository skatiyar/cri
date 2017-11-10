package types

type DOM_NodeId int
type DOM_BackendNodeId int
type DOM_BackendNode struct {
	NodeType      int               `json:"nodeType"`
	NodeName      string            `json:"nodeName"`
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
}
type DOM_PseudoType string
type DOM_ShadowRootType string
type DOM_Node struct {
	NodeId           DOM_NodeId          `json:"nodeId"`
	ParentId         *DOM_NodeId         `json:"parentId,omitempty"`
	BackendNodeId    DOM_BackendNodeId   `json:"backendNodeId"`
	NodeType         int                 `json:"nodeType"`
	NodeName         string              `json:"nodeName"`
	LocalName        string              `json:"localName"`
	NodeValue        string              `json:"nodeValue"`
	ChildNodeCount   *int                `json:"childNodeCount,omitempty"`
	Children         []*DOM_Node         `json:"children,omitempty"`
	Attributes       []string            `json:"attributes,omitempty"`
	DocumentURL      *string             `json:"documentURL,omitempty"`
	BaseURL          *string             `json:"baseURL,omitempty"`
	PublicId         *string             `json:"publicId,omitempty"`
	SystemId         *string             `json:"systemId,omitempty"`
	InternalSubset   *string             `json:"internalSubset,omitempty"`
	XmlVersion       *string             `json:"xmlVersion,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Value            *string             `json:"value,omitempty"`
	PseudoType       *DOM_PseudoType     `json:"pseudoType,omitempty"`
	ShadowRootType   *DOM_ShadowRootType `json:"shadowRootType,omitempty"`
	FrameId          *Page_FrameId       `json:"frameId,omitempty"`
	ContentDocument  *DOM_Node           `json:"contentDocument,omitempty"`
	ShadowRoots      []*DOM_Node         `json:"shadowRoots,omitempty"`
	TemplateContent  *DOM_Node           `json:"templateContent,omitempty"`
	PseudoElements   []*DOM_Node         `json:"pseudoElements,omitempty"`
	ImportedDocument *DOM_Node           `json:"importedDocument,omitempty"`
	DistributedNodes []DOM_BackendNode   `json:"distributedNodes,omitempty"`
	IsSVG            *bool               `json:"isSVG,omitempty"`
}
type DOM_RGBA struct {
	R int      `json:"r"`
	G int      `json:"g"`
	B int      `json:"b"`
	A *float32 `json:"a,omitempty"`
}
type DOM_Quad []float32
type DOM_BoxModel struct {
	Content      DOM_Quad              `json:"content"`
	Padding      DOM_Quad              `json:"padding"`
	Border       DOM_Quad              `json:"border"`
	Margin       DOM_Quad              `json:"margin"`
	Width        int                   `json:"width"`
	Height       int                   `json:"height"`
	ShapeOutside *DOM_ShapeOutsideInfo `json:"shapeOutside,omitempty"`
}
type DOM_ShapeOutsideInfo struct {
	Bounds      DOM_Quad      `json:"bounds"`
	Shape       []interface{} `json:"shape"`
	MarginShape []interface{} `json:"marginShape"`
}
type DOM_Rect struct {
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}
