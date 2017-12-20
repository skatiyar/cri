package cri_test

import (
	"crypto/tls"
	"testing"

	"github.com/SKatiyar/cri"
	"github.com/SKatiyar/cri/browser"
	"github.com/SKatiyar/cri/emulation"
	"github.com/SKatiyar/cri/page"
	"github.com/stretchr/testify/assert"
)

func TestNewConnection(t *testing.T) {
	assert := assert.New(t)

	defaultConn, defaultConnErr := cri.NewConnection()
	assert.Nil(defaultConnErr)
	assert.NotNil(defaultConn, "connection object should be returned for default address")

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
	cri.NewConnection(cri.SetTargetID("<target-id>"), cri.SetTLSConfig(config))
}

func ExampleConnection_Send() {
	// create new connection
	conn, _ := cri.NewConnection()

	// send Page.enable command
	conn.Send(page.Enable, nil, nil)

	// send Emulation.setDeviceMetricsOverride command with parameters
	conn.Send(emulation.SetDeviceMetricsOverride, emulation.SetDeviceMetricsOverrideRequest{}, nil)

	// send Browser.getVersion command and decode response
	var data browser.GetVersionResponse
	conn.Send(browser.GetVersion, nil, &data)
}

func ExampleConnection_On() {
	// create new connection
	conn, _ := cri.NewConnection()

	// listen for page load
	var params page.LoadEventFiredParams
	conn.On(page.LoadEventFired, &params)
}
