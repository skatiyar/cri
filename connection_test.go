package cri_test

import (
	"crypto/tls"
	"testing"

	"github.com/SKatiyar/cri"
	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	assert := assert.New(t)

	defaultConn, defaultConnErr := cri.NewConnection()
	assert.Nil(defaultConnErr)
	assert.NotNil(defaultConn, "connection object should returned for default address")

	defaultCloseErr := defaultConn.Close()
	assert.Nil(defaultCloseErr)
}

func ExampleNewConnection() {
	// create connection with default options
	cri.NewConnection()
	// override the remote address
	cri.NewConnection(cri.SetAddress("192.168.0.7:1290"))
	// select a target with ssl connection
	config := &tls.Config{ /* ...params */ }
	cri.NewConnection(cri.SetAddress("192.168.0.7:1290"), cri.SetTargetID("<target-id>"), cri.SetTLSConfig(config))
}
