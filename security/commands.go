package security

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Security struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Security {
	return &Security{conn}
}
func (obj *Security) Enable() (err error) {
	err = obj.conn.Send("Security.enable", nil, nil)
	return
}
func (obj *Security) Disable() (err error) {
	err = obj.conn.Send("Security.disable", nil, nil)
	return
}

type HandleCertificateErrorRequest struct {
	EventId	int					`json:"eventId"`// The ID of the event.
	Action	types.Security_CertificateErrorAction	`json:"action"`// The action to take on the certificate error.
}

func (obj *Security) HandleCertificateError(request *HandleCertificateErrorRequest) (err error) {
	err = obj.conn.Send("Security.handleCertificateError", request, nil)
	return
}

type SetOverrideCertificateErrorsRequest struct {
	Override bool `json:"override"`// If true, certificate errors will be overridden.
}

func (obj *Security) SetOverrideCertificateErrors(request *SetOverrideCertificateErrorsRequest) (err error) {
	err = obj.conn.Send("Security.setOverrideCertificateErrors", request, nil)
	return
}
