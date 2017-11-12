/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

//Unique Layer identifier.
type LayerTree_LayerId string

//Unique snapshot identifier.
type LayerTree_SnapshotId string

//Rectangle where scrolling happens on the main thread.
type LayerTree_ScrollRect struct {
	// Rectangle itself.
	Rect DOM_Rect `json:"rect"`
	// Reason for rectangle to force scrolling on the main thread
	Type string `json:"type"`
}

//Sticky position constraints.
type LayerTree_StickyPositionConstraint struct {
	// Layout rectangle of the sticky element before being shifted
	StickyBoxRect DOM_Rect `json:"stickyBoxRect"`
	// Layout rectangle of the containing block of the sticky element
	ContainingBlockRect DOM_Rect `json:"containingBlockRect"`
	// The nearest sticky layer that shifts the sticky box
	NearestLayerShiftingStickyBox *LayerTree_LayerId `json:"nearestLayerShiftingStickyBox,omitempty"`
	// The nearest sticky layer that shifts the containing block
	NearestLayerShiftingContainingBlock *LayerTree_LayerId `json:"nearestLayerShiftingContainingBlock,omitempty"`
}

//Serialized fragment of layer picture along with its offset within the layer.
type LayerTree_PictureTile struct {
	// Offset from owning layer left boundary
	X float32 `json:"x"`
	// Offset from owning layer top boundary
	Y float32 `json:"y"`
	// Base64-encoded snapshot data.
	Picture string `json:"picture"`
}

//Information about a compositing layer.
type LayerTree_Layer struct {
	// The unique id for this layer.
	LayerId LayerTree_LayerId `json:"layerId"`
	// The id of parent (not present for root).
	ParentLayerId *LayerTree_LayerId `json:"parentLayerId,omitempty"`
	// The backend id for the node associated with this layer.
	BackendNodeId *DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// Offset from parent layer, X coordinate.
	OffsetX float32 `json:"offsetX"`
	// Offset from parent layer, Y coordinate.
	OffsetY float32 `json:"offsetY"`
	// Layer width.
	Width float32 `json:"width"`
	// Layer height.
	Height float32 `json:"height"`
	// Transformation matrix for layer, default is identity matrix
	Transform []float32 `json:"transform,omitempty"`
	// Transform anchor point X, absent if no transform specified
	AnchorX *float32 `json:"anchorX,omitempty"`
	// Transform anchor point Y, absent if no transform specified
	AnchorY *float32 `json:"anchorY,omitempty"`
	// Transform anchor point Z, absent if no transform specified
	AnchorZ *float32 `json:"anchorZ,omitempty"`
	// Indicates how many time this layer has painted.
	PaintCount int `json:"paintCount"`
	// Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
	DrawsContent bool `json:"drawsContent"`
	// Set if layer is not visible.
	Invisible *bool `json:"invisible,omitempty"`
	// Rectangles scrolling on main thread only.
	ScrollRects []LayerTree_ScrollRect `json:"scrollRects,omitempty"`
	// Sticky position constraint information
	StickyPositionConstraint *LayerTree_StickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}

//Array of timings, one per paint step.
type LayerTree_PaintProfile []float32
