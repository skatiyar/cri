package cri

// Connector interface defines methods required by command packages.
type Connector interface {
	// Send sends a command with suplied request params to underlying
	// connection. It decodes result to response if provided.
	// Implementation should wait for response or timeout.
	Send(command string, request, response interface{}) error

	// On listens for an event. It takes event name and channel to signal close
	// as arguments. It returns a decoder function, which blocks till event is received.
	// Params received are decoded in params argument, returning an error if it fails.
	On(event string, closeChn chan struct{}) func(params interface{}) error
}
