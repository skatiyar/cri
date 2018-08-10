/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Describes a single graphics processor (GPU).
type SystemInfo_GPUDevice struct {
	// PCI ID of the GPU vendor, if available; 0 otherwise.
	VendorId float32 `json:"vendorId"`
	// PCI ID of the GPU device, if available; 0 otherwise.
	DeviceId float32 `json:"deviceId"`
	// String description of the GPU vendor, if the PCI ID is not available.
	VendorString string `json:"vendorString"`
	// String description of the GPU device, if the PCI ID is not available.
	DeviceString string `json:"deviceString"`
}

// Provides information about the GPU(s) on the system.
type SystemInfo_GPUInfo struct {
	// The graphics devices on the system. Element 0 is the primary GPU.
	Devices []SystemInfo_GPUDevice `json:"devices"`
	// An optional dictionary of additional GPU related attributes.
	AuxAttributes map[string]interface{} `json:"auxAttributes,omitempty"`
	// An optional dictionary of graphics features and their status.
	FeatureStatus map[string]interface{} `json:"featureStatus,omitempty"`
	// An optional array of GPU driver bug workarounds.
	DriverBugWorkarounds []string `json:"driverBugWorkarounds"`
}
