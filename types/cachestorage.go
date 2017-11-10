package types

type CacheStorage_CacheId string
type CacheStorage_DataEntry struct {
	RequestURL		string			`json:"requestURL"`// Request URL.
	RequestMethod		string			`json:"requestMethod"`// Request method.
	RequestHeaders		[]CacheStorage_Header	`json:"requestHeaders"`// Request headers
	ResponseTime		float32			`json:"responseTime"`// Number of seconds since epoch.
	ResponseStatus		int			`json:"responseStatus"`// HTTP response status code.
	ResponseStatusText	string			`json:"responseStatusText"`// HTTP response status text.
	ResponseHeaders		[]CacheStorage_Header	`json:"responseHeaders"`// Response headers
}
type CacheStorage_Cache struct {
	CacheId		CacheStorage_CacheId	`json:"cacheId"`// An opaque unique id of the cache.
	SecurityOrigin	string			`json:"securityOrigin"`// Security origin of the cache.
	CacheName	string			`json:"cacheName"`// The name of the cache.
}
type CacheStorage_Header struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
type CacheStorage_CachedResponse struct {
	Body string `json:"body"`// Entry content, base64-encoded.
}
