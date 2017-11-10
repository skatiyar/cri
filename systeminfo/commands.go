package systeminfo

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type SystemInfo struct {
	conn cri.Connector
}

func New(conn cri.Connector) *SystemInfo {
	return &SystemInfo{conn}
}
func (obj *SystemInfo) GetInfo() (response struct {
	Gpu          types.SystemInfo_GPUInfo `json:"gpu"`
	ModelName    string                   `json:"modelName"`
	ModelVersion string                   `json:"modelVersion"`
	CommandLine  string                   `json:"commandLine"`
}, err error) {
	err = obj.conn.Send("SystemInfo.getInfo", nil, &response)
	return
}
