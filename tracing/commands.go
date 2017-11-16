/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package tracing provides commands and events for Tracing domain.

package tracing

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Tracing struct {
	conn cri.Connector
}

// New creates a Tracing instance
func New(conn cri.Connector) *Tracing {
	return &Tracing{conn}
}

type StartRequest struct {
	// Category/tag filter
	Categories *string `json:"categories,omitempty"`
	// Tracing options
	Options *string `json:"options,omitempty"`
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval *float32 `json:"bufferUsageReportingInterval,omitempty"`
	// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>).
	TransferMode *string                    `json:"transferMode,omitempty"`
	TraceConfig  *types.Tracing_TraceConfig `json:"traceConfig,omitempty"`
}

// Start trace events collection.
func (obj *Tracing) Start(request *StartRequest) (err error) {
	err = obj.conn.Send("Tracing.start", request, nil)
	return
}

// Stop trace events collection.
func (obj *Tracing) End() (err error) {
	err = obj.conn.Send("Tracing.end", nil, nil)
	return
}

type GetCategoriesResponse struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

// Gets supported tracing categories.
func (obj *Tracing) GetCategories() (response GetCategoriesResponse, err error) {
	err = obj.conn.Send("Tracing.getCategories", nil, &response)
	return
}

type RequestMemoryDumpResponse struct {
	// GUID of the resulting global memory dump.
	DumpGuid string `json:"dumpGuid"`
	// True iff the global memory dump succeeded.
	Success bool `json:"success"`
}

// Request a global memory dump.
func (obj *Tracing) RequestMemoryDump() (response RequestMemoryDumpResponse, err error) {
	err = obj.conn.Send("Tracing.requestMemoryDump", nil, &response)
	return
}

type RecordClockSyncMarkerRequest struct {
	// The ID of this clock sync marker
	SyncId string `json:"syncId"`
}

// Record a clock sync marker in the trace.
func (obj *Tracing) RecordClockSyncMarker(request *RecordClockSyncMarkerRequest) (err error) {
	err = obj.conn.Send("Tracing.recordClockSyncMarker", request, nil)
	return
}

type DataCollectedParams struct {
	Value []map[string]interface{} `json:"value"`
}

// Contains an bucket of collected trace events. When tracing is stopped collected events will be send as a sequence of dataCollected events followed by tracingComplete event.
func (obj *Tracing) DataCollected(fn func(params *DataCollectedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Tracing.dataCollected", closeChn)
	go func() {
		for {
			params := DataCollectedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

type TracingCompleteParams struct {
	// A handle of the stream that holds resulting trace data.
	Stream *types.IO_StreamHandle `json:"stream,omitempty"`
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
func (obj *Tracing) TracingComplete(fn func(params *TracingCompleteParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Tracing.tracingComplete", closeChn)
	go func() {
		for {
			params := TracingCompleteParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}

type BufferUsageParams struct {
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	PercentFull *float32 `json:"percentFull,omitempty"`
	// An approximate number of events in the trace log.
	EventCount *float32 `json:"eventCount,omitempty"`
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	Value *float32 `json:"value,omitempty"`
}

func (obj *Tracing) BufferUsage(fn func(params *BufferUsageParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Tracing.bufferUsage", closeChn)
	go func() {
		for {
			params := BufferUsageParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
