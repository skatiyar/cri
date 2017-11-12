/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Input/Output operations for streams produced by DevTools.
package io

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type IO struct {
	conn cri.Connector
}

// New creates a IO instance
func New(conn cri.Connector) *IO {
	return &IO{conn}
}

type ReadRequest struct {
	// Handle of the stream to read.
	Handle types.IO_StreamHandle `json:"handle"`
	// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
	Offset *int `json:"offset,omitempty"`
	// Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size *int `json:"size,omitempty"`
}

type ReadResponse struct {
	// Set if the data is base64-encoded
	Base64Encoded *bool `json:"base64Encoded,omitempty"`
	// Data that were read.
	Data string `json:"data"`
	// Set if the end-of-file condition occured while reading.
	Eof bool `json:"eof"`
}

// Read a chunk of the stream
func (obj *IO) Read(request *ReadRequest) (response ReadResponse, err error) {
	err = obj.conn.Send("IO.read", request, &response)
	return
}

type CloseRequest struct {
	// Handle of the stream to close.
	Handle types.IO_StreamHandle `json:"handle"`
}

// Close the stream, discard any temporary backing storage.
func (obj *IO) Close(request *CloseRequest) (err error) {
	err = obj.conn.Send("IO.close", request, nil)
	return
}

type ResolveBlobRequest struct {
	// Object id of a Blob object wrapper.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

type ResolveBlobResponse struct {
	// UUID of the specified Blob.
	Uuid string `json:"uuid"`
}

// Return UUID of Blob object specified by a remote object id.
func (obj *IO) ResolveBlob(request *ResolveBlobRequest) (response ResolveBlobResponse, err error) {
	err = obj.conn.Send("IO.resolveBlob", request, &response)
	return
}
