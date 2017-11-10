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
	MaxTotalBufferSize    *int `json:"maxTotalBufferSize,omitempty"`
	MaxResourceBufferSize *int `json:"maxResourceBufferSize,omitempty"`
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
	UserAgent string `json:"userAgent"`
}

func (obj *Network) SetUserAgentOverride(request *SetUserAgentOverrideRequest) (err error) {
	err = obj.conn.Send("Network.setUserAgentOverride", request, nil)
	return
}

type SetExtraHTTPHeadersRequest struct {
	Headers types.Network_Headers `json:"headers"`
}

func (obj *Network) SetExtraHTTPHeaders(request *SetExtraHTTPHeadersRequest) (err error) {
	err = obj.conn.Send("Network.setExtraHTTPHeaders", request, nil)
	return
}

type GetResponseBodyRequest struct {
	RequestId types.Network_RequestId `json:"requestId"`
}

func (obj *Network) GetResponseBody(request *GetResponseBodyRequest) (response struct {
	Body          string `json:"body"`
	Base64Encoded bool   `json:"base64Encoded"`
}, err error) {
	err = obj.conn.Send("Network.getResponseBody", request, &response)
	return
}

type SetBlockedURLsRequest struct {
	Urls []string `json:"urls"`
}

func (obj *Network) SetBlockedURLs(request *SetBlockedURLsRequest) (err error) {
	err = obj.conn.Send("Network.setBlockedURLs", request, nil)
	return
}

type ReplayXHRRequest struct {
	RequestId types.Network_RequestId `json:"requestId"`
}

func (obj *Network) ReplayXHR(request *ReplayXHRRequest) (err error) {
	err = obj.conn.Send("Network.replayXHR", request, nil)
	return
}
func (obj *Network) CanClearBrowserCache() (response struct {
	Result bool `json:"result"`
}, err error) {
	err = obj.conn.Send("Network.canClearBrowserCache", nil, &response)
	return
}
func (obj *Network) ClearBrowserCache() (err error) {
	err = obj.conn.Send("Network.clearBrowserCache", nil, nil)
	return
}
func (obj *Network) CanClearBrowserCookies() (response struct {
	Result bool `json:"result"`
}, err error) {
	err = obj.conn.Send("Network.canClearBrowserCookies", nil, &response)
	return
}
func (obj *Network) ClearBrowserCookies() (err error) {
	err = obj.conn.Send("Network.clearBrowserCookies", nil, nil)
	return
}

type GetCookiesRequest struct {
	Urls []string `json:"urls,omitempty"`
}

func (obj *Network) GetCookies(request *GetCookiesRequest) (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`
}, err error) {
	err = obj.conn.Send("Network.getCookies", request, &response)
	return
}
func (obj *Network) GetAllCookies() (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`
}, err error) {
	err = obj.conn.Send("Network.getAllCookies", nil, &response)
	return
}

type DeleteCookiesRequest struct {
	Name   string  `json:"name"`
	Url    *string `json:"url,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Path   *string `json:"path,omitempty"`
}

func (obj *Network) DeleteCookies(request *DeleteCookiesRequest) (err error) {
	err = obj.conn.Send("Network.deleteCookies", request, nil)
	return
}

type SetCookieRequest struct {
	Name     string                        `json:"name"`
	Value    string                        `json:"value"`
	Url      *string                       `json:"url,omitempty"`
	Domain   *string                       `json:"domain,omitempty"`
	Path     *string                       `json:"path,omitempty"`
	Secure   *bool                         `json:"secure,omitempty"`
	HttpOnly *bool                         `json:"httpOnly,omitempty"`
	SameSite *types.Network_CookieSameSite `json:"sameSite,omitempty"`
	Expires  *types.Network_TimeSinceEpoch `json:"expires,omitempty"`
}

func (obj *Network) SetCookie(request *SetCookieRequest) (response struct {
	Success bool `json:"success"`
}, err error) {
	err = obj.conn.Send("Network.setCookie", request, &response)
	return
}

type SetCookiesRequest struct {
	Cookies []types.Network_CookieParam `json:"cookies"`
}

func (obj *Network) SetCookies(request *SetCookiesRequest) (err error) {
	err = obj.conn.Send("Network.setCookies", request, nil)
	return
}
func (obj *Network) CanEmulateNetworkConditions() (response struct {
	Result bool `json:"result"`
}, err error) {
	err = obj.conn.Send("Network.canEmulateNetworkConditions", nil, &response)
	return
}

type EmulateNetworkConditionsRequest struct {
	Offline            bool                          `json:"offline"`
	Latency            float32                       `json:"latency"`
	DownloadThroughput float32                       `json:"downloadThroughput"`
	UploadThroughput   float32                       `json:"uploadThroughput"`
	ConnectionType     *types.Network_ConnectionType `json:"connectionType,omitempty"`
}

func (obj *Network) EmulateNetworkConditions(request *EmulateNetworkConditionsRequest) (err error) {
	err = obj.conn.Send("Network.emulateNetworkConditions", request, nil)
	return
}

type SetCacheDisabledRequest struct {
	CacheDisabled bool `json:"cacheDisabled"`
}

func (obj *Network) SetCacheDisabled(request *SetCacheDisabledRequest) (err error) {
	err = obj.conn.Send("Network.setCacheDisabled", request, nil)
	return
}

type SetBypassServiceWorkerRequest struct {
	Bypass bool `json:"bypass"`
}

func (obj *Network) SetBypassServiceWorker(request *SetBypassServiceWorkerRequest) (err error) {
	err = obj.conn.Send("Network.setBypassServiceWorker", request, nil)
	return
}

type SetDataSizeLimitsForTestRequest struct {
	MaxTotalSize    int `json:"maxTotalSize"`
	MaxResourceSize int `json:"maxResourceSize"`
}

func (obj *Network) SetDataSizeLimitsForTest(request *SetDataSizeLimitsForTestRequest) (err error) {
	err = obj.conn.Send("Network.setDataSizeLimitsForTest", request, nil)
	return
}

type GetCertificateRequest struct {
	Origin string `json:"origin"`
}

func (obj *Network) GetCertificate(request *GetCertificateRequest) (response struct {
	TableNames []string `json:"tableNames"`
}, err error) {
	err = obj.conn.Send("Network.getCertificate", request, &response)
	return
}

type SetRequestInterceptionRequest struct {
	Patterns []types.Network_RequestPattern `json:"patterns"`
}

func (obj *Network) SetRequestInterception(request *SetRequestInterceptionRequest) (err error) {
	err = obj.conn.Send("Network.setRequestInterception", request, nil)
	return
}

type ContinueInterceptedRequestRequest struct {
	InterceptionId        types.Network_InterceptionId         `json:"interceptionId"`
	ErrorReason           *types.Network_ErrorReason           `json:"errorReason,omitempty"`
	RawResponse           *string                              `json:"rawResponse,omitempty"`
	Url                   *string                              `json:"url,omitempty"`
	Method                *string                              `json:"method,omitempty"`
	PostData              *string                              `json:"postData,omitempty"`
	Headers               *types.Network_Headers               `json:"headers,omitempty"`
	AuthChallengeResponse *types.Network_AuthChallengeResponse `json:"authChallengeResponse,omitempty"`
}

func (obj *Network) ContinueInterceptedRequest(request *ContinueInterceptedRequestRequest) (err error) {
	err = obj.conn.Send("Network.continueInterceptedRequest", request, nil)
	return
}
