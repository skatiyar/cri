/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package applicationcache provides commands and events for ApplicationCache domain.
package applicationcache

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in ApplicationCache domain
const (
	Enable                      = "ApplicationCache.enable"
	GetApplicationCacheForFrame = "ApplicationCache.getApplicationCacheForFrame"
	GetFramesWithManifests      = "ApplicationCache.getFramesWithManifests"
	GetManifestForFrame         = "ApplicationCache.getManifestForFrame"
)

// List of events in ApplicationCache domain
const (
	ApplicationCacheStatusUpdated = "ApplicationCache.applicationCacheStatusUpdated"
	NetworkStateUpdated           = "ApplicationCache.networkStateUpdated"
)

type ApplicationCache struct {
	conn cri.Connector
}

// New creates a ApplicationCache instance
func New(conn cri.Connector) *ApplicationCache {
	return &ApplicationCache{conn}
}

// Enables application cache domain notifications.
func (obj *ApplicationCache) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
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
	err = obj.conn.Send(GetApplicationCacheForFrame, request, &response)
	return
}

type GetFramesWithManifestsResponse struct {
	// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
	FrameIds []types.ApplicationCache_FrameWithManifest `json:"frameIds"`
}

// Returns array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
func (obj *ApplicationCache) GetFramesWithManifests() (response GetFramesWithManifestsResponse, err error) {
	err = obj.conn.Send(GetFramesWithManifests, nil, &response)
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
	err = obj.conn.Send(GetManifestForFrame, request, &response)
	return
}

type ApplicationCacheStatusUpdatedParams struct {
	// Identifier of the frame containing document whose application cache updated status.
	FrameId types.Page_FrameId `json:"frameId"`
	// Manifest URL.
	ManifestURL string `json:"manifestURL"`
	// Updated application cache status.
	Status int `json:"status"`
}

func (obj *ApplicationCache) ApplicationCacheStatusUpdated(fn func(event string, params ApplicationCacheStatusUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(ApplicationCacheStatusUpdated)
	go func() {
		defer closer()
		for {
			var params ApplicationCacheStatusUpdatedParams
			if !fn(ApplicationCacheStatusUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type NetworkStateUpdatedParams struct {
	IsNowOnline bool `json:"isNowOnline"`
}

func (obj *ApplicationCache) NetworkStateUpdated(fn func(event string, params NetworkStateUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(NetworkStateUpdated)
	go func() {
		defer closer()
		for {
			var params NetworkStateUpdatedParams
			if !fn(NetworkStateUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}
