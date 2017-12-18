/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package inspector provides commands and events for Inspector domain.
package inspector

import (
	"github.com/SKatiyar/cri"
)

// List of commands in Inspector domain
const (
	Disable = "Inspector.disable"
	Enable  = "Inspector.enable"
)

// List of events in Inspector domain
const (
	Detached      = "Inspector.detached"
	TargetCrashed = "Inspector.targetCrashed"
)

type Inspector struct {
	conn cri.Connector
}

// New creates a Inspector instance
func New(conn cri.Connector) *Inspector {
	return &Inspector{conn}
}

// Disables inspector domain notifications.
func (obj *Inspector) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables inspector domain notifications.
func (obj *Inspector) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type DetachedParams struct {
	// The reason why connection has been terminated.
	Reason string `json:"reason"`
}

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
func (obj *Inspector) Detached() (params DetachedParams, err error) {
	err = obj.conn.On(Detached, &params)
	return
}

// Fired when debugging target has crashed
func (obj *Inspector) TargetCrashed() (err error) {
	err = obj.conn.On(TargetCrashed, nil)
	return
}
