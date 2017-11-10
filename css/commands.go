package css

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type CSS struct {
	conn cri.Connector
}

func New(conn cri.Connector) *CSS {
	return &CSS{conn}
}
func (obj *CSS) Enable() (err error) {
	err = obj.conn.Send("CSS.enable", nil, nil)
	return
}
func (obj *CSS) Disable() (err error) {
	err = obj.conn.Send("CSS.disable", nil, nil)
	return
}

type GetMatchedStylesForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetMatchedStylesForNode(request *GetMatchedStylesForNodeRequest) (response struct {
	InlineStyle		*types.CSS_CSSStyle			`json:"inlineStyle,omitempty"`// Inline style for the specified DOM node.
	AttributesStyle		*types.CSS_CSSStyle			`json:"attributesStyle,omitempty"`// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules		[]types.CSS_RuleMatch			`json:"matchedCSSRules,omitempty"`// CSS rules matching this node, from all applicable stylesheets.
	PseudoElements		[]types.CSS_PseudoElementMatches	`json:"pseudoElements,omitempty"`// Pseudo style matches for this node.
	Inherited		[]types.CSS_InheritedStyleEntry		`json:"inherited,omitempty"`// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	CssKeyframesRules	[]types.CSS_CSSKeyframesRule		`json:"cssKeyframesRules,omitempty"`// A list of CSS keyframed animations matching this node.
}, err error) {
	err = obj.conn.Send("CSS.getMatchedStylesForNode", request, &response)
	return
}

type GetInlineStylesForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetInlineStylesForNode(request *GetInlineStylesForNodeRequest) (response struct {
	InlineStyle	*types.CSS_CSSStyle	`json:"inlineStyle,omitempty"`// Inline style for the specified DOM node.
	AttributesStyle	*types.CSS_CSSStyle	`json:"attributesStyle,omitempty"`// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}, err error) {
	err = obj.conn.Send("CSS.getInlineStylesForNode", request, &response)
	return
}

type GetComputedStyleForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetComputedStyleForNode(request *GetComputedStyleForNodeRequest) (response struct {
	ComputedStyle []types.CSS_CSSComputedStyleProperty `json:"computedStyle"`// Computed style for the specified DOM node.
}, err error) {
	err = obj.conn.Send("CSS.getComputedStyleForNode", request, &response)
	return
}

type GetPlatformFontsForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetPlatformFontsForNode(request *GetPlatformFontsForNodeRequest) (response struct {
	Fonts []types.CSS_PlatformFontUsage `json:"fonts"`// Usage statistics for every employed platform font.
}, err error) {
	err = obj.conn.Send("CSS.getPlatformFontsForNode", request, &response)
	return
}

type GetStyleSheetTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

func (obj *CSS) GetStyleSheetText(request *GetStyleSheetTextRequest) (response struct {
	Text string `json:"text"`// The stylesheet text.
}, err error) {
	err = obj.conn.Send("CSS.getStyleSheetText", request, &response)
	return
}

type CollectClassNamesRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

func (obj *CSS) CollectClassNames(request *CollectClassNamesRequest) (response struct {
	ClassNames []string `json:"classNames"`// Class name list.
}, err error) {
	err = obj.conn.Send("CSS.collectClassNames", request, &response)
	return
}

type SetStyleSheetTextRequest struct {
	StyleSheetId	types.CSS_StyleSheetId	`json:"styleSheetId"`
	Text		string			`json:"text"`
}

func (obj *CSS) SetStyleSheetText(request *SetStyleSheetTextRequest) (response struct {
	SourceMapURL *string `json:"sourceMapURL,omitempty"`// URL of source map associated with script (if any).
}, err error) {
	err = obj.conn.Send("CSS.setStyleSheetText", request, &response)
	return
}

type SetRuleSelectorRequest struct {
	StyleSheetId	types.CSS_StyleSheetId	`json:"styleSheetId"`
	Range		types.CSS_SourceRange	`json:"range"`
	Selector	string			`json:"selector"`
}

func (obj *CSS) SetRuleSelector(request *SetRuleSelectorRequest) (response struct {
	SelectorList types.CSS_SelectorList `json:"selectorList"`// The resulting selector list after modification.
}, err error) {
	err = obj.conn.Send("CSS.setRuleSelector", request, &response)
	return
}

type SetKeyframeKeyRequest struct {
	StyleSheetId	types.CSS_StyleSheetId	`json:"styleSheetId"`
	Range		types.CSS_SourceRange	`json:"range"`
	KeyText		string			`json:"keyText"`
}

func (obj *CSS) SetKeyframeKey(request *SetKeyframeKeyRequest) (response struct {
	KeyText types.CSS_Value `json:"keyText"`// The resulting key text after modification.
}, err error) {
	err = obj.conn.Send("CSS.setKeyframeKey", request, &response)
	return
}

type SetStyleTextsRequest struct {
	Edits []types.CSS_StyleDeclarationEdit `json:"edits"`
}

func (obj *CSS) SetStyleTexts(request *SetStyleTextsRequest) (response struct {
	Styles []types.CSS_CSSStyle `json:"styles"`// The resulting styles after modification.
}, err error) {
	err = obj.conn.Send("CSS.setStyleTexts", request, &response)
	return
}

type SetMediaTextRequest struct {
	StyleSheetId	types.CSS_StyleSheetId	`json:"styleSheetId"`
	Range		types.CSS_SourceRange	`json:"range"`
	Text		string			`json:"text"`
}

func (obj *CSS) SetMediaText(request *SetMediaTextRequest) (response struct {
	Media types.CSS_CSSMedia `json:"media"`// The resulting CSS media rule after modification.
}, err error) {
	err = obj.conn.Send("CSS.setMediaText", request, &response)
	return
}

type CreateStyleSheetRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`// Identifier of the frame where "via-inspector" stylesheet should be created.
}

func (obj *CSS) CreateStyleSheet(request *CreateStyleSheetRequest) (response struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`// Identifier of the created "via-inspector" stylesheet.
}, err error) {
	err = obj.conn.Send("CSS.createStyleSheet", request, &response)
	return
}

