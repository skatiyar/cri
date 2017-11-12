/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// The Tethering domain defines methods and events for browser port binding.
package tethering

import (
	"github.com/SKatiyar/cri"
)

type Tethering struct {
	conn cri.Connector
}

// New creates a Tethering instance
func New(conn cri.Connector) *Tethering {
	return &Tethering{conn}
}

type BindRequest struct {
	// Port number to bind.
	Port int `json:"port"`
}

// Request browser port binding.
func (obj *Tethering) Bind(request *BindRequest) (err error) {
	err = obj.conn.Send("Tethering.bind", request, nil)
	return
}

type UnbindRequest struct {
	// Port number to unbind.
	Port int `json:"port"`
}

// Request browser port unbinding.
func (obj *Tethering) Unbind(request *UnbindRequest) (err error) {
	err = obj.conn.Send("Tethering.unbind", request, nil)
	return
}

type AcceptedParams struct {
	// Port number that was successfully bound.
	Port int `json:"port"`
	// Connection id to be used.
	ConnectionId string `json:"connectionId"`
}

// Informs that port was successfully bound and got a specified connection id.
func (obj *Tethering) Accepted(fn func(params *AcceptedParams) bool) {
	params := AcceptedParams{}
	closeChn := make(chan struct{})
	go func() {
		for closeChn != nil {
			obj.conn.On("Tethering.accepted", closeChn, &params)
			if !fn(&params) {
				closeChn <- struct{}{}
				close(closeChn)
			}
		}
	}()
}
