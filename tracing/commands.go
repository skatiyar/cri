package tracing

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Tracing struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Tracing {
	return &Tracing{conn}
}

type StartRequest struct {
	Categories			*string				`json:"categories,omitempty"`// Category/tag filter
	Options				*string				`json:"options,omitempty"`// Tracing options
	BufferUsageReportingInterval	*float32			`json:"bufferUsageReportingInterval,omitempty"`// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	TransferMode			*string				`json:"transferMode,omitempty"`// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>).
	TraceConfig			*types.Tracing_TraceConfig	`json:"traceConfig,omitempty"`
}

func (obj *Tracing) Start(request *StartRequest) (err error) {
	err = obj.conn.Send("Tracing.start", request, nil)
	return
}
func (obj *Tracing) End() (err error) {
	err = obj.conn.Send("Tracing.end", nil, nil)
	return
}
func (obj *Tracing) GetCategories() (response struct {
	Categories []string `json:"categories"`// A list of supported tracing categories.
}, err error) {
	err = obj.conn.Send("Tracing.getCategories", nil, &response)
	return
}
func (obj *Tracing) RequestMemoryDump() (response struct {
	DumpGuid	string	`json:"dumpGuid"`// GUID of the resulting global memory dump.
	Success		bool	`json:"success"`// True iff the global memory dump succeeded.
}, err error) {
	err = obj.conn.Send("Tracing.requestMemoryDump", nil, &response)
	return
}

type RecordClockSyncMarkerRequest struct {
	SyncId string `json:"syncId"`// The ID of this clock sync marker
}

func (obj *Tracing) RecordClockSyncMarker(request *RecordClockSyncMarkerRequest) (err error) {
	err = obj.conn.Send("Tracing.recordClockSyncMarker", request, nil)
	return
}
