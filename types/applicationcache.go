package types

type ApplicationCache_ApplicationCacheResource struct {
	Url	string	`json:"url"`// Resource url.
	Size	int	`json:"size"`// Resource size.
	Type	string	`json:"type"`// Resource type.
}
type ApplicationCache_ApplicationCache struct {
	ManifestURL	string						`json:"manifestURL"`// Manifest URL.
	Size		float32						`json:"size"`// Application cache size.
	CreationTime	float32						`json:"creationTime"`// Application cache creation time.
	UpdateTime	float32						`json:"updateTime"`// Application cache update time.
	Resources	[]ApplicationCache_ApplicationCacheResource	`json:"resources"`// Application cache resources.
}
type ApplicationCache_FrameWithManifest struct {
	FrameId		Page_FrameId	`json:"frameId"`// Frame identifier.
	ManifestURL	string		`json:"manifestURL"`// Manifest URL.
	Status		int		`json:"status"`// Application cache status.
}
