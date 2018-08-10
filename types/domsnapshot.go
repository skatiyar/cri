/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// A Node in the DOM tree.
type DOMSnapshot_DOMNode struct {
	// `Node`'s nodeType.
	NodeType int `json:"nodeType"`
	// `Node`'s nodeName.
	NodeName string `json:"nodeName"`
	// `Node`'s nodeValue.
	NodeValue string `json:"nodeValue"`
	// Only set for textarea elements, contains the text value.
	TextValue *string `json:"textValue,omitempty"`
	// Only set for input elements, contains the input's associated text value.
	InputValue *string `json:"inputValue,omitempty"`
	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked *bool `json:"inputChecked,omitempty"`
	// Only set for option elements, indicates if the element has been selected
	OptionSelected *bool `json:"optionSelected,omitempty"`
	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeId DOM_BackendNodeId `json:"backendNodeId"`
	// The indexes of the node's child nodes in the `domNodes` array returned by `getSnapshot`, if any.
	ChildNodeIndexes []int `json:"childNodeIndexes,omitempty"`
	// Attributes of an `Element` node.
	Attributes []DOMSnapshot_NameValue `json:"attributes,omitempty"`
	// Indexes of pseudo elements associated with this node in the `domNodes` array returned by `getSnapshot`, if any.
	PseudoElementIndexes []int `json:"pseudoElementIndexes,omitempty"`
	// The index of the node's related layout tree node in the `layoutTreeNodes` array returned by `getSnapshot`, if any.
	LayoutNodeIndex *int `json:"layoutNodeIndex,omitempty"`
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL *string `json:"documentURL,omitempty"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL *string `json:"baseURL,omitempty"`
	// Only set for documents, contains the document's content language.
	ContentLanguage *string `json:"contentLanguage,omitempty"`
	// Only set for documents, contains the document's character set encoding.
	DocumentEncoding *string `json:"documentEncoding,omitempty"`
	// `DocumentType` node's publicId.
	PublicId *string `json:"publicId,omitempty"`
	// `DocumentType` node's systemId.
	SystemId *string `json:"systemId,omitempty"`
	// Frame ID for frame owner elements and also for the document node.
	FrameId *Page_FrameId `json:"frameId,omitempty"`
	// The index of a frame owner element's content document in the `domNodes` array returned by `getSnapshot`, if any.
	ContentDocumentIndex *int `json:"contentDocumentIndex,omitempty"`
	// Type of a pseudo element node.
	PseudoType *DOM_PseudoType `json:"pseudoType,omitempty"`
	// Shadow root type.
	ShadowRootType *DOM_ShadowRootType `json:"shadowRootType,omitempty"`
	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	IsClickable *bool `json:"isClickable,omitempty"`
	// Details of the node's event listeners, if any.
	EventListeners []DOMDebugger_EventListener `json:"eventListeners,omitempty"`
	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL *string `json:"currentSourceURL,omitempty"`
	// The url of the script (if any) that generates this node.
	OriginURL *string `json:"originURL,omitempty"`
}

// Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type DOMSnapshot_InlineTextBox struct {
	// The absolute position bounding box.
	BoundingBox DOM_Rect `json:"boundingBox"`
	// The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	StartCharacterIndex int `json:"startCharacterIndex"`
	// The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters int `json:"numCharacters"`
}

// Details of an element in the DOM tree with a LayoutObject.
type DOMSnapshot_LayoutTreeNode struct {
	// The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	DomNodeIndex int `json:"domNodeIndex"`
	// The absolute position bounding box.
	BoundingBox DOM_Rect `json:"boundingBox"`
	// Contents of the LayoutText, if any.
	LayoutText *string `json:"layoutText,omitempty"`
	// The post-layout inline text nodes, if any.
	InlineTextNodes []DOMSnapshot_InlineTextBox `json:"inlineTextNodes,omitempty"`
	// Index into the `computedStyles` array returned by `getSnapshot`.
	StyleIndex *int `json:"styleIndex,omitempty"`
	// Global paint order index, which is determined by the stacking order of the nodes. Nodes that are painted together will have the same index. Only provided if includePaintOrder in getSnapshot was true.
	PaintOrder *int `json:"paintOrder,omitempty"`
}

// A subset of the full ComputedStyle as defined by the request whitelist.
type DOMSnapshot_ComputedStyle struct {
	// Name/value pairs of computed style properties.
	Properties []DOMSnapshot_NameValue `json:"properties"`
}

// A name/value pair.
type DOMSnapshot_NameValue struct {
	// Attribute/property name.
	Name string `json:"name"`
	// Attribute/property value.
	Value string `json:"value"`
}

// Index of the string in the strings table.
type DOMSnapshot_StringIndex int

// Index of the string in the strings table.
type DOMSnapshot_ArrayOfStrings []DOMSnapshot_StringIndex

// Data that is only present on rare nodes.
type DOMSnapshot_RareStringData struct {
	Index []int                     `json:"index"`
	Value []DOMSnapshot_StringIndex `json:"value"`
}

type DOMSnapshot_RareBooleanData struct {
	Index []int `json:"index"`
}

type DOMSnapshot_RareIntegerData struct {
	Index []int `json:"index"`
	Value []int `json:"value"`
}

type DOMSnapshot_Rectangle []float32

// Document snapshot.
type DOMSnapshot_DocumentSnapshot struct {
	// Document URL that `Document` or `FrameOwner` node points to.
	DocumentURL DOMSnapshot_StringIndex `json:"documentURL"`
	// Base URL that `Document` or `FrameOwner` node uses for URL completion.
	BaseURL DOMSnapshot_StringIndex `json:"baseURL"`
	// Contains the document's content language.
	ContentLanguage DOMSnapshot_StringIndex `json:"contentLanguage"`
	// Contains the document's character set encoding.
	EncodingName DOMSnapshot_StringIndex `json:"encodingName"`
	// `DocumentType` node's publicId.
	PublicId DOMSnapshot_StringIndex `json:"publicId"`
	// `DocumentType` node's systemId.
	SystemId DOMSnapshot_StringIndex `json:"systemId"`
	// Frame ID for frame owner elements and also for the document node.
	FrameId DOMSnapshot_StringIndex `json:"frameId"`
	// A table with dom nodes.
	Nodes DOMSnapshot_NodeTreeSnapshot `json:"nodes"`
	// The nodes in the layout tree.
	Layout DOMSnapshot_LayoutTreeSnapshot `json:"layout"`
	// The post-layout inline text nodes.
	TextBoxes DOMSnapshot_TextBoxSnapshot `json:"textBoxes"`
}

// Table containing nodes.
type DOMSnapshot_NodeTreeSnapshot struct {
	// Parent node index.
	ParentIndex []int `json:"parentIndex,omitempty"`
	// `Node`'s nodeType.
	NodeType []int `json:"nodeType,omitempty"`
	// `Node`'s nodeName.
	NodeName []DOMSnapshot_StringIndex `json:"nodeName,omitempty"`
	// `Node`'s nodeValue.
	NodeValue []DOMSnapshot_StringIndex `json:"nodeValue,omitempty"`
	// `Node`'s id, corresponds to DOM.Node.backendNodeId.
	BackendNodeId []DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// Attributes of an `Element` node. Flatten name, value pairs.
	Attributes []DOMSnapshot_ArrayOfStrings `json:"attributes,omitempty"`
	// Only set for textarea elements, contains the text value.
	TextValue *DOMSnapshot_RareStringData `json:"textValue,omitempty"`
	// Only set for input elements, contains the input's associated text value.
	InputValue *DOMSnapshot_RareStringData `json:"inputValue,omitempty"`
	// Only set for radio and checkbox input elements, indicates if the element has been checked
	InputChecked *DOMSnapshot_RareBooleanData `json:"inputChecked,omitempty"`
	// Only set for option elements, indicates if the element has been selected
	OptionSelected *DOMSnapshot_RareBooleanData `json:"optionSelected,omitempty"`
	// The index of the document in the list of the snapshot documents.
	ContentDocumentIndex *DOMSnapshot_RareIntegerData `json:"contentDocumentIndex,omitempty"`
	// Type of a pseudo element node.
	PseudoType *DOMSnapshot_RareStringData `json:"pseudoType,omitempty"`
	// Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	IsClickable *DOMSnapshot_RareBooleanData `json:"isClickable,omitempty"`
	// The selected url for nodes with a srcset attribute.
	CurrentSourceURL *DOMSnapshot_RareStringData `json:"currentSourceURL,omitempty"`
	// The url of the script (if any) that generates this node.
	OriginURL *DOMSnapshot_RareStringData `json:"originURL,omitempty"`
}

// Details of an element in the DOM tree with a LayoutObject.
type DOMSnapshot_LayoutTreeSnapshot struct {
	// The index of the related DOM node in the `domNodes` array returned by `getSnapshot`.
	NodeIndex []int `json:"nodeIndex"`
	// Index into the `computedStyles` array returned by `captureSnapshot`.
	Styles []DOMSnapshot_ArrayOfStrings `json:"styles"`
	// The absolute position bounding box.
	Bounds []DOMSnapshot_Rectangle `json:"bounds"`
	// Contents of the LayoutText, if any.
	Text []DOMSnapshot_StringIndex `json:"text"`
}

// Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type DOMSnapshot_TextBoxSnapshot struct {
	// Intex of th elayout tree node that owns this box collection.
	LayoutIndex []int `json:"layoutIndex"`
	// The absolute position bounding box.
	Bounds []DOMSnapshot_Rectangle `json:"bounds"`
	// The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	Start []int `json:"start"`
	// The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	Length []int `json:"length"`
}
