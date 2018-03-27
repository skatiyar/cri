/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package security provides commands and events for Security domain.
package security

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Security domain
const (
	Disable                      = "Security.disable"
	Enable                       = "Security.enable"
	SetIgnoreCertificateErrors   = "Security.setIgnoreCertificateErrors"
	HandleCertificateError       = "Security.handleCertificateError"
	SetOverrideCertificateErrors = "Security.setOverrideCertificateErrors"
)

// List of events in Security domain
const (
	CertificateError     = "Security.certificateError"
	SecurityStateChanged = "Security.securityStateChanged"
)

// Security
type Security struct {
	conn cri.Connector
}

// New creates a Security instance
func New(conn cri.Connector) *Security {
	return &Security{conn}
}

// Disables tracking security state changes.
func (obj *Security) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables tracking security state changes.
func (obj *Security) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type SetIgnoreCertificateErrorsRequest struct {
	// If true, all certificate errors will be ignored.
	Ignore bool `json:"ignore"`
}

// Enable/disable whether all certificate errors should be ignored.
func (obj *Security) SetIgnoreCertificateErrors(request *SetIgnoreCertificateErrorsRequest) (err error) {
	err = obj.conn.Send(SetIgnoreCertificateErrors, request, nil)
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
	err = obj.conn.Send(HandleCertificateError, request, nil)
	return
}

type SetOverrideCertificateErrorsRequest struct {
	// If true, certificate errors will be overridden.
	Override bool `json:"override"`
}

// Enable/disable overriding certificate errors. If enabled, all certificate error events need to be handled by the DevTools client and should be answered with handleCertificateError commands.
func (obj *Security) SetOverrideCertificateErrors(request *SetOverrideCertificateErrorsRequest) (err error) {
	err = obj.conn.Send(SetOverrideCertificateErrors, request, nil)
	return
}

type CertificateErrorParams struct {
	// The ID of the event.
	EventId int `json:"eventId"`
	// The type of the error.
	ErrorType string `json:"errorType"`
	// The url that was requested.
	RequestURL string `json:"requestURL"`
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the handleCertificateError command. Note: this event does not fire if the certificate error has been allowed internally. Only one client per target should override certificate errors at the same time.
func (obj *Security) CertificateError(fn func(event string, params CertificateErrorParams, err error) bool) {
	listen, closer := obj.conn.On(CertificateError)
	go func() {
		defer closer()
		for {
			var params CertificateErrorParams
			if !fn(CertificateError, params, listen(&params)) {
				return
			}
		}
	}()
}

type SecurityStateChangedParams struct {
	// Security state.
	SecurityState types.Security_SecurityState `json:"securityState"`
	// True if the page was loaded over cryptographic transport such as HTTPS.
	SchemeIsCryptographic bool `json:"schemeIsCryptographic"`
	// List of explanations for the security state. If the overall security state is `insecure` or `warning`, at least one corresponding explanation should be included.
	Explanations []types.Security_SecurityStateExplanation `json:"explanations"`
	// Information about insecure content on the page.
	InsecureContentStatus types.Security_InsecureContentStatus `json:"insecureContentStatus"`
	// Overrides user-visible description of the state.
	Summary *string `json:"summary,omitempty"`
}

// The security state of the page changed.
func (obj *Security) SecurityStateChanged(fn func(event string, params SecurityStateChangedParams, err error) bool) {
	listen, closer := obj.conn.On(SecurityStateChanged)
	go func() {
		defer closer()
		for {
			var params SecurityStateChangedParams
			if !fn(SecurityStateChanged, params, listen(&params)) {
				return
			}
		}
	}()
}
