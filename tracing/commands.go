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
	Categories                   *string                    `json:"categories,omitempty"`
	Options                      *string                    `json:"options,omitempty"`
	BufferUsageReportingInterval *float32                   `json:"bufferUsageReportingInterval,omitempty"`
	TransferMode                 *string                    `json:"transferMode,omitempty"`
	TraceConfig                  *types.Tracing_TraceConfig `json:"traceConfig,omitempty"`
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
	Categories []string `json:"categories"`
}, err error) {
	err = obj.conn.Send("Tracing.getCategories", nil, &response)
	return
}
func (obj *Tracing) RequestMemoryDump() (response struct {
	DumpGuid string `json:"dumpGuid"`
	Success  bool   `json:"success"`
}, err error) {
	err = obj.conn.Send("Tracing.requestMemoryDump", nil, &response)
	return
}

type RecordClockSyncMarkerRequest struct {
	SyncId string `json:"syncId"`
}

func (obj *Tracing) RecordClockSyncMarker(request *RecordClockSyncMarkerRequest) (err error) {
	err = obj.conn.Send("Tracing.recordClockSyncMarker", request, nil)
	return
}
