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
