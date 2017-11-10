package types

type Animation_Animation struct {
	Id           string                     `json:"id"`
	Name         string                     `json:"name"`
	PausedState  bool                       `json:"pausedState"`
	PlayState    string                     `json:"playState"`
	PlaybackRate float32                    `json:"playbackRate"`
	StartTime    float32                    `json:"startTime"`
	CurrentTime  float32                    `json:"currentTime"`
	Type         string                     `json:"type"`
	Source       *Animation_AnimationEffect `json:"source,omitempty"`
	CssId        *string                    `json:"cssId,omitempty"`
}
type Animation_AnimationEffect struct {
	Delay          float32                  `json:"delay"`
	EndDelay       float32                  `json:"endDelay"`
	IterationStart float32                  `json:"iterationStart"`
	Iterations     float32                  `json:"iterations"`
	Duration       float32                  `json:"duration"`
	Direction      string                   `json:"direction"`
	Fill           string                   `json:"fill"`
	BackendNodeId  *DOM_BackendNodeId       `json:"backendNodeId,omitempty"`
	KeyframesRule  *Animation_KeyframesRule `json:"keyframesRule,omitempty"`
	Easing         string                   `json:"easing"`
}
type Animation_KeyframesRule struct {
	Name      *string                   `json:"name,omitempty"`
	Keyframes []Animation_KeyframeStyle `json:"keyframes"`
}
type Animation_KeyframeStyle struct {
	Offset string `json:"offset"`
	Easing string `json:"easing"`
}
