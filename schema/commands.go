package schema

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Schema struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Schema {
	return &Schema{conn}
}
func (obj *Schema) GetDomains() (response struct {
	Domains []types.Schema_Domain `json:"domains"`
}, err error) {
	err = obj.conn.Send("Schema.getDomains", nil, &response)
	return
}
