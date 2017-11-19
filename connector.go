package cri

// Connector defines the methods required by the protocols to work.
type Connector interface {
	Send(command string, request, response interface{}) error
	On(event string, closeChn chan struct{}) func(params interface{}) error
}
