package storage

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Storage struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Storage {
	return &Storage{conn}
}

type ClearDataForOriginRequest struct {
	Origin		string	`json:"origin"`// Security origin.
	StorageTypes	string	`json:"storageTypes"`// Comma separated origin names.
}

func (obj *Storage) ClearDataForOrigin(request *ClearDataForOriginRequest) (err error) {
	err = obj.conn.Send("Storage.clearDataForOrigin", request, nil)
	return
}

type GetUsageAndQuotaRequest struct {
	Origin string `json:"origin"`// Security origin.
}

func (obj *Storage) GetUsageAndQuota(request *GetUsageAndQuotaRequest) (response struct {
	Usage		float32				`json:"usage"`// Storage usage (bytes).
	Quota		float32				`json:"quota"`// Storage quota (bytes).
	UsageBreakdown	[]types.Storage_UsageForType	`json:"usageBreakdown"`// Storage usage per type (bytes).
}, err error) {
	err = obj.conn.Send("Storage.getUsageAndQuota", request, &response)
	return
}

type TrackCacheStorageForOriginRequest struct {
	Origin string `json:"origin"`// Security origin.
}

func (obj *Storage) TrackCacheStorageForOrigin(request *TrackCacheStorageForOriginRequest) (err error) {
	err = obj.conn.Send("Storage.trackCacheStorageForOrigin", request, nil)
	return
}

type UntrackCacheStorageForOriginRequest struct {
	Origin string `json:"origin"`// Security origin.
}

func (obj *Storage) UntrackCacheStorageForOrigin(request *UntrackCacheStorageForOriginRequest) (err error) {
	err = obj.conn.Send("Storage.untrackCacheStorageForOrigin", request, nil)
	return
}

type TrackIndexedDBForOriginRequest struct {
	Origin string `json:"origin"`// Security origin.
}

func (obj *Storage) TrackIndexedDBForOrigin(request *TrackIndexedDBForOriginRequest) (err error) {
	err = obj.conn.Send("Storage.trackIndexedDBForOrigin", request, nil)
	return
}

type UntrackIndexedDBForOriginRequest struct {
	Origin string `json:"origin"`// Security origin.
}

func (obj *Storage) UntrackIndexedDBForOrigin(request *UntrackIndexedDBForOriginRequest) (err error) {
	err = obj.conn.Send("Storage.untrackIndexedDBForOrigin", request, nil)
	return
}
