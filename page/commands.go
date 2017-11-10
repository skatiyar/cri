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
	Identifier types.Page_ScriptIdentifier `json:"identifier"`// Identifier of the added script.
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
	Identifier types.Page_ScriptIdentifier `json:"identifier"`// Identifier of the added script.
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
	AutoAttach bool `json:"autoAttach"`// If true, browser will open a new inspector window for every page created from this one.
}

func (obj *Page) SetAutoAttachToCreatedPages(request *SetAutoAttachToCreatedPagesRequest) (err error) {
	err = obj.conn.Send("Page.setAutoAttachToCreatedPages", request, nil)
	return
}

type ReloadRequest struct {
	IgnoreCache		*bool	`json:"ignoreCache,omitempty"`// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	ScriptToEvaluateOnLoad	*string	`json:"scriptToEvaluateOnLoad,omitempty"`// If set, the script will be injected into all frames of the inspected page after reload.
}

func (obj *Page) Reload(request *ReloadRequest) (err error) {
	err = obj.conn.Send("Page.reload", request, nil)
	return
}

type SetAdBlockingEnabledRequest struct {
	Enabled bool `json:"enabled"`// Whether to block ads.
}

func (obj *Page) SetAdBlockingEnabled(request *SetAdBlockingEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setAdBlockingEnabled", request, nil)
	return
}

type NavigateRequest struct {
	Url		string				`json:"url"`// URL to navigate the page to.
	Referrer	*string				`json:"referrer,omitempty"`// Referrer URL.
	TransitionType	*types.Page_TransitionType	`json:"transitionType,omitempty"`// Intended transition type.
}

func (obj *Page) Navigate(request *NavigateRequest) (response struct {
	FrameId types.Page_FrameId `json:"frameId"`// Frame id that will be navigated.
}, err error) {
	err = obj.conn.Send("Page.navigate", request, &response)
	return
}
func (obj *Page) StopLoading() (err error) {
	err = obj.conn.Send("Page.stopLoading", nil, nil)
	return
}
func (obj *Page) GetNavigationHistory() (response struct {
	CurrentIndex	int				`json:"currentIndex"`// Index of the current navigation history entry.
	Entries		[]types.Page_NavigationEntry	`json:"entries"`// Array of navigation history entries.
}, err error) {
	err = obj.conn.Send("Page.getNavigationHistory", nil, &response)
	return
}

type NavigateToHistoryEntryRequest struct {
	EntryId int `json:"entryId"`// Unique id of the entry to navigate to.
}

func (obj *Page) NavigateToHistoryEntry(request *NavigateToHistoryEntryRequest) (err error) {
	err = obj.conn.Send("Page.navigateToHistoryEntry", request, nil)
	return
}
func (obj *Page) GetCookies() (response struct {
	Cookies []types.Network_Cookie `json:"cookies"`// Array of cookie objects.
}, err error) {
	err = obj.conn.Send("Page.getCookies", nil, &response)
	return
}

type DeleteCookieRequest struct {
	CookieName	string	`json:"cookieName"`// Name of the cookie to remove.
	Url		string	`json:"url"`// URL to match cooke domain and path.
}

func (obj *Page) DeleteCookie(request *DeleteCookieRequest) (err error) {
	err = obj.conn.Send("Page.deleteCookie", request, nil)
	return
}
func (obj *Page) GetResourceTree() (response struct {
	FrameTree types.Page_FrameResourceTree `json:"frameTree"`// Present frame / resource tree structure.
}, err error) {
	err = obj.conn.Send("Page.getResourceTree", nil, &response)
	return
}

type GetResourceContentRequest struct {
	FrameId	types.Page_FrameId	`json:"frameId"`// Frame id to get resource for.
	Url	string			`json:"url"`// URL of the resource to get content for.
}

func (obj *Page) GetResourceContent(request *GetResourceContentRequest) (response struct {
	Content		string	`json:"content"`// Resource content.
	Base64Encoded	bool	`json:"base64Encoded"`// True, if content was served as base64.
}, err error) {
	err = obj.conn.Send("Page.getResourceContent", request, &response)
	return
}

