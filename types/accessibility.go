/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Unique accessibility node identifier.
type Accessibility_AXNodeId string

//Enum of possible property types.
type Accessibility_AXValueType string

//Enum of possible property sources.
type Accessibility_AXValueSourceType string

//Enum of possible native property sources (as a subtype of a particular AXValueSourceType).
type Accessibility_AXValueNativeSourceType string

//A single source for a computed AX property.
type Accessibility_AXValueSource struct {
	// What type of source this is.
	Type Accessibility_AXValueSourceType `json:"type"`
	// The value of this property source.
	Value *Accessibility_AXValue `json:"value,omitempty"`
	// The name of the relevant attribute, if any.
	Attribute *string `json:"attribute,omitempty"`
	// The value of the relevant attribute, if any.
	AttributeValue *Accessibility_AXValue `json:"attributeValue,omitempty"`
	// Whether this source is superseded by a higher priority source.
	Superseded *bool `json:"superseded,omitempty"`
	// The native markup source for this value, e.g. a <label> element.
	NativeSource *Accessibility_AXValueNativeSourceType `json:"nativeSource,omitempty"`
	// The value, such as a node or node list, of the native source.
	NativeSourceValue *Accessibility_AXValue `json:"nativeSourceValue,omitempty"`
	// Whether the value for this property is invalid.
	Invalid *bool `json:"invalid,omitempty"`
	// Reason for the value being invalid, if it is.
	InvalidReason *string `json:"invalidReason,omitempty"`
}

type Accessibility_AXRelatedNode struct {
	// The BackendNodeId of the related DOM node.
	BackendDOMNodeId DOM_BackendNodeId `json:"backendDOMNodeId"`
	// The IDRef value provided, if any.
	Idref *string `json:"idref,omitempty"`
	// The text alternative of this node in the current context.
	Text *string `json:"text,omitempty"`
}

type Accessibility_AXProperty struct {
	// The name of this property.
	Name string `json:"name"`
	// The value of this property.
	Value Accessibility_AXValue `json:"value"`
}

//A single computed AX property.
type Accessibility_AXValue struct {
	// The type of this value.
	Type Accessibility_AXValueType `json:"type"`
	// The computed value of this property.
	Value interface{} `json:"value,omitempty"`
	// One or more related nodes, if applicable.
	RelatedNodes []Accessibility_AXRelatedNode `json:"relatedNodes,omitempty"`
	// The sources which contributed to the computation of this property.
	Sources []Accessibility_AXValueSource `json:"sources,omitempty"`
}

//States which apply to every AX node.
type Accessibility_AXGlobalStates string

//Attributes which apply to nodes in live regions.
type Accessibility_AXLiveRegionAttributes string

//Attributes which apply to widgets.
type Accessibility_AXWidgetAttributes string

//States which apply to widgets.
type Accessibility_AXWidgetStates string

//Relationships between elements other than parent/child/sibling.
type Accessibility_AXRelationshipAttributes string

//A node in the accessibility tree.
type Accessibility_AXNode struct {
	// Unique identifier for this node.
	NodeId Accessibility_AXNodeId `json:"nodeId"`
	// Whether this node is ignored for accessibility
	Ignored bool `json:"ignored"`
	// Collection of reasons why this node is hidden.
	IgnoredReasons []Accessibility_AXProperty `json:"ignoredReasons,omitempty"`
	// This <code>Node</code>'s role, whether explicit or implicit.
	Role *Accessibility_AXValue `json:"role,omitempty"`
	// The accessible name for this <code>Node</code>.
	Name *Accessibility_AXValue `json:"name,omitempty"`
	// The accessible description for this <code>Node</code>.
	Description *Accessibility_AXValue `json:"description,omitempty"`
	// The value for this <code>Node</code>.
	Value *Accessibility_AXValue `json:"value,omitempty"`
	// All other properties
	Properties []Accessibility_AXProperty `json:"properties,omitempty"`
	// IDs for each of this node's child nodes.
	ChildIds []Accessibility_AXNodeId `json:"childIds,omitempty"`
	// The backend ID for the associated DOM node, if any.
	BackendDOMNodeId *DOM_BackendNodeId `json:"backendDOMNodeId,omitempty"`
}
