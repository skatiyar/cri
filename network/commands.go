/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package network provides commands and events for Network domain.
package network

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Network domain
const (
	CanClearBrowserCache                    = "Network.canClearBrowserCache"
	CanClearBrowserCookies                  = "Network.canClearBrowserCookies"
	CanEmulateNetworkConditions             = "Network.canEmulateNetworkConditions"
	ClearBrowserCache                       = "Network.clearBrowserCache"
	ClearBrowserCookies                     = "Network.clearBrowserCookies"
	ContinueInterceptedRequest              = "Network.continueInterceptedRequest"
	DeleteCookies                           = "Network.deleteCookies"
	Disable                                 = "Network.disable"
	EmulateNetworkConditions                = "Network.emulateNetworkConditions"
	Enable                                  = "Network.enable"
	GetAllCookies                           = "Network.getAllCookies"
	GetCertificate                          = "Network.getCertificate"
	GetCookies                              = "Network.getCookies"
	GetResponseBody                         = "Network.getResponseBody"
	GetRequestPostData                      = "Network.getRequestPostData"
	GetResponseBodyForInterception          = "Network.getResponseBodyForInterception"
	TakeResponseBodyForInterceptionAsStream = "Network.takeResponseBodyForInterceptionAsStream"
	ReplayXHR                               = "Network.replayXHR"
	SearchInResponseBody                    = "Network.searchInResponseBody"
	SetBlockedURLs                          = "Network.setBlockedURLs"
	SetBypassServiceWorker                  = "Network.setBypassServiceWorker"
	SetCacheDisabled                        = "Network.setCacheDisabled"
	SetCookie                               = "Network.setCookie"
	SetCookies                              = "Network.setCookies"
	SetDataSizeLimitsForTest                = "Network.setDataSizeLimitsForTest"
	SetExtraHTTPHeaders                     = "Network.setExtraHTTPHeaders"
	SetRequestInterception                  = "Network.setRequestInterception"
	SetUserAgentOverride                    = "Network.setUserAgentOverride"
)

// List of events in Network domain
const (
	DataReceived                       = "Network.dataReceived"
	EventSourceMessageReceived         = "Network.eventSourceMessageReceived"
	LoadingFailed                      = "Network.loadingFailed"
	LoadingFinished                    = "Network.loadingFinished"
	RequestIntercepted                 = "Network.requestIntercepted"
	RequestServedFromCache             = "Network.requestServedFromCache"
	RequestWillBeSent                  = "Network.requestWillBeSent"
	ResourceChangedPriority            = "Network.resourceChangedPriority"
	SignedExchangeReceived             = "Network.signedExchangeReceived"
	ResponseReceived                   = "Network.responseReceived"
	WebSocketClosed                    = "Network.webSocketClosed"
	WebSocketCreated                   = "Network.webSocketCreated"
	WebSocketFrameError                = "Network.webSocketFrameError"
	WebSocketFrameReceived             = "Network.webSocketFrameReceived"
	WebSocketFrameSent                 = "Network.webSocketFrameSent"
	WebSocketHandshakeResponseReceived = "Network.webSocketHandshakeResponseReceived"
	WebSocketWillSendHandshakeRequest  = "Network.webSocketWillSendHandshakeRequest"
)

// Network domain allows tracking network activities of the page. It exposes information about http, file, data and other requests and responses, their headers, bodies, timing, etc.
type Network struct {
	conn cri.Connector
}

// New creates a Network instance
func New(conn cri.Connector) *Network {
	return &Network{conn}
}

