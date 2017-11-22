/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Animation instance.
type Animation_Animation struct {
	// <code>Animation</code>'s id.
	Id string `json:"id"`
	// <code>Animation</code>'s name.
	Name string `json:"name"`
	// <code>Animation</code>'s internal paused state.
	PausedState bool `json:"pausedState"`
	// <code>Animation</code>'s play state.
	PlayState string `json:"playState"`
	// <code>Animation</code>'s playback rate.
	PlaybackRate float32 `json:"playbackRate"`
	// <code>Animation</code>'s start time.
	StartTime float32 `json:"startTime"`
	// <code>Animation</code>'s current time.
	CurrentTime float32 `json:"currentTime"`
	// Animation type of <code>Animation</code>.
	Type string `json:"type"`
	// <code>Animation</code>'s source animation node.
	Source *Animation_AnimationEffect `json:"source,omitempty"`
	// A unique ID for <code>Animation</code> representing the sources that triggered this CSS animation/transition.
	CssId *string `json:"cssId,omitempty"`
}

// AnimationEffect instance
type Animation_AnimationEffect struct {
	// <code>AnimationEffect</code>'s delay.
	Delay float32 `json:"delay"`
	// <code>AnimationEffect</code>'s end delay.
	EndDelay float32 `json:"endDelay"`
	// <code>AnimationEffect</code>'s iteration start.
	IterationStart float32 `json:"iterationStart"`
	// <code>AnimationEffect</code>'s iterations.
	Iterations float32 `json:"iterations"`
	// <code>AnimationEffect</code>'s iteration duration.
	Duration float32 `json:"duration"`
	// <code>AnimationEffect</code>'s playback direction.
	Direction string `json:"direction"`
	// <code>AnimationEffect</code>'s fill mode.
	Fill string `json:"fill"`
	// <code>AnimationEffect</code>'s target node.
	BackendNodeId *DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// <code>AnimationEffect</code>'s keyframes.
	KeyframesRule *Animation_KeyframesRule `json:"keyframesRule,omitempty"`
	// <code>AnimationEffect</code>'s timing function.
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
	// <code>AnimationEffect</code>'s timing function.
	Easing string `json:"easing"`
}
