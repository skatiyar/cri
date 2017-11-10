package console

import "github.com/SKatiyar/cri"

type Console struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Console {
	return &Console{conn}
}
func (obj *Console) Enable() (err error) {
	err = obj.conn.Send("Console.enable", nil, nil)
	return
}
func (obj *Console) Disable() (err error) {
	err = obj.conn.Send("Console.disable", nil, nil)
	return
}
func (obj *Console) ClearMessages() (err error) {
	err = obj.conn.Send("Console.clearMessages", nil, nil)
	return
}
