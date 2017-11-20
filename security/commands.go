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

// Security
type Security struct {
	conn cri.Connector
}

// New creates a Security instance
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
func (obj *Security) SecurityStateChanged(fn func(params *SecurityStateChangedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Security.securityStateChanged", closeChn)
	go func() {
		for {
			params := SecurityStateChangedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type CertificateErrorParams struct {
	// The ID of the event.
	EventId int `json:"eventId"`
	// The type of the error.
	ErrorType string `json:"errorType"`
	// The url that was requested.
	RequestURL string `json:"requestURL"`
}

// There is a certificate error. If overriding certificate errors is enabled, then it should be handled with the handleCertificateError command. Note: this event does not fire if the certificate error has been allowed internally.
func (obj *Security) CertificateError(fn func(params *CertificateErrorParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Security.certificateError", closeChn)
	go func() {
		for {
			params := CertificateErrorParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
