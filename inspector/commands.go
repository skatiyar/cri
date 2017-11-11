package inspector

import "github.com/SKatiyar/cri"

type Inspector struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Inspector {
	return &Inspector{conn}
}

// Enables inspector domain notifications.
func (obj *Inspector) Enable() (err error) {
	err = obj.conn.Send("Inspector.enable", nil, nil)
	return
}

// Disables inspector domain notifications.
func (obj *Inspector) Disable() (err error) {
	err = obj.conn.Send("Inspector.disable", nil, nil)
	return
}
