package animation

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Animation struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Animation {
	return &Animation{conn}
}

// Enables animation domain notifications.
func (obj *Animation) Enable() (err error) {
	err = obj.conn.Send("Animation.enable", nil, nil)
	return
}

// Disables animation domain notifications.
func (obj *Animation) Disable() (err error) {
	err = obj.conn.Send("Animation.disable", nil, nil)
	return
}

type GetPlaybackRateResponse struct {
	// Playback rate for animations on page.
	PlaybackRate float32 `json:"playbackRate"`
}

// Gets the playback rate of the document timeline.
func (obj *Animation) GetPlaybackRate() (response GetPlaybackRateResponse, err error) {
	err = obj.conn.Send("Animation.getPlaybackRate", nil, &response)
	return
}

type SetPlaybackRateRequest struct {
	// Playback rate for animations on page
	PlaybackRate float32 `json:"playbackRate"`
}

// Sets the playback rate of the document timeline.
func (obj *Animation) SetPlaybackRate(request *SetPlaybackRateRequest) (err error) {
	err = obj.conn.Send("Animation.setPlaybackRate", request, nil)
	return
}

type GetCurrentTimeRequest struct {
	// Id of animation.
	Id string `json:"id"`
}
type GetCurrentTimeResponse struct {
	// Current time of the page.
	CurrentTime float32 `json:"currentTime"`
}

// Returns the current time of the an animation.
func (obj *Animation) GetCurrentTime(request *GetCurrentTimeRequest) (response GetCurrentTimeResponse, err error) {
	err = obj.conn.Send("Animation.getCurrentTime", request, &response)
	return
}

type SetPausedRequest struct {
	// Animations to set the pause state of.
	Animations []string `json:"animations"`
	// Paused state to set to.
	Paused bool `json:"paused"`
}

// Sets the paused state of a set of animations.
func (obj *Animation) SetPaused(request *SetPausedRequest) (err error) {
	err = obj.conn.Send("Animation.setPaused", request, nil)
	return
}

type SetTimingRequest struct {
	// Animation id.
	AnimationId string `json:"animationId"`
	// Duration of the animation.
	Duration float32 `json:"duration"`
	// Delay of the animation.
	Delay float32 `json:"delay"`
}

// Sets the timing of an animation node.
func (obj *Animation) SetTiming(request *SetTimingRequest) (err error) {
	err = obj.conn.Send("Animation.setTiming", request, nil)
	return
}

type SeekAnimationsRequest struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
	// Set the current time of each animation.
	CurrentTime float32 `json:"currentTime"`
}

// Seek a set of animations to a particular time within each animation.
func (obj *Animation) SeekAnimations(request *SeekAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.seekAnimations", request, nil)
	return
}

type ReleaseAnimationsRequest struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

// Releases a set of animations to no longer be manipulated.
func (obj *Animation) ReleaseAnimations(request *ReleaseAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.releaseAnimations", request, nil)
	return
}

type ResolveAnimationRequest struct {
	// Animation id.
	AnimationId string `json:"animationId"`
}
type ResolveAnimationResponse struct {
	// Corresponding remote object.
	RemoteObject types.Runtime_RemoteObject `json:"remoteObject"`
}

// Gets the remote object of the Animation.
func (obj *Animation) ResolveAnimation(request *ResolveAnimationRequest) (response ResolveAnimationResponse, err error) {
	err = obj.conn.Send("Animation.resolveAnimation", request, &response)
	return
}
