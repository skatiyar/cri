package browser

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Browser struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Browser {
	return &Browser{conn}
}
func (obj *Browser) Close() (err error) {
	err = obj.conn.Send("Browser.close", nil, nil)
	return
}

type GetWindowForTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`// Devtools agent host id.
}

func (obj *Browser) GetWindowForTarget(request *GetWindowForTargetRequest) (response struct {
	WindowId	types.Browser_WindowID	`json:"windowId"`// Browser window id.
	Bounds		types.Browser_Bounds	`json:"bounds"`// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}, err error) {
	err = obj.conn.Send("Browser.getWindowForTarget", request, &response)
	return
}
func (obj *Browser) GetVersion() (response struct {
	ProtocolVersion	string	`json:"protocolVersion"`// Protocol version.
	Product		string	`json:"product"`// Product name.
	Revision	string	`json:"revision"`// Product revision.
	UserAgent	string	`json:"userAgent"`// User-Agent.
	JsVersion	string	`json:"jsVersion"`// V8 version.
}, err error) {
	err = obj.conn.Send("Browser.getVersion", nil, &response)
	return
}

type SetWindowBoundsRequest struct {
	WindowId	types.Browser_WindowID	`json:"windowId"`// Browser window id.
	Bounds		types.Browser_Bounds	`json:"bounds"`// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
}

func (obj *Browser) SetWindowBounds(request *SetWindowBoundsRequest) (err error) {
	err = obj.conn.Send("Browser.setWindowBounds", request, nil)
	return
}

type GetWindowBoundsRequest struct {
	WindowId types.Browser_WindowID `json:"windowId"`// Browser window id.
}

func (obj *Browser) GetWindowBounds(request *GetWindowBoundsRequest) (response struct {
	Bounds types.Browser_Bounds `json:"bounds"`// Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
}, err error) {
	err = obj.conn.Send("Browser.getWindowBounds", request, &response)
	return
}
