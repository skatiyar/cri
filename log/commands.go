package log

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Log struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Log {
	return &Log{conn}
}
func (obj *Log) Enable() (err error) {
	err = obj.conn.Send("Log.enable", nil, nil)
	return
}
func (obj *Log) Disable() (err error) {
	err = obj.conn.Send("Log.disable", nil, nil)
	return
}
func (obj *Log) Clear() (err error) {
	err = obj.conn.Send("Log.clear", nil, nil)
	return
}

type StartViolationsReportRequest struct {
	Config []types.Log_ViolationSetting `json:"config"`// Configuration for violations.
}

func (obj *Log) StartViolationsReport(request *StartViolationsReportRequest) (err error) {
	err = obj.conn.Send("Log.startViolationsReport", request, nil)
	return
}
func (obj *Log) StopViolationsReport() (err error) {
	err = obj.conn.Send("Log.stopViolationsReport", nil, nil)
	return
}
