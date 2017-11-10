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
	RequestId	types.Network_RequestId	`json:"requestId"`// Identifier of the network request to get content for.
	Encoding	string			`json:"encoding"`// The encoding to use.
	Quality		*float32		`json:"quality,omitempty"`// The quality of the encoding (0-1). (defaults to 1)
	SizeOnly	*bool			`json:"sizeOnly,omitempty"`// Whether to only return the size information (defaults to false).
}

func (obj *Audits) GetEncodedResponse(request *GetEncodedResponseRequest) (response struct {
	Body		*string	`json:"body,omitempty"`// The encoded body as a base64 string. Omitted if sizeOnly is true.
	OriginalSize	int	`json:"originalSize"`// Size before re-encoding.
	EncodedSize	int	`json:"encodedSize"`// Size after re-encoding.
}, err error) {
	err = obj.conn.Send("Audits.getEncodedResponse", request, &response)
	return
}