type SearchInResourceRequest struct {
	FrameId		types.Page_FrameId	`json:"frameId"`// Frame id for resource to search in.
	Url		string			`json:"url"`// URL of the resource to search in.
	Query		string			`json:"query"`// String to search for.
	CaseSensitive	*bool			`json:"caseSensitive,omitempty"`// If true, search is case sensitive.
	IsRegex		*bool			`json:"isRegex,omitempty"`// If true, treats string parameter as regex.
}

func (obj *Page) SearchInResource(request *SearchInResourceRequest) (response struct {
	Result []types.Debugger_SearchMatch `json:"result"`// List of search matches.
}, err error) {
	err = obj.conn.Send("Page.searchInResource", request, &response)
	return
}

type SetDocumentContentRequest struct {
	FrameId	types.Page_FrameId	`json:"frameId"`// Frame id to set HTML for.
	Html	string			`json:"html"`// HTML content to set.
}

func (obj *Page) SetDocumentContent(request *SetDocumentContentRequest) (err error) {
	err = obj.conn.Send("Page.setDocumentContent", request, nil)
	return
}

type SetDeviceMetricsOverrideRequest struct {
	Width			int					`json:"width"`// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height			int					`json:"height"`// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor	float32					`json:"deviceScaleFactor"`// Overriding device scale factor value. 0 disables the override.
	Mobile			bool					`json:"mobile"`// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Scale			*float32				`json:"scale,omitempty"`// Scale to apply to resulting view image.
	ScreenWidth		*int					`json:"screenWidth,omitempty"`// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenHeight		*int					`json:"screenHeight,omitempty"`// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	PositionX		*int					`json:"positionX,omitempty"`// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionY		*int					`json:"positionY,omitempty"`// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	DontSetVisibleSize	*bool					`json:"dontSetVisibleSize,omitempty"`// Do not set visible view size, rely upon explicit setVisibleSize call.
	ScreenOrientation	*types.Emulation_ScreenOrientation	`json:"screenOrientation,omitempty"`// Screen orientation override.
	Viewport		*types.Page_Viewport			`json:"viewport,omitempty"`// The viewport dimensions and scale. If not set, the override is cleared.
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
	Latitude	*float32	`json:"latitude,omitempty"`// Mock latitude
	Longitude	*float32	`json:"longitude,omitempty"`// Mock longitude
	Accuracy	*float32	`json:"accuracy,omitempty"`// Mock accuracy
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
	Alpha	float32	`json:"alpha"`// Mock alpha
	Beta	float32	`json:"beta"`// Mock beta
	Gamma	float32	`json:"gamma"`// Mock gamma
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
	Enabled		bool	`json:"enabled"`// Whether the touch event emulation should be enabled.
	Configuration	*string	`json:"configuration,omitempty"`// Touch/gesture events configuration. Default: current platform.
}

func (obj *Page) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setTouchEmulationEnabled", request, nil)
	return
}

type CaptureScreenshotRequest struct {
	Format		*string			`json:"format,omitempty"`// Image compression format (defaults to png).
	Quality		*int			`json:"quality,omitempty"`// Compression quality from range [0..100] (jpeg only).
	Clip		*types.Page_Viewport	`json:"clip,omitempty"`// Capture the screenshot of a given region only.
	FromSurface	*bool			`json:"fromSurface,omitempty"`// Capture the screenshot from the surface, rather than the view. Defaults to true.
}

func (obj *Page) CaptureScreenshot(request *CaptureScreenshotRequest) (response struct {
	Data string `json:"data"`// Base64-encoded image data.
}, err error) {
	err = obj.conn.Send("Page.captureScreenshot", request, &response)
	return
}

