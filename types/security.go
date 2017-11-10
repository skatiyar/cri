package types

type Security_CertificateId int
type Security_MixedContentType string
type Security_SecurityState string
type Security_SecurityStateExplanation struct {
	SecurityState    Security_SecurityState    `json:"securityState"`
	Summary          string                    `json:"summary"`
	Description      string                    `json:"description"`
	MixedContentType Security_MixedContentType `json:"mixedContentType"`
	Certificate      []string                  `json:"certificate"`
}
type Security_InsecureContentStatus struct {
	RanMixedContent                bool                   `json:"ranMixedContent"`
	DisplayedMixedContent          bool                   `json:"displayedMixedContent"`
	ContainedMixedForm             bool                   `json:"containedMixedForm"`
	RanContentWithCertErrors       bool                   `json:"ranContentWithCertErrors"`
	DisplayedContentWithCertErrors bool                   `json:"displayedContentWithCertErrors"`
	RanInsecureContentStyle        Security_SecurityState `json:"ranInsecureContentStyle"`
	DisplayedInsecureContentStyle  Security_SecurityState `json:"displayedInsecureContentStyle"`
}
type Security_CertificateErrorAction string
