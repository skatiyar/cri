/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package types



type CSS_StyleSheetId string

//Stylesheet type: "injected" for stylesheets injected via extension, "user-agent" for user-agent stylesheets, "inspector" for stylesheets created by the inspector (i.e. those holding the "via inspector" rules), "regular" for regular stylesheets.
type CSS_StyleSheetOrigin string

//CSS rule collection for a single pseudo style.
type CSS_PseudoElementMatches struct {
	// Pseudo element type.
	PseudoType DOM_PseudoType `json:"pseudoType"`
	// Matches of CSS rules applicable to the pseudo style.
	Matches []CSS_RuleMatch `json:"matches"`
}

//Inherited CSS rule collection from ancestor node.
type CSS_InheritedStyleEntry struct {
	// The ancestor node's inline style, if any, in the style inheritance chain.
	InlineStyle *CSS_CSSStyle `json:"inlineStyle,omitempty"`
	// Matches of CSS rules matching the ancestor node in the style inheritance chain.
	MatchedCSSRules []CSS_RuleMatch `json:"matchedCSSRules"`
}

//Match data for a CSS rule.
type CSS_RuleMatch struct {
	// CSS rule in the match.
	Rule CSS_CSSRule `json:"rule"`
	// Matching selector indices in the rule's selectorList selectors (0-based).
	MatchingSelectors []int `json:"matchingSelectors"`
}

//Data for a simple selector (these are delimited by commas in a selector list).
type CSS_Value struct {
	// Value text.
	Text string `json:"text"`
	// Value range in the underlying resource (if available).
	Range *CSS_SourceRange `json:"range,omitempty"`
}

//Selector list data.
type CSS_SelectorList struct {
	// Selectors in the list.
	Selectors []CSS_Value `json:"selectors"`
	// Rule selector text.
	Text string `json:"text"`
}

//CSS stylesheet metainformation.
type CSS_CSSStyleSheetHeader struct {
	// The stylesheet identifier.
	StyleSheetId CSS_StyleSheetId `json:"styleSheetId"`
	// Owner frame identifier.
	FrameId Page_FrameId `json:"frameId"`
	// Stylesheet resource URL.
	SourceURL string `json:"sourceURL"`
	// URL of source map associated with the stylesheet (if any).
	SourceMapURL *string `json:"sourceMapURL,omitempty"`
	// Stylesheet origin.
	Origin CSS_StyleSheetOrigin `json:"origin"`
	// Stylesheet title.
	Title string `json:"title"`
	// The backend id for the owner node of the stylesheet.
	OwnerNode *DOM_BackendNodeId `json:"ownerNode,omitempty"`
	// Denotes whether the stylesheet is disabled.
	Disabled bool `json:"disabled"`
	// Whether the sourceURL field value comes from the sourceURL comment.
	HasSourceURL *bool `json:"hasSourceURL,omitempty"`
	// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	IsInline bool `json:"isInline"`
	// Line offset of the stylesheet within the resource (zero based).
	StartLine float32 `json:"startLine"`
	// Column offset of the stylesheet within the resource (zero based).
	StartColumn float32 `json:"startColumn"`
	// Size of the content (in characters).
	// NOTE Experimental
	Length float32 `json:"length"`
}

//CSS rule representation.
type CSS_CSSRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetId *CSS_StyleSheetId `json:"styleSheetId,omitempty"`
	// Rule selector data.
	SelectorList CSS_SelectorList `json:"selectorList"`
	// Parent stylesheet's origin.
	Origin CSS_StyleSheetOrigin `json:"origin"`
	// Associated style declaration.
	Style CSS_CSSStyle `json:"style"`
	// Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
	Media []CSS_CSSMedia `json:"media,omitempty"`
}

//CSS coverage information.
type CSS_RuleUsage struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetId CSS_StyleSheetId `json:"styleSheetId"`
	// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	StartOffset float32 `json:"startOffset"`
	// Offset of the end of the rule body from the beginning of the stylesheet.
	EndOffset float32 `json:"endOffset"`
	// Indicates whether the rule was actually used by some element in the page.
	Used bool `json:"used"`
}

//Text range within a resource. All numbers are zero-based.
type CSS_SourceRange struct {
	// Start line of range.
	StartLine int `json:"startLine"`
	// Start column of range (inclusive).
	StartColumn int `json:"startColumn"`
	// End line of range
	EndLine int `json:"endLine"`
	// End column of range (exclusive).
	EndColumn int `json:"endColumn"`
}


type CSS_ShorthandEntry struct {
	// Shorthand name.
	Name string `json:"name"`
	// Shorthand value.
	Value string `json:"value"`
	// Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Important *bool `json:"important,omitempty"`
}


