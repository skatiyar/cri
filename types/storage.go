package types

type Storage_StorageType string
type Storage_UsageForType struct {
	StorageType Storage_StorageType `json:"storageType"`
	Usage       float32             `json:"usage"`
}
