/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

// Actions and events related to the inspected page belong to the page domain.
package page

import (
    "github.com/SKatiyar/cri"
    types "github.com/SKatiyar/cri/types"
)

type Page struct {
	conn cri.Connector
}

// New creates a Page instance
func New(conn cri.Connector) *Page {
	return &Page{conn}
}
// Enables page domain notifications.
func (obj *Page) Enable() (err error) {
	err = obj.conn.Send("Page.enable", nil, nil)
	return
}

// Disables page domain notifications.
func (obj *Page) Disable() (err error) {
	err = obj.conn.Send("Page.disable", nil, nil)
	return
}


type AddScriptToEvaluateOnLoadRequest struct {
	ScriptSource string `json:"scriptSource"`
}


type AddScriptToEvaluateOnLoadResponse struct {
	// Identifier of the added script.
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

// Deprecated, please use addScriptToEvaluateOnNewDocument instead.
func (obj *Page) AddScriptToEvaluateOnLoad(request *AddScriptToEvaluateOnLoadRequest) (response AddScriptToEvaluateOnLoadResponse, err error) {
	err = obj.conn.Send("Page.addScriptToEvaluateOnLoad", request, &response)
	return
}


type RemoveScriptToEvaluateOnLoadRequest struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

// Deprecated, please use removeScriptToEvaluateOnNewDocument instead.
func (obj *Page) RemoveScriptToEvaluateOnLoad(request *RemoveScriptToEvaluateOnLoadRequest) (err error) {
	err = obj.conn.Send("Page.removeScriptToEvaluateOnLoad", request, nil)
	return
}


type AddScriptToEvaluateOnNewDocumentRequest struct {
	Source string `json:"source"`
}


type AddScriptToEvaluateOnNewDocumentResponse struct {
	// Identifier of the added script.
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

// Evaluates given script in every frame upon creation (before loading frame's scripts).
func (obj *Page) AddScriptToEvaluateOnNewDocument(request *AddScriptToEvaluateOnNewDocumentRequest) (response AddScriptToEvaluateOnNewDocumentResponse, err error) {
	err = obj.conn.Send("Page.addScriptToEvaluateOnNewDocument", request, &response)
	return
}


type RemoveScriptToEvaluateOnNewDocumentRequest struct {
	Identifier types.Page_ScriptIdentifier `json:"identifier"`
}

// Removes given script from the list.
func (obj *Page) RemoveScriptToEvaluateOnNewDocument(request *RemoveScriptToEvaluateOnNewDocumentRequest) (err error) {
	err = obj.conn.Send("Page.removeScriptToEvaluateOnNewDocument", request, nil)
	return
}


type SetAutoAttachToCreatedPagesRequest struct {
	// If true, browser will open a new inspector window for every page created from this one.
	AutoAttach bool `json:"autoAttach"`
}

// Controls whether browser will open a new inspector window for connected pages.
func (obj *Page) SetAutoAttachToCreatedPages(request *SetAutoAttachToCreatedPagesRequest) (err error) {
	err = obj.conn.Send("Page.setAutoAttachToCreatedPages", request, nil)
	return
}


type ReloadRequest struct {
	// If true, browser cache is ignored (as if the user pressed Shift+refresh).
	IgnoreCache *bool `json:"ignoreCache,omitempty"`
	// If set, the script will be injected into all frames of the inspected page after reload.
	ScriptToEvaluateOnLoad *string `json:"scriptToEvaluateOnLoad,omitempty"`
}

// Reloads given page optionally ignoring the cache.
func (obj *Page) Reload(request *ReloadRequest) (err error) {
	err = obj.conn.Send("Page.reload", request, nil)
	return
}


type SetAdBlockingEnabledRequest struct {
	// Whether to block ads.
	Enabled bool `json:"enabled"`
}

// Enable Chrome's experimental ad filter on all sites.
func (obj *Page) SetAdBlockingEnabled(request *SetAdBlockingEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setAdBlockingEnabled", request, nil)
	return
}


type NavigateRequest struct {
	// URL to navigate the page to.
	Url string `json:"url"`
	// Referrer URL.
	// NOTE Experimental
	Referrer *string `json:"referrer,omitempty"`
	// Intended transition type.
	// NOTE Experimental
	TransitionType *types.Page_TransitionType `json:"transitionType,omitempty"`
}


type NavigateResponse struct {
	// Frame id that will be navigated.
	// NOTE Experimental
	FrameId types.Page_FrameId `json:"frameId"`
}

// Navigates current page to the given URL.
func (obj *Page) Navigate(request *NavigateRequest) (response NavigateResponse, err error) {
	err = obj.conn.Send("Page.navigate", request, &response)
	return
}

// Force the page stop all navigations and pending resource fetches.
func (obj *Page) StopLoading() (err error) {
	err = obj.conn.Send("Page.stopLoading", nil, nil)
	return
}


type GetNavigationHistoryResponse struct {
	// Index of the current navigation history entry.
	CurrentIndex int `json:"currentIndex"`
	// Array of navigation history entries.
	Entries []types.Page_NavigationEntry `json:"entries"`
}

// Returns navigation history for the current page.
func (obj *Page) GetNavigationHistory() (response GetNavigationHistoryResponse, err error) {
	err = obj.conn.Send("Page.getNavigationHistory", nil, &response)
	return
}


type NavigateToHistoryEntryRequest struct {
	// Unique id of the entry to navigate to.
	EntryId int `json:"entryId"`
}

// Navigates current page to the given history entry.
func (obj *Page) NavigateToHistoryEntry(request *NavigateToHistoryEntryRequest) (err error) {
	err = obj.conn.Send("Page.navigateToHistoryEntry", request, nil)
	return
}


type GetCookiesResponse struct {
	// Array of cookie objects.
	Cookies []types.Network_Cookie `json:"cookies"`
}

// Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the <code>cookies</code> field.
func (obj *Page) GetCookies() (response GetCookiesResponse, err error) {
	err = obj.conn.Send("Page.getCookies", nil, &response)
	return
}


type DeleteCookieRequest struct {
	// Name of the cookie to remove.
	CookieName string `json:"cookieName"`
	// URL to match cooke domain and path.
	Url string `json:"url"`
}

// Deletes browser cookie with given name, domain and path.
func (obj *Page) DeleteCookie(request *DeleteCookieRequest) (err error) {
	err = obj.conn.Send("Page.deleteCookie", request, nil)
	return
}


type GetResourceTreeResponse struct {
	// Present frame / resource tree structure.
	FrameTree types.Page_FrameResourceTree `json:"frameTree"`
}

// Returns present frame / resource tree structure.
func (obj *Page) GetResourceTree() (response GetResourceTreeResponse, err error) {
	err = obj.conn.Send("Page.getResourceTree", nil, &response)
	return
}


type GetResourceContentRequest struct {
	// Frame id to get resource for.
	FrameId types.Page_FrameId `json:"frameId"`
	// URL of the resource to get content for.
	Url string `json:"url"`
}


type GetResourceContentResponse struct {
	// Resource content.
	Content string `json:"content"`
	// True, if content was served as base64.
	Base64Encoded bool `json:"base64Encoded"`
}

// Returns content of the given resource.
func (obj *Page) GetResourceContent(request *GetResourceContentRequest) (response GetResourceContentResponse, err error) {
	err = obj.conn.Send("Page.getResourceContent", request, &response)
	return
}


type SearchInResourceRequest struct {
	// Frame id for resource to search in.
	FrameId types.Page_FrameId `json:"frameId"`
	// URL of the resource to search in.
	Url string `json:"url"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex *bool `json:"isRegex,omitempty"`
}


type SearchInResourceResponse struct {
	// List of search matches.
	Result []types.Debugger_SearchMatch `json:"result"`
}

// Searches for given string in resource content.
func (obj *Page) SearchInResource(request *SearchInResourceRequest) (response SearchInResourceResponse, err error) {
	err = obj.conn.Send("Page.searchInResource", request, &response)
	return
}


type SetDocumentContentRequest struct {
	// Frame id to set HTML for.
	FrameId types.Page_FrameId `json:"frameId"`
	// HTML content to set.
	Html string `json:"html"`
}

// Sets given markup as the document's HTML.
func (obj *Page) SetDocumentContent(request *SetDocumentContentRequest) (err error) {
	err = obj.conn.Send("Page.setDocumentContent", request, nil)
	return
}


type SetDeviceMetricsOverrideRequest struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float32 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`
	// Scale to apply to resulting view image.
	Scale *float32 `json:"scale,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000).
	ScreenWidth *int `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000).
	ScreenHeight *int `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000).
	PositionX *int `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000).
	PositionY *int `json:"positionY,omitempty"`
	// Do not set visible view size, rely upon explicit setVisibleSize call.
	DontSetVisibleSize *bool `json:"dontSetVisibleSize,omitempty"`
	// Screen orientation override.
	ScreenOrientation *types.Emulation_ScreenOrientation `json:"screenOrientation,omitempty"`
	// The viewport dimensions and scale. If not set, the override is cleared.
	Viewport *types.Page_Viewport `json:"viewport,omitempty"`
}

// Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (obj *Page) SetDeviceMetricsOverride(request *SetDeviceMetricsOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setDeviceMetricsOverride", request, nil)
	return
}

// Clears the overriden device metrics.
func (obj *Page) ClearDeviceMetricsOverride() (err error) {
	err = obj.conn.Send("Page.clearDeviceMetricsOverride", nil, nil)
	return
}


type SetGeolocationOverrideRequest struct {
	// Mock latitude
	Latitude *float32 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude *float32 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy *float32 `json:"accuracy,omitempty"`
}

// Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (obj *Page) SetGeolocationOverride(request *SetGeolocationOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setGeolocationOverride", request, nil)
	return
}

// Clears the overriden Geolocation Position and Error.
func (obj *Page) ClearGeolocationOverride() (err error) {
	err = obj.conn.Send("Page.clearGeolocationOverride", nil, nil)
	return
}


type SetDeviceOrientationOverrideRequest struct {
	// Mock alpha
	Alpha float32 `json:"alpha"`
	// Mock beta
	Beta float32 `json:"beta"`
	// Mock gamma
	Gamma float32 `json:"gamma"`
}

// Overrides the Device Orientation.
func (obj *Page) SetDeviceOrientationOverride(request *SetDeviceOrientationOverrideRequest) (err error) {
	err = obj.conn.Send("Page.setDeviceOrientationOverride", request, nil)
	return
}