type CSS_CSSComputedStyleProperty struct {
	// Computed style property name.
	Name string `json:"name"`
	// Computed style property value.
	Value string `json:"value"`
}

//CSS style representation.
type CSS_CSSStyle struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetId *CSS_StyleSheetId `json:"styleSheetId,omitempty"`
	// CSS properties in the style.
	CssProperties []CSS_CSSProperty `json:"cssProperties"`
	// Computed values for all shorthands found in the style.
	ShorthandEntries []CSS_ShorthandEntry `json:"shorthandEntries"`
	// Style declaration text (if available).
	CssText *string `json:"cssText,omitempty"`
	// Style declaration range in the enclosing stylesheet (if available).
	Range *CSS_SourceRange `json:"range,omitempty"`
}

//CSS property declaration data.
type CSS_CSSProperty struct {
	// The property name.
	Name string `json:"name"`
	// The property value.
	Value string `json:"value"`
	// Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Important *bool `json:"important,omitempty"`
	// Whether the property is implicit (implies <code>false</code> if absent).
	Implicit *bool `json:"implicit,omitempty"`
	// The full property text as specified in the style.
	Text *string `json:"text,omitempty"`
	// Whether the property is understood by the browser (implies <code>true</code> if absent).
	ParsedOk *bool `json:"parsedOk,omitempty"`
	// Whether the property is disabled by the user (present for source-based properties only).
	Disabled *bool `json:"disabled,omitempty"`
	// The entire property range in the enclosing style declaration (if available).
	Range *CSS_SourceRange `json:"range,omitempty"`
}

//CSS media rule descriptor.
type CSS_CSSMedia struct {
	// Media query text.
	Text string `json:"text"`
	// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	Source string `json:"source"`
	// URL of the document containing the media query description.
	SourceURL *string `json:"sourceURL,omitempty"`
	// The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	Range *CSS_SourceRange `json:"range,omitempty"`
	// Identifier of the stylesheet containing this object (if exists).
	StyleSheetId *CSS_StyleSheetId `json:"styleSheetId,omitempty"`
	// Array of media queries.
	// NOTE Experimental
	MediaList []CSS_MediaQuery `json:"mediaList,omitempty"`
}

//Media query descriptor.
type CSS_MediaQuery struct {
	// Array of media query expressions.
	Expressions []CSS_MediaQueryExpression `json:"expressions"`
	// Whether the media query condition is satisfied.
	Active bool `json:"active"`
}

//Media query expression descriptor.
type CSS_MediaQueryExpression struct {
	// Media query expression value.
	Value float32 `json:"value"`
	// Media query expression units.
	Unit string `json:"unit"`
	// Media query expression feature.
	Feature string `json:"feature"`
	// The associated range of the value text in the enclosing stylesheet (if available).
	ValueRange *CSS_SourceRange `json:"valueRange,omitempty"`
	// Computed length of media query expression (if applicable).
	ComputedLength *float32 `json:"computedLength,omitempty"`
}

//Information about amount of glyphs that were rendered with given font.
type CSS_PlatformFontUsage struct {
	// Font's family name reported by platform.
	FamilyName string `json:"familyName"`
	// Indicates if the font was downloaded or resolved locally.
	IsCustomFont bool `json:"isCustomFont"`
	// Amount of glyphs that were rendered with this font.
	GlyphCount float32 `json:"glyphCount"`
}

//CSS keyframes rule representation.
type CSS_CSSKeyframesRule struct {
	// Animation name.
	AnimationName CSS_Value `json:"animationName"`
	// List of keyframes.
	Keyframes []CSS_CSSKeyframeRule `json:"keyframes"`
}

//CSS keyframe rule representation.
type CSS_CSSKeyframeRule struct {
	// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StyleSheetId *CSS_StyleSheetId `json:"styleSheetId,omitempty"`
	// Parent stylesheet's origin.
	Origin CSS_StyleSheetOrigin `json:"origin"`
	// Associated key text.
	KeyText CSS_Value `json:"keyText"`
	// Associated style declaration.
	Style CSS_CSSStyle `json:"style"`
}

//A descriptor of operation to mutate style declaration text.
type CSS_StyleDeclarationEdit struct {
	// The css style sheet identifier.
	StyleSheetId CSS_StyleSheetId `json:"styleSheetId"`
	// The range of the style text in the enclosing stylesheet.
	Range CSS_SourceRange `json:"range"`
	// New style text.
	Text string `json:"text"`
}

//Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type CSS_InlineTextBox struct {
	// The absolute position bounding box.
	BoundingBox DOM_Rect `json:"boundingBox"`
	// The starting index in characters, for this post layout textbox substring.
	StartCharacterIndex int `json:"startCharacterIndex"`
	// The number of characters in this post layout textbox substring.
	NumCharacters int `json:"numCharacters"`
}

