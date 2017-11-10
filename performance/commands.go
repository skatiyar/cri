package performance

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Performance struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Performance {
	return &Performance{conn}
}
func (obj *Performance) Enable() (err error) {
	err = obj.conn.Send("Performance.enable", nil, nil)
	return
}
func (obj *Performance) Disable() (err error) {
	err = obj.conn.Send("Performance.disable", nil, nil)
	return
}
func (obj *Performance) GetMetrics() (response struct {
	Metrics []types.Performance_Metric `json:"metrics"`// Current values for run-time metrics.
}, err error) {
	err = obj.conn.Send("Performance.getMetrics", nil, &response)
	return
}
