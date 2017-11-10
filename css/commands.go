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
	InlineStyle       *types.CSS_CSSStyle              `json:"inlineStyle,omitempty"`
	AttributesStyle   *types.CSS_CSSStyle              `json:"attributesStyle,omitempty"`
	MatchedCSSRules   []types.CSS_RuleMatch            `json:"matchedCSSRules,omitempty"`
	PseudoElements    []types.CSS_PseudoElementMatches `json:"pseudoElements,omitempty"`
	Inherited         []types.CSS_InheritedStyleEntry  `json:"inherited,omitempty"`
	CssKeyframesRules []types.CSS_CSSKeyframesRule     `json:"cssKeyframesRules,omitempty"`
}, err error) {
	err = obj.conn.Send("CSS.getMatchedStylesForNode", request, &response)
	return
}

type GetInlineStylesForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetInlineStylesForNode(request *GetInlineStylesForNodeRequest) (response struct {
	InlineStyle     *types.CSS_CSSStyle `json:"inlineStyle,omitempty"`
	AttributesStyle *types.CSS_CSSStyle `json:"attributesStyle,omitempty"`
}, err error) {
	err = obj.conn.Send("CSS.getInlineStylesForNode", request, &response)
	return
}

type GetComputedStyleForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetComputedStyleForNode(request *GetComputedStyleForNodeRequest) (response struct {
	ComputedStyle []types.CSS_CSSComputedStyleProperty `json:"computedStyle"`
}, err error) {
	err = obj.conn.Send("CSS.getComputedStyleForNode", request, &response)
	return
}

type GetPlatformFontsForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetPlatformFontsForNode(request *GetPlatformFontsForNodeRequest) (response struct {
	Fonts []types.CSS_PlatformFontUsage `json:"fonts"`
}, err error) {
	err = obj.conn.Send("CSS.getPlatformFontsForNode", request, &response)
	return
}

type GetStyleSheetTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

func (obj *CSS) GetStyleSheetText(request *GetStyleSheetTextRequest) (response struct {
	Text string `json:"text"`
}, err error) {
	err = obj.conn.Send("CSS.getStyleSheetText", request, &response)
	return
}

type CollectClassNamesRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

func (obj *CSS) CollectClassNames(request *CollectClassNamesRequest) (response struct {
	ClassNames []string `json:"classNames"`
}, err error) {
	err = obj.conn.Send("CSS.collectClassNames", request, &response)
	return
}

type SetStyleSheetTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Text         string                 `json:"text"`
}

func (obj *CSS) SetStyleSheetText(request *SetStyleSheetTextRequest) (response struct {
	SourceMapURL *string `json:"sourceMapURL,omitempty"`
}, err error) {
	err = obj.conn.Send("CSS.setStyleSheetText", request, &response)
	return
}

type SetRuleSelectorRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	Selector     string                 `json:"selector"`
}

func (obj *CSS) SetRuleSelector(request *SetRuleSelectorRequest) (response struct {
	SelectorList types.CSS_SelectorList `json:"selectorList"`
}, err error) {
	err = obj.conn.Send("CSS.setRuleSelector", request, &response)
	return
}

type SetKeyframeKeyRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	KeyText      string                 `json:"keyText"`
}

func (obj *CSS) SetKeyframeKey(request *SetKeyframeKeyRequest) (response struct {
	KeyText types.CSS_Value `json:"keyText"`
}, err error) {
	err = obj.conn.Send("CSS.setKeyframeKey", request, &response)
	return
}

type SetStyleTextsRequest struct {
	Edits []types.CSS_StyleDeclarationEdit `json:"edits"`
}

func (obj *CSS) SetStyleTexts(request *SetStyleTextsRequest) (response struct {
	Styles []types.CSS_CSSStyle `json:"styles"`
}, err error) {
	err = obj.conn.Send("CSS.setStyleTexts", request, &response)
	return
}

type SetMediaTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	Text         string                 `json:"text"`
}

func (obj *CSS) SetMediaText(request *SetMediaTextRequest) (response struct {
	Media types.CSS_CSSMedia `json:"media"`
}, err error) {
	err = obj.conn.Send("CSS.setMediaText", request, &response)
	return
}

type CreateStyleSheetRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`
}

func (obj *CSS) CreateStyleSheet(request *CreateStyleSheetRequest) (response struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}, err error) {
	err = obj.conn.Send("CSS.createStyleSheet", request, &response)
	return
}

type AddRuleRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	RuleText     string                 `json:"ruleText"`
	Location     types.CSS_SourceRange  `json:"location"`
}

func (obj *CSS) AddRule(request *AddRuleRequest) (response struct {
	Rule types.CSS_CSSRule `json:"rule"`
}, err error) {
	err = obj.conn.Send("CSS.addRule", request, &response)
	return
}

type ForcePseudoStateRequest struct {
	NodeId              types.DOM_NodeId `json:"nodeId"`
	ForcedPseudoClasses []string         `json:"forcedPseudoClasses"`
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
	NodeId       types.DOM_NodeId `json:"nodeId"`
	PropertyName string           `json:"propertyName"`
	Value        string           `json:"value"`
}

func (obj *CSS) SetEffectivePropertyValueForNode(request *SetEffectivePropertyValueForNodeRequest) (err error) {
	err = obj.conn.Send("CSS.setEffectivePropertyValueForNode", request, nil)
	return
}

type GetBackgroundColorsRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *CSS) GetBackgroundColors(request *GetBackgroundColorsRequest) (response struct {
	BackgroundColors     []string `json:"backgroundColors,omitempty"`
	ComputedFontSize     *string  `json:"computedFontSize,omitempty"`
	ComputedFontWeight   *string  `json:"computedFontWeight,omitempty"`
	ComputedBodyFontSize *string  `json:"computedBodyFontSize,omitempty"`
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
