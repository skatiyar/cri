/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package css provides commands and events for CSS domain.
package css

import (
	"github.com/skatiyar/cri"
	types "github.com/skatiyar/cri/types"
)

// List of commands in CSS domain
const (
	AddRule                          = "CSS.addRule"
	CollectClassNames                = "CSS.collectClassNames"
	CreateStyleSheet                 = "CSS.createStyleSheet"
	Disable                          = "CSS.disable"
	Enable                           = "CSS.enable"
	ForcePseudoState                 = "CSS.forcePseudoState"
	GetBackgroundColors              = "CSS.getBackgroundColors"
	GetComputedStyleForNode          = "CSS.getComputedStyleForNode"
	GetInlineStylesForNode           = "CSS.getInlineStylesForNode"
	GetMatchedStylesForNode          = "CSS.getMatchedStylesForNode"
	GetMediaQueries                  = "CSS.getMediaQueries"
	GetPlatformFontsForNode          = "CSS.getPlatformFontsForNode"
	GetStyleSheetText                = "CSS.getStyleSheetText"
	SetEffectivePropertyValueForNode = "CSS.setEffectivePropertyValueForNode"
	SetKeyframeKey                   = "CSS.setKeyframeKey"
	SetMediaText                     = "CSS.setMediaText"
	SetRuleSelector                  = "CSS.setRuleSelector"
	SetStyleSheetText                = "CSS.setStyleSheetText"
	SetStyleTexts                    = "CSS.setStyleTexts"
	StartRuleUsageTracking           = "CSS.startRuleUsageTracking"
	StopRuleUsageTracking            = "CSS.stopRuleUsageTracking"
	TakeCoverageDelta                = "CSS.takeCoverageDelta"
)

// List of events in CSS domain
const (
	FontsUpdated            = "CSS.fontsUpdated"
	MediaQueryResultChanged = "CSS.mediaQueryResultChanged"
	StyleSheetAdded         = "CSS.styleSheetAdded"
	StyleSheetChanged       = "CSS.styleSheetChanged"
	StyleSheetRemoved       = "CSS.styleSheetRemoved"
)

// This domain exposes CSS read/write operations. All CSS objects (stylesheets, rules, and styles) have an associated `id` used in subsequent operations on the related object. Each object type has a specific `id` structure, and those are not interchangeable between objects of different kinds. CSS objects can be loaded using the `get*ForNode()` calls (which accept a DOM node id). A client can also keep track of stylesheets via the `styleSheetAdded`/`styleSheetRemoved` events and subsequently load the required stylesheet contents using the `getStyleSheet[Text]()` methods.
type CSS struct {
	conn cri.Connector
}

// New creates a CSS instance
func New(conn cri.Connector) *CSS {
	return &CSS{conn}
}

type AddRuleRequest struct {
	// The css style sheet identifier where a new rule should be inserted.
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	// The text of a new rule.
	RuleText string `json:"ruleText"`
	// Text position of a new rule in the target style sheet.
	Location types.CSS_SourceRange `json:"location"`
}

type AddRuleResponse struct {
	// The newly created rule.
	Rule types.CSS_CSSRule `json:"rule"`
}

// Inserts a new rule with the given `ruleText` in a stylesheet with given `styleSheetId`, at the position specified by `location`.
func (obj *CSS) AddRule(request *AddRuleRequest) (response AddRuleResponse, err error) {
	err = obj.conn.Send(AddRule, request, &response)
	return
}

type CollectClassNamesRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

type CollectClassNamesResponse struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Returns all class names from specified stylesheet.
func (obj *CSS) CollectClassNames(request *CollectClassNamesRequest) (response CollectClassNamesResponse, err error) {
	err = obj.conn.Send(CollectClassNames, request, &response)
	return
}

type CreateStyleSheetRequest struct {
	// Identifier of the frame where "via-inspector" stylesheet should be created.
	FrameId types.Page_FrameId `json:"frameId"`
}

