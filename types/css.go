package types

type CSS_StyleSheetId string
type CSS_StyleSheetOrigin string
type CSS_PseudoElementMatches struct {
	PseudoType	DOM_PseudoType	`json:"pseudoType"`// Pseudo element type.
	Matches		[]CSS_RuleMatch	`json:"matches"`// Matches of CSS rules applicable to the pseudo style.
}
type CSS_InheritedStyleEntry struct {
	InlineStyle	*CSS_CSSStyle	`json:"inlineStyle,omitempty"`// The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules	[]CSS_RuleMatch	`json:"matchedCSSRules"`// Matches of CSS rules matching the ancestor node in the style inheritance chain.
}
type CSS_RuleMatch struct {
	Rule			CSS_CSSRule	`json:"rule"`// CSS rule in the match.
	MatchingSelectors	[]int		`json:"matchingSelectors"`// Matching selector indices in the rule's selectorList selectors (0-based).
}
type CSS_Value struct {
	Text	string			`json:"text"`// Value text.
	Range	*CSS_SourceRange	`json:"range,omitempty"`// Value range in the underlying resource (if available).
}
type CSS_SelectorList struct {
	Selectors	[]CSS_Value	`json:"selectors"`// Selectors in the list.
	Text		string		`json:"text"`// Rule selector text.
}
type CSS_CSSStyleSheetHeader struct {
	StyleSheetId	CSS_StyleSheetId	`json:"styleSheetId"`// The stylesheet identifier.
	FrameId		Page_FrameId		`json:"frameId"`// Owner frame identifier.
	SourceURL	string			`json:"sourceURL"`// Stylesheet resource URL.
	SourceMapURL	*string			`json:"sourceMapURL,omitempty"`// URL of source map associated with the stylesheet (if any).
	Origin		CSS_StyleSheetOrigin	`json:"origin"`// Stylesheet origin.
	Title		string			`json:"title"`// Stylesheet title.
	OwnerNode	*DOM_BackendNodeId	`json:"ownerNode,omitempty"`// The backend id for the owner node of the stylesheet.
	Disabled	bool			`json:"disabled"`// Denotes whether the stylesheet is disabled.
	HasSourceURL	*bool			`json:"hasSourceURL,omitempty"`// Whether the sourceURL field value comes from the sourceURL comment.
	IsInline	bool			`json:"isInline"`// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine	float32			`json:"startLine"`// Line offset of the stylesheet within the resource (zero based).
	StartColumn	float32			`json:"startColumn"`// Column offset of the stylesheet within the resource (zero based).
	Length		float32			`json:"length"`// Size of the content (in characters).
}
type CSS_CSSRule struct {
	StyleSheetId	*CSS_StyleSheetId	`json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList	CSS_SelectorList	`json:"selectorList"`// Rule selector data.
	Origin		CSS_StyleSheetOrigin	`json:"origin"`// Parent stylesheet's origin.
	Style		CSS_CSSStyle		`json:"style"`// Associated style declaration.
	Media		[]CSS_CSSMedia		`json:"media,omitempty"`// Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}
type CSS_RuleUsage struct {
	StyleSheetId	CSS_StyleSheetId	`json:"styleSheetId"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset	float32			`json:"startOffset"`// Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset	float32			`json:"endOffset"`// Offset of the end of the rule body from the beginning of the stylesheet.
	Used		bool			`json:"used"`// Indicates whether the rule was actually used by some element in the page.
}
type CSS_SourceRange struct {
	StartLine	int	`json:"startLine"`// Start line of range.
	StartColumn	int	`json:"startColumn"`// Start column of range (inclusive).
	EndLine		int	`json:"endLine"`// End line of range
	EndColumn	int	`json:"endColumn"`// End column of range (exclusive).
}
type CSS_ShorthandEntry struct {
	Name		string	`json:"name"`// Shorthand name.
	Value		string	`json:"value"`// Shorthand value.
	Important	*bool	`json:"important,omitempty"`// Whether the property has "!important" annotation (implies <code>false</code> if absent).
}
type CSS_CSSComputedStyleProperty struct {
	Name	string	`json:"name"`// Computed style property name.
	Value	string	`json:"value"`// Computed style property value.
}
type CSS_CSSStyle struct {
	StyleSheetId		*CSS_StyleSheetId	`json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CssProperties		[]CSS_CSSProperty	`json:"cssProperties"`// CSS properties in the style.
	ShorthandEntries	[]CSS_ShorthandEntry	`json:"shorthandEntries"`// Computed values for all shorthands found in the style.
	CssText			*string			`json:"cssText,omitempty"`// Style declaration text (if available).
	Range			*CSS_SourceRange	`json:"range,omitempty"`// Style declaration range in the enclosing stylesheet (if available).
}
type CSS_CSSProperty struct {
	Name		string			`json:"name"`// The property name.
	Value		string			`json:"value"`// The property value.
	Important	*bool			`json:"important,omitempty"`// Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Implicit	*bool			`json:"implicit,omitempty"`// Whether the property is implicit (implies <code>false</code> if absent).
	Text		*string			`json:"text,omitempty"`// The full property text as specified in the style.
	ParsedOk	*bool			`json:"parsedOk,omitempty"`// Whether the property is understood by the browser (implies <code>true</code> if absent).
	Disabled	*bool			`json:"disabled,omitempty"`// Whether the property is disabled by the user (present for source-based properties only).
	Range		*CSS_SourceRange	`json:"range,omitempty"`// The entire property range in the enclosing style declaration (if available).
}
type CSS_CSSMedia struct {
	Text		string			`json:"text"`// Media query text.
	Source		string			`json:"source"`// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL	*string			`json:"sourceURL,omitempty"`// URL of the document containing the media query description.
	Range		*CSS_SourceRange	`json:"range,omitempty"`// The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetId	*CSS_StyleSheetId	`json:"styleSheetId,omitempty"`// Identifier of the stylesheet containing this object (if exists).
	MediaList	[]CSS_MediaQuery	`json:"mediaList,omitempty"`// Array of media queries.
}
type CSS_MediaQuery struct {
	Expressions	[]CSS_MediaQueryExpression	`json:"expressions"`// Array of media query expressions.
	Active		bool				`json:"active"`// Whether the media query condition is satisfied.
}
type CSS_MediaQueryExpression struct {
	Value		float32			`json:"value"`// Media query expression value.
	Unit		string			`json:"unit"`// Media query expression units.
	Feature		string			`json:"feature"`// Media query expression feature.
	ValueRange	*CSS_SourceRange	`json:"valueRange,omitempty"`// The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength	*float32		`json:"computedLength,omitempty"`// Computed length of media query expression (if applicable).
}
type CSS_PlatformFontUsage struct {
	FamilyName	string	`json:"familyName"`// Font's family name reported by platform.
	IsCustomFont	bool	`json:"isCustomFont"`// Indicates if the font was downloaded or resolved locally.
	GlyphCount	float32	`json:"glyphCount"`// Amount of glyphs that were rendered with this font.
}
type CSS_CSSKeyframesRule struct {
	AnimationName	CSS_Value		`json:"animationName"`// Animation name.
	Keyframes	[]CSS_CSSKeyframeRule	`json:"keyframes"`// List of keyframes.
}
type CSS_CSSKeyframeRule struct {
	StyleSheetId	*CSS_StyleSheetId	`json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin		CSS_StyleSheetOrigin	`json:"origin"`// Parent stylesheet's origin.
	KeyText		CSS_Value		`json:"keyText"`// Associated key text.
	Style		CSS_CSSStyle		`json:"style"`// Associated style declaration.
}
type CSS_StyleDeclarationEdit struct {
	StyleSheetId	CSS_StyleSheetId	`json:"styleSheetId"`// The css style sheet identifier.
	Range		CSS_SourceRange		`json:"range"`// The range of the style text in the enclosing stylesheet.
	Text		string			`json:"text"`// New style text.
}
type CSS_InlineTextBox struct {
	BoundingBox		DOM_Rect	`json:"boundingBox"`// The absolute position bounding box.
	StartCharacterIndex	int		`json:"startCharacterIndex"`// The starting index in characters, for this post layout textbox substring.
	NumCharacters		int		`json:"numCharacters"`// The number of characters in this post layout textbox substring.
}
