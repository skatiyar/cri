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
	TargetId types.Target_TargetID `json:"targetId"`
}

func (obj *Browser) GetWindowForTarget(request *GetWindowForTargetRequest) (response struct {
	WindowId types.Browser_WindowID `json:"windowId"`
	Bounds   types.Browser_Bounds   `json:"bounds"`
}, err error) {
	err = obj.conn.Send("Browser.getWindowForTarget", request, &response)
	return
}
func (obj *Browser) GetVersion() (response struct {
	ProtocolVersion string `json:"protocolVersion"`
	Product         string `json:"product"`
	Revision        string `json:"revision"`
	UserAgent       string `json:"userAgent"`
	JsVersion       string `json:"jsVersion"`
}, err error) {
	err = obj.conn.Send("Browser.getVersion", nil, &response)
	return
}

type SetWindowBoundsRequest struct {
	WindowId types.Browser_WindowID `json:"windowId"`
	Bounds   types.Browser_Bounds   `json:"bounds"`
}

func (obj *Browser) SetWindowBounds(request *SetWindowBoundsRequest) (err error) {
	err = obj.conn.Send("Browser.setWindowBounds", request, nil)
	return
}

type GetWindowBoundsRequest struct {
	WindowId types.Browser_WindowID `json:"windowId"`
}

func (obj *Browser) GetWindowBounds(request *GetWindowBoundsRequest) (response struct {
	Bounds types.Browser_Bounds `json:"bounds"`
}, err error) {
	err = obj.conn.Send("Browser.getWindowBounds", request, &response)
	return
}
