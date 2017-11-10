package types

type Network_LoaderId string
type Network_RequestId string
type Network_InterceptionId string
type Network_ErrorReason string
type Network_TimeSinceEpoch float32
type Network_MonotonicTime float32
type Network_Headers struct {
}
type Network_ConnectionType string
type Network_CookieSameSite string
type Network_ResourceTiming struct {
	RequestTime       float32 `json:"requestTime"`
	ProxyStart        float32 `json:"proxyStart"`
	ProxyEnd          float32 `json:"proxyEnd"`
	DnsStart          float32 `json:"dnsStart"`
	DnsEnd            float32 `json:"dnsEnd"`
	ConnectStart      float32 `json:"connectStart"`
	ConnectEnd        float32 `json:"connectEnd"`
	SslStart          float32 `json:"sslStart"`
	SslEnd            float32 `json:"sslEnd"`
	WorkerStart       float32 `json:"workerStart"`
	WorkerReady       float32 `json:"workerReady"`
	SendStart         float32 `json:"sendStart"`
	SendEnd           float32 `json:"sendEnd"`
	PushStart         float32 `json:"pushStart"`
	PushEnd           float32 `json:"pushEnd"`
	ReceiveHeadersEnd float32 `json:"receiveHeadersEnd"`
}
type Network_ResourcePriority string
type Network_Request struct {
	Url              string                     `json:"url"`
	Method           string                     `json:"method"`
	Headers          Network_Headers            `json:"headers"`
	PostData         *string                    `json:"postData,omitempty"`
	MixedContentType *Security_MixedContentType `json:"mixedContentType,omitempty"`
	InitialPriority  Network_ResourcePriority   `json:"initialPriority"`
	ReferrerPolicy   string                     `json:"referrerPolicy"`
	IsLinkPreload    *bool                      `json:"isLinkPreload,omitempty"`
}
type Network_SignedCertificateTimestamp struct {
	Status             string                 `json:"status"`
	Origin             string                 `json:"origin"`
	LogDescription     string                 `json:"logDescription"`
	LogId              string                 `json:"logId"`
	Timestamp          Network_TimeSinceEpoch `json:"timestamp"`
	HashAlgorithm      string                 `json:"hashAlgorithm"`
	SignatureAlgorithm string                 `json:"signatureAlgorithm"`
	SignatureData      string                 `json:"signatureData"`
}
type Network_SecurityDetails struct {
	Protocol                       string                               `json:"protocol"`
	KeyExchange                    string                               `json:"keyExchange"`
	KeyExchangeGroup               *string                              `json:"keyExchangeGroup,omitempty"`
	Cipher                         string                               `json:"cipher"`
	Mac                            *string                              `json:"mac,omitempty"`
	CertificateId                  Security_CertificateId               `json:"certificateId"`
	SubjectName                    string                               `json:"subjectName"`
	SanList                        []string                             `json:"sanList"`
	Issuer                         string                               `json:"issuer"`
	ValidFrom                      Network_TimeSinceEpoch               `json:"validFrom"`
	ValidTo                        Network_TimeSinceEpoch               `json:"validTo"`
	SignedCertificateTimestampList []Network_SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
}
type Network_BlockedReason string
type Network_Response struct {
	Url                string                   `json:"url"`
	Status             int                      `json:"status"`
	StatusText         string                   `json:"statusText"`
	Headers            Network_Headers          `json:"headers"`
	HeadersText        *string                  `json:"headersText,omitempty"`
	MimeType           string                   `json:"mimeType"`
	RequestHeaders     *Network_Headers         `json:"requestHeaders,omitempty"`
	RequestHeadersText *string                  `json:"requestHeadersText,omitempty"`
	ConnectionReused   bool                     `json:"connectionReused"`
	ConnectionId       float32                  `json:"connectionId"`
	RemoteIPAddress    *string                  `json:"remoteIPAddress,omitempty"`
	RemotePort         *int                     `json:"remotePort,omitempty"`
	FromDiskCache      *bool                    `json:"fromDiskCache,omitempty"`
	FromServiceWorker  *bool                    `json:"fromServiceWorker,omitempty"`
	EncodedDataLength  float32                  `json:"encodedDataLength"`
	Timing             *Network_ResourceTiming  `json:"timing,omitempty"`
	Protocol           *string                  `json:"protocol,omitempty"`
	SecurityState      Security_SecurityState   `json:"securityState"`
	SecurityDetails    *Network_SecurityDetails `json:"securityDetails,omitempty"`
}
type Network_WebSocketRequest struct {
	Headers Network_Headers `json:"headers"`
}
type Network_WebSocketResponse struct {
	Status             int              `json:"status"`
	StatusText         string           `json:"statusText"`
	Headers            Network_Headers  `json:"headers"`
	HeadersText        *string          `json:"headersText,omitempty"`
	RequestHeaders     *Network_Headers `json:"requestHeaders,omitempty"`
	RequestHeadersText *string          `json:"requestHeadersText,omitempty"`
}
type Network_WebSocketFrame struct {
	Opcode      float32 `json:"opcode"`
	Mask        bool    `json:"mask"`
	PayloadData string  `json:"payloadData"`
}
type Network_CachedResource struct {
	Url      string            `json:"url"`
	Type     Page_ResourceType `json:"type"`
	Response *Network_Response `json:"response,omitempty"`
	BodySize float32           `json:"bodySize"`
}
type Network_Initiator struct {
	Type       string              `json:"type"`
	Stack      *Runtime_StackTrace `json:"stack,omitempty"`
	Url        *string             `json:"url,omitempty"`
	LineNumber *float32            `json:"lineNumber,omitempty"`
}
type Network_Cookie struct {
	Name     string                  `json:"name"`
	Value    string                  `json:"value"`
	Domain   string                  `json:"domain"`
	Path     string                  `json:"path"`
	Expires  float32                 `json:"expires"`
	Size     int                     `json:"size"`
	HttpOnly bool                    `json:"httpOnly"`
	Secure   bool                    `json:"secure"`
	Session  bool                    `json:"session"`
	SameSite *Network_CookieSameSite `json:"sameSite,omitempty"`
}
type Network_CookieParam struct {
	Name     string                  `json:"name"`
	Value    string                  `json:"value"`
	Url      *string                 `json:"url,omitempty"`
	Domain   *string                 `json:"domain,omitempty"`
	Path     *string                 `json:"path,omitempty"`
	Secure   *bool                   `json:"secure,omitempty"`
	HttpOnly *bool                   `json:"httpOnly,omitempty"`
	SameSite *Network_CookieSameSite `json:"sameSite,omitempty"`
	Expires  *Network_TimeSinceEpoch `json:"expires,omitempty"`
}
type Network_AuthChallenge struct {
	Source *string `json:"source,omitempty"`
	Origin string  `json:"origin"`
	Scheme string  `json:"scheme"`
	Realm  string  `json:"realm"`
}
type Network_AuthChallengeResponse struct {
	Response string  `json:"response"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}
type Network_RequestPattern struct {
	UrlPattern   *string            `json:"urlPattern,omitempty"`
	ResourceType *Page_ResourceType `json:"resourceType,omitempty"`
}
