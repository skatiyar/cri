package tethering

import "github.com/SKatiyar/cri"

type Tethering struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Tethering {
	return &Tethering{conn}
}

type BindRequest struct {
	Port int `json:"port"`
}

func (obj *Tethering) Bind(request *BindRequest) (err error) {
	err = obj.conn.Send("Tethering.bind", request, nil)
	return
}

type UnbindRequest struct {
	Port int `json:"port"`
}

func (obj *Tethering) Unbind(request *UnbindRequest) (err error) {
	err = obj.conn.Send("Tethering.unbind", request, nil)
	return
}
