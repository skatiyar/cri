package types

type DOMSnapshot_DOMNode struct {
	// <code>Node</code>'s nodeType.
	NodeType int `json:"nodeType"`
	// <code>Node</code>'s nodeName.
	NodeName string `json:"nodeName"`
	// <code>Node</code>'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Only set for textarea elements, contains the text value.
	TextValue *string `json:"textValue,omitempty"`
	// Only set for input elements, contains the input's associated text value.
	InputValue *string `json:"inputValue,omitempty"`
	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked *bool `json:"inputChecked,omitempty"`
	// Only set for option elements, indicates if the element has been selected
	OptionSelected *bool `json:"optionSelected,omitempty"`
	// <code>Node</code>'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
	// The indexes of the node's child nodes in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	ChildNodeIndexes []int `json:"childNodeIndexes,omitempty"`
	// Attributes of an <code>Element</code> node.
	Attributes []DOMSnapshot_NameValue `json:"attributes,omitempty"`
	// Indexes of pseudo elements associated with this node in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	PseudoElementIndexes []int `json:"pseudoElementIndexes,omitempty"`
	// The index of the node's related layout tree node in the <code>layoutTreeNodes</code> array returned by <code>getSnapshot</code>, if any.
	LayoutNodeIndex *int `json:"layoutNodeIndex,omitempty"`
	// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	DocumentURL *string `json:"documentURL,omitempty"`
	// Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	BaseURL *string `json:"baseURL,omitempty"`
	// Only set for documents, contains the document's content language.
	ContentLanguage *string `json:"contentLanguage,omitempty"`
	// Only set for documents, contains the document's character set encoding.
	DocumentEncoding *string `json:"documentEncoding,omitempty"`
	// <code>DocumentType</code> node's publicId.
	PublicId *string `json:"publicId,omitempty"`
	// <code>DocumentType</code> node's systemId.
	SystemId *string `json:"systemId,omitempty"`
	// Frame ID for frame owner elements and also for the document node.
	FrameId *Page_FrameId `json:"frameId,omitempty"`
	// The index of a frame owner element's content document in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	ContentDocumentIndex *int `json:"contentDocumentIndex,omitempty"`
	// Index of the imported document's node of a link element in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	ImportedDocumentIndex *int `json:"importedDocumentIndex,omitempty"`
	// Index of the content node of a template element in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	TemplateContentIndex *int `json:"templateContentIndex,omitempty"`
	// Type of a pseudo element node.
	PseudoType *DOM_PseudoType `json:"pseudoType,omitempty"`
	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	IsClickable *bool `json:"isClickable,omitempty"`
}
type DOMSnapshot_LayoutTreeNode struct {
	// The index of the related DOM node in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	DomNodeIndex int `json:"domNodeIndex"`
	// The absolute position bounding box.
	BoundingBox DOM_Rect `json:"boundingBox"`
	// Contents of the LayoutText, if any.
	LayoutText *string `json:"layoutText,omitempty"`
	// The post-layout inline text nodes, if any.
	InlineTextNodes []CSS_InlineTextBox `json:"inlineTextNodes,omitempty"`
	// Index into the <code>computedStyles</code> array returned by <code>getSnapshot</code>.
	StyleIndex *int `json:"styleIndex,omitempty"`
}
type DOMSnapshot_ComputedStyle struct {
	// Name/value pairs of computed style properties.
	Properties []DOMSnapshot_NameValue `json:"properties"`
}
type DOMSnapshot_NameValue struct {
	// Attribute/property name.
	Name string `json:"name"`
	// Attribute/property value.
	Value string `json:"value"`
}
