/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types


//Configuration data for the highlighting of page elements.
type Overlay_HighlightConfig struct {
	// Whether the node info tooltip should be shown (default: false).
	ShowInfo *bool `json:"showInfo,omitempty"`
	// Whether the rulers should be shown (default: false).
	ShowRulers *bool `json:"showRulers,omitempty"`
	// Whether the extension lines from node to the rulers should be shown (default: false).
	ShowExtensionLines *bool `json:"showExtensionLines,omitempty"`
	DisplayAsMaterial *bool `json:"displayAsMaterial,omitempty"`
	// The content box highlight fill color (default: transparent).
	ContentColor *DOM_RGBA `json:"contentColor,omitempty"`
	// The padding highlight fill color (default: transparent).
	PaddingColor *DOM_RGBA `json:"paddingColor,omitempty"`
	// The border highlight fill color (default: transparent).
	BorderColor *DOM_RGBA `json:"borderColor,omitempty"`
	// The margin highlight fill color (default: transparent).
	MarginColor *DOM_RGBA `json:"marginColor,omitempty"`
	// The event target element highlight fill color (default: transparent).
	EventTargetColor *DOM_RGBA `json:"eventTargetColor,omitempty"`
	// The shape outside fill color (default: transparent).
	ShapeColor *DOM_RGBA `json:"shapeColor,omitempty"`
	// The shape margin fill color (default: transparent).
	ShapeMarginColor *DOM_RGBA `json:"shapeMarginColor,omitempty"`
	// Selectors to highlight relevant nodes.
	SelectorList *string `json:"selectorList,omitempty"`
	// The grid layout color (default: transparent).
	CssGridColor *DOM_RGBA `json:"cssGridColor,omitempty"`
}


type Overlay_InspectMode string

