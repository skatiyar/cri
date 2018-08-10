/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Animation instance.
type Animation_Animation struct {
	// `Animation`'s id.
	Id string `json:"id"`
	// `Animation`'s name.
	Name string `json:"name"`
	// `Animation`'s internal paused state.
	PausedState bool `json:"pausedState"`
	// `Animation`'s play state.
	PlayState string `json:"playState"`
	// `Animation`'s playback rate.
	PlaybackRate float32 `json:"playbackRate"`
	// `Animation`'s start time.
	StartTime float32 `json:"startTime"`
	// `Animation`'s current time.
	CurrentTime float32 `json:"currentTime"`
	// Animation type of `Animation`.
	Type string `json:"type"`
	// `Animation`'s source animation node.
	Source *Animation_AnimationEffect `json:"source,omitempty"`
	// A unique ID for `Animation` representing the sources that triggered this CSS animation/transition.
	CssId *string `json:"cssId,omitempty"`
}

// AnimationEffect instance
type Animation_AnimationEffect struct {
	// `AnimationEffect`'s delay.
	Delay float32 `json:"delay"`
	// `AnimationEffect`'s end delay.
	EndDelay float32 `json:"endDelay"`
	// `AnimationEffect`'s iteration start.
	IterationStart float32 `json:"iterationStart"`
	// `AnimationEffect`'s iterations.
	Iterations float32 `json:"iterations"`
	// `AnimationEffect`'s iteration duration.
	Duration float32 `json:"duration"`
	// `AnimationEffect`'s playback direction.
	Direction string `json:"direction"`
	// `AnimationEffect`'s fill mode.
	Fill string `json:"fill"`
	// `AnimationEffect`'s target node.
	BackendNodeId *DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// `AnimationEffect`'s keyframes.
	KeyframesRule *Animation_KeyframesRule `json:"keyframesRule,omitempty"`
	// `AnimationEffect`'s timing function.
	Easing string `json:"easing"`
}

// Keyframes Rule
type Animation_KeyframesRule struct {
	// CSS keyframed animation's name.
	Name *string `json:"name,omitempty"`
	// List of animation keyframes.
	Keyframes []Animation_KeyframeStyle `json:"keyframes"`
}

// Keyframe Style
type Animation_KeyframeStyle struct {
	// Keyframe's time offset.
	Offset string `json:"offset"`
	// `AnimationEffect`'s timing function.
	Easing string `json:"easing"`
}
