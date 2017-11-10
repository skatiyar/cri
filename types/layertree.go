package types

type LayerTree_LayerId string
type LayerTree_SnapshotId string
type LayerTree_ScrollRect struct {
	Rect	DOM_Rect	`json:"rect"`// Rectangle itself.
	Type	string		`json:"type"`// Reason for rectangle to force scrolling on the main thread
}
type LayerTree_StickyPositionConstraint struct {
	StickyBoxRect				DOM_Rect		`json:"stickyBoxRect"`// Layout rectangle of the sticky element before being shifted
	ContainingBlockRect			DOM_Rect		`json:"containingBlockRect"`// Layout rectangle of the containing block of the sticky element
	NearestLayerShiftingStickyBox		*LayerTree_LayerId	`json:"nearestLayerShiftingStickyBox,omitempty"`// The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingContainingBlock	*LayerTree_LayerId	`json:"nearestLayerShiftingContainingBlock,omitempty"`// The nearest sticky layer that shifts the containing block
}
type LayerTree_PictureTile struct {
	X	float32	`json:"x"`// Offset from owning layer left boundary
	Y	float32	`json:"y"`// Offset from owning layer top boundary
	Picture	string	`json:"picture"`// Base64-encoded snapshot data.
}
type LayerTree_Layer struct {
	LayerId				LayerTree_LayerId			`json:"layerId"`// The unique id for this layer.
	ParentLayerId			*LayerTree_LayerId			`json:"parentLayerId,omitempty"`// The id of parent (not present for root).
	BackendNodeId			*DOM_BackendNodeId			`json:"backendNodeId,omitempty"`// The backend id for the node associated with this layer.
	OffsetX				float32					`json:"offsetX"`// Offset from parent layer, X coordinate.
	OffsetY				float32					`json:"offsetY"`// Offset from parent layer, Y coordinate.
	Width				float32					`json:"width"`// Layer width.
	Height				float32					`json:"height"`// Layer height.
	Transform			[]float32				`json:"transform,omitempty"`// Transformation matrix for layer, default is identity matrix
	AnchorX				*float32				`json:"anchorX,omitempty"`// Transform anchor point X, absent if no transform specified
	AnchorY				*float32				`json:"anchorY,omitempty"`// Transform anchor point Y, absent if no transform specified
	AnchorZ				*float32				`json:"anchorZ,omitempty"`// Transform anchor point Z, absent if no transform specified
	PaintCount			int					`json:"paintCount"`// Indicates how many time this layer has painted.
	DrawsContent			bool					`json:"drawsContent"`// Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	Invisible			*bool					`json:"invisible,omitempty"`// Set if layer is not visible.
	ScrollRects			[]LayerTree_ScrollRect			`json:"scrollRects,omitempty"`// Rectangles scrolling on main thread only.
	StickyPositionConstraint	*LayerTree_StickyPositionConstraint	`json:"stickyPositionConstraint,omitempty"`// Sticky position constraint information
}
type LayerTree_PaintProfile []float32
