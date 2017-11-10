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
	Handle types.IO_StreamHandle `json:"handle"`
	Offset *int                  `json:"offset,omitempty"`
	Size   *int                  `json:"size,omitempty"`
}

func (obj *IO) Read(request *ReadRequest) (response struct {
	Base64Encoded *bool  `json:"base64Encoded,omitempty"`
	Data          string `json:"data"`
	Eof           bool   `json:"eof"`
}, err error) {
	err = obj.conn.Send("IO.read", request, &response)
	return
}

type CloseRequest struct {
	Handle types.IO_StreamHandle `json:"handle"`
}

func (obj *IO) Close(request *CloseRequest) (err error) {
	err = obj.conn.Send("IO.close", request, nil)
	return
}

type ResolveBlobRequest struct {
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

func (obj *IO) ResolveBlob(request *ResolveBlobRequest) (response struct {
	Uuid string `json:"uuid"`
}, err error) {
	err = obj.conn.Send("IO.resolveBlob", request, &response)
	return
}
