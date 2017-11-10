package types

type Accessibility_AXNodeId string
type Accessibility_AXValueType string
type Accessibility_AXValueSourceType string
type Accessibility_AXValueNativeSourceType string
type Accessibility_AXValueSource struct {
	Type			Accessibility_AXValueSourceType		`json:"type"`// What type of source this is.
	Value			*Accessibility_AXValue			`json:"value,omitempty"`// The value of this property source.
	Attribute		*string					`json:"attribute,omitempty"`// The name of the relevant attribute, if any.
	AttributeValue		*Accessibility_AXValue			`json:"attributeValue,omitempty"`// The value of the relevant attribute, if any.
	Superseded		*bool					`json:"superseded,omitempty"`// Whether this source is superseded by a higher priority source.
	NativeSource		*Accessibility_AXValueNativeSourceType	`json:"nativeSource,omitempty"`// The native markup source for this value, e.g. a <label> element.
	NativeSourceValue	*Accessibility_AXValue			`json:"nativeSourceValue,omitempty"`// The value, such as a node or node list, of the native source.
	Invalid			*bool					`json:"invalid,omitempty"`// Whether the value for this property is invalid.
	InvalidReason		*string					`json:"invalidReason,omitempty"`// Reason for the value being invalid, if it is.
}
type Accessibility_AXRelatedNode struct {
	BackendDOMNodeId	DOM_BackendNodeId	`json:"backendDOMNodeId"`// The BackendNodeId of the related DOM node.
	Idref			*string			`json:"idref,omitempty"`// The IDRef value provided, if any.
	Text			*string			`json:"text,omitempty"`// The text alternative of this node in the current context.
}
type Accessibility_AXProperty struct {
	Name	string			`json:"name"`// The name of this property.
	Value	Accessibility_AXValue	`json:"value"`// The value of this property.
}
type Accessibility_AXValue struct {
	Type		Accessibility_AXValueType	`json:"type"`// The type of this value.
	Value		interface{}			`json:"value,omitempty"`// The computed value of this property.
	RelatedNodes	[]Accessibility_AXRelatedNode	`json:"relatedNodes,omitempty"`// One or more related nodes, if applicable.
	Sources		[]Accessibility_AXValueSource	`json:"sources,omitempty"`// The sources which contributed to the computation of this property.
}
type Accessibility_AXGlobalStates string
type Accessibility_AXLiveRegionAttributes string
type Accessibility_AXWidgetAttributes string
type Accessibility_AXWidgetStates string
type Accessibility_AXRelationshipAttributes string
type Accessibility_AXNode struct {
	NodeId			Accessibility_AXNodeId		`json:"nodeId"`// Unique identifier for this node.
	Ignored			bool				`json:"ignored"`// Whether this node is ignored for accessibility
	IgnoredReasons		[]Accessibility_AXProperty	`json:"ignoredReasons,omitempty"`// Collection of reasons why this node is hidden.
	Role			*Accessibility_AXValue		`json:"role,omitempty"`// This <code>Node</code>'s role, whether explicit or implicit.
	Name			*Accessibility_AXValue		`json:"name,omitempty"`// The accessible name for this <code>Node</code>.
	Description		*Accessibility_AXValue		`json:"description,omitempty"`// The accessible description for this <code>Node</code>.
	Value			*Accessibility_AXValue		`json:"value,omitempty"`// The value for this <code>Node</code>.
	Properties		[]Accessibility_AXProperty	`json:"properties,omitempty"`// All other properties
	ChildIds		[]Accessibility_AXNodeId	`json:"childIds,omitempty"`// IDs for each of this node's child nodes.
	BackendDOMNodeId	*DOM_BackendNodeId		`json:"backendDOMNodeId,omitempty"`// The backend ID for the associated DOM node, if any.
}
