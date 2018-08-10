/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Detailed application cache resource information.
type ApplicationCache_ApplicationCacheResource struct {
	// Resource url.
	Url string `json:"url"`
	// Resource size.
	Size int `json:"size"`
	// Resource type.
	Type string `json:"type"`
}

// Detailed application cache information.
type ApplicationCache_ApplicationCache struct {
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Application cache size.
	Size float32 `json:"size"`
	// Application cache creation time.
	CreationTime float32 `json:"creationTime"`
	// Application cache update time.
	UpdateTime float32 `json:"updateTime"`
	// Application cache resources.
	Resources []ApplicationCache_ApplicationCacheResource `json:"resources"`
}

// Frame identifier - manifest URL pair.
type ApplicationCache_FrameWithManifest struct {
	// Frame identifier.
	FrameId Page_FrameId `json:"frameId"`
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Application cache status.
	Status int `json:"status"`
}
