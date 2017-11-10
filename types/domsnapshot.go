package types

type DOMSnapshot_DOMNode struct {
	NodeType              int                     `json:"nodeType"`
	NodeName              string                  `json:"nodeName"`
	NodeValue             string                  `json:"nodeValue"`
	TextValue             *string                 `json:"textValue,omitempty"`
	InputValue            *string                 `json:"inputValue,omitempty"`
	InputChecked          *bool                   `json:"inputChecked,omitempty"`
	OptionSelected        *bool                   `json:"optionSelected,omitempty"`
	BackendNodeId         DOM_BackendNodeId       `json:"backendNodeId"`
	ChildNodeIndexes      []int                   `json:"childNodeIndexes,omitempty"`
	Attributes            []DOMSnapshot_NameValue `json:"attributes,omitempty"`
	PseudoElementIndexes  []int                   `json:"pseudoElementIndexes,omitempty"`
	LayoutNodeIndex       *int                    `json:"layoutNodeIndex,omitempty"`
	DocumentURL           *string                 `json:"documentURL,omitempty"`
	BaseURL               *string                 `json:"baseURL,omitempty"`
	ContentLanguage       *string                 `json:"contentLanguage,omitempty"`
	DocumentEncoding      *string                 `json:"documentEncoding,omitempty"`
	PublicId              *string                 `json:"publicId,omitempty"`
	SystemId              *string                 `json:"systemId,omitempty"`
	FrameId               *Page_FrameId           `json:"frameId,omitempty"`
	ContentDocumentIndex  *int                    `json:"contentDocumentIndex,omitempty"`
	ImportedDocumentIndex *int                    `json:"importedDocumentIndex,omitempty"`
	TemplateContentIndex  *int                    `json:"templateContentIndex,omitempty"`
	PseudoType            *DOM_PseudoType         `json:"pseudoType,omitempty"`
	IsClickable           *bool                   `json:"isClickable,omitempty"`
}
type DOMSnapshot_LayoutTreeNode struct {
	DomNodeIndex    int                 `json:"domNodeIndex"`
	BoundingBox     DOM_Rect            `json:"boundingBox"`
	LayoutText      *string             `json:"layoutText,omitempty"`
	InlineTextNodes []CSS_InlineTextBox `json:"inlineTextNodes,omitempty"`
	StyleIndex      *int                `json:"styleIndex,omitempty"`
}
type DOMSnapshot_ComputedStyle struct {
	Properties []DOMSnapshot_NameValue `json:"properties"`
}
type DOMSnapshot_NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
