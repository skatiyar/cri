package types

type Security_CertificateId int
type Security_MixedContentType string
type Security_SecurityState string
type Security_SecurityStateExplanation struct {
	SecurityState		Security_SecurityState		`json:"securityState"`// Security state representing the severity of the factor being explained.
	Summary			string				`json:"summary"`// Short phrase describing the type of factor.
	Description		string				`json:"description"`// Full text explanation of the factor.
	MixedContentType	Security_MixedContentType	`json:"mixedContentType"`// The type of mixed content described by the explanation.
	Certificate		[]string			`json:"certificate"`// Page certificate.
}
type Security_InsecureContentStatus struct {
	RanMixedContent			bool			`json:"ranMixedContent"`// True if the page was loaded over HTTPS and ran mixed (HTTP) content such as scripts.
	DisplayedMixedContent		bool			`json:"displayedMixedContent"`// True if the page was loaded over HTTPS and displayed mixed (HTTP) content such as images.
	ContainedMixedForm		bool			`json:"containedMixedForm"`// True if the page was loaded over HTTPS and contained a form targeting an insecure url.
	RanContentWithCertErrors	bool			`json:"ranContentWithCertErrors"`// True if the page was loaded over HTTPS without certificate errors, and ran content such as scripts that were loaded with certificate errors.
	DisplayedContentWithCertErrors	bool			`json:"displayedContentWithCertErrors"`// True if the page was loaded over HTTPS without certificate errors, and displayed content such as images that were loaded with certificate errors.
	RanInsecureContentStyle		Security_SecurityState	`json:"ranInsecureContentStyle"`// Security state representing a page that ran insecure content.
	DisplayedInsecureContentStyle	Security_SecurityState	`json:"displayedInsecureContentStyle"`// Security state representing a page that displayed insecure content.
}
type Security_CertificateErrorAction string
