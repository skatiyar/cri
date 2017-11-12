/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

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
func (obj *Inspector) Detached(fn func(params *DetachedParams) bool) {
	params := DetachedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Inspector.detached", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}

// Fired when debugging target has crashed
func (obj *Inspector) TargetCrashed(fn func() bool) {

	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Inspector.targetCrashed", closeChn, nil)
			if !fn() {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}
