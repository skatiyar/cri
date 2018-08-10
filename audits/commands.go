/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package audits provides commands and events for Audits domain.
package audits

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Audits domain
const (
	GetEncodedResponse = "Audits.getEncodedResponse"
)

// Audits domain allows investigation of page violations and possible improvements.
type Audits struct {
	conn cri.Connector
}

// New creates a Audits instance
func New(conn cri.Connector) *Audits {
	return &Audits{conn}
}

type GetEncodedResponseRequest struct {
	// Identifier of the network request to get content for.
	RequestId types.Network_RequestId `json:"requestId"`
	// The encoding to use.
	Encoding string `json:"encoding"`
	// The quality of the encoding (0-1). (defaults to 1)
	Quality *float32 `json:"quality,omitempty"`
	// Whether to only return the size information (defaults to false).
	SizeOnly *bool `json:"sizeOnly,omitempty"`
}

type GetEncodedResponseResponse struct {
	// The encoded body as a base64 string. Omitted if sizeOnly is true.
	Body *string `json:"body,omitempty"`
	// Size before re-encoding.
	OriginalSize int `json:"originalSize"`
	// Size after re-encoding.
	EncodedSize int `json:"encodedSize"`
}

// Returns the response body and size if it were re-encoded with the specified settings. Only applies to images.
func (obj *Audits) GetEncodedResponse(request *GetEncodedResponseRequest) (response GetEncodedResponseResponse, err error) {
	err = obj.conn.Send(GetEncodedResponse, request, &response)
	return
}
