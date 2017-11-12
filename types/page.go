/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Resource type as it was perceived by the rendering engine.
type Page_ResourceType string

//Unique frame identifier.
type Page_FrameId string

//Information about the Frame on the page.
type Page_Frame struct {
	// Frame unique identifier.
	Id string `json:"id"`
	// Parent frame identifier.
	ParentId *string `json:"parentId,omitempty"`
	// Identifier of the loader associated with this frame.
	LoaderId Network_LoaderId `json:"loaderId"`
	// Frame's name as specified in the tag.
	Name *string `json:"name,omitempty"`
	// Frame document's URL.
	Url string `json:"url"`
	// Frame document's security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Frame document's mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// If the frame failed to load, this contains the URL that could not be loaded.
	// NOTE Experimental
	UnreachableUrl *string `json:"unreachableUrl,omitempty"`
}

//Information about the Resource on the page.
type Page_FrameResource struct {
	// Resource URL.
	Url string `json:"url"`
	// Type of this resource.
	Type Page_ResourceType `json:"type"`
	// Resource mimeType as determined by the browser.
	MimeType string `json:"mimeType"`
	// last-modified timestamp as reported by server.
	LastModified *Network_TimeSinceEpoch `json:"lastModified,omitempty"`
	// Resource content size.
	ContentSize *float32 `json:"contentSize,omitempty"`
	// True if the resource failed to load.
	Failed *bool `json:"failed,omitempty"`
	// True if the resource was canceled during loading.
	Canceled *bool `json:"canceled,omitempty"`
}

//Information about the Frame hierarchy along with their cached resources.
type Page_FrameResourceTree struct {
	// Frame information for this tree item.
	Frame Page_Frame `json:"frame"`
	// Child frames.
	ChildFrames []*Page_FrameResourceTree `json:"childFrames,omitempty"`
	// Information about frame resources.
	Resources []Page_FrameResource `json:"resources"`
}

//Unique script identifier.
type Page_ScriptIdentifier string

//Transition type.
type Page_TransitionType string

//Navigation history entry.
type Page_NavigationEntry struct {
	// Unique id of the navigation history entry.
	Id int `json:"id"`
	// URL of the navigation history entry.
	Url string `json:"url"`
	// URL that the user typed in the url bar.
	UserTypedURL string `json:"userTypedURL"`
	// Title of the navigation history entry.
	Title string `json:"title"`
	// Transition type.
	TransitionType Page_TransitionType `json:"transitionType"`
}

//Screencast frame metadata.
type Page_ScreencastFrameMetadata struct {
	// Top offset in DIP.
	// NOTE Experimental
	OffsetTop float32 `json:"offsetTop"`
	// Page scale factor.
	// NOTE Experimental
	PageScaleFactor float32 `json:"pageScaleFactor"`
	// Device screen width in DIP.
	// NOTE Experimental
	DeviceWidth float32 `json:"deviceWidth"`
	// Device screen height in DIP.
	// NOTE Experimental
	DeviceHeight float32 `json:"deviceHeight"`
	// Position of horizontal scroll in CSS pixels.
	// NOTE Experimental
	ScrollOffsetX float32 `json:"scrollOffsetX"`
	// Position of vertical scroll in CSS pixels.
	// NOTE Experimental
	ScrollOffsetY float32 `json:"scrollOffsetY"`
	// Frame swap timestamp.
	// NOTE Experimental
	Timestamp *Network_TimeSinceEpoch `json:"timestamp,omitempty"`
}

//Javascript dialog type.
type Page_DialogType string

//Error while paring app manifest.
type Page_AppManifestError struct {
	// Error message.
	Message string `json:"message"`
	// If criticial, this is a non-recoverable parse error.
	Critical int `json:"critical"`
	// Error line.
	Line int `json:"line"`
	// Error column.
	Column int `json:"column"`
}

//Proceed: allow the navigation; Cancel: cancel the navigation; CancelAndIgnore: cancels the navigation and makes the requester of the navigation acts like the request was never made.
type Page_NavigationResponse string

//Layout viewport position and dimensions.
type Page_LayoutViewport struct {
	// Horizontal offset relative to the document (CSS pixels).
	PageX int `json:"pageX"`
	// Vertical offset relative to the document (CSS pixels).
	PageY int `json:"pageY"`
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth int `json:"clientWidth"`
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"`
}

//Visual viewport position, dimensions, and scale.
type Page_VisualViewport struct {
	// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetX float32 `json:"offsetX"`
	// Vertical offset relative to the layout viewport (CSS pixels).
	OffsetY float32 `json:"offsetY"`
	// Horizontal offset relative to the document (CSS pixels).
	PageX float32 `json:"pageX"`
	// Vertical offset relative to the document (CSS pixels).
	PageY float32 `json:"pageY"`
	// Width (CSS pixels), excludes scrollbar if present.
	ClientWidth float32 `json:"clientWidth"`
	// Height (CSS pixels), excludes scrollbar if present.
	ClientHeight float32 `json:"clientHeight"`
	// Scale relative to the ideal viewport (size at width=device-width).
	Scale float32 `json:"scale"`
}

//Viewport for capturing screenshot.
type Page_Viewport struct {
	// X offset in CSS pixels.
	X float32 `json:"x"`
	// Y offset in CSS pixels
	Y float32 `json:"y"`
	// Rectangle width in CSS pixels
	Width float32 `json:"width"`
	// Rectangle height in CSS pixels
	Height float32 `json:"height"`
	// Page scale factor.
	Scale float32 `json:"scale"`
}
