package page

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Page struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Page {
	return &Page{conn}
}
func (obj *Page) Enable() (err error) {
	err = obj.conn.Send("Page.enable", nil, nil)
	return
}
func (obj *Page) Disable() (err error) {
	err = obj.conn.Send("Page.disable", nil, nil)
	return
}

type AddScriptToEvaluateOnLoadRequest struct {
	ScriptSource string `json:"scriptSource"`
}

func (obj *Page) AddScriptToEvaluateOnLoad(request *AddScriptToEvaluateOnLoadRequest) (response struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}, err error) {
	err = obj.conn.Send("Page.addScriptToEvaluateOnLoad", request, &response)
	return
}

type RemoveScriptToEvaluateOnLoadRequest struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

func (obj *Page) RemoveScriptToEvaluateOnLoad(request *RemoveScriptToEvaluateOnLoadRequest) (err error) {
	err = obj.conn.Send("Page.removeScriptToEvaluateOnLoad", request, nil)
	return
}

type AddScriptToEvaluateOnNewDocumentRequest struct {
	Source string `json:"source"`
}

func (obj *Page) AddScriptToEvaluateOnNewDocument(request *AddScriptToEvaluateOnNewDocumentRequest) (response struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}, err error) {
	err = obj.conn.Send("Page.addScriptToEvaluateOnNewDocument", request, &response)
	return
}

type RemoveScriptToEvaluateOnNewDocumentRequest struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

func (obj *Page) RemoveScriptToEvaluateOnNewDocument(request *RemoveScriptToEvaluateOnNewDocumentRequest) (err error) {
	err = obj.conn.Send("Page.removeScriptToEvaluateOnNewDocument", request, nil)
	return
}

type SetAutoAttachToCreatedPagesRequest struct {
	AutoAttach bool `json:"autoAttach"`
}

func (obj *Page) SetAutoAttachToCreatedPages(request *SetAutoAttachToCreatedPagesRequest) (err error) {
	err = obj.conn.Send("Page.setAutoAttachToCreatedPages", request, nil)
	return
}

type ReloadRequest struct {
	IgnoreCache            *bool   `json:"ignoreCache,omitempty"`
	ScriptToEvaluateOnLoad *string `json:"scriptToEvaluateOnLoad,omitempty"`
}

func (obj *Page) Reload(request *ReloadRequest) (err error) {
	err = obj.conn.Send("Page.reload", request, nil)
	return
}

type SetAdBlockingEnabledRequest struct {
	Enabled bool `json:"enabled"`
}

func (obj *Page) SetAdBlockingEnabled(request *SetAdBlockingEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setAdBlockingEnabled", request, nil)
	return
}

type NavigateRequest struct {
	Url            string                     `json:"url"`
	Referrer       *string                    `json:"referrer,omitempty"`
	TransitionType *types.Page_TransitionType `json:"transitionType,omitempty"`
}

func (obj *Page) Navigate(request *NavigateRequest) (response struct {
	FrameId types.Page_FrameId `json:"frameId"`
}, err error) {
	err = obj.conn.Send("Page.navigate", request, &response)
	return
}
func (obj *Page) StopLoading() (err error) {
	err = obj.conn.Send("Page.stopLoading", nil, nil)
	return
}
func (obj *Page) GetNavigationHistory() (response struct {
	CurrentIndex int                          `json:"currentIndex"`
	Entries      []types.Page_NavigationEntry `json:"entries"`
}, err error) {
	err = obj.conn.Send("Page.getNavigationHistory", nil, &response)
	return
}

type NavigateToHistoryEntryRequest struct {
	EntryId int `json:"entryId"`
}

func (obj *Page) NavigateToHistoryEntry(request *NavigateToHistoryEntryRequest) (err error) {
	err = obj.conn.Send("Page.navigateToHistoryEntry", request, nil)
	return
}
func (obj *Page) GetCookies() (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`
}, err error) {
	err = obj.conn.Send("Page.getCookies", nil, &response)
	return
}

type DeleteCookieRequest struct {
	CookieName string `json:"cookieName"`
	Url        string `json:"url"`
}

func (obj *Page) DeleteCookie(request *DeleteCookieRequest) (err error) {
	err = obj.conn.Send("Page.deleteCookie", request, nil)
	return
}
func (obj *Page) GetResourceTree() (response struct {
	FrameTree types.Page_FrameResourceTree `json:"frameTree"`
}, err error) {
	err = obj.conn.Send("Page.getResourceTree", nil, &response)
	return
}

type GetResourceContentRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`
	Url     string             `json:"url"`
}

