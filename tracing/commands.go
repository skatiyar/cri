/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/


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
	TransferMode *string `json:"transferMode,omitempty"`
	TraceConfig *types.Tracing_TraceConfig `json:"traceConfig,omitempty"`
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
