/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package inspector provides commands and events for Inspector domain.
package inspector

import (
	"github.com/skatiyar/cri"
)

// List of commands in Inspector domain
const (
	Disable = "Inspector.disable"
	Enable  = "Inspector.enable"
)

// List of events in Inspector domain
const (
	Detached                 = "Inspector.detached"
	TargetCrashed            = "Inspector.targetCrashed"
	TargetReloadedAfterCrash = "Inspector.targetReloadedAfterCrash"
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
func (obj *Inspector) Detached(fn func(event string, params DetachedParams, err error) bool) {
	listen, closer := obj.conn.On(Detached)
	go func() {
		defer closer()
		for {
			var params DetachedParams
			if !fn(Detached, params, listen(&params)) {
				return
			}
		}
	}()
}

// Fired when debugging target has crashed
func (obj *Inspector) TargetCrashed(fn func(event string, err error) bool) {
	listen, closer := obj.conn.On(TargetCrashed)
	go func() {
		defer closer()
		for {
			if !fn(TargetCrashed, listen(nil)) {
				return
			}
		}
	}()
}

// Fired when debugging target has reloaded after crash
func (obj *Inspector) TargetReloadedAfterCrash(fn func(event string, err error) bool) {
	listen, closer := obj.conn.On(TargetReloadedAfterCrash)
	go func() {
		defer closer()
		for {
			if !fn(TargetReloadedAfterCrash, listen(nil)) {
				return
			}
		}
	}()
}
