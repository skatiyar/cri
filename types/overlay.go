package types

type Overlay_HighlightConfig struct {
	ShowInfo           *bool     `json:"showInfo,omitempty"`
	ShowRulers         *bool     `json:"showRulers,omitempty"`
	ShowExtensionLines *bool     `json:"showExtensionLines,omitempty"`
	DisplayAsMaterial  *bool     `json:"displayAsMaterial,omitempty"`
	ContentColor       *DOM_RGBA `json:"contentColor,omitempty"`
	PaddingColor       *DOM_RGBA `json:"paddingColor,omitempty"`
	BorderColor        *DOM_RGBA `json:"borderColor,omitempty"`
	MarginColor        *DOM_RGBA `json:"marginColor,omitempty"`
	EventTargetColor   *DOM_RGBA `json:"eventTargetColor,omitempty"`
	ShapeColor         *DOM_RGBA `json:"shapeColor,omitempty"`
	ShapeMarginColor   *DOM_RGBA `json:"shapeMarginColor,omitempty"`
	SelectorList       *string   `json:"selectorList,omitempty"`
	CssGridColor       *DOM_RGBA `json:"cssGridColor,omitempty"`
}
type Overlay_InspectMode string