// Clears the overridden Device Orientation.
func (obj *Page) ClearDeviceOrientationOverride() (err error) {
	err = obj.conn.Send("Page.clearDeviceOrientationOverride", nil, nil)
	return
}


type SetTouchEmulationEnabledRequest struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration *string `json:"configuration,omitempty"`
}

// Toggles mouse event-based touch event emulation.
func (obj *Page) SetTouchEmulationEnabled(request *SetTouchEmulationEnabledRequest) (err error) {
	err = obj.conn.Send("Page.setTouchEmulationEnabled", request, nil)
	return
}


type CaptureScreenshotRequest struct {
	// Image compression format (defaults to png).
	Format *string `json:"format,omitempty"`
	// Compression quality from range [0..100] (jpeg only).
	Quality *int `json:"quality,omitempty"`
	// Capture the screenshot of a given region only.
	// NOTE Experimental
	Clip *types.Page_Viewport `json:"clip,omitempty"`
	// Capture the screenshot from the surface, rather than the view. Defaults to true.
	// NOTE Experimental
	FromSurface *bool `json:"fromSurface,omitempty"`
}


type CaptureScreenshotResponse struct {
	// Base64-encoded image data.
	Data string `json:"data"`
}

// Capture page screenshot.
func (obj *Page) CaptureScreenshot(request *CaptureScreenshotRequest) (response CaptureScreenshotResponse, err error) {
	err = obj.conn.Send("Page.captureScreenshot", request, &response)
	return
}


type PrintToPDFRequest struct {
	// Paper orientation. Defaults to false.
	Landscape *bool `json:"landscape,omitempty"`
	// Display header and footer. Defaults to false.
	DisplayHeaderFooter *bool `json:"displayHeaderFooter,omitempty"`
	// Print background graphics. Defaults to false.
	PrintBackground *bool `json:"printBackground,omitempty"`
	// Scale of the webpage rendering. Defaults to 1.
	Scale *float32 `json:"scale,omitempty"`
	// Paper width in inches. Defaults to 8.5 inches.
	PaperWidth *float32 `json:"paperWidth,omitempty"`
	// Paper height in inches. Defaults to 11 inches.
	PaperHeight *float32 `json:"paperHeight,omitempty"`
	// Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginTop *float32 `json:"marginTop,omitempty"`
	// Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom *float32 `json:"marginBottom,omitempty"`
	// Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft *float32 `json:"marginLeft,omitempty"`
	// Right margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight *float32 `json:"marginRight,omitempty"`
	// Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	PageRanges *string `json:"pageRanges,omitempty"`
	// Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
	IgnoreInvalidPageRanges *bool `json:"ignoreInvalidPageRanges,omitempty"`
}


type PrintToPDFResponse struct {
	// Base64-encoded pdf data.
	Data string `json:"data"`
}

