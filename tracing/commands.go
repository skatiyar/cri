/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package tracing provides commands and events for Tracing domain.
package tracing

import (
	"github.com/skatiyar/cri"
	types "github.com/skatiyar/cri/types"
)

// List of commands in Tracing domain
const (
	End                   = "Tracing.end"
	GetCategories         = "Tracing.getCategories"
	RecordClockSyncMarker = "Tracing.recordClockSyncMarker"
	RequestMemoryDump     = "Tracing.requestMemoryDump"
	Start                 = "Tracing.start"
)

// List of events in Tracing domain
const (
	BufferUsage     = "Tracing.bufferUsage"
	DataCollected   = "Tracing.dataCollected"
	TracingComplete = "Tracing.tracingComplete"
)

type Tracing struct {
	conn cri.Connector
}

// New creates a Tracing instance
func New(conn cri.Connector) *Tracing {
	return &Tracing{conn}
}

// Stop trace events collection.
func (obj *Tracing) End() (err error) {
	err = obj.conn.Send(End, nil, nil)
	return
}

type GetCategoriesResponse struct {
	// A list of supported tracing categories.
	Categories []string `json:"categories"`
}

// Gets supported tracing categories.
func (obj *Tracing) GetCategories() (response GetCategoriesResponse, err error) {
	err = obj.conn.Send(GetCategories, nil, &response)
	return
}

type RecordClockSyncMarkerRequest struct {
	// The ID of this clock sync marker
	SyncId string `json:"syncId"`
}

// Record a clock sync marker in the trace.
func (obj *Tracing) RecordClockSyncMarker(request *RecordClockSyncMarkerRequest) (err error) {
	err = obj.conn.Send(RecordClockSyncMarker, request, nil)
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
	err = obj.conn.Send(RequestMemoryDump, nil, &response)
	return
}

type StartRequest struct {
	// Category/tag filter
	Categories *string `json:"categories,omitempty"`
	// Tracing options
	Options *string `json:"options,omitempty"`
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval *float32 `json:"bufferUsageReportingInterval,omitempty"`
	// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to `ReportEvents`).
	TransferMode *string `json:"transferMode,omitempty"`
	// Compression format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `none`)
	StreamCompression *types.Tracing_StreamCompression `json:"streamCompression,omitempty"`
	TraceConfig       *types.Tracing_TraceConfig       `json:"traceConfig,omitempty"`
}

// Start trace events collection.
func (obj *Tracing) Start(request *StartRequest) (err error) {
	err = obj.conn.Send(Start, request, nil)
	return
}

type BufferUsageParams struct {
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	PercentFull *float32 `json:"percentFull,omitempty"`
	// An approximate number of events in the trace log.
	EventCount *float32 `json:"eventCount,omitempty"`
	// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	Value *float32 `json:"value,omitempty"`
}

func (obj *Tracing) BufferUsage(fn func(event string, params BufferUsageParams, err error) bool) {
	listen, closer := obj.conn.On(BufferUsage)
	go func() {
		defer closer()
		for {
			var params BufferUsageParams
			if !fn(BufferUsage, params, listen(&params)) {
				return
			}
		}
	}()
}

type DataCollectedParams struct {
	Value []map[string]interface{} `json:"value"`
}

// Contains an bucket of collected trace events. When tracing is stopped collected events will be send as a sequence of dataCollected events followed by tracingComplete event.
func (obj *Tracing) DataCollected(fn func(event string, params DataCollectedParams, err error) bool) {
	listen, closer := obj.conn.On(DataCollected)
	go func() {
		defer closer()
		for {
			var params DataCollectedParams
			if !fn(DataCollected, params, listen(&params)) {
				return
			}
		}
	}()
}

type TracingCompleteParams struct {
	// A handle of the stream that holds resulting trace data.
	Stream *types.IO_StreamHandle `json:"stream,omitempty"`
	// Compression format of returned stream.
	StreamCompression *types.Tracing_StreamCompression `json:"streamCompression,omitempty"`
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
func (obj *Tracing) TracingComplete(fn func(event string, params TracingCompleteParams, err error) bool) {
	listen, closer := obj.conn.On(TracingComplete)
	go func() {
		defer closer()
		for {
			var params TracingCompleteParams
			if !fn(TracingComplete, params, listen(&params)) {
				return
			}
		}
	}()
}
