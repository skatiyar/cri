package types

type Security_CertificateId int
type Security_MixedContentType string
type Security_SecurityState string
type Security_SecurityStateExplanation struct {
	// Security state representing the severity of the factor being explained.
	SecurityState Security_SecurityState `json:"securityState"`
	// Short phrase describing the type of factor.
	Summary string `json:"summary"`
	// Full text explanation of the factor.
	Description string `json:"description"`
	// The type of mixed content described by the explanation.
	MixedContentType Security_MixedContentType `json:"mixedContentType"`
	// Page certificate.
	Certificate []string `json:"certificate"`
}
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
type Security_CertificateErrorAction string