type PrintToPDFRequest struct {
	Landscape		*bool		`json:"landscape,omitempty"`// Paper orientation. Defaults to false.
	DisplayHeaderFooter	*bool		`json:"displayHeaderFooter,omitempty"`// Display header and footer. Defaults to false.
	PrintBackground		*bool		`json:"printBackground,omitempty"`// Print background graphics. Defaults to false.
	Scale			*float32	`json:"scale,omitempty"`// Scale of the webpage rendering. Defaults to 1.
	PaperWidth		*float32	`json:"paperWidth,omitempty"`// Paper width in inches. Defaults to 8.5 inches.
	PaperHeight		*float32	`json:"paperHeight,omitempty"`// Paper height in inches. Defaults to 11 inches.
	MarginTop		*float32	`json:"marginTop,omitempty"`// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom		*float32	`json:"marginBottom,omitempty"`// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft		*float32	`json:"marginLeft,omitempty"`// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight		*float32	`json:"marginRight,omitempty"`// Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges		*string		`json:"pageRanges,omitempty"`// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges	*bool		`json:"ignoreInvalidPageRanges,omitempty"`// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
}

func (obj *Page) PrintToPDF(request *PrintToPDFRequest) (response struct {
	Data string `json:"data"`// Base64-encoded pdf data.
}, err error) {
	err = obj.conn.Send("Page.printToPDF", request, &response)
	return
}

type StartScreencastRequest struct {
	Format		*string	`json:"format,omitempty"`// Image compression format.
	Quality		*int	`json:"quality,omitempty"`// Compression quality from range [0..100].
	MaxWidth	*int	`json:"maxWidth,omitempty"`// Maximum screenshot width.
	MaxHeight	*int	`json:"maxHeight,omitempty"`// Maximum screenshot height.
	EveryNthFrame	*int	`json:"everyNthFrame,omitempty"`// Send every n-th frame.
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
	SessionId int `json:"sessionId"`// Frame number.
}

func (obj *Page) ScreencastFrameAck(request *ScreencastFrameAckRequest) (err error) {
	err = obj.conn.Send("Page.screencastFrameAck", request, nil)
	return
}

type HandleJavaScriptDialogRequest struct {
	Accept		bool	`json:"accept"`// Whether to accept or dismiss the dialog.
	PromptText	*string	`json:"promptText,omitempty"`// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}

func (obj *Page) HandleJavaScriptDialog(request *HandleJavaScriptDialogRequest) (err error) {
	err = obj.conn.Send("Page.handleJavaScriptDialog", request, nil)
	return
}
func (obj *Page) GetAppManifest() (response struct {
	Url	string				`json:"url"`// Manifest location.
	Errors	[]types.Page_AppManifestError	`json:"errors"`
	Data	*string				`json:"data,omitempty"`// Manifest content.
}, err error) {
	err = obj.conn.Send("Page.getAppManifest", nil, &response)
	return
}
func (obj *Page) RequestAppBanner() (err error) {
	err = obj.conn.Send("Page.requestAppBanner", nil, nil)
	return
}
func (obj *Page) GetLayoutMetrics() (response struct {
	LayoutViewport	types.Page_LayoutViewport	`json:"layoutViewport"`// Metrics relating to the layout viewport.
	VisualViewport	types.Page_VisualViewport	`json:"visualViewport"`// Metrics relating to the visual viewport.
	ContentSize	types.DOM_Rect			`json:"contentSize"`// Size of scrollable area.
}, err error) {
	err = obj.conn.Send("Page.getLayoutMetrics", nil, &response)
	return
}

type CreateIsolatedWorldRequest struct {
	FrameId			types.Page_FrameId	`json:"frameId"`// Id of the frame in which the isolated world should be created.
	WorldName		*string			`json:"worldName,omitempty"`// An optional name which is reported in the Execution Context.
	GrantUniveralAccess	*bool			`json:"grantUniveralAccess,omitempty"`// Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
}

func (obj *Page) CreateIsolatedWorld(request *CreateIsolatedWorldRequest) (response struct {
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`// Execution context of the isolated world.
}, err error) {
	err = obj.conn.Send("Page.createIsolatedWorld", request, &response)
	return
}
func (obj *Page) BringToFront() (err error) {
	err = obj.conn.Send("Page.bringToFront", nil, nil)
	return
}

type SetDownloadBehaviorRequest struct {
	Behavior	string	`json:"behavior"`// Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
	DownloadPath	*string	`json:"downloadPath,omitempty"`// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
}

func (obj *Page) SetDownloadBehavior(request *SetDownloadBehaviorRequest) (err error) {
	err = obj.conn.Send("Page.setDownloadBehavior", request, nil)
	return
}
