/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package systeminfo provides commands and events for SystemInfo domain.
package systeminfo

import (
	"github.com/skatiyar/cri"
	types "github.com/skatiyar/cri/types"
)

// List of commands in SystemInfo domain
const (
	GetInfo = "SystemInfo.getInfo"
)

// The SystemInfo domain defines methods and events for querying low-level system information.
type SystemInfo struct {
	conn cri.Connector
}

// New creates a SystemInfo instance
func New(conn cri.Connector) *SystemInfo {
	return &SystemInfo{conn}
}

type GetInfoResponse struct {
	// Information about the GPUs on the system.
	Gpu types.SystemInfo_GPUInfo `json:"gpu"`
	// A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelName string `json:"modelName"`
	// A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
	ModelVersion string `json:"modelVersion"`
	// The command line string used to launch the browser. Will be the empty string if not supported.
	CommandLine string `json:"commandLine"`
}

// Returns information about the system.
func (obj *SystemInfo) GetInfo() (response GetInfoResponse, err error) {
	err = obj.conn.Send(GetInfo, nil, &response)
	return
}
