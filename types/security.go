/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// An internal certificate ID value.
type Security_CertificateId int

// A description of mixed content (HTTP resources on HTTPS pages), as defined by https://www.w3.org/TR/mixed-content/#categories
type Security_MixedContentType string

// The security level of a page or resource.
type Security_SecurityState string

// An explanation of an factor contributing to the security state.
type Security_SecurityStateExplanation struct {
	// Security state representing the severity of the factor being explained.
	SecurityState Security_SecurityState `json:"securityState"`
	// Title describing the type of factor.
	Title string `json:"title"`
	// Short phrase describing the type of factor.
	Summary string `json:"summary"`
	// Full text explanation of the factor.
	Description string `json:"description"`
	// The type of mixed content described by the explanation.
	MixedContentType Security_MixedContentType `json:"mixedContentType"`
	// Page certificate.
	Certificate []string `json:"certificate"`
}

// Information about insecure content on the page.
type Security_InsecureContentStatus struct {
	// True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	RanMixedContent bool `json:"ranMixedContent"`
	// True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	DisplayedMixedContent bool `json:"displayedMixedContent"`
	// True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	ContainedMixedForm bool `json:"containedMixedForm"`
	// True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	RanContentWithCertErrors bool `json:"ranContentWithCertErrors"`
	// True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	DisplayedContentWithCertErrors bool `json:"displayedContentWithCertErrors"`
	// Security state representing a page that ran insecure content.
	RanInsecureContentStyle Security_SecurityState `json:"ranInsecureContentStyle"`
	// Security state representing a page that displayed insecure content.
	DisplayedInsecureContentStyle Security_SecurityState `json:"displayedInsecureContentStyle"`
}

// The action to take when a certificate error occurs. continue will continue processing the request and cancel will cancel the request.
type Security_CertificateErrorAction string