func (obj *Page) GetResourceContent(request *GetResourceContentRequest) (response struct {
	Content       string `json:"content"`
	Base64Encoded bool   `json:"base64Encoded"`
}, err error) {
	err = obj.conn.Send("Page.getResourceContent", request, &response)
	return
}

type SearchInResourceRequest struct {
	FrameId       types.Page_FrameId `json:"frameId"`
	Url           string             `json:"url"`
	Query         string             `json:"query"`
	CaseSensitive *bool              `json:"caseSensitive,omitempty"`
	IsRegex       *bool              `json:"isRegex,omitempty"`
}

func (obj *Page) SearchInResource(request *SearchInResourceRequest) (response struct {
	Result []types.Debugger_SearchMatch `json:"result"`
}, err error) {
	err = obj.conn.Send("Page.searchInResource", request, &response)
	return
}

type SetDocumentContentRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`
	Html    string             `json:"html"`
}

func (obj *Page) SetDocumentContent(request *SetDocumentContentRequest) (err error) {
	err = obj.conn.Send("Page.setDocumentContent", request, nil)
	return
}

type SetDeviceMetricsOverrideRequest struct {
	Width              int                                `json:"width"`
	Height             int                                `json:"height"`
	DeviceScaleFactor  float32                            `json:"deviceScaleFactor"`
	Mobile             bool                               `json:"mobile"`
	Scale              *float32                           `json:"scale,omitempty"`
	ScreenWidth        *int                               `json:"screenWidth,omitempty"`
	ScreenHeight       *int                               `json:"screenHeight,omitempty"`
	PositionX          *int                               `json:"positionX,omitempty"`
	PositionY          *int                               `json:"positionY,omitempty"`
	DontSetVisibleSize *bool                              `json:"dontSetVisibleSize,omitempty"`
	ScreenOrientation  *types.Emulation_ScreenOrientation `json:"screenOrientation,omitempty"`
	Viewport           *types.Page_Viewport               `json:"viewport,omitempty"`
}

func (obj *Page) SetDeviceMetricsOverride(request *SetDeviceMetricsOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setDeviceMetricsOverride", request, nil)
	return
}
func (obj *Page) ClearDeviceMetricsOverride() (err error) {
	err = obj.conn.Send("Page.clearDeviceMetricsOverride", nil, nil)
	return
}

type SetGeolocationOverrideRequest struct {
	Latitude  *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Accuracy  *float32 `json:"accuracy,omitempty"`
}

func (obj *Page) SetGeolocationOverride(request *SetGeolocationOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setGeolocationOverride", request, nil)
	return
}
func (obj *Page) ClearGeolocationOverride() (err error) {
	err = obj.conn.Send("Page.clearGeolocationOverride", nil, nil)
	return
}

type SetDeviceOrientationOverrideRequest struct {
	Alpha float32 `json:"alpha"`
	Beta  float32 `json:"beta"`
	Gamma float32 `json:"gamma"`
}

func (obj *Page) SetDeviceOrientationOverride(request *SetDeviceOrientationOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setDeviceOrientationOverride", request, nil)
	return
}
func (obj *Page) ClearDeviceOrientationOverride() (err error) {
	err = obj.conn.Send("Page.clearDeviceOrientationOverride", nil, nil)
	return
}

type SetTouchEmulationEnabledRequest struct {
	Enabled       bool    `json:"enabled"`
	Configuration *string `json:"configuration,omitempty"`
}

func (obj *Page) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setTouchEmulationEnabled", request, nil)
	return
}

type CaptureScreenshotRequest struct {
	Format      *string              `json:"format,omitempty"`
	Quality     *int                 `json:"quality,omitempty"`
	Clip        *types.Page_Viewport `json:"clip,omitempty"`
	FromSurface *bool                `json:"fromSurface,omitempty"`
}

func (obj *Page) CaptureScreenshot(request *CaptureScreenshotRequest) (response struct {
	Data string `json:"data"`
}, err error) {
	err = obj.conn.Send("Page.captureScreenshot", request, &response)
	return
}

type PrintToPDFRequest struct {
	Landscape               *bool    `json:"landscape,omitempty"`
	DisplayHeaderFooter     *bool    `json:"displayHeaderFooter,omitempty"`
	PrintBackground         *bool    `json:"printBackground,omitempty"`
	Scale                   *float32 `json:"scale,omitempty"`
	PaperWidth              *float32 `json:"paperWidth,omitempty"`
	PaperHeight             *float32 `json:"paperHeight,omitempty"`
	MarginTop               *float32 `json:"marginTop,omitempty"`
	MarginBottom            *float32 `json:"marginBottom,omitempty"`
	MarginLeft              *float32 `json:"marginLeft,omitempty"`
	MarginRight             *float32 `json:"marginRight,omitempty"`
	PageRanges              *string  `json:"pageRanges,omitempty"`
	IgnoreInvalidPageRanges *bool    `json:"ignoreInvalidPageRanges,omitempty"`
}

func (obj *Page) PrintToPDF(request *PrintToPDFRequest) (response struct {
	Data string `json:"data"`
}, err error) {
	err = obj.conn.Send("Page.printToPDF", request, &response)
	return
}

type StartScreencastRequest struct {
	Format        *string `json:"format,omitempty"`
	Quality       *int    `json:"quality,omitempty"`
	MaxWidth      *int    `json:"maxWidth,omitempty"`
	MaxHeight     *int    `json:"maxHeight,omitempty"`
	EveryNthFrame *int    `json:"everyNthFrame,omitempty"`
}

func (obj *Page) StartScreencast(request *StartScreencastRequest) (err error) {
	err = obj.conn.Send("Page.startScreencast", request, nil)
	return
}
func (obj *Page) StopScreencast() (err error) {
	err = obj.conn.Send("Page.stopScreencast", nil, nil)
	return
}

type ScreencastFrameAckRequest struct {
	SessionId int `json:"sessionId"`
}

func (obj *Page) ScreencastFrameAck(request *ScreencastFrameAckRequest) (err error) {
	err = obj.conn.Send("Page.screencastFrameAck", request, nil)
	return
}

type HandleJavaScriptDialogRequest struct {
	Accept     bool    `json:"accept"`
	PromptText *string `json:"promptText,omitempty"`
}

func (obj *Page) HandleJavaScriptDialog(request *HandleJavaScriptDialogRequest) (err error) {
	err = obj.conn.Send("Page.handleJavaScriptDialog", request, nil)
	return
}
func (obj *Page) GetAppManifest() (response struct {
	Url    string                        `json:"url"`
	Errors []types.Page_AppManifestError `json:"errors"`
	Data   *string                       `json:"data,omitempty"`
}, err error) {
	err = obj.conn.Send("Page.getAppManifest", nil, &response)
	return
}
func (obj *Page) RequestAppBanner() (err error) {
	err = obj.conn.Send("Page.requestAppBanner", nil, nil)
	return
}
func (obj *Page) GetLayoutMetrics() (response struct {
	LayoutViewport types.Page_LayoutViewport `json:"layoutViewport"`
	VisualViewport types.Page_VisualViewport `json:"visualViewport"`
	ContentSize    types.DOM_Rect            `json:"contentSize"`
}, err error) {
	err = obj.conn.Send("Page.getLayoutMetrics", nil, &response)
	return
}

type CreateIsolatedWorldRequest struct {
	FrameId             types.Page_FrameId `json:"frameId"`
	WorldName           *string            `json:"worldName,omitempty"`
	GrantUniveralAccess *bool              `json:"grantUniveralAccess,omitempty"`
}

func (obj *Page) CreateIsolatedWorld(request *CreateIsolatedWorldRequest) (response struct {
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
}, err error) {
	err = obj.conn.Send("Page.createIsolatedWorld", request, &response)
	return
}
func (obj *Page) BringToFront() (err error) {
	err = obj.conn.Send("Page.bringToFront", nil, nil)
	return
}

type SetDownloadBehaviorRequest struct {
	Behavior     string  `json:"behavior"`
	DownloadPath *string `json:"downloadPath,omitempty"`
}

func (obj *Page) SetDownloadBehavior(request *SetDownloadBehaviorRequest) (err error) {
	err = obj.conn.Send("Page.setDownloadBehavior", request, nil)
	return
}
