/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package schema provides commands and events for Schema domain.
// Provides information about the protocol schema.
package schema

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Schema struct {
	conn cri.Connector
}

// New creates a Schema instance
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