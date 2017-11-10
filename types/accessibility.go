package types

type Accessibility_AXNodeId string
type Accessibility_AXValueType string
type Accessibility_AXValueSourceType string
type Accessibility_AXValueNativeSourceType string
type Accessibility_AXValueSource struct {
	Type              Accessibility_AXValueSourceType        `json:"type"`
	Value             *Accessibility_AXValue                 `json:"value,omitempty"`
	Attribute         *string                                `json:"attribute,omitempty"`
	AttributeValue    *Accessibility_AXValue                 `json:"attributeValue,omitempty"`
	Superseded        *bool                                  `json:"superseded,omitempty"`
	NativeSource      *Accessibility_AXValueNativeSourceType `json:"nativeSource,omitempty"`
	NativeSourceValue *Accessibility_AXValue                 `json:"nativeSourceValue,omitempty"`
	Invalid           *bool                                  `json:"invalid,omitempty"`
	InvalidReason     *string                                `json:"invalidReason,omitempty"`
}
type Accessibility_AXRelatedNode struct {
	BackendDOMNodeId DOM_BackendNodeId `json:"backendDOMNodeId"`
	Idref            *string           `json:"idref,omitempty"`
	Text             *string           `json:"text,omitempty"`
}
type Accessibility_AXProperty struct {
	Name  string                `json:"name"`
	Value Accessibility_AXValue `json:"value"`
}
type Accessibility_AXValue struct {
	Type         Accessibility_AXValueType     `json:"type"`
	Value        interface{}                   `json:"value,omitempty"`
	RelatedNodes []Accessibility_AXRelatedNode `json:"relatedNodes,omitempty"`
	Sources      []Accessibility_AXValueSource `json:"sources,omitempty"`
}
type Accessibility_AXGlobalStates string
type Accessibility_AXLiveRegionAttributes string
type Accessibility_AXWidgetAttributes string
type Accessibility_AXWidgetStates string
type Accessibility_AXRelationshipAttributes string
type Accessibility_AXNode struct {
	NodeId           Accessibility_AXNodeId     `json:"nodeId"`
	Ignored          bool                       `json:"ignored"`
	IgnoredReasons   []Accessibility_AXProperty `json:"ignoredReasons,omitempty"`
	Role             *Accessibility_AXValue     `json:"role,omitempty"`
	Name             *Accessibility_AXValue     `json:"name,omitempty"`
	Description      *Accessibility_AXValue     `json:"description,omitempty"`
	Value            *Accessibility_AXValue     `json:"value,omitempty"`
	Properties       []Accessibility_AXProperty `json:"properties,omitempty"`
	ChildIds         []Accessibility_AXNodeId   `json:"childIds,omitempty"`
	BackendDOMNodeId *DOM_BackendNodeId         `json:"backendDOMNodeId,omitempty"`
}
