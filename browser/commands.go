/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package browser provides commands and events for Browser domain.
package browser

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Browser domain
const (
	Close                 = "Browser.close"
	GetVersion            = "Browser.getVersion"
	GetBrowserCommandLine = "Browser.getBrowserCommandLine"
	GetHistograms         = "Browser.getHistograms"
	GetHistogram          = "Browser.getHistogram"
	GetWindowBounds       = "Browser.getWindowBounds"
	GetWindowForTarget    = "Browser.getWindowForTarget"
	SetWindowBounds       = "Browser.setWindowBounds"
)

// The Browser domain defines methods and events for browser managing.
type Browser struct {
	conn cri.Connector
}

// New creates a Browser instance
func New(conn cri.Connector) *Browser {
	return &Browser{conn}
}

// Close browser gracefully.
func (obj *Browser) Close() (err error) {
	err = obj.conn.Send(Close, nil, nil)
	return
}

type GetVersionResponse struct {
	// Protocol version.
	ProtocolVersion string `json:"protocolVersion"`
	// Product name.
	Product string `json:"product"`
	// Product revision.
	Revision string `json:"revision"`
	// User-Agent.
	UserAgent string `json:"userAgent"`
	// V8 version.
	JsVersion string `json:"jsVersion"`
}

// Returns version information.
func (obj *Browser) GetVersion() (response GetVersionResponse, err error) {
	err = obj.conn.Send(GetVersion, nil, &response)
	return
}

type GetBrowserCommandLineResponse struct {
	// Commandline parameters
	Arguments []string `json:"arguments"`
}

// Returns the command line switches for the browser process if, and only if --enable-automation is on the commandline.
func (obj *Browser) GetBrowserCommandLine() (response GetBrowserCommandLineResponse, err error) {
	err = obj.conn.Send(GetBrowserCommandLine, nil, &response)
	return
}

type GetHistogramsRequest struct {
	// Requested substring in name. Only histograms which have query as a substring in their name are extracted. An empty or absent query returns all histograms.
	Query *string `json:"query,omitempty"`
	// If true, retrieve delta since last call.
	Delta *bool `json:"delta,omitempty"`
}

type GetHistogramsResponse struct {
	// Histograms.
	Histograms []types.Browser_Histogram `json:"histograms"`
}

// Get Chrome histograms.
func (obj *Browser) GetHistograms(request *GetHistogramsRequest) (response GetHistogramsResponse, err error) {
	err = obj.conn.Send(GetHistograms, request, &response)
	return
}

type GetHistogramRequest struct {
	// Requested histogram name.
	Name string `json:"name"`
	// If true, retrieve delta since last call.
	Delta *bool `json:"delta,omitempty"`
}

type GetHistogramResponse struct {
	// Histogram.
	Histogram types.Browser_Histogram `json:"histogram"`
}

// Get a Chrome histogram by name.
func (obj *Browser) GetHistogram(request *GetHistogramRequest) (response GetHistogramResponse, err error) {
	err = obj.conn.Send(GetHistogram, request, &response)
	return
}

type GetWindowBoundsRequest struct {
	// Browser window id.
	WindowId types.Browser_WindowID `json:"windowId"`
}

type GetWindowBoundsResponse struct {
	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds types.Browser_Bounds `json:"bounds"`
}

// Get position and size of the browser window.
func (obj *Browser) GetWindowBounds(request *GetWindowBoundsRequest) (response GetWindowBoundsResponse, err error) {
	err = obj.conn.Send(GetWindowBounds, request, &response)
	return
}

type GetWindowForTargetRequest struct {
	// Devtools agent host id.
	TargetId types.Target_TargetID `json:"targetId"`
}

type GetWindowForTargetResponse struct {
	// Browser window id.
	WindowId types.Browser_WindowID `json:"windowId"`
	// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
	Bounds types.Browser_Bounds `json:"bounds"`
}

// Get the browser window that contains the devtools target.
func (obj *Browser) GetWindowForTarget(request *GetWindowForTargetRequest) (response GetWindowForTargetResponse, err error) {
	err = obj.conn.Send(GetWindowForTarget, request, &response)
	return
}

type SetWindowBoundsRequest struct {
	// Browser window id.
	WindowId types.Browser_WindowID `json:"windowId"`
	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds types.Browser_Bounds `json:"bounds"`
}

// Set position and/or size of the browser window.
func (obj *Browser) SetWindowBounds(request *SetWindowBoundsRequest) (err error) {
	err = obj.conn.Send(SetWindowBounds, request, nil)
	return
}
