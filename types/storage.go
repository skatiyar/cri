package types

type Storage_StorageType string
type Storage_UsageForType struct {
	// Name of storage type.
	StorageType Storage_StorageType `json:"storageType"`
	// Storage usage (bytes).
	Usage float32 `json:"usage"`
}
