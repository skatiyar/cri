package network

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Network struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Network {
	return &Network{conn}
}

type EnableRequest struct {
	MaxTotalBufferSize	*int	`json:"maxTotalBufferSize,omitempty"`// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	MaxResourceBufferSize	*int	`json:"maxResourceBufferSize,omitempty"`// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
}

func (obj *Network) Enable(request *EnableRequest) (err error) {
	err = obj.conn.Send("Network.enable", request, nil)
	return
}
func (obj *Network) Disable() (err error) {
	err = obj.conn.Send("Network.disable", nil, nil)
	return
}

type SetUserAgentOverrideRequest struct {
	UserAgent string `json:"userAgent"`// User agent to use.
}

func (obj *Network) SetUserAgentOverride(request *SetUserAgentOverrideRequest) (err error) {
	err = obj.conn.Send("Network.setUserAgentOverride", request, nil)
	return
}

type SetExtraHTTPHeadersRequest struct {
	Headers types.Network_Headers `json:"headers"`// Map with extra HTTP headers.
}

func (obj *Network) SetExtraHTTPHeaders(request *SetExtraHTTPHeadersRequest) (err error) {
	err = obj.conn.Send("Network.setExtraHTTPHeaders", request, nil)
	return
}

type GetResponseBodyRequest struct {
	RequestId types.Network_RequestId `json:"requestId"`// Identifier of the network request to get content for.
}

func (obj *Network) GetResponseBody(request *GetResponseBodyRequest) (response struct {
	Body		string	`json:"body"`// Response body.
	Base64Encoded	bool	`json:"base64Encoded"`// True, if content was sent as base64.
}, err error) {
	err = obj.conn.Send("Network.getResponseBody", request, &response)
	return
}

type SetBlockedURLsRequest struct {
	Urls []string `json:"urls"`// URL patterns to block. Wildcards ('*') are allowed.
}

func (obj *Network) SetBlockedURLs(request *SetBlockedURLsRequest) (err error) {
	err = obj.conn.Send("Network.setBlockedURLs", request, nil)
	return
}

type ReplayXHRRequest struct {
	RequestId types.Network_RequestId `json:"requestId"`// Identifier of XHR to replay.
}

func (obj *Network) ReplayXHR(request *ReplayXHRRequest) (err error) {
	err = obj.conn.Send("Network.replayXHR", request, nil)
	return
}
func (obj *Network) CanClearBrowserCache() (response struct {
	Result bool `json:"result"`// True if browser cache can be cleared.
}, err error) {
	err = obj.conn.Send("Network.canClearBrowserCache", nil, &response)
	return
}
func (obj *Network) ClearBrowserCache() (err error) {
	err = obj.conn.Send("Network.clearBrowserCache", nil, nil)
	return
}
func (obj *Network) CanClearBrowserCookies() (response struct {
	Result bool `json:"result"`// True if browser cookies can be cleared.
}, err error) {
	err = obj.conn.Send("Network.canClearBrowserCookies", nil, &response)
	return
}
func (obj *Network) ClearBrowserCookies() (err error) {
	err = obj.conn.Send("Network.clearBrowserCookies", nil, nil)
	return
}

type GetCookiesRequest struct {
	Urls []string `json:"urls,omitempty"`// The list of URLs for which applicable cookies will be fetched
}

func (obj *Network) GetCookies(request *GetCookiesRequest) (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`// Array of cookie objects.
}, err error) {
	err = obj.conn.Send("Network.getCookies", request, &response)
	return
}
func (obj *Network) GetAllCookies() (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`// Array of cookie objects.
}, err error) {
	err = obj.conn.Send("Network.getAllCookies", nil, &response)
	return
}

type DeleteCookiesRequest struct {
	Name	string	`json:"name"`// Name of the cookies to remove.
	Url	*string	`json:"url,omitempty"`// If specified, deletes all the cookies with the given name where domain and path match provided URL.
	Domain	*string	`json:"domain,omitempty"`// If specified, deletes only cookies with the exact domain.
	Path	*string	`json:"path,omitempty"`// If specified, deletes only cookies with the exact path.
}

func (obj *Network) DeleteCookies(request *DeleteCookiesRequest) (err error) {
	err = obj.conn.Send("Network.deleteCookies", request, nil)
	return
}

type SetCookieRequest struct {
	Name		string				`json:"name"`// Cookie name.
	Value		string				`json:"value"`// Cookie value.
	Url		*string				`json:"url,omitempty"`// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain		*string				`json:"domain,omitempty"`// Cookie domain.
	Path		*string				`json:"path,omitempty"`// Cookie path.
	Secure		*bool				`json:"secure,omitempty"`// True if cookie is secure.
	HttpOnly	*bool				`json:"httpOnly,omitempty"`// True if cookie is http-only.
	SameSite	*types.Network_CookieSameSite	`json:"sameSite,omitempty"`// Cookie SameSite type.
	Expires		*types.Network_TimeSinceEpoch	`json:"expires,omitempty"`// Cookie expiration date, session cookie if not set
}

