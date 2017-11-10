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
	FrameIds []types.ApplicationCache_FrameWithManifest `json:"frameIds"`
}, err error) {
	err = obj.conn.Send("ApplicationCache.getFramesWithManifests", nil, &response)
	return
}
func (obj *ApplicationCache) Enable() (err error) {
	err = obj.conn.Send("ApplicationCache.enable", nil, nil)
	return
}

type GetManifestForFrameRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`
}

func (obj *ApplicationCache) GetManifestForFrame(request *GetManifestForFrameRequest) (response struct {
	ManifestURL string `json:"manifestURL"`
}, err error) {
	err = obj.conn.Send("ApplicationCache.getManifestForFrame", request, &response)
	return
}

type GetApplicationCacheForFrameRequest struct {
	FrameId types.Page_FrameId `json:"frameId"`
}

func (obj *ApplicationCache) GetApplicationCacheForFrame(request *GetApplicationCacheForFrameRequest) (response struct {
	ApplicationCache types.ApplicationCache_ApplicationCache `json:"applicationCache"`
}, err error) {
	err = obj.conn.Send("ApplicationCache.getApplicationCacheForFrame", request, &response)
	return
}
