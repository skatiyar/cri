/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package console provides commands and events for Console domain.
// This domain is deprecated - use Runtime or Log instead.
package console

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Console struct {
	conn cri.Connector
}

// New creates a Console instance
func New(conn cri.Connector) *Console {
	return &Console{conn}
}

// Enables console domain, sends the messages collected so far to the client by means of the <code>messageAdded</code> notification.
func (obj *Console) Enable() (err error) {
	err = obj.conn.Send("Console.enable", nil, nil)
	return
}

// Disables console domain, prevents further console messages from being reported to the client.
func (obj *Console) Disable() (err error) {
	err = obj.conn.Send("Console.disable", nil, nil)
	return
}

// Does nothing.
func (obj *Console) ClearMessages() (err error) {
	err = obj.conn.Send("Console.clearMessages", nil, nil)
	return
}

type MessageAddedParams struct {
	// Console message that has been added.
	Message types.Console_ConsoleMessage `json:"message"`
}

// Issued when new console message is added.
func (obj *Console) MessageAdded(fn func(params *MessageAddedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Console.messageAdded", closeChn)
	go func() {
		for {
			params := MessageAddedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