func (obj *Network) SetCookie(request *SetCookieRequest) (response struct {
	Success bool `json:"success"`// True if successfully set cookie.
}, err error) {
	err = obj.conn.Send("Network.setCookie", request, &response)
	return
}

type SetCookiesRequest struct {
	Cookies []types.Network_CookieParam `json:"cookies"`// Cookies to be set.
}

func (obj *Network) SetCookies(request *SetCookiesRequest) (err error) {
	err = obj.conn.Send("Network.setCookies", request, nil)
	return
}
func (obj *Network) CanEmulateNetworkConditions() (response struct {
	Result bool `json:"result"`// True if emulation of network conditions is supported.
}, err error) {
	err = obj.conn.Send("Network.canEmulateNetworkConditions", nil, &response)
	return
}

type EmulateNetworkConditionsRequest struct {
	Offline			bool				`json:"offline"`// True to emulate internet disconnection.
	Latency			float32				`json:"latency"`// Minimum latency from request sent to response headers received (ms).
	DownloadThroughput	float32				`json:"downloadThroughput"`// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	UploadThroughput	float32				`json:"uploadThroughput"`// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	ConnectionType		*types.Network_ConnectionType	`json:"connectionType,omitempty"`// Connection type if known.
}

func (obj *Network) EmulateNetworkConditions(request *EmulateNetworkConditionsRequest) (err error) {
	err = obj.conn.Send("Network.emulateNetworkConditions", request, nil)
	return
}

type SetCacheDisabledRequest struct {
	CacheDisabled bool `json:"cacheDisabled"`// Cache disabled state.
}

func (obj *Network) SetCacheDisabled(request *SetCacheDisabledRequest) (err error) {
	err = obj.conn.Send("Network.setCacheDisabled", request, nil)
	return
}

type SetBypassServiceWorkerRequest struct {
	Bypass bool `json:"bypass"`// Bypass service worker and load from network.
}

func (obj *Network) SetBypassServiceWorker(request *SetBypassServiceWorkerRequest) (err error) {
	err = obj.conn.Send("Network.setBypassServiceWorker", request, nil)
	return
}

type SetDataSizeLimitsForTestRequest struct {
	MaxTotalSize	int	`json:"maxTotalSize"`// Maximum total buffer size.
	MaxResourceSize	int	`json:"maxResourceSize"`// Maximum per-resource size.
}

func (obj *Network) SetDataSizeLimitsForTest(request *SetDataSizeLimitsForTestRequest) (err error) {
	err = obj.conn.Send("Network.setDataSizeLimitsForTest", request, nil)
	return
}

type GetCertificateRequest struct {
	Origin string `json:"origin"`// Origin to get certificate for.
}

func (obj *Network) GetCertificate(request *GetCertificateRequest) (response struct {
	TableNames []string `json:"tableNames"`
}, err error) {
	err = obj.conn.Send("Network.getCertificate", request, &response)
	return
}

type SetRequestInterceptionRequest struct {
	Patterns []types.Network_RequestPattern `json:"patterns"`// Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
}

func (obj *Network) SetRequestInterception(request *SetRequestInterceptionRequest) (err error) {
	err = obj.conn.Send("Network.setRequestInterception", request, nil)
	return
}

type ContinueInterceptedRequestRequest struct {
	InterceptionId		types.Network_InterceptionId		`json:"interceptionId"`
	ErrorReason		*types.Network_ErrorReason		`json:"errorReason,omitempty"`// If set this causes the request to fail with the given reason. Passing <code>Aborted</code> for requests marked with <code>isNavigationRequest</code> also cancels the navigation. Must not be set in response to an authChallenge.
	RawResponse		*string					`json:"rawResponse,omitempty"`// If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
	Url			*string					`json:"url,omitempty"`// If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
	Method			*string					`json:"method,omitempty"`// If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
	PostData		*string					`json:"postData,omitempty"`// If set this allows postData to be set. Must not be set in response to an authChallenge.
	Headers			*types.Network_Headers			`json:"headers,omitempty"`// If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
	AuthChallengeResponse	*types.Network_AuthChallengeResponse	`json:"authChallengeResponse,omitempty"`// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
}

func (obj *Network) ContinueInterceptedRequest(request *ContinueInterceptedRequestRequest) (err error) {
	err = obj.conn.Send("Network.continueInterceptedRequest", request, nil)
	return
}
