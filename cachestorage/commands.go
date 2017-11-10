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
	SecurityOrigin string `json:"securityOrigin"`// Security origin.
}

func (obj *CacheStorage) RequestCacheNames(request *RequestCacheNamesRequest) (response struct {
	Caches []types.CacheStorage_Cache `json:"caches"`// Caches for the security origin.
}, err error) {
	err = obj.conn.Send("CacheStorage.requestCacheNames", request, &response)
	return
}

type RequestEntriesRequest struct {
	CacheId		types.CacheStorage_CacheId	`json:"cacheId"`// ID of cache to get entries from.
	SkipCount	int				`json:"skipCount"`// Number of records to skip.
	PageSize	int				`json:"pageSize"`// Number of records to fetch.
}

func (obj *CacheStorage) RequestEntries(request *RequestEntriesRequest) (response struct {
	CacheDataEntries	[]types.CacheStorage_DataEntry	`json:"cacheDataEntries"`// Array of object store data entries.
	HasMore			bool				`json:"hasMore"`// If true, there are more entries to fetch in the given range.
}, err error) {
	err = obj.conn.Send("CacheStorage.requestEntries", request, &response)
	return
}

type DeleteCacheRequest struct {
	CacheId types.CacheStorage_CacheId `json:"cacheId"`// Id of cache for deletion.
}

func (obj *CacheStorage) DeleteCache(request *DeleteCacheRequest) (err error) {
	err = obj.conn.Send("CacheStorage.deleteCache", request, nil)
	return
}

type DeleteEntryRequest struct {
	CacheId	types.CacheStorage_CacheId	`json:"cacheId"`// Id of cache where the entry will be deleted.
	Request	string				`json:"request"`// URL spec of the request.
}

func (obj *CacheStorage) DeleteEntry(request *DeleteEntryRequest) (err error) {
	err = obj.conn.Send("CacheStorage.deleteEntry", request, nil)
	return
}

type RequestCachedResponseRequest struct {
	CacheId		types.CacheStorage_CacheId	`json:"cacheId"`// Id of cache that contains the enty.
	RequestURL	string				`json:"requestURL"`// URL spec of the request.
}

func (obj *CacheStorage) RequestCachedResponse(request *RequestCachedResponseRequest) (response struct {
	Response types.CacheStorage_CachedResponse `json:"response"`// Response read from the cache.
}, err error) {
	err = obj.conn.Send("CacheStorage.requestCachedResponse", request, &response)
	return
}
