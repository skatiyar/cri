/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package cachestorage provides commands and events for CacheStorage domain.
package cachestorage

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in CacheStorage domain
const (
	DeleteCache           = "CacheStorage.deleteCache"
	DeleteEntry           = "CacheStorage.deleteEntry"
	RequestCacheNames     = "CacheStorage.requestCacheNames"
	RequestCachedResponse = "CacheStorage.requestCachedResponse"
	RequestEntries        = "CacheStorage.requestEntries"
)

type CacheStorage struct {
	conn cri.Connector
}

// New creates a CacheStorage instance
func New(conn cri.Connector) *CacheStorage {
	return &CacheStorage{conn}
}

type DeleteCacheRequest struct {
	// Id of cache for deletion.
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
}

// Deletes a cache.
func (obj *CacheStorage) DeleteCache(request *DeleteCacheRequest) (err error) {
	err = obj.conn.Send(DeleteCache, request, nil)
	return
}

type DeleteEntryRequest struct {
	// Id of cache where the entry will be deleted.
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
	// URL spec of the request.
	Request string `json:"request"`
}

// Deletes a cache entry.
func (obj *CacheStorage) DeleteEntry(request *DeleteEntryRequest) (err error) {
	err = obj.conn.Send(DeleteEntry, request, nil)
	return
}

type RequestCacheNamesRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

type RequestCacheNamesResponse struct {
	// Caches for the security origin.
	Caches []types.CacheStorage_Cache `json:"caches"`
}

// Requests cache names.
func (obj *CacheStorage) RequestCacheNames(request *RequestCacheNamesRequest) (response RequestCacheNamesResponse, err error) {
	err = obj.conn.Send(RequestCacheNames, request, &response)
	return
}

type RequestCachedResponseRequest struct {
	// Id of cache that contains the enty.
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
	// URL spec of the request.
	RequestURL string `json:"requestURL"`
}

type RequestCachedResponseResponse struct {
	// Response read from the cache.
	Response types.CacheStorage_CachedResponse `json:"response"`
}

// Fetches cache entry.
func (obj *CacheStorage) RequestCachedResponse(request *RequestCachedResponseRequest) (response RequestCachedResponseResponse, err error) {
	err = obj.conn.Send(RequestCachedResponse, request, &response)
	return
}

type RequestEntriesRequest struct {
	// ID of cache to get entries from.
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
	// Number of records to skip.
	SkipCount int `json:"skipCount"`
	// Number of records to fetch.
	PageSize int `json:"pageSize"`
}

type RequestEntriesResponse struct {
	// Array of object store data entries.
	CacheDataEntries []types.CacheStorage_DataEntry `json:"cacheDataEntries"`
	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Requests data from cache.
func (obj *CacheStorage) RequestEntries(request *RequestEntriesRequest) (response RequestEntriesResponse, err error) {
	err = obj.conn.Send(RequestEntries, request, &response)
	return
}
