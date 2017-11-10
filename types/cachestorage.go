package types

type CacheStorage_CacheId string
type CacheStorage_DataEntry struct {
	RequestURL         string                `json:"requestURL"`
	RequestMethod      string                `json:"requestMethod"`
	RequestHeaders     []CacheStorage_Header `json:"requestHeaders"`
	ResponseTime       float32               `json:"responseTime"`
	ResponseStatus     int                   `json:"responseStatus"`
	ResponseStatusText string                `json:"responseStatusText"`
	ResponseHeaders    []CacheStorage_Header `json:"responseHeaders"`
}
type CacheStorage_Cache struct {
	CacheId        CacheStorage_CacheId `json:"cacheId"`
	SecurityOrigin string               `json:"securityOrigin"`
	CacheName      string               `json:"cacheName"`
}
type CacheStorage_Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type CacheStorage_CachedResponse struct {
	Body string `json:"body"`
}
