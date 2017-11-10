package applicationcache

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type ApplicationCache struct {
	conn cri.Connector
}

func New(conn cri.Connector) *ApplicationCache {
	return &ApplicationCache{conn}
}
func (obj *ApplicationCache) GetFramesWithManifests() (response struct {
	FrameIds []types.ApplicationCache_FrameWithManifest `json:"frameIds"`// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}, err error) {
	err = obj.conn.Send("ApplicationCache.getFramesWithManifests", nil, &response)
	return
}
func (obj *ApplicationCache) Enable() (err error) {
	err = obj.conn.Send("ApplicationCache.enable", nil, nil)
	return
}

type GetManifestForFrameRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`// Identifier of the frame containing document whose manifest is retrieved.
}

func (obj *ApplicationCache) GetManifestForFrame(request *GetManifestForFrameRequest) (response struct {
	ManifestURL string `json:"manifestURL"`// Manifest URL for document in the given frame.
}, err error) {
	err = obj.conn.Send("ApplicationCache.getManifestForFrame", request, &response)
	return
}

type GetApplicationCacheForFrameRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`// Identifier of the frame containing document whose application cache is retrieved.
}

func (obj *ApplicationCache) GetApplicationCacheForFrame(request *GetApplicationCacheForFrameRequest) (response struct {
	ApplicationCache types.ApplicationCache_ApplicationCache `json:"applicationCache"`// Relevant application cache data for the document in given frame.
}, err error) {
	err = obj.conn.Send("ApplicationCache.getApplicationCacheForFrame", request, &response)
	return
}
