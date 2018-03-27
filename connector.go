package cri

type Decoder func(params interface{}) error
type Closer func() error

// Connector interface defines methods required by command packages.
type Connector interface {
	// Send sends a command with suplied request params to underlying
	// connection. It decodes result to response if provided.
	// Implementation should wait for response or timeout.
	Send(command string, request, response interface{}) error

	// On listens for an event. It takes event name as an argument.
	// It waits for event and decodes received event in params argument.
	// Returns an error if it fails.
	On(event string) (Decoder, Closer)
}
