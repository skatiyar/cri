package schema

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Schema struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Schema {
	return &Schema{conn}
}

type GetDomainsResponse struct {
	// List of supported domains.
	Domains []types.Schema_Domain `json:"domains"`
}

// Returns supported domains.
func (obj *Schema) GetDomains() (response GetDomainsResponse, err error) {
	err = obj.conn.Send("Schema.getDomains", nil, &response)
	return
}