type AddRuleRequest struct {
	StyleSheetId	types.CSS_StyleSheetId	`json:"styleSheetId"`// The css style sheet identifier where a new rule should be inserted.
	RuleText	string			`json:"ruleText"`// The text of a new rule.
	Location	types.CSS_SourceRange	`json:"location"`// Text position of a new rule in the target style sheet.
}

func (obj *CSS) AddRule(request *AddRuleRequest) (response struct {
	Rule types.CSS_CSSRule `json:"rule"`// The newly created rule.
}, err error) {
	err = obj.conn.Send("CSS.addRule", request, &response)
	return
}

type ForcePseudoStateRequest struct {
	NodeId			types.DOM_NodeId	`json:"nodeId"`// The element id for which to force the pseudo state.
	ForcedPseudoClasses	[]string		`json:"forcedPseudoClasses"`// Element pseudo classes to force when computing the element's style.
}

func (obj *CSS) ForcePseudoState(request *ForcePseudoStateRequest) (err error) {
	err = obj.conn.Send("CSS.forcePseudoState", request, nil)
	return
}
func (obj *CSS) GetMediaQueries() (response struct {
	Medias []types.CSS_CSSMedia `json:"medias"`
}, err error) {
	err = obj.conn.Send("CSS.getMediaQueries", nil, &response)
	return
}

type SetEffectivePropertyValueForNodeRequest struct {
	NodeId		types.DOM_NodeId	`json:"nodeId"`// The element id for which to set property.
	PropertyName	string			`json:"propertyName"`
	Value		string			`json:"value"`
}

func (obj *CSS) SetEffectivePropertyValueForNode(request *SetEffectivePropertyValueForNodeRequest) (err error) {
	err = obj.conn.Send("CSS.setEffectivePropertyValueForNode", request, nil)
	return
}

type GetBackgroundColorsRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node to get background colors for.
}

func (obj *CSS) GetBackgroundColors(request *GetBackgroundColorsRequest) (response struct {
	BackgroundColors	[]string	`json:"backgroundColors,omitempty"`// The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	ComputedFontSize	*string		`json:"computedFontSize,omitempty"`// The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontWeight	*string		`json:"computedFontWeight,omitempty"`// The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
	ComputedBodyFontSize	*string		`json:"computedBodyFontSize,omitempty"`// The computed font size for the document body, as a computed CSS value string (e.g. '16px').
}, err error) {
	err = obj.conn.Send("CSS.getBackgroundColors", request, &response)
	return
}
func (obj *CSS) StartRuleUsageTracking() (err error) {
	err = obj.conn.Send("CSS.startRuleUsageTracking", nil, nil)
	return
}
func (obj *CSS) TakeCoverageDelta() (response struct {
	Coverage []types.CSS_RuleUsage `json:"coverage"`
}, err error) {
	err = obj.conn.Send("CSS.takeCoverageDelta", nil, &response)
	return
}
func (obj *CSS) StopRuleUsageTracking() (response struct {
	RuleUsage []types.CSS_RuleUsage `json:"ruleUsage"`
}, err error) {
	err = obj.conn.Send("CSS.stopRuleUsageTracking", nil, &response)
	return
}
