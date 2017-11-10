package types

type Storage_StorageType string
type Storage_UsageForType struct {
	StorageType	Storage_StorageType	`json:"storageType"`// Name of storage type.
	Usage		float32			`json:"usage"`// Storage usage (bytes).
}
