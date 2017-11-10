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
	RequestTime		float32	`json:"requestTime"`// Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart		float32	`json:"proxyStart"`// Started resolving proxy.
	ProxyEnd		float32	`json:"proxyEnd"`// Finished resolving proxy.
	DnsStart		float32	`json:"dnsStart"`// Started DNS address resolve.
	DnsEnd			float32	`json:"dnsEnd"`// Finished DNS address resolve.
	ConnectStart		float32	`json:"connectStart"`// Started connecting to the remote host.
	ConnectEnd		float32	`json:"connectEnd"`// Connected to the remote host.
	SslStart		float32	`json:"sslStart"`// Started SSL handshake.
	SslEnd			float32	`json:"sslEnd"`// Finished SSL handshake.
	WorkerStart		float32	`json:"workerStart"`// Started running ServiceWorker.
	WorkerReady		float32	`json:"workerReady"`// Finished Starting ServiceWorker.
	SendStart		float32	`json:"sendStart"`// Started sending request.
	SendEnd			float32	`json:"sendEnd"`// Finished sending request.
	PushStart		float32	`json:"pushStart"`// Time the server started pushing request.
	PushEnd			float32	`json:"pushEnd"`// Time the server finished pushing request.
	ReceiveHeadersEnd	float32	`json:"receiveHeadersEnd"`// Finished receiving response headers.
}
type Network_ResourcePriority string
type Network_Request struct {
	Url			string				`json:"url"`// Request URL.
	Method			string				`json:"method"`// HTTP request method.
	Headers			Network_Headers			`json:"headers"`// HTTP request headers.
	PostData		*string				`json:"postData,omitempty"`// HTTP POST request data.
	MixedContentType	*Security_MixedContentType	`json:"mixedContentType,omitempty"`// The mixed content type of the request.
	InitialPriority		Network_ResourcePriority	`json:"initialPriority"`// Priority of the resource request at the time request is sent.
	ReferrerPolicy		string				`json:"referrerPolicy"`// The referrer policy of the request, as defined in https://www.w3.org/TR/referrer-policy/
	IsLinkPreload		*bool				`json:"isLinkPreload,omitempty"`// Whether is loaded via link preload.
}
type Network_SignedCertificateTimestamp struct {
	Status			string			`json:"status"`// Validation status.
	Origin			string			`json:"origin"`// Origin.
	LogDescription		string			`json:"logDescription"`// Log name / description.
	LogId			string			`json:"logId"`// Log ID.
	Timestamp		Network_TimeSinceEpoch	`json:"timestamp"`// Issuance date.
	HashAlgorithm		string			`json:"hashAlgorithm"`// Hash algorithm.
	SignatureAlgorithm	string			`json:"signatureAlgorithm"`// Signature algorithm.
	SignatureData		string			`json:"signatureData"`// Signature data.
}
type Network_SecurityDetails struct {
	Protocol			string					`json:"protocol"`// Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange			string					`json:"keyExchange"`// Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup		*string					`json:"keyExchangeGroup,omitempty"`// (EC)DH group used by the connection, if applicable.
	Cipher				string					`json:"cipher"`// Cipher name.
	Mac				*string					`json:"mac,omitempty"`// TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateId			Security_CertificateId			`json:"certificateId"`// Certificate ID value.
	SubjectName			string					`json:"subjectName"`// Certificate subject name.
	SanList				[]string				`json:"sanList"`// Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer				string					`json:"issuer"`// Name of the issuing CA.
	ValidFrom			Network_TimeSinceEpoch			`json:"validFrom"`// Certificate valid from date.
	ValidTo				Network_TimeSinceEpoch			`json:"validTo"`// Certificate valid to (expiration) date
	SignedCertificateTimestampList	[]Network_SignedCertificateTimestamp	`json:"signedCertificateTimestampList"`// List of signed certificate timestamps (SCTs).
}
type Network_BlockedReason string
type Network_Response struct {
	Url			string				`json:"url"`// Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status			int				`json:"status"`// HTTP response status code.
	StatusText		string				`json:"statusText"`// HTTP response status text.
	Headers			Network_Headers			`json:"headers"`// HTTP response headers.
	HeadersText		*string				`json:"headersText,omitempty"`// HTTP response headers text.
	MimeType		string				`json:"mimeType"`// Resource mimeType as determined by the browser.
	RequestHeaders		*Network_Headers		`json:"requestHeaders,omitempty"`// Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText	*string				`json:"requestHeadersText,omitempty"`// HTTP request headers text.
	ConnectionReused	bool				`json:"connectionReused"`// Specifies whether physical connection was actually reused for this request.
	ConnectionId		float32				`json:"connectionId"`// Physical connection id that was actually used for this request.
	RemoteIPAddress		*string				`json:"remoteIPAddress,omitempty"`// Remote IP address.
	RemotePort		*int				`json:"remotePort,omitempty"`// Remote port.
	FromDiskCache		*bool				`json:"fromDiskCache,omitempty"`// Specifies that the request was served from the disk cache.
	FromServiceWorker	*bool				`json:"fromServiceWorker,omitempty"`// Specifies that the request was served from the ServiceWorker.
	EncodedDataLength	float32				`json:"encodedDataLength"`// Total number of bytes received for this request so far.
	Timing			*Network_ResourceTiming		`json:"timing,omitempty"`// Timing information for the given request.
	Protocol		*string				`json:"protocol,omitempty"`// Protocol used to fetch this request.
	SecurityState		Security_SecurityState		`json:"securityState"`// Security state of the request resource.
	SecurityDetails		*Network_SecurityDetails	`json:"securityDetails,omitempty"`// Security details for the request.
}
type Network_WebSocketRequest struct {
	Headers Network_Headers `json:"headers"`// HTTP request headers.
}
type Network_WebSocketResponse struct {
	Status			int			`json:"status"`// HTTP response status code.
	StatusText		string			`json:"statusText"`// HTTP response status text.
	Headers			Network_Headers		`json:"headers"`// HTTP response headers.
	HeadersText		*string			`json:"headersText,omitempty"`// HTTP response headers text.
	RequestHeaders		*Network_Headers	`json:"requestHeaders,omitempty"`// HTTP request headers.
	RequestHeadersText	*string			`json:"requestHeadersText,omitempty"`// HTTP request headers text.
}
type Network_WebSocketFrame struct {
	Opcode		float32	`json:"opcode"`// WebSocket frame opcode.
	Mask		bool	`json:"mask"`// WebSocke frame mask.
	PayloadData	string	`json:"payloadData"`// WebSocke frame payload data.
}
type Network_CachedResource struct {
	Url		string			`json:"url"`// Resource URL. This is the url of the original network request.
	Type		Page_ResourceType	`json:"type"`// Type of this resource.
	Response	*Network_Response	`json:"response,omitempty"`// Cached response data.
	BodySize	float32			`json:"bodySize"`// Cached response body size.
}
type Network_Initiator struct {
	Type		string			`json:"type"`// Type of this initiator.
	Stack		*Runtime_StackTrace	`json:"stack,omitempty"`// Initiator JavaScript stack trace, set for Script only.
	Url		*string			`json:"url,omitempty"`// Initiator URL, set for Parser type or for Script type (when script is importing module).
	LineNumber	*float32		`json:"lineNumber,omitempty"`// Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}
type Network_Cookie struct {
	Name		string			`json:"name"`// Cookie name.
	Value		string			`json:"value"`// Cookie value.
	Domain		string			`json:"domain"`// Cookie domain.
	Path		string			`json:"path"`// Cookie path.
	Expires		float32			`json:"expires"`// Cookie expiration date as the number of seconds since the UNIX epoch.
	Size		int			`json:"size"`// Cookie size.
	HttpOnly	bool			`json:"httpOnly"`// True if cookie is http-only.
	Secure		bool			`json:"secure"`// True if cookie is secure.
	Session		bool			`json:"session"`// True in case of session cookie.
	SameSite	*Network_CookieSameSite	`json:"sameSite,omitempty"`// Cookie SameSite type.
}
type Network_CookieParam struct {
	Name		string			`json:"name"`// Cookie name.
	Value		string			`json:"value"`// Cookie value.
	Url		*string			`json:"url,omitempty"`// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain		*string			`json:"domain,omitempty"`// Cookie domain.
	Path		*string			`json:"path,omitempty"`// Cookie path.
	Secure		*bool			`json:"secure,omitempty"`// True if cookie is secure.
	HttpOnly	*bool			`json:"httpOnly,omitempty"`// True if cookie is http-only.
	SameSite	*Network_CookieSameSite	`json:"sameSite,omitempty"`// Cookie SameSite type.
	Expires		*Network_TimeSinceEpoch	`json:"expires,omitempty"`// Cookie expiration date, session cookie if not set
}
type Network_AuthChallenge struct {
	Source	*string	`json:"source,omitempty"`// Source of the authentication challenge.
	Origin	string	`json:"origin"`// Origin of the challenger.
	Scheme	string	`json:"scheme"`// The authentication scheme used, such as basic or digest
	Realm	string	`json:"realm"`// The realm of the challenge. May be empty.
}
type Network_AuthChallengeResponse struct {
	Response	string	`json:"response"`// The decision on what to do in response to the authorization challenge.  Default means deferring to the default behavior of the net stack, which will likely either the Cancel authentication or display a popup dialog box.
	Username	*string	`json:"username,omitempty"`// The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password	*string	`json:"password,omitempty"`// The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}
type Network_RequestPattern struct {
	UrlPattern	*string			`json:"urlPattern,omitempty"`// Wildcards ('*' -> zero or more, '?' -> exactly one) are allowed. Escape character is backslash. Omitting is equivalent to "*".
	ResourceType	*Page_ResourceType	`json:"resourceType,omitempty"`// If set, only requests for matching resource types will be intercepted.
}
