package overlay

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Overlay struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Overlay {
	return &Overlay{conn}
}
func (obj *Overlay) Enable() (err error) {
	err = obj.conn.Send("Overlay.enable", nil, nil)
	return
}
func (obj *Overlay) Disable() (err error) {
	err = obj.conn.Send("Overlay.disable", nil, nil)
	return
}

type SetShowPaintRectsRequest struct {
	Result bool `json:"result"`
}

func (obj *Overlay) SetShowPaintRects(request *SetShowPaintRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowPaintRects", request, nil)
	return
}

type SetShowDebugBordersRequest struct {
	Show bool `json:"show"`
}

func (obj *Overlay) SetShowDebugBorders(request *SetShowDebugBordersRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowDebugBorders", request, nil)
	return
}

type SetShowFPSCounterRequest struct {
	Show bool `json:"show"`
}

func (obj *Overlay) SetShowFPSCounter(request *SetShowFPSCounterRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowFPSCounter", request, nil)
	return
}

type SetShowScrollBottleneckRectsRequest struct {
	Show bool `json:"show"`
}

func (obj *Overlay) SetShowScrollBottleneckRects(request *SetShowScrollBottleneckRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowScrollBottleneckRects", request, nil)
	return
}

type SetShowViewportSizeOnResizeRequest struct {
	Show bool `json:"show"`
}

func (obj *Overlay) SetShowViewportSizeOnResize(request *SetShowViewportSizeOnResizeRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowViewportSizeOnResize", request, nil)
	return
}

type SetPausedInDebuggerMessageRequest struct {
	Message *string `json:"message,omitempty"`
}

func (obj *Overlay) SetPausedInDebuggerMessage(request *SetPausedInDebuggerMessageRequest) (err error) {
	err = obj.conn.Send("Overlay.setPausedInDebuggerMessage", request, nil)
	return
}

type SetSuspendedRequest struct {
	Suspended bool `json:"suspended"`
}

func (obj *Overlay) SetSuspended(request *SetSuspendedRequest) (err error) {
	err = obj.conn.Send("Overlay.setSuspended", request, nil)
	return
}

type SetInspectModeRequest struct {
	Mode            types.Overlay_InspectMode      `json:"mode"`
	HighlightConfig *types.Overlay_HighlightConfig `json:"highlightConfig,omitempty"`
}

func (obj *Overlay) SetInspectMode(request *SetInspectModeRequest) (err error) {
	err = obj.conn.Send("Overlay.setInspectMode", request, nil)
	return
}

type HighlightRectRequest struct {
	X            int             `json:"x"`
	Y            int             `json:"y"`
	Width        int             `json:"width"`
	Height       int             `json:"height"`
	Color        *types.DOM_RGBA `json:"color,omitempty"`
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

func (obj *Overlay) HighlightRect(request *HighlightRectRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightRect", request, nil)
	return
}

type HighlightQuadRequest struct {
	Quad         types.DOM_Quad  `json:"quad"`
	Color        *types.DOM_RGBA `json:"color,omitempty"`
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

func (obj *Overlay) HighlightQuad(request *HighlightQuadRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightQuad", request, nil)
	return
}

type HighlightNodeRequest struct {
	HighlightConfig types.Overlay_HighlightConfig `json:"highlightConfig"`
	NodeId          *types.DOM_NodeId             `json:"nodeId,omitempty"`
	BackendNodeId   *types.DOM_BackendNodeId      `json:"backendNodeId,omitempty"`
	ObjectId        *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

func (obj *Overlay) HighlightNode(request *HighlightNodeRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightNode", request, nil)
	return
}

type HighlightFrameRequest struct {
	FrameId             types.Page_FrameId `json:"frameId"`
	ContentColor        *types.DOM_RGBA    `json:"contentColor,omitempty"`
	ContentOutlineColor *types.DOM_RGBA    `json:"contentOutlineColor,omitempty"`
}

func (obj *Overlay) HighlightFrame(request *HighlightFrameRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightFrame", request, nil)
	return
}
func (obj *Overlay) HideHighlight() (err error) {
	err = obj.conn.Send("Overlay.hideHighlight", nil, nil)
	return
}

type GetHighlightObjectForTestRequest struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

func (obj *Overlay) GetHighlightObjectForTest(request *GetHighlightObjectForTestRequest) (response struct {
	Highlight map[string]interface{} `json:"highlight"`
}, err error) {
	err = obj.conn.Send("Overlay.getHighlightObjectForTest", request, &response)
	return
}
