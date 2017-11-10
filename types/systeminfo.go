package types

type SystemInfo_GPUDevice struct {
	VendorId     float32 `json:"vendorId"`
	DeviceId     float32 `json:"deviceId"`
	VendorString string  `json:"vendorString"`
	DeviceString string  `json:"deviceString"`
}
type SystemInfo_GPUInfo struct {
	Devices              []SystemInfo_GPUDevice  `json:"devices"`
	AuxAttributes        *map[string]interface{} `json:"auxAttributes,omitempty"`
	FeatureStatus        *map[string]interface{} `json:"featureStatus,omitempty"`
	DriverBugWorkarounds []string                `json:"driverBugWorkarounds"`
}
