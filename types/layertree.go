package types

type LayerTree_LayerId string
type LayerTree_SnapshotId string
type LayerTree_ScrollRect struct {
	Rect DOM_Rect `json:"rect"`
	Type string   `json:"type"`
}
type LayerTree_StickyPositionConstraint struct {
	StickyBoxRect                       DOM_Rect           `json:"stickyBoxRect"`
	ContainingBlockRect                 DOM_Rect           `json:"containingBlockRect"`
	NearestLayerShiftingStickyBox       *LayerTree_LayerId `json:"nearestLayerShiftingStickyBox,omitempty"`
	NearestLayerShiftingContainingBlock *LayerTree_LayerId `json:"nearestLayerShiftingContainingBlock,omitempty"`
}
type LayerTree_PictureTile struct {
	X       float32 `json:"x"`
	Y       float32 `json:"y"`
	Picture string  `json:"picture"`
}
type LayerTree_Layer struct {
	LayerId                  LayerTree_LayerId                   `json:"layerId"`
	ParentLayerId            *LayerTree_LayerId                  `json:"parentLayerId,omitempty"`
	BackendNodeId            *DOM_BackendNodeId                  `json:"backendNodeId,omitempty"`
	OffsetX                  float32                             `json:"offsetX"`
	OffsetY                  float32                             `json:"offsetY"`
	Width                    float32                             `json:"width"`
	Height                   float32                             `json:"height"`
	Transform                []float32                           `json:"transform,omitempty"`
	AnchorX                  *float32                            `json:"anchorX,omitempty"`
	AnchorY                  *float32                            `json:"anchorY,omitempty"`
	AnchorZ                  *float32                            `json:"anchorZ,omitempty"`
	PaintCount               int                                 `json:"paintCount"`
	DrawsContent             bool                                `json:"drawsContent"`
	Invisible                *bool                               `json:"invisible,omitempty"`
	ScrollRects              []LayerTree_ScrollRect              `json:"scrollRects,omitempty"`
	StickyPositionConstraint *LayerTree_StickyPositionConstraint `json:"stickyPositionConstraint,omitempty"`
}
type LayerTree_PaintProfile []float32
