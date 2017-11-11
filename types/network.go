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
	// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	RequestTime float32 `json:"requestTime"`
	// Started resolving proxy.
	ProxyStart float32 `json:"proxyStart"`
	// Finished resolving proxy.
	ProxyEnd float32 `json:"proxyEnd"`
	// Started DNS address resolve.
	DnsStart float32 `json:"dnsStart"`
	// Finished DNS address resolve.
	DnsEnd float32 `json:"dnsEnd"`
	// Started connecting to the remote host.
	ConnectStart float32 `json:"connectStart"`
	// Connected to the remote host.
	ConnectEnd float32 `json:"connectEnd"`
	// Started SSL handshake.
	SslStart float32 `json:"sslStart"`
	// Finished SSL handshake.
	SslEnd float32 `json:"sslEnd"`
	// Started running ServiceWorker.
	// NOTE Experimental
	WorkerStart float32 `json:"workerStart"`
	// Finished Starting ServiceWorker.
	// NOTE Experimental
	WorkerReady float32 `json:"workerReady"`
	// Started sending request.
	SendStart float32 `json:"sendStart"`
	// Finished sending request.
	SendEnd float32 `json:"sendEnd"`
	// Time the server started pushing request.
	// NOTE Experimental
	PushStart float32 `json:"pushStart"`
	// Time the server finished pushing request.
	// NOTE Experimental
	PushEnd float32 `json:"pushEnd"`
	// Finished receiving response headers.
	ReceiveHeadersEnd float32 `json:"receiveHeadersEnd"`
}
type Network_ResourcePriority string
type Network_Request struct {
	// Request URL.
	Url string `json:"url"`
	// HTTP request method.
	Method string `json:"method"`
	// HTTP request headers.
	Headers Network_Headers `json:"headers"`
	// HTTP POST request data.
	PostData *string `json:"postData,omitempty"`
	// The mixed content type of the request.
	MixedContentType *Security_MixedContentType `json:"mixedContentType,omitempty"`
	// Priority of the resource request at the time request is sent.
	InitialPriority Network_ResourcePriority `json:"initialPriority"`
	// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	ReferrerPolicy string `json:"referrerPolicy"`
	// Whether is loaded via link preload.
	IsLinkPreload *bool `json:"isLinkPreload,omitempty"`
}
type Network_SignedCertificateTimestamp struct {
	// Validation status.
	Status string `json:"status"`
	// Origin.
	Origin string `json:"origin"`
	// Log name / description.
	LogDescription string `json:"logDescription"`
	// Log ID.
	LogId string `json:"logId"`
	// Issuance date.
	Timestamp Network_TimeSinceEpoch `json:"timestamp"`
	// Hash algorithm.
	HashAlgorithm string `json:"hashAlgorithm"`
	// Signature algorithm.
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	// Signature data.
	SignatureData string `json:"signatureData"`
}
type Network_SecurityDetails struct {
	// Protocol name (e.g. "TLS 1.2" or "QUIC").
	Protocol string `json:"protocol"`
	// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchange string `json:"keyExchange"`
	// (EC)DH group used by the connection, if applicable.
	KeyExchangeGroup *string `json:"keyExchangeGroup,omitempty"`
	// Cipher name.
	Cipher string `json:"cipher"`
	// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	Mac *string `json:"mac,omitempty"`
	// Certificate ID value.
	CertificateId Security_CertificateId `json:"certificateId"`
	// Certificate subject name.
	SubjectName string `json:"subjectName"`
	// Subject Alternative Name (SAN) DNS names and IP addresses.
	SanList []string `json:"sanList"`
	// Name of the issuing CA.
	Issuer string `json:"issuer"`
	// Certificate valid from date.
	ValidFrom Network_TimeSinceEpoch `json:"validFrom"`
	// Certificate valid to (expiration) date
	ValidTo Network_TimeSinceEpoch `json:"validTo"`
	// List of signed certificate timestamps (SCTs).
	SignedCertificateTimestampList []Network_SignedCertificateTimestamp `json:"signedCertificateTimestampList"`
}
type Network_BlockedReason string
type Network_Response struct {
	// Response URL. This URL can be different from CachedResource.url in case of redirect.
	Url string `json:"url"`
	// HTTP response status code.
	Status int `json:"status"`
	// HTTP response status text.
	StatusText string `json:"statusText"`
	// HTTP response headers.
	Headers Network_Headers `json:"headers"`
	// HTTP response headers text.
	HeadersText *string `json:"headersText,omitempty"`
	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// Refined HTTP request headers that were actually transmitted over the network.
	RequestHeaders *Network_Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text.
	RequestHeadersText *string `json:"requestHeadersText,omitempty"`
	// Specifies whether physical connection was actually reused for this request.
	ConnectionReused bool `json:"connectionReused"`
	// Physical connection id that was actually used for this request.
	ConnectionId float32 `json:"connectionId"`
	// Remote IP address.
	// NOTE Experimental
	RemoteIPAddress *string `json:"remoteIPAddress,omitempty"`
	// Remote port.
	// NOTE Experimental
	RemotePort *int `json:"remotePort,omitempty"`
	// Specifies that the request was served from the disk cache.
	FromDiskCache *bool `json:"fromDiskCache,omitempty"`
	// Specifies that the request was served from the ServiceWorker.
	FromServiceWorker *bool `json:"fromServiceWorker,omitempty"`
	// Total number of bytes received for this request so far.
	EncodedDataLength float32 `json:"encodedDataLength"`
	// Timing information for the given request.
	Timing *Network_ResourceTiming `json:"timing,omitempty"`
	// Protocol used to fetch this request.
	Protocol *string `json:"protocol,omitempty"`
	// Security state of the request resource.
	SecurityState Security_SecurityState `json:"securityState"`
	// Security details for the request.
	SecurityDetails *Network_SecurityDetails `json:"securityDetails,omitempty"`
}
type Network_WebSocketRequest struct {
	// HTTP request headers.
	Headers Network_Headers `json:"headers"`
}
type Network_WebSocketResponse struct {
	// HTTP response status code.
	Status int `json:"status"`
	// HTTP response status text.
	StatusText string `json:"statusText"`
	// HTTP response headers.
	Headers Network_Headers `json:"headers"`
	// HTTP response headers text.
	HeadersText *string `json:"headersText,omitempty"`
	// HTTP request headers.
	RequestHeaders *Network_Headers `json:"requestHeaders,omitempty"`
	// HTTP request headers text.
	RequestHeadersText *string `json:"requestHeadersText,omitempty"`
}
type Network_WebSocketFrame struct {
	// WebSocket frame opcode.
	Opcode float32 `json:"opcode"`
	// WebSocke frame mask.
	Mask bool `json:"mask"`
	// WebSocke frame payload data.
	PayloadData string `json:"payloadData"`
}
type Network_CachedResource struct {
	// Resource URL. This is the url of the original network request.
	Url string `json:"url"`
	// Type of this resource.
	Type Page_ResourceType `json:"type"`
	// Cached response data.
	Response *Network_Response `json:"response,omitempty"`
	// Cached response body size.
	BodySize float32 `json:"bodySize"`
}
type Network_Initiator struct {
	// Type of this initiator.
	Type string `json:"type"`
	// Initiator JavaScript stack trace, set for Script only.
	Stack *Runtime_StackTrace `json:"stack,omitempty"`
	// Initiator URL, set for Parser type or for Script type (when script is importing module).
	Url *string `json:"url,omitempty"`
	// Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
	LineNumber *float32 `json:"lineNumber,omitempty"`
}
type Network_Cookie struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
	// Cookie domain.
	Domain string `json:"domain"`
	// Cookie path.
	Path string `json:"path"`
	// Cookie expiration date as the number of seconds since the UNIX epoch.
	Expires float32 `json:"expires"`
	// Cookie size.
	Size int `json:"size"`
	// True if cookie is http-only.
	HttpOnly bool `json:"httpOnly"`
	// True if cookie is secure.
	Secure bool `json:"secure"`
	// True in case of session cookie.
	Session bool `json:"session"`
	// Cookie SameSite type.
	SameSite *Network_CookieSameSite `json:"sameSite,omitempty"`
}
type Network_CookieParam struct {
	// Cookie name.
	Name string `json:"name"`
	// Cookie value.
	Value string `json:"value"`
	// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Url *string `json:"url,omitempty"`
	// Cookie domain.
	Domain *string `json:"domain,omitempty"`
	// Cookie path.
	Path *string `json:"path,omitempty"`
	// True if cookie is secure.
	Secure *bool `json:"secure,omitempty"`
	// True if cookie is http-only.
	HttpOnly *bool `json:"httpOnly,omitempty"`
	// Cookie SameSite type.
	SameSite *Network_CookieSameSite `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires *Network_TimeSinceEpoch `json:"expires,omitempty"`
}
type Network_AuthChallenge struct {
	// Source of the authentication challenge.
	Source *string `json:"source,omitempty"`
	// Origin of the challenger.
	Origin string `json:"origin"`
	// The authentication scheme used, such as basic or digest
	Scheme string `json:"scheme"`
	// The realm of the challenge. May be empty.
	Realm string `json:"realm"`
}
type Network_AuthChallengeResponse struct {
	// The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Response string `json:"response"`
	// The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Username *string `json:"username,omitempty"`
	// The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password *string `json:"password,omitempty"`
}
type Network_RequestPattern struct {
	// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	UrlPattern *string `json:"urlPattern,omitempty"`
	// If set, only requests for matching resource types will be intercepted.
	ResourceType *Page_ResourceType `json:"resourceType,omitempty"`
}
