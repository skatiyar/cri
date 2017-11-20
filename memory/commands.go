/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package memory provides commands and events for Memory domain.
package memory

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Memory domain
const (
	GetDOMCounters                     = "Memory.getDOMCounters"
	PrepareForLeakDetection            = "Memory.prepareForLeakDetection"
	SetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"
	SimulatePressureNotification       = "Memory.simulatePressureNotification"
)

type Memory struct {
	conn cri.Connector
}

// New creates a Memory instance
func New(conn cri.Connector) *Memory {
	return &Memory{conn}
}

type GetDOMCountersResponse struct {
	Documents        int `json:"documents"`
	Nodes            int `json:"nodes"`
	JsEventListeners int `json:"jsEventListeners"`
}

func (obj *Memory) GetDOMCounters() (response GetDOMCountersResponse, err error) {
	err = obj.conn.Send(GetDOMCounters, nil, &response)
	return
}

func (obj *Memory) PrepareForLeakDetection() (err error) {
	err = obj.conn.Send(PrepareForLeakDetection, nil, nil)
	return
}

type SetPressureNotificationsSuppressedRequest struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// Enable/disable suppressing memory pressure notifications in all processes.
func (obj *Memory) SetPressureNotificationsSuppressed(request *SetPressureNotificationsSuppressedRequest) (err error) {
	err = obj.conn.Send(SetPressureNotificationsSuppressed, request, nil)
	return
}

type SimulatePressureNotificationRequest struct {
	// Memory pressure level of the notification.
	Level types.Memory_PressureLevel `json:"level"`
}

// Simulate a memory pressure notification in all processes.
func (obj *Memory) SimulatePressureNotification(request *SimulatePressureNotificationRequest) (err error) {
	err = obj.conn.Send(SimulatePressureNotification, request, nil)
	return
}
