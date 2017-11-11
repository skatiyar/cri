package types

type DOMStorage_StorageId struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`
	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}
type DOMStorage_Item []string
