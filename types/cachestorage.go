/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Unique identifier of the Cache object.
type CacheStorage_CacheId string

//Data entry.
type CacheStorage_DataEntry struct {
	// Request URL.
	RequestURL string `json:"requestURL"`
	// Request method.
	RequestMethod string `json:"requestMethod"`
	// Request headers
	RequestHeaders []CacheStorage_Header `json:"requestHeaders"`
	// Number of seconds since epoch.
	ResponseTime float32 `json:"responseTime"`
	// HTTP response status code.
	ResponseStatus int `json:"responseStatus"`
	// HTTP response status text.
	ResponseStatusText string `json:"responseStatusText"`
	// Response headers
	ResponseHeaders []CacheStorage_Header `json:"responseHeaders"`
}

//Cache identifier.
type CacheStorage_Cache struct {
	// An opaque unique id of the cache.
	CacheId CacheStorage_CacheId `json:"cacheId"`
	// Security origin of the cache.
	SecurityOrigin string `json:"securityOrigin"`
	// The name of the cache.
	CacheName string `json:"cacheName"`
}

type CacheStorage_Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Cached response
type CacheStorage_CachedResponse struct {
	// Entry content, base64-encoded.
	Body string `json:"body"`
}
