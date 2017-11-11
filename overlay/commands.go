package overlay

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Overlay struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Overlay {
	return &Overlay{conn}
}

// Enables domain notifications.
func (obj *Overlay) Enable() (err error) {
	err = obj.conn.Send("Overlay.enable", nil, nil)
	return
}

// Disables domain notifications.
func (obj *Overlay) Disable() (err error) {
	err = obj.conn.Send("Overlay.disable", nil, nil)
	return
}

type SetShowPaintRectsRequest struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// Requests that backend shows paint rectangles
func (obj *Overlay) SetShowPaintRects(request *SetShowPaintRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowPaintRects", request, nil)
	return
}

type SetShowDebugBordersRequest struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// Requests that backend shows debug borders on layers
func (obj *Overlay) SetShowDebugBorders(request *SetShowDebugBordersRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowDebugBorders", request, nil)
	return
}

type SetShowFPSCounterRequest struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// Requests that backend shows the FPS counter
func (obj *Overlay) SetShowFPSCounter(request *SetShowFPSCounterRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowFPSCounter", request, nil)
	return
}

type SetShowScrollBottleneckRectsRequest struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// Requests that backend shows scroll bottleneck rects
func (obj *Overlay) SetShowScrollBottleneckRects(request *SetShowScrollBottleneckRectsRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowScrollBottleneckRects", request, nil)
	return
}

type SetShowViewportSizeOnResizeRequest struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// Paints viewport size upon main frame resize.
func (obj *Overlay) SetShowViewportSizeOnResize(request *SetShowViewportSizeOnResizeRequest) (err error) {
	err = obj.conn.Send("Overlay.setShowViewportSizeOnResize", request, nil)
	return
}

type SetPausedInDebuggerMessageRequest struct {
	// The message to display, also triggers resume and step over controls.
	Message *string `json:"message,omitempty"`
}

func (obj *Overlay) SetPausedInDebuggerMessage(request *SetPausedInDebuggerMessageRequest) (err error) {
	err = obj.conn.Send("Overlay.setPausedInDebuggerMessage", request, nil)
	return
}

type SetSuspendedRequest struct {
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

func (obj *Overlay) SetSuspended(request *SetSuspendedRequest) (err error) {
	err = obj.conn.Send("Overlay.setSuspended", request, nil)
	return
}

type SetInspectModeRequest struct {
	// Set an inspection mode.
	Mode types.Overlay_InspectMode `json:"mode"`
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
	HighlightConfig *types.Overlay_HighlightConfig `json:"highlightConfig,omitempty"`
}

// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
func (obj *Overlay) SetInspectMode(request *SetInspectModeRequest) (err error) {
	err = obj.conn.Send("Overlay.setInspectMode", request, nil)
	return
}

type HighlightRectRequest struct {
	// X coordinate
	X int `json:"x"`
	// Y coordinate
	Y int `json:"y"`
	// Rectangle width
	Width int `json:"width"`
	// Rectangle height
	Height int `json:"height"`
	// The highlight fill color (default: transparent).
	Color *types.DOM_RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (obj *Overlay) HighlightRect(request *HighlightRectRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightRect", request, nil)
	return
}

type HighlightQuadRequest struct {
	// Quad to highlight
	Quad types.DOM_Quad `json:"quad"`
	// The highlight fill color (default: transparent).
	Color *types.DOM_RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
func (obj *Overlay) HighlightQuad(request *HighlightQuadRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightQuad", request, nil)
	return
}

type HighlightNodeRequest struct {
	// A descriptor for the highlight appearance.
	HighlightConfig types.Overlay_HighlightConfig `json:"highlightConfig"`
	// Identifier of the node to highlight.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (obj *Overlay) HighlightNode(request *HighlightNodeRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightNode", request, nil)
	return
}

type HighlightFrameRequest struct {
	// Identifier of the frame to highlight.
	FrameId types.Page_FrameId `json:"frameId"`
	// The content box highlight fill color (default: transparent).
	ContentColor *types.DOM_RGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor *types.DOM_RGBA `json:"contentOutlineColor,omitempty"`
}

// Highlights owner element of the frame with given id.
func (obj *Overlay) HighlightFrame(request *HighlightFrameRequest) (err error) {
	err = obj.conn.Send("Overlay.highlightFrame", request, nil)
	return
}

// Hides any highlight.
func (obj *Overlay) HideHighlight() (err error) {
	err = obj.conn.Send("Overlay.hideHighlight", nil, nil)
	return
}

type GetHighlightObjectForTestRequest struct {
	// Id of the node to get highlight object for.
	NodeId types.DOM_NodeId `json:"nodeId"`
}
type GetHighlightObjectForTestResponse struct {
	// Highlight data for the node.
	Highlight map[string]interface{} `json:"highlight"`
}

// For testing.
func (obj *Overlay) GetHighlightObjectForTest(request *GetHighlightObjectForTestRequest) (response GetHighlightObjectForTestResponse, err error) {
	err = obj.conn.Send("Overlay.getHighlightObjectForTest", request, &response)
	return
}
