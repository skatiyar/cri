package cri

type Connector interface {
	Send(command string, request, response interface{}) error
}
