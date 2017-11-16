/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package inspector provides commands and events for Inspector domain.

package inspector

import (
	"github.com/SKatiyar/cri"
)

type Inspector struct {
	conn cri.Connector
}

// New creates a Inspector instance
func New(conn cri.Connector) *Inspector {
	return &Inspector{conn}
}

// Enables inspector domain notifications.
func (obj *Inspector) Enable() (err error) {
	err = obj.conn.Send("Inspector.enable", nil, nil)
	return
}

// Disables inspector domain notifications.
func (obj *Inspector) Disable() (err error) {
	err = obj.conn.Send("Inspector.disable", nil, nil)
	return
}

type DetachedParams struct {
	// The reason why connection has been terminated.
	Reason string `json:"reason"`
}

// Fired when remote debugging connection is about to be terminated. Contains detach reason.
func (obj *Inspector) Detached(fn func(params *DetachedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Inspector.detached", closeChn)
	go func() {
		for {
			params := DetachedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

// Fired when debugging target has crashed
func (obj *Inspector) TargetCrashed(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Inspector.targetCrashed", closeChn)
	go func() {
		for {

			readErr := decoder(nil)
			if !fn(readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
