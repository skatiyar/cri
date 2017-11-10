package types

type Page_ResourceType string
type Page_FrameId string
type Page_Frame struct {
	Id		string			`json:"id"`// Frame unique identifier.
	ParentId	*string			`json:"parentId,omitempty"`// Parent frame identifier.
	LoaderId	Network_LoaderId	`json:"loaderId"`// Identifier of the loader associated with this frame.
	Name		*string			`json:"name,omitempty"`// Frame's name as specified in the tag.
	Url		string			`json:"url"`// Frame document's URL.
	SecurityOrigin	string			`json:"securityOrigin"`// Frame document's security origin.
	MimeType	string			`json:"mimeType"`// Frame document's mimeType as determined by the browser.
	UnreachableUrl	*string			`json:"unreachableUrl,omitempty"`// If the frame failed to load, this contains the URL that could not be loaded.
}
type Page_FrameResource struct {
	Url		string			`json:"url"`// Resource URL.
	Type		Page_ResourceType	`json:"type"`// Type of this resource.
	MimeType	string			`json:"mimeType"`// Resource mimeType as determined by the browser.
	LastModified	*Network_TimeSinceEpoch	`json:"lastModified,omitempty"`// last-modified timestamp as reported by server.
	ContentSize	*float32		`json:"contentSize,omitempty"`// Resource content size.
	Failed		*bool			`json:"failed,omitempty"`// True if the resource failed to load.
	Canceled	*bool			`json:"canceled,omitempty"`// True if the resource was canceled during loading.
}
type Page_FrameResourceTree struct {
	Frame		Page_Frame			`json:"frame"`// Frame information for this tree item.
	ChildFrames	[]*Page_FrameResourceTree	`json:"childFrames,omitempty"`// Child frames.
	Resources	[]Page_FrameResource		`json:"resources"`// Information about frame resources.
}
type Page_ScriptIdentifier string
type Page_TransitionType string
type Page_NavigationEntry struct {
	Id		int			`json:"id"`// Unique id of the navigation history entry.
	Url		string			`json:"url"`// URL of the navigation history entry.
	UserTypedURL	string			`json:"userTypedURL"`// URL that the user typed in the url bar.
	Title		string			`json:"title"`// Title of the navigation history entry.
	TransitionType	Page_TransitionType	`json:"transitionType"`// Transition type.
}
type Page_ScreencastFrameMetadata struct {
	OffsetTop	float32			`json:"offsetTop"`// Top offset in DIP.
	PageScaleFactor	float32			`json:"pageScaleFactor"`// Page scale factor.
	DeviceWidth	float32			`json:"deviceWidth"`// Device screen width in DIP.
	DeviceHeight	float32			`json:"deviceHeight"`// Device screen height in DIP.
	ScrollOffsetX	float32			`json:"scrollOffsetX"`// Position of horizontal scroll in CSS pixels.
	ScrollOffsetY	float32			`json:"scrollOffsetY"`// Position of vertical scroll in CSS pixels.
	Timestamp	*Network_TimeSinceEpoch	`json:"timestamp,omitempty"`// Frame swap timestamp.
}
type Page_DialogType string
type Page_AppManifestError struct {
	Message		string	`json:"message"`// Error message.
	Critical	int	`json:"critical"`// If criticial, this is a non-recoverable parse error.
	Line		int	`json:"line"`// Error line.
	Column		int	`json:"column"`// Error column.
}
type Page_NavigationResponse string
type Page_LayoutViewport struct {
	PageX		int	`json:"pageX"`// Horizontal offset relative to the document (CSS pixels).
	PageY		int	`json:"pageY"`// Vertical offset relative to the document (CSS pixels).
	ClientWidth	int	`json:"clientWidth"`// Width (CSS pixels), excludes scrollbar if present.
	ClientHeight	int	`json:"clientHeight"`// Height (CSS pixels), excludes scrollbar if present.
}
type Page_VisualViewport struct {
	OffsetX		float32	`json:"offsetX"`// Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY		float32	`json:"offsetY"`// Vertical offset relative to the layout viewport (CSS pixels).
	PageX		float32	`json:"pageX"`// Horizontal offset relative to the document (CSS pixels).
	PageY		float32	`json:"pageY"`// Vertical offset relative to the document (CSS pixels).
	ClientWidth	float32	`json:"clientWidth"`// Width (CSS pixels), excludes scrollbar if present.
	ClientHeight	float32	`json:"clientHeight"`// Height (CSS pixels), excludes scrollbar if present.
	Scale		float32	`json:"scale"`// Scale relative to the ideal viewport (size at width=device-width).
}
type Page_Viewport struct {
	X	float32	`json:"x"`// X offset in CSS pixels.
	Y	float32	`json:"y"`// Y offset in CSS pixels
	Width	float32	`json:"width"`// Rectangle width in CSS pixels
	Height	float32	`json:"height"`// Rectangle height in CSS pixels
	Scale	float32	`json:"scale"`// Page scale factor.
}
