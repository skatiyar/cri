package io

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type IO struct {
	conn cri.Connector
}

func New(conn cri.Connector) *IO {
	return &IO{conn}
}

type ReadRequest struct {
	Handle	types.IO_StreamHandle	`json:"handle"`// Handle of the stream to read.
	Offset	*int			`json:"offset,omitempty"`// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
	Size	*int			`json:"size,omitempty"`// Maximum number of bytes to read (left upon the agent discretion if not specified).
}

func (obj *IO) Read(request *ReadRequest) (response struct {
	Base64Encoded	*bool	`json:"base64Encoded,omitempty"`// Set if the data is base64-encoded
	Data		string	`json:"data"`// Data that were read.
	Eof		bool	`json:"eof"`// Set if the end-of-file condition occured while reading.
}, err error) {
	err = obj.conn.Send("IO.read", request, &response)
	return
}

type CloseRequest struct {
	Handle types.IO_StreamHandle `json:"handle"`// Handle of the stream to close.
}

func (obj *IO) Close(request *CloseRequest) (err error) {
	err = obj.conn.Send("IO.close", request, nil)
	return
}

type ResolveBlobRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`// Object id of a Blob object wrapper.
}

func (obj *IO) ResolveBlob(request *ResolveBlobRequest) (response struct {
	Uuid string `json:"uuid"`// UUID of the specified Blob.
}, err error) {
	err = obj.conn.Send("IO.resolveBlob", request, &response)
	return
}
