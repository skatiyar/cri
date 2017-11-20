/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package tethering provides commands and events for Tethering domain.
package tethering

import (
	"github.com/SKatiyar/cri"
)

// List of commands in Tethering domain
const (
	Bind   = "Tethering.bind"
	Unbind = "Tethering.unbind"
)

// List of events in Tethering domain
const (
	Accepted = "Tethering.accepted"
)

// The Tethering domain defines methods and events for browser port binding.
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
	err = obj.conn.Send(Bind, request, nil)
	return
}

type UnbindRequest struct {
	// Port number to unbind.
	Port int `json:"port"`
}

// Request browser port unbinding.
func (obj *Tethering) Unbind(request *UnbindRequest) (err error) {
	err = obj.conn.Send(Unbind, request, nil)
	return
}

type AcceptedParams struct {
	// Port number that was successfully bound.
	Port int `json:"port"`
	// Connection id to be used.
	ConnectionId string `json:"connectionId"`
}

// Informs that port was successfully bound and got a specified connection id.
func (obj *Tethering) Accepted(fn func(params *AcceptedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(Accepted, closeChn)
	go func() {
		for {
			params := AcceptedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
