package types

type DOMSnapshot_DOMNode struct {
	NodeType		int			`json:"nodeType"`// <code>Node</code>'s nodeType.
	NodeName		string			`json:"nodeName"`// <code>Node</code>'s nodeName.
	NodeValue		string			`json:"nodeValue"`// <code>Node</code>'s nodeValue.
	TextValue		*string			`json:"textValue,omitempty"`// Only set for textarea elements, contains the text value.
	InputValue		*string			`json:"inputValue,omitempty"`// Only set for input elements, contains the input's associated text value.
	InputChecked		*bool			`json:"inputChecked,omitempty"`// Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected		*bool			`json:"optionSelected,omitempty"`// Only set for option elements, indicates if the element has been selected
	BackendNodeId		DOM_BackendNodeId	`json:"backendNodeId"`// <code>Node</code>'s id, corresponds to DOM.Node.backendNodeId.
	ChildNodeIndexes	[]int			`json:"childNodeIndexes,omitempty"`// The indexes of the node's child nodes in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	Attributes		[]DOMSnapshot_NameValue	`json:"attributes,omitempty"`// Attributes of an <code>Element</code> node.
	PseudoElementIndexes	[]int			`json:"pseudoElementIndexes,omitempty"`// Indexes of pseudo elements associated with this node in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	LayoutNodeIndex		*int			`json:"layoutNodeIndex,omitempty"`// The index of the node's related layout tree node in the <code>layoutTreeNodes</code> array returned by <code>getSnapshot</code>, if any.
	DocumentURL		*string			`json:"documentURL,omitempty"`// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL			*string			`json:"baseURL,omitempty"`// Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	ContentLanguage		*string			`json:"contentLanguage,omitempty"`// Only set for documents, contains the document's content language.
	DocumentEncoding	*string			`json:"documentEncoding,omitempty"`// Only set for documents, contains the document's character set encoding.
	PublicId		*string			`json:"publicId,omitempty"`// <code>DocumentType</code> node's publicId.
	SystemId		*string			`json:"systemId,omitempty"`// <code>DocumentType</code> node's systemId.
	FrameId			*Page_FrameId		`json:"frameId,omitempty"`// Frame ID for frame owner elements and also for the document node.
	ContentDocumentIndex	*int			`json:"contentDocumentIndex,omitempty"`// The index of a frame owner element's content document in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	ImportedDocumentIndex	*int			`json:"importedDocumentIndex,omitempty"`// Index of the imported document's node of a link element in the <code>domNodes</code> array returned by <code>getSnapshot</code>, if any.
	TemplateContentIndex	*int			`json:"templateContentIndex,omitempty"`// Index of the content node of a template element in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	PseudoType		*DOM_PseudoType		`json:"pseudoType,omitempty"`// Type of a pseudo element node.
	IsClickable		*bool			`json:"isClickable,omitempty"`// Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
}
type DOMSnapshot_LayoutTreeNode struct {
	DomNodeIndex	int			`json:"domNodeIndex"`// The index of the related DOM node in the <code>domNodes</code> array returned by <code>getSnapshot</code>.
	BoundingBox	DOM_Rect		`json:"boundingBox"`// The absolute position bounding box.
	LayoutText	*string			`json:"layoutText,omitempty"`// Contents of the LayoutText, if any.
	InlineTextNodes	[]CSS_InlineTextBox	`json:"inlineTextNodes,omitempty"`// The post-layout inline text nodes, if any.
	StyleIndex	*int			`json:"styleIndex,omitempty"`// Index into the <code>computedStyles</code> array returned by <code>getSnapshot</code>.
}
type DOMSnapshot_ComputedStyle struct {
	Properties []DOMSnapshot_NameValue `json:"properties"`// Name/value pairs of computed style properties.
}
type DOMSnapshot_NameValue struct {
	Name	string	`json:"name"`// Attribute/property name.
	Value	string	`json:"value"`// Attribute/property value.
}
