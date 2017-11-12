/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//DOM Storage identifier.
type DOMStorage_StorageId struct {
	// Security origin for the storage.
	SecurityOrigin string `json:"securityOrigin"`
	// Whether the storage is local storage (not session storage).
	IsLocalStorage bool `json:"isLocalStorage"`
}

//DOM Storage item.
type DOMStorage_Item []string