type CanClearBrowserCacheResponse struct {
	// True if browser cache can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cache is supported.
func (obj *Network) CanClearBrowserCache() (response CanClearBrowserCacheResponse, err error) {
	err = obj.conn.Send(CanClearBrowserCache, nil, &response)
	return
}

type CanClearBrowserCookiesResponse struct {
	// True if browser cookies can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cookies is supported.
func (obj *Network) CanClearBrowserCookies() (response CanClearBrowserCookiesResponse, err error) {
	err = obj.conn.Send(CanClearBrowserCookies, nil, &response)
	return
}

type CanEmulateNetworkConditionsResponse struct {
	// True if emulation of network conditions is supported.
	Result bool `json:"result"`
}

// Tells whether emulation of network conditions is supported.
func (obj *Network) CanEmulateNetworkConditions() (response CanEmulateNetworkConditionsResponse, err error) {
	err = obj.conn.Send(CanEmulateNetworkConditions, nil, &response)
	return
}

// Clears browser cache.
func (obj *Network) ClearBrowserCache() (err error) {
	err = obj.conn.Send(ClearBrowserCache, nil, nil)
	return
}

// Clears browser cookies.
func (obj *Network) ClearBrowserCookies() (err error) {
	err = obj.conn.Send(ClearBrowserCookies, nil, nil)
	return
}

type ContinueInterceptedRequestRequest struct {
	InterceptionId types.Network_InterceptionId `json:"interceptionId"`
	// If set this causes the request to fail with the given reason. Passing `Aborted` for requests marked with `isNavigationRequest` also cancels the navigation. Must not be set in response to an authChallenge.
	ErrorReason *types.Network_ErrorReason `json:"errorReason,omitempty"`
	// If set the requests completes using with the provided base64 encoded raw response, including HTTP status line and headers etc... Must not be set in response to an authChallenge.
	RawResponse *string `json:"rawResponse,omitempty"`
	// If set the request url will be modified in a way that's not observable by page. Must not be set in response to an authChallenge.
	Url *string `json:"url,omitempty"`
	// If set this allows the request method to be overridden. Must not be set in response to an authChallenge.
	Method *string `json:"method,omitempty"`
	// If set this allows postData to be set. Must not be set in response to an authChallenge.
	PostData *string `json:"postData,omitempty"`
	// If set this allows the request headers to be changed. Must not be set in response to an authChallenge.
	Headers *types.Network_Headers `json:"headers,omitempty"`
	// Response to a requestIntercepted with an authChallenge. Must not be set otherwise.
	AuthChallengeResponse *types.Network_AuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

// Response to Network.requestIntercepted which either modifies the request to continue with any modifications, or blocks it, or completes it with the provided response bytes. If a network fetch occurs as a result which encounters a redirect an additional Network.requestIntercepted event will be sent with the same InterceptionId.
func (obj *Network) ContinueInterceptedRequest(request *ContinueInterceptedRequestRequest) (err error) {
	err = obj.conn.Send(ContinueInterceptedRequest, request, nil)
	return
}

type DeleteCookiesRequest struct {
	// Name of the cookies to remove.
	Name string `json:"name"`
	// If specified, deletes all the cookies with the given name where domain and path match provided URL.
	Url *string `json:"url,omitempty"`
	// If specified, deletes only cookies with the exact domain.
	Domain *string `json:"domain,omitempty"`
	// If specified, deletes only cookies with the exact path.
	Path *string `json:"path,omitempty"`
}

// Deletes browser cookies with matching name and url or domain/path pair.
func (obj *Network) DeleteCookies(request *DeleteCookiesRequest) (err error) {
	err = obj.conn.Send(DeleteCookies, request, nil)
	return
}

// Disables network tracking, prevents network events from being sent to the client.
func (obj *Network) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type EmulateNetworkConditionsRequest struct {
	// True to emulate internet disconnection.
	Offline bool `json:"offline"`
	// Minimum latency from request sent to response headers received (ms).
	Latency float32 `json:"latency"`
	// Maximal aggregated download throughput (bytes/sec). -1 disables download throttling.
	DownloadThroughput float32 `json:"downloadThroughput"`
	// Maximal aggregated upload throughput (bytes/sec).  -1 disables upload throttling.
	UploadThroughput float32 `json:"uploadThroughput"`
	// Connection type if known.
	ConnectionType *types.Network_ConnectionType `json:"connectionType,omitempty"`
}

// Activates emulation of network conditions.
func (obj *Network) EmulateNetworkConditions(request *EmulateNetworkConditionsRequest) (err error) {
	err = obj.conn.Send(EmulateNetworkConditions, request, nil)
	return
}

type EnableRequest struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	// NOTE Experimental
	MaxTotalBufferSize *int `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	// NOTE Experimental
	MaxResourceBufferSize *int `json:"maxResourceBufferSize,omitempty"`
	// Longest post body size (in bytes) that would be included in requestWillBeSent notification
	MaxPostDataSize *int `json:"maxPostDataSize,omitempty"`
}

// Enables network tracking, network events will now be delivered to the client.
func (obj *Network) Enable(request *EnableRequest) (err error) {
	err = obj.conn.Send(Enable, request, nil)
	return
}

type GetAllCookiesResponse struct {
	// Array of cookie objects.
	Cookies []types.Network_Cookie `json:"cookies"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the `cookies` field.
func (obj *Network) GetAllCookies() (response GetAllCookiesResponse, err error) {
	err = obj.conn.Send(GetAllCookies, nil, &response)
	return
}

type GetCertificateRequest struct {
	// Origin to get certificate for.
	Origin string `json:"origin"`
}

type GetCertificateResponse struct {
	TableNames []string `json:"tableNames"`
}

// Returns the DER-encoded certificate.
func (obj *Network) GetCertificate(request *GetCertificateRequest) (response GetCertificateResponse, err error) {
	err = obj.conn.Send(GetCertificate, request, &response)
	return
}

type GetCookiesRequest struct {
	// The list of URLs for which applicable cookies will be fetched
	Urls []string `json:"urls,omitempty"`
}

type GetCookiesResponse struct {
	// Array of cookie objects.
	Cookies []types.Network_Cookie `json:"cookies"`
}

// Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the `cookies` field.
func (obj *Network) GetCookies(request *GetCookiesRequest) (response GetCookiesResponse, err error) {
	err = obj.conn.Send(GetCookies, request, &response)
	return
}

type GetResponseBodyRequest struct {
	// Identifier of the network request to get content for.
	RequestId types.Network_RequestId `json:"requestId"`
}

type GetResponseBodyResponse struct {
	// Response body.
	Body string `json:"body"`
	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Returns content served for the given request.
func (obj *Network) GetResponseBody(request *GetResponseBodyRequest) (response GetResponseBodyResponse, err error) {
	err = obj.conn.Send(GetResponseBody, request, &response)
	return
}

type GetRequestPostDataRequest struct {
	// Identifier of the network request to get content for.
	RequestId types.Network_RequestId `json:"requestId"`
}

type GetRequestPostDataResponse struct {
	// Base64-encoded request body.
	PostData string `json:"postData"`
}

// Returns post data sent with the request. Returns an error when no data was sent with the request.
func (obj *Network) GetRequestPostData(request *GetRequestPostDataRequest) (response GetRequestPostDataResponse, err error) {
	err = obj.conn.Send(GetRequestPostData, request, &response)
	return
}

type GetResponseBodyForInterceptionRequest struct {
	// Identifier for the intercepted request to get body for.
	InterceptionId types.Network_InterceptionId `json:"interceptionId"`
}

type GetResponseBodyForInterceptionResponse struct {
	// Response body.
	Body string `json:"body"`
	// True, if content was sent as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Returns content served for the given currently intercepted request.
func (obj *Network) GetResponseBodyForInterception(request *GetResponseBodyForInterceptionRequest) (response GetResponseBodyForInterceptionResponse, err error) {
	err = obj.conn.Send(GetResponseBodyForInterception, request, &response)
	return
}

type TakeResponseBodyForInterceptionAsStreamRequest struct {
	InterceptionId types.Network_InterceptionId `json:"interceptionId"`
}

type TakeResponseBodyForInterceptionAsStreamResponse struct {
	Stream types.IO_StreamHandle `json:"stream"`
}

// Returns a handle to the stream representing the response body. Note that after this command, the intercepted request can't be continued as is -- you either need to cancel it or to provide the response body. The stream only supports sequential read, IO.read will fail if the position is specified.
func (obj *Network) TakeResponseBodyForInterceptionAsStream(request *TakeResponseBodyForInterceptionAsStreamRequest) (response TakeResponseBodyForInterceptionAsStreamResponse, err error) {
	err = obj.conn.Send(TakeResponseBodyForInterceptionAsStream, request, &response)
	return
}

type ReplayXHRRequest struct {
	// Identifier of XHR to replay.
	RequestId types.Network_RequestId `json:"requestId"`
}

// This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
func (obj *Network) ReplayXHR(request *ReplayXHRRequest) (err error) {
	err = obj.conn.Send(ReplayXHR, request, nil)
	return
}

type SearchInResponseBodyRequest struct {
	// Identifier of the network response to search.
	RequestId types.Network_RequestId `json:"requestId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex *bool `json:"isRegex,omitempty"`
}

type SearchInResponseBodyResponse struct {
	// List of search matches.
	Result []types.Debugger_SearchMatch `json:"result"`
}

// Searches for given string in response content.
func (obj *Network) SearchInResponseBody(request *SearchInResponseBodyRequest) (response SearchInResponseBodyResponse, err error) {
	err = obj.conn.Send(SearchInResponseBody, request, &response)
	return
}

type SetBlockedURLsRequest struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls []string `json:"urls"`
}

// Blocks URLs from loading.
func (obj *Network) SetBlockedURLs(request *SetBlockedURLsRequest) (err error) {
	err = obj.conn.Send(SetBlockedURLs, request, nil)
	return
}

type SetBypassServiceWorkerRequest struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

// Toggles ignoring of service worker for each request.
func (obj *Network) SetBypassServiceWorker(request *SetBypassServiceWorkerRequest) (err error) {
	err = obj.conn.Send(SetBypassServiceWorker, request, nil)
	return
}

type SetCacheDisabledRequest struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// Toggles ignoring cache for each request. If `true`, cache will not be used.
func (obj *Network) SetCacheDisabled(request *SetCacheDisabledRequest) (err error) {
	err = obj.conn.Send(SetCacheDisabled, request, nil)
	return
}

type SetCookieRequest struct {
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
	SameSite *types.Network_CookieSameSite `json:"sameSite,omitempty"`
	// Cookie expiration date, session cookie if not set
	Expires *types.Network_TimeSinceEpoch `json:"expires,omitempty"`
}

type SetCookieResponse struct {
	// True if successfully set cookie.
	Success bool `json:"success"`
}

// Sets a cookie with the given cookie data; may overwrite equivalent cookies if they exist.
func (obj *Network) SetCookie(request *SetCookieRequest) (response SetCookieResponse, err error) {
	err = obj.conn.Send(SetCookie, request, &response)
	return
}

type SetCookiesRequest struct {
	// Cookies to be set.
	Cookies []types.Network_CookieParam `json:"cookies"`
}

// Sets given cookies.
func (obj *Network) SetCookies(request *SetCookiesRequest) (err error) {
	err = obj.conn.Send(SetCookies, request, nil)
	return
}

type SetDataSizeLimitsForTestRequest struct {
	// Maximum total buffer size.
	MaxTotalSize int `json:"maxTotalSize"`
	// Maximum per-resource size.
	MaxResourceSize int `json:"maxResourceSize"`
}

// For testing.
func (obj *Network) SetDataSizeLimitsForTest(request *SetDataSizeLimitsForTestRequest) (err error) {
	err = obj.conn.Send(SetDataSizeLimitsForTest, request, nil)
	return
}

type SetExtraHTTPHeadersRequest struct {
	// Map with extra HTTP headers.
	Headers types.Network_Headers `json:"headers"`
}

// Specifies whether to always send extra HTTP headers with the requests from this page.
func (obj *Network) SetExtraHTTPHeaders(request *SetExtraHTTPHeadersRequest) (err error) {
	err = obj.conn.Send(SetExtraHTTPHeaders, request, nil)
	return
}

type SetRequestInterceptionRequest struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
	Patterns []types.Network_RequestPattern `json:"patterns"`
}

// Sets the requests to intercept that match a the provided patterns and optionally resource types.
func (obj *Network) SetRequestInterception(request *SetRequestInterceptionRequest) (err error) {
	err = obj.conn.Send(SetRequestInterception, request, nil)
	return
}

type SetUserAgentOverrideRequest struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
	// Browser langugage to emulate.
	AcceptLanguage *string `json:"acceptLanguage,omitempty"`
	// The platform navigator.platform should return.
	Platform *string `json:"platform,omitempty"`
}

// Allows overriding user agent with the given string.
func (obj *Network) SetUserAgentOverride(request *SetUserAgentOverrideRequest) (err error) {
	err = obj.conn.Send(SetUserAgentOverride, request, nil)
	return
}

type DataReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Data chunk length.
	DataLength int `json:"dataLength"`
	// Actual bytes received (might be less than dataLength for compressed encodings).
	EncodedDataLength int `json:"encodedDataLength"`
}

// Fired when data chunk was received over the network.
func (obj *Network) DataReceived(fn func(event string, params DataReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(DataReceived)
	go func() {
		defer closer()
		for {
			var params DataReceivedParams
			if !fn(DataReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type EventSourceMessageReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Message type.
	EventName string `json:"eventName"`
	// Message identifier.
	EventId string `json:"eventId"`
	// Message content.
	Data string `json:"data"`
}

// Fired when EventSource message is received.
func (obj *Network) EventSourceMessageReceived(fn func(event string, params EventSourceMessageReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(EventSourceMessageReceived)
	go func() {
		defer closer()
		for {
			var params EventSourceMessageReceivedParams
			if !fn(EventSourceMessageReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type LoadingFailedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Resource type.
	Type types.Page_ResourceType `json:"type"`
	// User friendly error message.
	ErrorText string `json:"errorText"`
	// True if loading was canceled.
	Canceled *bool `json:"canceled,omitempty"`
	// The reason why loading was blocked, if any.
	BlockedReason *types.Network_BlockedReason `json:"blockedReason,omitempty"`
}

// Fired when HTTP request has failed to load.
func (obj *Network) LoadingFailed(fn func(event string, params LoadingFailedParams, err error) bool) {
	listen, closer := obj.conn.On(LoadingFailed)
	go func() {
		defer closer()
		for {
			var params LoadingFailedParams
			if !fn(LoadingFailed, params, listen(&params)) {
				return
			}
		}
	}()
}

type LoadingFinishedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Total number of bytes received for this request.
	EncodedDataLength float32 `json:"encodedDataLength"`
	// Set when 1) response was blocked by Cross-Origin Read Blocking and also 2) this needs to be reported to the DevTools console.
	ShouldReportCorbBlocking *bool `json:"shouldReportCorbBlocking,omitempty"`
}

// Fired when HTTP request has finished loading.
func (obj *Network) LoadingFinished(fn func(event string, params LoadingFinishedParams, err error) bool) {
	listen, closer := obj.conn.On(LoadingFinished)
	go func() {
		defer closer()
		for {
			var params LoadingFinishedParams
			if !fn(LoadingFinished, params, listen(&params)) {
				return
			}
		}
	}()
}

type RequestInterceptedParams struct {
	// Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	InterceptionId types.Network_InterceptionId `json:"interceptionId"`
	Request        types.Network_Request        `json:"request"`
	// The id of the frame that initiated the request.
	FrameId types.Page_FrameId `json:"frameId"`
	// How the requested resource will be used.
	ResourceType types.Page_ResourceType `json:"resourceType"`
	// Whether this is a navigation request, which can abort the navigation completely.
	IsNavigationRequest bool `json:"isNavigationRequest"`
	// Set if the request is a navigation that will result in a download. Only present after response is received from the server (i.e. HeadersReceived stage).
	IsDownload *bool `json:"isDownload,omitempty"`
	// Redirect location, only sent if a redirect was intercepted.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge *types.Network_AuthChallenge `json:"authChallenge,omitempty"`
	// Response error if intercepted at response stage or if redirect occurred while intercepting request.
	ResponseErrorReason *types.Network_ErrorReason `json:"responseErrorReason,omitempty"`
	// Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseStatusCode *int `json:"responseStatusCode,omitempty"`
	// Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders *types.Network_Headers `json:"responseHeaders,omitempty"`
}

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or mocked.
// NOTE Experimental
func (obj *Network) RequestIntercepted(fn func(event string, params RequestInterceptedParams, err error) bool) {
	listen, closer := obj.conn.On(RequestIntercepted)
	go func() {
		defer closer()
		for {
			var params RequestInterceptedParams
			if !fn(RequestIntercepted, params, listen(&params)) {
				return
			}
		}
	}()
}

type RequestServedFromCacheParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
}

// Fired if request ended up loading from cache.
func (obj *Network) RequestServedFromCache(fn func(event string, params RequestServedFromCacheParams, err error) bool) {
	listen, closer := obj.conn.On(RequestServedFromCache)
	go func() {
		defer closer()
		for {
			var params RequestServedFromCacheParams
			if !fn(RequestServedFromCache, params, listen(&params)) {
				return
			}
		}
	}()
}

type RequestWillBeSentParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId types.Network_LoaderId `json:"loaderId"`
	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`
	// Request data.
	Request types.Network_Request `json:"request"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Timestamp.
	WallTime types.Network_TimeSinceEpoch `json:"wallTime"`
	// Request initiator.
	Initiator types.Network_Initiator `json:"initiator"`
	// Redirect response data.
	RedirectResponse *types.Network_Response `json:"redirectResponse,omitempty"`
	// Type of this resource.
	Type *types.Page_ResourceType `json:"type,omitempty"`
	// Frame identifier.
	FrameId *types.Page_FrameId `json:"frameId,omitempty"`
	// Whether the request is initiated by a user gesture. Defaults to false.
	HasUserGesture *bool `json:"hasUserGesture,omitempty"`
}

// Fired when page is about to send HTTP request.
func (obj *Network) RequestWillBeSent(fn func(event string, params RequestWillBeSentParams, err error) bool) {
	listen, closer := obj.conn.On(RequestWillBeSent)
	go func() {
		defer closer()
		for {
			var params RequestWillBeSentParams
			if !fn(RequestWillBeSent, params, listen(&params)) {
				return
			}
		}
	}()
}

type ResourceChangedPriorityParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// New priority
	NewPriority types.Network_ResourcePriority `json:"newPriority"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
}

// Fired when resource loading priority is changed
// NOTE Experimental
func (obj *Network) ResourceChangedPriority(fn func(event string, params ResourceChangedPriorityParams, err error) bool) {
	listen, closer := obj.conn.On(ResourceChangedPriority)
	go func() {
		defer closer()
		for {
			var params ResourceChangedPriorityParams
			if !fn(ResourceChangedPriority, params, listen(&params)) {
				return
			}
		}
	}()
}

type SignedExchangeReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Information about the signed exchange response.
	Info types.Network_SignedExchangeInfo `json:"info"`
}

// Fired when a signed exchange was received over the network
// NOTE Experimental
func (obj *Network) SignedExchangeReceived(fn func(event string, params SignedExchangeReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(SignedExchangeReceived)
	go func() {
		defer closer()
		for {
			var params SignedExchangeReceivedParams
			if !fn(SignedExchangeReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type ResponseReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched from worker.
	LoaderId types.Network_LoaderId `json:"loaderId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Resource type.
	Type types.Page_ResourceType `json:"type"`
	// Response data.
	Response types.Network_Response `json:"response"`
	// Frame identifier.
	FrameId *types.Page_FrameId `json:"frameId,omitempty"`
}

// Fired when HTTP response is available.
func (obj *Network) ResponseReceived(fn func(event string, params ResponseReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(ResponseReceived)
	go func() {
		defer closer()
		for {
			var params ResponseReceivedParams
			if !fn(ResponseReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketClosedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
}

// Fired when WebSocket is closed.
func (obj *Network) WebSocketClosed(fn func(event string, params WebSocketClosedParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketClosed)
	go func() {
		defer closer()
		for {
			var params WebSocketClosedParams
			if !fn(WebSocketClosed, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketCreatedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// WebSocket request URL.
	Url string `json:"url"`
	// Request initiator.
	Initiator *types.Network_Initiator `json:"initiator,omitempty"`
}

// Fired upon WebSocket creation.
func (obj *Network) WebSocketCreated(fn func(event string, params WebSocketCreatedParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketCreated)
	go func() {
		defer closer()
		for {
			var params WebSocketCreatedParams
			if !fn(WebSocketCreated, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketFrameErrorParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// WebSocket frame error message.
	ErrorMessage string `json:"errorMessage"`
}

// Fired when WebSocket frame error occurs.
func (obj *Network) WebSocketFrameError(fn func(event string, params WebSocketFrameErrorParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketFrameError)
	go func() {
		defer closer()
		for {
			var params WebSocketFrameErrorParams
			if !fn(WebSocketFrameError, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketFrameReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// WebSocket response data.
	Response types.Network_WebSocketFrame `json:"response"`
}

// Fired when WebSocket frame is received.
func (obj *Network) WebSocketFrameReceived(fn func(event string, params WebSocketFrameReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketFrameReceived)
	go func() {
		defer closer()
		for {
			var params WebSocketFrameReceivedParams
			if !fn(WebSocketFrameReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketFrameSentParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// WebSocket response data.
	Response types.Network_WebSocketFrame `json:"response"`
}

// Fired when WebSocket frame is sent.
func (obj *Network) WebSocketFrameSent(fn func(event string, params WebSocketFrameSentParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketFrameSent)
	go func() {
		defer closer()
		for {
			var params WebSocketFrameSentParams
			if !fn(WebSocketFrameSent, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketHandshakeResponseReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// WebSocket response data.
	Response types.Network_WebSocketResponse `json:"response"`
}

// Fired when WebSocket handshake response becomes available.
func (obj *Network) WebSocketHandshakeResponseReceived(fn func(event string, params WebSocketHandshakeResponseReceivedParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketHandshakeResponseReceived)
	go func() {
		defer closer()
		for {
			var params WebSocketHandshakeResponseReceivedParams
			if !fn(WebSocketHandshakeResponseReceived, params, listen(&params)) {
				return
			}
		}
	}()
}

type WebSocketWillSendHandshakeRequestParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// UTC Timestamp.
	WallTime types.Network_TimeSinceEpoch `json:"wallTime"`
	// WebSocket request data.
	Request types.Network_WebSocketRequest `json:"request"`
}

// Fired when WebSocket is about to initiate handshake.
func (obj *Network) WebSocketWillSendHandshakeRequest(fn func(event string, params WebSocketWillSendHandshakeRequestParams, err error) bool) {
	listen, closer := obj.conn.On(WebSocketWillSendHandshakeRequest)
	go func() {
		defer closer()
		for {
			var params WebSocketWillSendHandshakeRequestParams
			if !fn(WebSocketWillSendHandshakeRequest, params, listen(&params)) {
				return
			}
		}
	}()
}
