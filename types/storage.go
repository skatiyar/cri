/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Enum of possible storage types.
type Storage_StorageType string

//Usage for a storage type.
type Storage_UsageForType struct {
	// Name of storage type.
	StorageType Storage_StorageType `json:"storageType"`
	// Storage usage (bytes).
	Usage float32 `json:"usage"`
}
