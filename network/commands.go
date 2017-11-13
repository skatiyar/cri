/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Network domain allows tracking network activities of the page. It exposes information about http, file, data and other requests and responses, their headers, bodies, timing, etc.
package network

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Network struct {
	conn cri.Connector
}

// New creates a Network instance
func New(conn cri.Connector) *Network {
	return &Network{conn}
}

type EnableRequest struct {
	// Buffer size in bytes to use when preserving network payloads (XHRs, etc).
	// NOTE Experimental
	MaxTotalBufferSize *int `json:"maxTotalBufferSize,omitempty"`
	// Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
	// NOTE Experimental
	MaxResourceBufferSize *int `json:"maxResourceBufferSize,omitempty"`
}

// Enables network tracking, network events will now be delivered to the client.
func (obj *Network) Enable(request *EnableRequest) (err error) {
	err = obj.conn.Send("Network.enable", request, nil)
	return
}

// Disables network tracking, prevents network events from being sent to the client.
func (obj *Network) Disable() (err error) {
	err = obj.conn.Send("Network.disable", nil, nil)
	return
}

type SetUserAgentOverrideRequest struct {
	// User agent to use.
	UserAgent string `json:"userAgent"`
}

// Allows overriding user agent with the given string.
func (obj *Network) SetUserAgentOverride(request *SetUserAgentOverrideRequest) (err error) {
	err = obj.conn.Send("Network.setUserAgentOverride", request, nil)
	return
}

type SetExtraHTTPHeadersRequest struct {
	// Map with extra HTTP headers.
	Headers types.Network_Headers `json:"headers"`
}

// Specifies whether to always send extra HTTP headers with the requests from this page.
func (obj *Network) SetExtraHTTPHeaders(request *SetExtraHTTPHeadersRequest) (err error) {
	err = obj.conn.Send("Network.setExtraHTTPHeaders", request, nil)
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
	err = obj.conn.Send("Network.getResponseBody", request, &response)
	return
}

type SetBlockedURLsRequest struct {
	// URL patterns to block. Wildcards ('*') are allowed.
	Urls []string `json:"urls"`
}

// Blocks URLs from loading.
func (obj *Network) SetBlockedURLs(request *SetBlockedURLsRequest) (err error) {
	err = obj.conn.Send("Network.setBlockedURLs", request, nil)
	return
}

type ReplayXHRRequest struct {
	// Identifier of XHR to replay.
	RequestId types.Network_RequestId `json:"requestId"`
}

// This method sends a new XMLHttpRequest which is identical to the original one. The following parameters should be identical: method, url, async, request body, extra headers, withCredentials attribute, user, password.
func (obj *Network) ReplayXHR(request *ReplayXHRRequest) (err error) {
	err = obj.conn.Send("Network.replayXHR", request, nil)
	return
}

