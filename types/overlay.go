package types

type Overlay_HighlightConfig struct {
	ShowInfo		*bool		`json:"showInfo,omitempty"`// Whether the node info tooltip should be shown (default: false).
	ShowRulers		*bool		`json:"showRulers,omitempty"`// Whether the rulers should be shown (default: false).
	ShowExtensionLines	*bool		`json:"showExtensionLines,omitempty"`// Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial	*bool		`json:"displayAsMaterial,omitempty"`
	ContentColor		*DOM_RGBA	`json:"contentColor,omitempty"`// The content box highlight fill color (default: transparent).
	PaddingColor		*DOM_RGBA	`json:"paddingColor,omitempty"`// The padding highlight fill color (default: transparent).
	BorderColor		*DOM_RGBA	`json:"borderColor,omitempty"`// The border highlight fill color (default: transparent).
	MarginColor		*DOM_RGBA	`json:"marginColor,omitempty"`// The margin highlight fill color (default: transparent).
	EventTargetColor	*DOM_RGBA	`json:"eventTargetColor,omitempty"`// The event target element highlight fill color (default: transparent).
	ShapeColor		*DOM_RGBA	`json:"shapeColor,omitempty"`// The shape outside fill color (default: transparent).
	ShapeMarginColor	*DOM_RGBA	`json:"shapeMarginColor,omitempty"`// The shape margin fill color (default: transparent).
	SelectorList		*string		`json:"selectorList,omitempty"`// Selectors to highlight relevant nodes.
	CssGridColor		*DOM_RGBA	`json:"cssGridColor,omitempty"`// The grid layout color (default: transparent).
}
type Overlay_InspectMode string
