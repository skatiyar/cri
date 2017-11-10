package types

type Animation_Animation struct {
	Id		string				`json:"id"`// <code>Animation</code>'s id.
	Name		string				`json:"name"`// <code>Animation</code>'s name.
	PausedState	bool				`json:"pausedState"`// <code>Animation</code>'s internal paused state.
	PlayState	string				`json:"playState"`// <code>Animation</code>'s play state.
	PlaybackRate	float32				`json:"playbackRate"`// <code>Animation</code>'s playback rate.
	StartTime	float32				`json:"startTime"`// <code>Animation</code>'s start time.
	CurrentTime	float32				`json:"currentTime"`// <code>Animation</code>'s current time.
	Type		string				`json:"type"`// Animation type of <code>Animation</code>.
	Source		*Animation_AnimationEffect	`json:"source,omitempty"`// <code>Animation</code>'s source animation node.
	CssId		*string				`json:"cssId,omitempty"`// A unique ID for <code>Animation</code> representing the sources that triggered this CSS animation/transition.
}
type Animation_AnimationEffect struct {
	Delay		float32				`json:"delay"`// <code>AnimationEffect</code>'s delay.
	EndDelay	float32				`json:"endDelay"`// <code>AnimationEffect</code>'s end delay.
	IterationStart	float32				`json:"iterationStart"`// <code>AnimationEffect</code>'s iteration start.
	Iterations	float32				`json:"iterations"`// <code>AnimationEffect</code>'s iterations.
	Duration	float32				`json:"duration"`// <code>AnimationEffect</code>'s iteration duration.
	Direction	string				`json:"direction"`// <code>AnimationEffect</code>'s playback direction.
	Fill		string				`json:"fill"`// <code>AnimationEffect</code>'s fill mode.
	BackendNodeId	*DOM_BackendNodeId		`json:"backendNodeId,omitempty"`// <code>AnimationEffect</code>'s target node.
	KeyframesRule	*Animation_KeyframesRule	`json:"keyframesRule,omitempty"`// <code>AnimationEffect</code>'s keyframes.
	Easing		string				`json:"easing"`// <code>AnimationEffect</code>'s timing function.
}
type Animation_KeyframesRule struct {
	Name		*string				`json:"name,omitempty"`// CSS keyframed animation's name.
	Keyframes	[]Animation_KeyframeStyle	`json:"keyframes"`// List of animation keyframes.
}
type Animation_KeyframeStyle struct {
	Offset	string	`json:"offset"`// Keyframe's time offset.
	Easing	string	`json:"easing"`// <code>AnimationEffect</code>'s timing function.
}