type CanClearBrowserCacheResponse struct {
	// True if browser cache can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cache is supported.
func (obj *Network) CanClearBrowserCache() (response CanClearBrowserCacheResponse, err error) {
	err = obj.conn.Send("Network.canClearBrowserCache", nil, &response)
	return
}

// Clears browser cache.
func (obj *Network) ClearBrowserCache() (err error) {
	err = obj.conn.Send("Network.clearBrowserCache", nil, nil)
	return
}

type CanClearBrowserCookiesResponse struct {
	// True if browser cookies can be cleared.
	Result bool `json:"result"`
}

// Tells whether clearing browser cookies is supported.
func (obj *Network) CanClearBrowserCookies() (response CanClearBrowserCookiesResponse, err error) {
	err = obj.conn.Send("Network.canClearBrowserCookies", nil, &response)
	return
}

// Clears browser cookies.
func (obj *Network) ClearBrowserCookies() (err error) {
	err = obj.conn.Send("Network.clearBrowserCookies", nil, nil)
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

// Returns all browser cookies for the current URL. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
func (obj *Network) GetCookies(request *GetCookiesRequest) (response GetCookiesResponse, err error) {
	err = obj.conn.Send("Network.getCookies", request, &response)
	return
}

type GetAllCookiesResponse struct {
	// Array of cookie objects.
	Cookies []types.Network_Cookie `json:"cookies"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
func (obj *Network) GetAllCookies() (response GetAllCookiesResponse, err error) {
	err = obj.conn.Send("Network.getAllCookies", nil, &response)
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
	err = obj.conn.Send("Network.deleteCookies", request, nil)
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
	err = obj.conn.Send("Network.setCookie", request, &response)
	return
}

type SetCookiesRequest struct {
	// Cookies to be set.
	Cookies []types.Network_CookieParam `json:"cookies"`
}

// Sets given cookies.
func (obj *Network) SetCookies(request *SetCookiesRequest) (err error) {
	err = obj.conn.Send("Network.setCookies", request, nil)
	return
}

type CanEmulateNetworkConditionsResponse struct {
	// True if emulation of network conditions is supported.
	Result bool `json:"result"`
}

// Tells whether emulation of network conditions is supported.
func (obj *Network) CanEmulateNetworkConditions() (response CanEmulateNetworkConditionsResponse, err error) {
	err = obj.conn.Send("Network.canEmulateNetworkConditions", nil, &response)
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
	err = obj.conn.Send("Network.emulateNetworkConditions", request, nil)
	return
}

type SetCacheDisabledRequest struct {
	// Cache disabled state.
	CacheDisabled bool `json:"cacheDisabled"`
}

// Toggles ignoring cache for each request. If <code>true</code>, cache will not be used.
func (obj *Network) SetCacheDisabled(request *SetCacheDisabledRequest) (err error) {
	err = obj.conn.Send("Network.setCacheDisabled", request, nil)
	return
}

type SetBypassServiceWorkerRequest struct {
	// Bypass service worker and load from network.
	Bypass bool `json:"bypass"`
}

// Toggles ignoring of service worker for each request.
func (obj *Network) SetBypassServiceWorker(request *SetBypassServiceWorkerRequest) (err error) {
	err = obj.conn.Send("Network.setBypassServiceWorker", request, nil)
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
	err = obj.conn.Send("Network.setDataSizeLimitsForTest", request, nil)
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
	err = obj.conn.Send("Network.getCertificate", request, &response)
	return
}

type SetRequestInterceptionRequest struct {
	// Requests matching any of these patterns will be forwarded and wait for the corresponding continueInterceptedRequest call.
	Patterns []types.Network_RequestPattern `json:"patterns"`
}

// Sets the requests to intercept that match a the provided patterns and optionally resource types.
func (obj *Network) SetRequestInterception(request *SetRequestInterceptionRequest) (err error) {
	err = obj.conn.Send("Network.setRequestInterception", request, nil)
	return
}

type ContinueInterceptedRequestRequest struct {
	InterceptionId types.Network_InterceptionId `json:"interceptionId"`
	// If set this causes the request to fail with the given reason. Passing <code>Aborted</code> for requests marked with <code>isNavigationRequest</code> also cancels the navigation. Must not be set in response to an authChallenge.
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
	err = obj.conn.Send("Network.continueInterceptedRequest", request, nil)
	return
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
func (obj *Network) ResourceChangedPriority(fn func(params *ResourceChangedPriorityParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.resourceChangedPriority", closeChn)
	go func() {
		for {
			params := ResourceChangedPriorityParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

type RequestWillBeSentParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched form worker.
	LoaderId types.Network_LoaderId `json:"loaderId"`
	// URL of the document this request is loaded for.
	DocumentURL string `json:"documentURL"`
	// Request data.
	Request types.Network_Request `json:"request"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Timestamp.
	// NOTE Experimental
	WallTime types.Network_TimeSinceEpoch `json:"wallTime"`
	// Request initiator.
	Initiator types.Network_Initiator `json:"initiator"`
	// Redirect response data.
	RedirectResponse *types.Network_Response `json:"redirectResponse,omitempty"`
	// Type of this resource.
	// NOTE Experimental
	Type *types.Page_ResourceType `json:"type,omitempty"`
	// Frame identifier.
	// NOTE Experimental
	FrameId *types.Page_FrameId `json:"frameId,omitempty"`
}

// Fired when page is about to send HTTP request.
func (obj *Network) RequestWillBeSent(fn func(params *RequestWillBeSentParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.requestWillBeSent", closeChn)
	go func() {
		for {
			params := RequestWillBeSentParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

type RequestServedFromCacheParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
}

// Fired if request ended up loading from cache.
func (obj *Network) RequestServedFromCache(fn func(params *RequestServedFromCacheParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.requestServedFromCache", closeChn)
	go func() {
		for {
			params := RequestServedFromCacheParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

type ResponseReceivedParams struct {
	// Request identifier.
	RequestId types.Network_RequestId `json:"requestId"`
	// Loader identifier. Empty string if the request is fetched form worker.
	LoaderId types.Network_LoaderId `json:"loaderId"`
	// Timestamp.
	Timestamp types.Network_MonotonicTime `json:"timestamp"`
	// Resource type.
	Type types.Page_ResourceType `json:"type"`
	// Response data.
	Response types.Network_Response `json:"response"`
	// Frame identifier.
	// NOTE Experimental
	FrameId *types.Page_FrameId `json:"frameId,omitempty"`
}

// Fired when HTTP response is available.
func (obj *Network) ResponseReceived(fn func(params *ResponseReceivedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.responseReceived", closeChn)
	go func() {
		for {
			params := ResponseReceivedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
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
func (obj *Network) DataReceived(fn func(params *DataReceivedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.dataReceived", closeChn)
	go func() {
		for {
			params := DataReceivedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
}

// Fired when HTTP request has finished loading.
func (obj *Network) LoadingFinished(fn func(params *LoadingFinishedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.loadingFinished", closeChn)
	go func() {
		for {
			params := LoadingFinishedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
	// NOTE Experimental
	BlockedReason *types.Network_BlockedReason `json:"blockedReason,omitempty"`
}

// Fired when HTTP request has failed to load.
func (obj *Network) LoadingFailed(fn func(params *LoadingFailedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.loadingFailed", closeChn)
	go func() {
		for {
			params := LoadingFailedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
	// NOTE Experimental
	WallTime types.Network_TimeSinceEpoch `json:"wallTime"`
	// WebSocket request data.
	Request types.Network_WebSocketRequest `json:"request"`
}

// Fired when WebSocket is about to initiate handshake.
// NOTE Experimental
func (obj *Network) WebSocketWillSendHandshakeRequest(fn func(params *WebSocketWillSendHandshakeRequestParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketWillSendHandshakeRequest", closeChn)
	go func() {
		for {
			params := WebSocketWillSendHandshakeRequestParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketHandshakeResponseReceived(fn func(params *WebSocketHandshakeResponseReceivedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketHandshakeResponseReceived", closeChn)
	go func() {
		for {
			params := WebSocketHandshakeResponseReceivedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketCreated(fn func(params *WebSocketCreatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketCreated", closeChn)
	go func() {
		for {
			params := WebSocketCreatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketClosed(fn func(params *WebSocketClosedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketClosed", closeChn)
	go func() {
		for {
			params := WebSocketClosedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketFrameReceived(fn func(params *WebSocketFrameReceivedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketFrameReceived", closeChn)
	go func() {
		for {
			params := WebSocketFrameReceivedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketFrameError(fn func(params *WebSocketFrameErrorParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketFrameError", closeChn)
	go func() {
		for {
			params := WebSocketFrameErrorParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) WebSocketFrameSent(fn func(params *WebSocketFrameSentParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.webSocketFrameSent", closeChn)
	go func() {
		for {
			params := WebSocketFrameSentParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
// NOTE Experimental
func (obj *Network) EventSourceMessageReceived(fn func(params *EventSourceMessageReceivedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.eventSourceMessageReceived", closeChn)
	go func() {
		for {
			params := EventSourceMessageReceivedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
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
	// HTTP response headers, only sent if a redirect was intercepted.
	RedirectHeaders *types.Network_Headers `json:"redirectHeaders,omitempty"`
	// HTTP response code, only sent if a redirect was intercepted.
	RedirectStatusCode *int `json:"redirectStatusCode,omitempty"`
	// Redirect location, only sent if a redirect was intercepted.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	AuthChallenge *types.Network_AuthChallenge `json:"authChallenge,omitempty"`
}

// Details of an intercepted HTTP request, which must be either allowed, blocked, modified or mocked.
// NOTE Experimental
func (obj *Network) RequestIntercepted(fn func(params *RequestInterceptedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Network.requestIntercepted", closeChn)
	go func() {
		for {
			params := RequestInterceptedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
