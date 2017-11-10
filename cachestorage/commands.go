package cachestorage

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type CacheStorage struct {
	conn cri.Connector
}

func New(conn cri.Connector) *CacheStorage {
	return &CacheStorage{conn}
}

type RequestCacheNamesRequest struct {
	SecurityOrigin string `json:"securityOrigin"`
}

func (obj *CacheStorage) RequestCacheNames(request *RequestCacheNamesRequest) (response struct {
	Caches []types.CacheStorage_Cache `json:"caches"`
}, err error) {
	err = obj.conn.Send("CacheStorage.requestCacheNames", request, &response)
	return
}

type RequestEntriesRequest struct {
	CacheId   types.CacheStorage_CacheId `json:"cacheId"`
	SkipCount int                        `json:"skipCount"`
	PageSize  int                        `json:"pageSize"`
}

func (obj *CacheStorage) RequestEntries(request *RequestEntriesRequest) (response struct {
	CacheDataEntries []types.CacheStorage_DataEntry `json:"cacheDataEntries"`
	HasMore          bool                           `json:"hasMore"`
}, err error) {
	err = obj.conn.Send("CacheStorage.requestEntries", request, &response)
	return
}

type DeleteCacheRequest struct {
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
}

func (obj *CacheStorage) DeleteCache(request *DeleteCacheRequest) (err error) {
	err = obj.conn.Send("CacheStorage.deleteCache", request, nil)
	return
}

type DeleteEntryRequest struct {
	CacheId types.CacheStorage_CacheId `json:"cacheId"`
	Request string                     `json:"request"`
}

func (obj *CacheStorage) DeleteEntry(request *DeleteEntryRequest) (err error) {
	err = obj.conn.Send("CacheStorage.deleteEntry", request, nil)
	return
}

type RequestCachedResponseRequest struct {
	CacheId    types.CacheStorage_CacheId `json:"cacheId"`
	RequestURL string                     `json:"requestURL"`
}

func (obj *CacheStorage) RequestCachedResponse(request *RequestCachedResponseRequest) (response struct {
	Response types.CacheStorage_CachedResponse `json:"response"`
}, err error) {
	err = obj.conn.Send("CacheStorage.requestCachedResponse", request, &response)
	return
}
