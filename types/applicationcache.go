package types

type ApplicationCache_ApplicationCacheResource struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
	Type string `json:"type"`
}
type ApplicationCache_ApplicationCache struct {
	ManifestURL  string                                      `json:"manifestURL"`
	Size         float32                                     `json:"size"`
	CreationTime float32                                     `json:"creationTime"`
	UpdateTime   float32                                     `json:"updateTime"`
	Resources    []ApplicationCache_ApplicationCacheResource `json:"resources"`
}
type ApplicationCache_FrameWithManifest struct {
	FrameId     Page_FrameId `json:"frameId"`
	ManifestURL string       `json:"manifestURL"`
	Status      int          `json:"status"`
}
