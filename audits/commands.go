package audits

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Audits struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Audits {
	return &Audits{conn}
}

type GetEncodedResponseRequest struct {
	RequestId types.Network_RequestId `json:"requestId"`
	Encoding  string                  `json:"encoding"`
	Quality   *float32                `json:"quality,omitempty"`
	SizeOnly  *bool                   `json:"sizeOnly,omitempty"`
}

func (obj *Audits) GetEncodedResponse(request *GetEncodedResponseRequest) (response struct {
	Body         *string `json:"body,omitempty"`
	OriginalSize int     `json:"originalSize"`
	EncodedSize  int     `json:"encodedSize"`
}, err error) {
	err = obj.conn.Send("Audits.getEncodedResponse", request, &response)
	return
}
