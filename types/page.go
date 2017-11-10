package types

type Page_ResourceType string
type Page_FrameId string
type Page_Frame struct {
	Id             string           `json:"id"`
	ParentId       *string          `json:"parentId,omitempty"`
	LoaderId       Network_LoaderId `json:"loaderId"`
	Name           *string          `json:"name,omitempty"`
	Url            string           `json:"url"`
	SecurityOrigin string           `json:"securityOrigin"`
	MimeType       string           `json:"mimeType"`
	UnreachableUrl *string          `json:"unreachableUrl,omitempty"`
}
type Page_FrameResource struct {
	Url          string                  `json:"url"`
	Type         Page_ResourceType       `json:"type"`
	MimeType     string                  `json:"mimeType"`
	LastModified *Network_TimeSinceEpoch `json:"lastModified,omitempty"`
	ContentSize  *float32                `json:"contentSize,omitempty"`
	Failed       *bool                   `json:"failed,omitempty"`
	Canceled     *bool                   `json:"canceled,omitempty"`
}
type Page_FrameResourceTree struct {
	Frame       Page_Frame                `json:"frame"`
	ChildFrames []*Page_FrameResourceTree `json:"childFrames,omitempty"`
	Resources   []Page_FrameResource      `json:"resources"`
}
type Page_ScriptIdentifier string
type Page_TransitionType string
type Page_NavigationEntry struct {
	Id             int                 `json:"id"`
	Url            string              `json:"url"`
	UserTypedURL   string              `json:"userTypedURL"`
	Title          string              `json:"title"`
	TransitionType Page_TransitionType `json:"transitionType"`
}
type Page_ScreencastFrameMetadata struct {
	OffsetTop       float32                 `json:"offsetTop"`
	PageScaleFactor float32                 `json:"pageScaleFactor"`
	DeviceWidth     float32                 `json:"deviceWidth"`
	DeviceHeight    float32                 `json:"deviceHeight"`
	ScrollOffsetX   float32                 `json:"scrollOffsetX"`
	ScrollOffsetY   float32                 `json:"scrollOffsetY"`
	Timestamp       *Network_TimeSinceEpoch `json:"timestamp,omitempty"`
}
type Page_DialogType string
type Page_AppManifestError struct {
	Message  string `json:"message"`
	Critical int    `json:"critical"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
}
type Page_NavigationResponse string
type Page_LayoutViewport struct {
	PageX        int `json:"pageX"`
	PageY        int `json:"pageY"`
	ClientWidth  int `json:"clientWidth"`
	ClientHeight int `json:"clientHeight"`
}
type Page_VisualViewport struct {
	OffsetX      float32 `json:"offsetX"`
	OffsetY      float32 `json:"offsetY"`
	PageX        float32 `json:"pageX"`
	PageY        float32 `json:"pageY"`
	ClientWidth  float32 `json:"clientWidth"`
	ClientHeight float32 `json:"clientHeight"`
	Scale        float32 `json:"scale"`
}
type Page_Viewport struct {
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
	Scale  float32 `json:"scale"`
}
