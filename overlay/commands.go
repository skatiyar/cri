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
	Result bool `json:"result"`// True for showing paint rectangles
}

func (obj *Overlay) SetShowPaintRects(request *SetShowPaintRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowPaintRects", request, nil)
	return
}

type SetShowDebugBordersRequest struct {
	Show bool `json:"show"`// True for showing debug borders
}

func (obj *Overlay) SetShowDebugBorders(request *SetShowDebugBordersRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowDebugBorders", request, nil)
	return
}

type SetShowFPSCounterRequest struct {
	Show bool `json:"show"`// True for showing the FPS counter
}

func (obj *Overlay) SetShowFPSCounter(request *SetShowFPSCounterRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowFPSCounter", request, nil)
	return
}

type SetShowScrollBottleneckRectsRequest struct {
	Show bool `json:"show"`// True for showing scroll bottleneck rects
}

func (obj *Overlay) SetShowScrollBottleneckRects(request *SetShowScrollBottleneckRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowScrollBottleneckRects", request, nil)
	return
}

type SetShowViewportSizeOnResizeRequest struct {
	Show bool `json:"show"`// Whether to paint size or not.
}

func (obj *Overlay) SetShowViewportSizeOnResize(request *SetShowViewportSizeOnResizeRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowViewportSizeOnResize", request, nil)
	return
}

type SetPausedInDebuggerMessageRequest struct {
	Message *string `json:"message,omitempty"`// The message to display, also triggers resume and step over controls.
}

func (obj *Overlay) SetPausedInDebuggerMessage(request *SetPausedInDebuggerMessageRequest) (err error) {
	err = obj.conn.Send("Overlay.setPausedInDebuggerMessage", request, nil)
	return
}

type SetSuspendedRequest struct {
	Suspended bool `json:"suspended"`// Whether overlay should be suspended and not consume any resources until resumed.
}

func (obj *Overlay) SetSuspended(request *SetSuspendedRequest) (err error) {
	err = obj.conn.Send("Overlay.setSuspended", request, nil)
	return
}

type SetInspectModeRequest struct {
	Mode		types.Overlay_InspectMode	`json:"mode"`// Set an inspection mode.
	HighlightConfig	*types.Overlay_HighlightConfig	`json:"highlightConfig,omitempty"`// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
}

func (obj *Overlay) SetInspectMode(request *SetInspectModeRequest) (err error) {
	err = obj.conn.Send("Overlay.setInspectMode", request, nil)
	return
}

type HighlightRectRequest struct {
	X		int		`json:"x"`// X coordinate
	Y		int		`json:"y"`// Y coordinate
	Width		int		`json:"width"`// Rectangle width
	Height		int		`json:"height"`// Rectangle height
	Color		*types.DOM_RGBA	`json:"color,omitempty"`// The highlight fill color (default: transparent).
	OutlineColor	*types.DOM_RGBA	`json:"outlineColor,omitempty"`// The highlight outline color (default: transparent).
}

func (obj *Overlay) HighlightRect(request *HighlightRectRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightRect", request, nil)
	return
}

type HighlightQuadRequest struct {
	Quad		types.DOM_Quad	`json:"quad"`// Quad to highlight
	Color		*types.DOM_RGBA	`json:"color,omitempty"`// The highlight fill color (default: transparent).
	OutlineColor	*types.DOM_RGBA	`json:"outlineColor,omitempty"`// The highlight outline color (default: transparent).
}

func (obj *Overlay) HighlightQuad(request *HighlightQuadRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightQuad", request, nil)
	return
}

type HighlightNodeRequest struct {
	HighlightConfig	types.Overlay_HighlightConfig	`json:"highlightConfig"`// A descriptor for the highlight appearance.
	NodeId		*types.DOM_NodeId		`json:"nodeId,omitempty"`// Identifier of the node to highlight.
	BackendNodeId	*types.DOM_BackendNodeId	`json:"backendNodeId,omitempty"`// Identifier of the backend node to highlight.
	ObjectId	*types.Runtime_RemoteObjectId	`json:"objectId,omitempty"`// JavaScript object id of the node to be highlighted.
}

func (obj *Overlay) HighlightNode(request *HighlightNodeRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightNode", request, nil)
	return
}

type HighlightFrameRequest struct {
	FrameId			types.Page_FrameId	`json:"frameId"`// Identifier of the frame to highlight.
	ContentColor		*types.DOM_RGBA		`json:"contentColor,omitempty"`// The content box highlight fill color (default: transparent).
	ContentOutlineColor	*types.DOM_RGBA		`json:"contentOutlineColor,omitempty"`// The content box highlight outline color (default: transparent).
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
	NodeId types.DOM_NodeId `json:"nodeId"`// Id of the node to get highlight object for.
}

func (obj *Overlay) GetHighlightObjectForTest(request *GetHighlightObjectForTestRequest) (response struct {
	Highlight map[string]interface{} `json:"highlight"`// Highlight data for the node.
}, err error) {
	err = obj.conn.Send("Overlay.getHighlightObjectForTest", request, &response)
	return
}
