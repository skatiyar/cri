package cri

type Connector interface {
	Send(command string, request, response interface{}) error
	On(event string, closeChn chan struct{}) func(params interface{}) error
}