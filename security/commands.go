package security

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Security struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Security {
	return &Security{conn}
}

// Enables tracking security state changes.
func (obj *Security) Enable() (err error) {
	err = obj.conn.Send("Security.enable", nil, nil)
	return
}

// Disables tracking security state changes.
func (obj *Security) Disable() (err error) {
	err = obj.conn.Send("Security.disable", nil, nil)
	return
}

type HandleCertificateErrorRequest struct {
	// The ID of the event.
	EventId int `json:"eventId"`
	// The action to take on the certificate error.
	Action types.Security_CertificateErrorAction `json:"action"`
}

// Handles a certificate error that fired a certificateError event.
func (obj *Security) HandleCertificateError(request *HandleCertificateErrorRequest) (err error) {
	err = obj.conn.Send("Security.handleCertificateError", request, nil)
	return
}

type SetOverrideCertificateErrorsRequest struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
func (obj *Security) SetOverrideCertificateErrors(request *SetOverrideCertificateErrorsRequest) (err error) {
	err = obj.conn.Send("Security.setOverrideCertificateErrors", request, nil)
	return
}
