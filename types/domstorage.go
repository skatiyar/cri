package types

type DOMStorage_StorageId struct {
	SecurityOrigin string `json:"securityOrigin"`
	IsLocalStorage bool   `json:"isLocalStorage"`
}
type DOMStorage_Item []string
