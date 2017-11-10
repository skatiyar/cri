package types

type CSS_StyleSheetId string
type CSS_StyleSheetOrigin string
type CSS_PseudoElementMatches struct {
	PseudoType DOM_PseudoType  `json:"pseudoType"`
	Matches    []CSS_RuleMatch `json:"matches"`
}
type CSS_InheritedStyleEntry struct {
	InlineStyle     *CSS_CSSStyle   `json:"inlineStyle,omitempty"`
	MatchedCSSRules []CSS_RuleMatch `json:"matchedCSSRules"`
}
type CSS_RuleMatch struct {
	Rule              CSS_CSSRule `json:"rule"`
	MatchingSelectors []int       `json:"matchingSelectors"`
}
type CSS_Value struct {
	Text  string           `json:"text"`
	Range *CSS_SourceRange `json:"range,omitempty"`
}
type CSS_SelectorList struct {
	Selectors []CSS_Value `json:"selectors"`
	Text      string      `json:"text"`
}
type CSS_CSSStyleSheetHeader struct {
	StyleSheetId CSS_StyleSheetId     `json:"styleSheetId"`
	FrameId      Page_FrameId         `json:"frameId"`
	SourceURL    string               `json:"sourceURL"`
	SourceMapURL *string              `json:"sourceMapURL,omitempty"`
	Origin       CSS_StyleSheetOrigin `json:"origin"`
	Title        string               `json:"title"`
	OwnerNode    *DOM_BackendNodeId   `json:"ownerNode,omitempty"`
	Disabled     bool                 `json:"disabled"`
	HasSourceURL *bool                `json:"hasSourceURL,omitempty"`
	IsInline     bool                 `json:"isInline"`
	StartLine    float32              `json:"startLine"`
	StartColumn  float32              `json:"startColumn"`
	Length       float32              `json:"length"`
}
type CSS_CSSRule struct {
	StyleSheetId *CSS_StyleSheetId    `json:"styleSheetId,omitempty"`
	SelectorList CSS_SelectorList     `json:"selectorList"`
	Origin       CSS_StyleSheetOrigin `json:"origin"`
	Style        CSS_CSSStyle         `json:"style"`
	Media        []CSS_CSSMedia       `json:"media,omitempty"`
}
type CSS_RuleUsage struct {
	StyleSheetId CSS_StyleSheetId `json:"styleSheetId"`
	StartOffset  float32          `json:"startOffset"`
	EndOffset    float32          `json:"endOffset"`
	Used         bool             `json:"used"`
}
type CSS_SourceRange struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn"`
	EndLine     int `json:"endLine"`
	EndColumn   int `json:"endColumn"`
}
type CSS_ShorthandEntry struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Important *bool  `json:"important,omitempty"`
}
type CSS_CSSComputedStyleProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type CSS_CSSStyle struct {
	StyleSheetId     *CSS_StyleSheetId    `json:"styleSheetId,omitempty"`
	CssProperties    []CSS_CSSProperty    `json:"cssProperties"`
	ShorthandEntries []CSS_ShorthandEntry `json:"shorthandEntries"`
	CssText          *string              `json:"cssText,omitempty"`
	Range            *CSS_SourceRange     `json:"range,omitempty"`
}
type CSS_CSSProperty struct {
	Name      string           `json:"name"`
	Value     string           `json:"value"`
	Important *bool            `json:"important,omitempty"`
	Implicit  *bool            `json:"implicit,omitempty"`
	Text      *string          `json:"text,omitempty"`
	ParsedOk  *bool            `json:"parsedOk,omitempty"`
	Disabled  *bool            `json:"disabled,omitempty"`
	Range     *CSS_SourceRange `json:"range,omitempty"`
}
type CSS_CSSMedia struct {
	Text         string            `json:"text"`
	Source       string            `json:"source"`
	SourceURL    *string           `json:"sourceURL,omitempty"`
	Range        *CSS_SourceRange  `json:"range,omitempty"`
	StyleSheetId *CSS_StyleSheetId `json:"styleSheetId,omitempty"`
	MediaList    []CSS_MediaQuery  `json:"mediaList,omitempty"`
}
type CSS_MediaQuery struct {
	Expressions []CSS_MediaQueryExpression `json:"expressions"`
	Active      bool                       `json:"active"`
}
type CSS_MediaQueryExpression struct {
	Value          float32          `json:"value"`
	Unit           string           `json:"unit"`
	Feature        string           `json:"feature"`
	ValueRange     *CSS_SourceRange `json:"valueRange,omitempty"`
	ComputedLength *float32         `json:"computedLength,omitempty"`
}
type CSS_PlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`
	IsCustomFont bool    `json:"isCustomFont"`
	GlyphCount   float32 `json:"glyphCount"`
}
type CSS_CSSKeyframesRule struct {
	AnimationName CSS_Value             `json:"animationName"`
	Keyframes     []CSS_CSSKeyframeRule `json:"keyframes"`
}
type CSS_CSSKeyframeRule struct {
	StyleSheetId *CSS_StyleSheetId    `json:"styleSheetId,omitempty"`
	Origin       CSS_StyleSheetOrigin `json:"origin"`
	KeyText      CSS_Value            `json:"keyText"`
	Style        CSS_CSSStyle         `json:"style"`
}
type CSS_StyleDeclarationEdit struct {
	StyleSheetId CSS_StyleSheetId `json:"styleSheetId"`
	Range        CSS_SourceRange  `json:"range"`
	Text         string           `json:"text"`
}
type CSS_InlineTextBox struct {
	BoundingBox         DOM_Rect `json:"boundingBox"`
	StartCharacterIndex int      `json:"startCharacterIndex"`
	NumCharacters       int      `json:"numCharacters"`
}
