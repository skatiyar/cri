package memory

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Memory struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Memory {
	return &Memory{conn}
}
func (obj *Memory) GetDOMCounters() (response struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}, err error) {
	err = obj.conn.Send("Memory.getDOMCounters", nil, &response)
	return
}
func (obj *Memory) PrepareForLeakDetection() (err error) {
	err = obj.conn.Send("Memory.prepareForLeakDetection", nil, nil)
	return
}

type SetPressureNotificationsSuppressedRequest struct {
	Suppressed bool `json:"suppressed"`
}

func (obj *Memory) SetPressureNotificationsSuppressed(request *SetPressureNotificationsSuppressedRequest) (err error) {
	err = obj.conn.Send("Memory.setPressureNotificationsSuppressed", request, nil)
	return
}

type SimulatePressureNotificationRequest struct {
	Level types.Memory_PressureLevel `json:"level"`
}

func (obj *Memory) SimulatePressureNotification(request *SimulatePressureNotificationRequest) (err error) {
	err = obj.conn.Send("Memory.simulatePressureNotification", request, nil)
	return
}