// Print page as PDF.
func (obj *Page) PrintToPDF(request *PrintToPDFRequest) (response PrintToPDFResponse, err error) {
	err = obj.conn.Send("Page.printToPDF", request, &response)
	return
}


type StartScreencastRequest struct {
	// Image compression format.
	Format *string `json:"format,omitempty"`
	// Compression quality from range [0..100].
	Quality *int `json:"quality,omitempty"`
	// Maximum screenshot width.
	MaxWidth *int `json:"maxWidth,omitempty"`
	// Maximum screenshot height.
	MaxHeight *int `json:"maxHeight,omitempty"`
	// Send every n-th frame.
	EveryNthFrame *int `json:"everyNthFrame,omitempty"`
}

// Starts sending each frame using the <code>screencastFrame</code> event.
func (obj *Page) StartScreencast(request *StartScreencastRequest) (err error) {
	err = obj.conn.Send("Page.startScreencast", request, nil)
	return
}

// Stops sending each frame in the <code>screencastFrame</code>.
func (obj *Page) StopScreencast() (err error) {
	err = obj.conn.Send("Page.stopScreencast", nil, nil)
	return
}


type ScreencastFrameAckRequest struct {
	// Frame number.
	SessionId int `json:"sessionId"`
}

// Acknowledges that a screencast frame has been received by the frontend.
func (obj *Page) ScreencastFrameAck(request *ScreencastFrameAckRequest) (err error) {
	err = obj.conn.Send("Page.screencastFrameAck", request, nil)
	return
}


type HandleJavaScriptDialogRequest struct {
	// Whether to accept or dismiss the dialog.
	Accept bool `json:"accept"`
	// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
	PromptText *string `json:"promptText,omitempty"`
}

// Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
func (obj *Page) HandleJavaScriptDialog(request *HandleJavaScriptDialogRequest) (err error) {
	err = obj.conn.Send("Page.handleJavaScriptDialog", request, nil)
	return
}


type GetAppManifestResponse struct {
	// Manifest location.
	Url string `json:"url"`
	Errors []types.Page_AppManifestError `json:"errors"`
	// Manifest content.
	Data *string `json:"data,omitempty"`
}


func (obj *Page) GetAppManifest() (response GetAppManifestResponse, err error) {
	err = obj.conn.Send("Page.getAppManifest", nil, &response)
	return
}


func (obj *Page) RequestAppBanner() (err error) {
	err = obj.conn.Send("Page.requestAppBanner", nil, nil)
	return
}


type GetLayoutMetricsResponse struct {
	// Metrics relating to the layout viewport.
	LayoutViewport types.Page_LayoutViewport `json:"layoutViewport"`
	// Metrics relating to the visual viewport.
	VisualViewport types.Page_VisualViewport `json:"visualViewport"`
	// Size of scrollable area.
	ContentSize types.DOM_Rect `json:"contentSize"`
}

// Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
func (obj *Page) GetLayoutMetrics() (response GetLayoutMetricsResponse, err error) {
	err = obj.conn.Send("Page.getLayoutMetrics", nil, &response)
	return
}


type CreateIsolatedWorldRequest struct {
	// Id of the frame in which the isolated world should be created.
	FrameId types.Page_FrameId `json:"frameId"`
	// An optional name which is reported in the Execution Context.
	WorldName *string `json:"worldName,omitempty"`
	// Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
	GrantUniveralAccess *bool `json:"grantUniveralAccess,omitempty"`
}


type CreateIsolatedWorldResponse struct {
	// Execution context of the isolated world.
	ExecutionContextId types.Runtime_ExecutionContextId `json:"executionContextId"`
}

// Creates an isolated world for the given frame.
func (obj *Page) CreateIsolatedWorld(request *CreateIsolatedWorldRequest) (response CreateIsolatedWorldResponse, err error) {
	err = obj.conn.Send("Page.createIsolatedWorld", request, &response)
	return
}

// Brings page to front (activates tab).
func (obj *Page) BringToFront() (err error) {
	err = obj.conn.Send("Page.bringToFront", nil, nil)
	return
}


type SetDownloadBehaviorRequest struct {
	// Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
	Behavior string `json:"behavior"`
	// The default path to save downloaded files to. This is requred if behavior is set to 'allow'
	DownloadPath *string `json:"downloadPath,omitempty"`
}

// Set the behavior when downloading a file.
func (obj *Page) SetDownloadBehavior(request *SetDownloadBehaviorRequest) (err error) {
	err = obj.conn.Send("Page.setDownloadBehavior", request, nil)
	return
}
