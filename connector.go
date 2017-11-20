package cri

// Connector interface defines the methods required by the command packages.
//
// Send sends a command with the suplied request params.
// It writes the json result to the response if provided.
// It returns an error.
//
// On
type Connector interface {
	Send(command string, request, response interface{}) error
	On(event string, closeChn chan struct{}) func(params interface{}) error
}
