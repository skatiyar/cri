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
	Gpu		types.SystemInfo_GPUInfo	`json:"gpu"`// Information about the GPUs on the system.
	ModelName	string				`json:"modelName"`// A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelVersion	string				`json:"modelVersion"`// A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
	CommandLine	string				`json:"commandLine"`// The command line string used to launch the browser. Will be the empty string if not supported.
}, err error) {
	err = obj.conn.Send("SystemInfo.getInfo", nil, &response)
	return
}
