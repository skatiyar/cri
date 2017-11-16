/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package log provides commands and events for Log domain.
package log

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// Provides access to log entries.
type Log struct {
	conn cri.Connector
}

// New creates a Log instance
func New(conn cri.Connector) *Log {
	return &Log{conn}
}

// Enables log domain, sends the entries collected so far to the client by means of the <code>entryAdded</code> notification.
func (obj *Log) Enable() (err error) {
	err = obj.conn.Send("Log.enable", nil, nil)
	return
}

// Disables log domain, prevents further log entries from being reported to the client.
func (obj *Log) Disable() (err error) {
	err = obj.conn.Send("Log.disable", nil, nil)
	return
}

// Clears the log.
func (obj *Log) Clear() (err error) {
	err = obj.conn.Send("Log.clear", nil, nil)
	return
}

type StartViolationsReportRequest struct {
	// Configuration for violations.
	Config []types.Log_ViolationSetting `json:"config"`
}

// start violation reporting.
func (obj *Log) StartViolationsReport(request *StartViolationsReportRequest) (err error) {
	err = obj.conn.Send("Log.startViolationsReport", request, nil)
	return
}

// Stop violation reporting.
func (obj *Log) StopViolationsReport() (err error) {
	err = obj.conn.Send("Log.stopViolationsReport", nil, nil)
	return
}

type EntryAddedParams struct {
	// The entry.
	Entry types.Log_LogEntry `json:"entry"`
}

// Issued when new message was logged.
func (obj *Log) EntryAdded(fn func(params *EntryAddedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Log.entryAdded", closeChn)
	go func() {
		for {
			params := EntryAddedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
