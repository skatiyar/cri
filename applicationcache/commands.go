/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/


package applicationcache

import (
    "github.com/SKatiyar/cri"
    types "github.com/SKatiyar/cri/types"
)

type ApplicationCache struct {
	conn cri.Connector
}

// New creates a ApplicationCache instance
func New(conn cri.Connector) *ApplicationCache {
	return &ApplicationCache{conn}
}

type GetFramesWithManifestsResponse struct {
	// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
	FrameIds []types.ApplicationCache_FrameWithManifest `json:"frameIds"`
}

// Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (obj *ApplicationCache) GetFramesWithManifests() (response GetFramesWithManifestsResponse, err error) {
	err = obj.conn.Send("ApplicationCache.getFramesWithManifests", nil, &response)
	return
}

// Enables application cache domain notifications.
func (obj *ApplicationCache) Enable() (err error) {
	err = obj.conn.Send("ApplicationCache.enable", nil, nil)
	return
}


type GetManifestForFrameRequest struct {
	// Identifier of the frame containing document whose manifest is retrieved.
	FrameId types.Page_FrameId `json:"frameId"`
}


type GetManifestForFrameResponse struct {
	// Manifest URL for document in the given frame.
	ManifestURL string `json:"manifestURL"`
}

// Returns manifest URL for document in the given frame.
func (obj *ApplicationCache) GetManifestForFrame(request *GetManifestForFrameRequest) (response GetManifestForFrameResponse, err error) {
	err = obj.conn.Send("ApplicationCache.getManifestForFrame", request, &response)
	return
}


type GetApplicationCacheForFrameRequest struct {
	// Identifier of the frame containing document whose application cache is retrieved.
	FrameId types.Page_FrameId `json:"frameId"`
}


type GetApplicationCacheForFrameResponse struct {
	// Relevant application cache data for the document in given frame.
	ApplicationCache types.ApplicationCache_ApplicationCache `json:"applicationCache"`
}

// Returns relevant application cache data for the document in given frame.
func (obj *ApplicationCache) GetApplicationCacheForFrame(request *GetApplicationCacheForFrameRequest) (response GetApplicationCacheForFrameResponse, err error) {
	err = obj.conn.Send("ApplicationCache.getApplicationCacheForFrame", request, &response)
	return
}