type CreateStyleSheetResponse struct {
	// Identifier of the created "via-inspector" stylesheet.
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

// Creates a new special "via-inspector" stylesheet in the frame with given `frameId`.
func (obj *CSS) CreateStyleSheet(request *CreateStyleSheetRequest) (response CreateStyleSheetResponse, err error) {
	err = obj.conn.Send(CreateStyleSheet, request, &response)
	return
}

// Disables the CSS agent for the given page.
func (obj *CSS) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (obj *CSS) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type ForcePseudoStateRequest struct {
	// The element id for which to force the pseudo state.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Element pseudo classes to force when computing the element's style.
	ForcedPseudoClasses []string `json:"forcedPseudoClasses"`
}

// Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
func (obj *CSS) ForcePseudoState(request *ForcePseudoStateRequest) (err error) {
	err = obj.conn.Send(ForcePseudoState, request, nil)
	return
}

type GetBackgroundColorsRequest struct {
	// Id of the node to get background colors for.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetBackgroundColorsResponse struct {
	// The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	BackgroundColors []string `json:"backgroundColors,omitempty"`
	// The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontSize *string `json:"computedFontSize,omitempty"`
	// The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
	ComputedFontWeight *string `json:"computedFontWeight,omitempty"`
	// The computed font size for the document body, as a computed CSS value string (e.g. '16px').
	ComputedBodyFontSize *string `json:"computedBodyFontSize,omitempty"`
}

func (obj *CSS) GetBackgroundColors(request *GetBackgroundColorsRequest) (response GetBackgroundColorsResponse, err error) {
	err = obj.conn.Send(GetBackgroundColors, request, &response)
	return
}

type GetComputedStyleForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetComputedStyleForNodeResponse struct {
	// Computed style for the specified DOM node.
	ComputedStyle []types.CSS_CSSComputedStyleProperty `json:"computedStyle"`
}

// Returns the computed style for a DOM node identified by `nodeId`.
func (obj *CSS) GetComputedStyleForNode(request *GetComputedStyleForNodeRequest) (response GetComputedStyleForNodeResponse, err error) {
	err = obj.conn.Send(GetComputedStyleForNode, request, &response)
	return
}

type GetInlineStylesForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetInlineStylesForNodeResponse struct {
	// Inline style for the specified DOM node.
	InlineStyle *types.CSS_CSSStyle `json:"inlineStyle,omitempty"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle *types.CSS_CSSStyle `json:"attributesStyle,omitempty"`
}

// Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by `nodeId`.
func (obj *CSS) GetInlineStylesForNode(request *GetInlineStylesForNodeRequest) (response GetInlineStylesForNodeResponse, err error) {
	err = obj.conn.Send(GetInlineStylesForNode, request, &response)
	return
}

type GetMatchedStylesForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetMatchedStylesForNodeResponse struct {
	// Inline style for the specified DOM node.
	InlineStyle *types.CSS_CSSStyle `json:"inlineStyle,omitempty"`
	// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	AttributesStyle *types.CSS_CSSStyle `json:"attributesStyle,omitempty"`
	// CSS rules matching this node, from all applicable stylesheets.
	MatchedCSSRules []types.CSS_RuleMatch `json:"matchedCSSRules,omitempty"`
	// Pseudo style matches for this node.
	PseudoElements []types.CSS_PseudoElementMatches `json:"pseudoElements,omitempty"`
	// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	Inherited []types.CSS_InheritedStyleEntry `json:"inherited,omitempty"`
	// A list of CSS keyframed animations matching this node.
	CssKeyframesRules []types.CSS_CSSKeyframesRule `json:"cssKeyframesRules,omitempty"`
}

// Returns requested styles for a DOM node identified by `nodeId`.
func (obj *CSS) GetMatchedStylesForNode(request *GetMatchedStylesForNodeRequest) (response GetMatchedStylesForNodeResponse, err error) {
	err = obj.conn.Send(GetMatchedStylesForNode, request, &response)
	return
}

type GetMediaQueriesResponse struct {
	Medias []types.CSS_CSSMedia `json:"medias"`
}

// Returns all media queries parsed by the rendering engine.
func (obj *CSS) GetMediaQueries() (response GetMediaQueriesResponse, err error) {
	err = obj.conn.Send(GetMediaQueries, nil, &response)
	return
}

type GetPlatformFontsForNodeRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetPlatformFontsForNodeResponse struct {
	// Usage statistics for every employed platform font.
	Fonts []types.CSS_PlatformFontUsage `json:"fonts"`
}

// Requests information about platform fonts which we used to render child TextNodes in the given node.
func (obj *CSS) GetPlatformFontsForNode(request *GetPlatformFontsForNodeRequest) (response GetPlatformFontsForNodeResponse, err error) {
	err = obj.conn.Send(GetPlatformFontsForNode, request, &response)
	return
}

type GetStyleSheetTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

type GetStyleSheetTextResponse struct {
	// The stylesheet text.
	Text string `json:"text"`
}

// Returns the current textual content for a stylesheet.
func (obj *CSS) GetStyleSheetText(request *GetStyleSheetTextRequest) (response GetStyleSheetTextResponse, err error) {
	err = obj.conn.Send(GetStyleSheetText, request, &response)
	return
}

type SetEffectivePropertyValueForNodeRequest struct {
	// The element id for which to set property.
	NodeId       types.DOM_NodeId `json:"nodeId"`
	PropertyName string           `json:"propertyName"`
	Value        string           `json:"value"`
}

// Find a rule with the given active property for the given node and set the new value for this property
func (obj *CSS) SetEffectivePropertyValueForNode(request *SetEffectivePropertyValueForNodeRequest) (err error) {
	err = obj.conn.Send(SetEffectivePropertyValueForNode, request, nil)
	return
}

type SetKeyframeKeyRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	KeyText      string                 `json:"keyText"`
}

type SetKeyframeKeyResponse struct {
	// The resulting key text after modification.
	KeyText types.CSS_Value `json:"keyText"`
}

// Modifies the keyframe rule key text.
func (obj *CSS) SetKeyframeKey(request *SetKeyframeKeyRequest) (response SetKeyframeKeyResponse, err error) {
	err = obj.conn.Send(SetKeyframeKey, request, &response)
	return
}

type SetMediaTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	Text         string                 `json:"text"`
}

type SetMediaTextResponse struct {
	// The resulting CSS media rule after modification.
	Media types.CSS_CSSMedia `json:"media"`
}

// Modifies the rule selector.
func (obj *CSS) SetMediaText(request *SetMediaTextRequest) (response SetMediaTextResponse, err error) {
	err = obj.conn.Send(SetMediaText, request, &response)
	return
}

type SetRuleSelectorRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Range        types.CSS_SourceRange  `json:"range"`
	Selector     string                 `json:"selector"`
}

type SetRuleSelectorResponse struct {
	// The resulting selector list after modification.
	SelectorList types.CSS_SelectorList `json:"selectorList"`
}

// Modifies the rule selector.
func (obj *CSS) SetRuleSelector(request *SetRuleSelectorRequest) (response SetRuleSelectorResponse, err error) {
	err = obj.conn.Send(SetRuleSelector, request, &response)
	return
}

type SetStyleSheetTextRequest struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
	Text         string                 `json:"text"`
}

type SetStyleSheetTextResponse struct {
	// URL of source map associated with script (if any).
	SourceMapURL *string `json:"sourceMapURL,omitempty"`
}

// Sets the new stylesheet text.
func (obj *CSS) SetStyleSheetText(request *SetStyleSheetTextRequest) (response SetStyleSheetTextResponse, err error) {
	err = obj.conn.Send(SetStyleSheetText, request, &response)
	return
}

type SetStyleTextsRequest struct {
	Edits []types.CSS_StyleDeclarationEdit `json:"edits"`
}

type SetStyleTextsResponse struct {
	// The resulting styles after modification.
	Styles []types.CSS_CSSStyle `json:"styles"`
}

// Applies specified style edits one after another in the given order.
func (obj *CSS) SetStyleTexts(request *SetStyleTextsRequest) (response SetStyleTextsResponse, err error) {
	err = obj.conn.Send(SetStyleTexts, request, &response)
	return
}

// Enables the selector recording.
func (obj *CSS) StartRuleUsageTracking() (err error) {
	err = obj.conn.Send(StartRuleUsageTracking, nil, nil)
	return
}

type StopRuleUsageTrackingResponse struct {
	RuleUsage []types.CSS_RuleUsage `json:"ruleUsage"`
}

// Stop tracking rule usage and return the list of rules that were used since last call to `takeCoverageDelta` (or since start of coverage instrumentation)
func (obj *CSS) StopRuleUsageTracking() (response StopRuleUsageTrackingResponse, err error) {
	err = obj.conn.Send(StopRuleUsageTracking, nil, &response)
	return
}

type TakeCoverageDeltaResponse struct {
	Coverage []types.CSS_RuleUsage `json:"coverage"`
}

// Obtain list of rules that became used since last call to this method (or since start of coverage instrumentation)
func (obj *CSS) TakeCoverageDelta() (response TakeCoverageDeltaResponse, err error) {
	err = obj.conn.Send(TakeCoverageDelta, nil, &response)
	return
}

type FontsUpdatedParams struct {
	// The web font that has loaded.
	Font *types.CSS_FontFace `json:"font,omitempty"`
}

// Fires whenever a web font is updated.  A non-empty font parameter indicates a successfully loaded web font
func (obj *CSS) FontsUpdated(fn func(event string, params FontsUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(FontsUpdated)
	go func() {
		defer closer()
		for {
			var params FontsUpdatedParams
			if !fn(FontsUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

// Fires whenever a MediaQuery result changes (for example, after a browser window has been resized.) The current implementation considers only viewport-dependent media features.
func (obj *CSS) MediaQueryResultChanged(fn func(event string, err error) bool) {
	listen, closer := obj.conn.On(MediaQueryResultChanged)
	go func() {
		defer closer()
		for {
			if !fn(MediaQueryResultChanged, listen(nil)) {
				return
			}
		}
	}()
}

type StyleSheetAddedParams struct {
	// Added stylesheet metainfo.
	Header types.CSS_CSSStyleSheetHeader `json:"header"`
}

// Fired whenever an active document stylesheet is added.
func (obj *CSS) StyleSheetAdded(fn func(event string, params StyleSheetAddedParams, err error) bool) {
	listen, closer := obj.conn.On(StyleSheetAdded)
	go func() {
		defer closer()
		for {
			var params StyleSheetAddedParams
			if !fn(StyleSheetAdded, params, listen(&params)) {
				return
			}
		}
	}()
}

type StyleSheetChangedParams struct {
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

// Fired whenever a stylesheet is changed as a result of the client operation.
func (obj *CSS) StyleSheetChanged(fn func(event string, params StyleSheetChangedParams, err error) bool) {
	listen, closer := obj.conn.On(StyleSheetChanged)
	go func() {
		defer closer()
		for {
			var params StyleSheetChangedParams
			if !fn(StyleSheetChanged, params, listen(&params)) {
				return
			}
		}
	}()
}

type StyleSheetRemovedParams struct {
	// Identifier of the removed stylesheet.
	StyleSheetId types.CSS_StyleSheetId `json:"styleSheetId"`
}

// Fired whenever an active document stylesheet is removed.
func (obj *CSS) StyleSheetRemoved(fn func(event string, params StyleSheetRemovedParams, err error) bool) {
	listen, closer := obj.conn.On(StyleSheetRemoved)
	go func() {
		defer closer()
		for {
			var params StyleSheetRemovedParams
			if !fn(StyleSheetRemoved, params, listen(&params)) {
				return
			}
		}
	}()
}
