package deviceorientation

import "github.com/SKatiyar/cri"

type DeviceOrientation struct {
	conn cri.Connector
}

func New(conn cri.Connector) *DeviceOrientation {
	return &DeviceOrientation{conn}
}

type SetDeviceOrientationOverrideRequest struct {
	Alpha	float32	`json:"alpha"`// Mock alpha
	Beta	float32	`json:"beta"`// Mock beta
	Gamma	float32	`json:"gamma"`// Mock gamma
}

func (obj *DeviceOrientation) SetDeviceOrientationOverride(request *SetDeviceOrientationOverrideRequest) (err error) {
	err = obj.conn.Send("DeviceOrientation.setDeviceOrientationOverride", request, nil)
	return
}
func (obj *DeviceOrientation) ClearDeviceOrientationOverride() (err error) {
	err = obj.conn.Send("DeviceOrientation.clearDeviceOrientationOverride", nil, nil)
	return
}
