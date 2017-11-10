package animation

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Animation struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Animation {
	return &Animation{conn}
}
func (obj *Animation) Enable() (err error) {
	err = obj.conn.Send("Animation.enable", nil, nil)
	return
}
func (obj *Animation) Disable() (err error) {
	err = obj.conn.Send("Animation.disable", nil, nil)
	return
}
func (obj *Animation) GetPlaybackRate() (response struct {
	PlaybackRate float32 `json:"playbackRate"`// Playback rate for animations on page.
}, err error) {
	err = obj.conn.Send("Animation.getPlaybackRate", nil, &response)
	return
}

type SetPlaybackRateRequest struct {
	PlaybackRate float32 `json:"playbackRate"`// Playback rate for animations on page
}

func (obj *Animation) SetPlaybackRate(request *SetPlaybackRateRequest) (err error) {
	err = obj.conn.Send("Animation.setPlaybackRate", request, nil)
	return
}

type GetCurrentTimeRequest struct {
	Id string `json:"id"`// Id of animation.
}

func (obj *Animation) GetCurrentTime(request *GetCurrentTimeRequest) (response struct {
	CurrentTime float32 `json:"currentTime"`// Current time of the page.
}, err error) {
	err = obj.conn.Send("Animation.getCurrentTime", request, &response)
	return
}

type SetPausedRequest struct {
	Animations	[]string	`json:"animations"`// Animations to set the pause state of.
	Paused		bool		`json:"paused"`// Paused state to set to.
}

func (obj *Animation) SetPaused(request *SetPausedRequest) (err error) {
	err = obj.conn.Send("Animation.setPaused", request, nil)
	return
}

type SetTimingRequest struct {
	AnimationId	string	`json:"animationId"`// Animation id.
	Duration	float32	`json:"duration"`// Duration of the animation.
	Delay		float32	`json:"delay"`// Delay of the animation.
}

func (obj *Animation) SetTiming(request *SetTimingRequest) (err error) {
	err = obj.conn.Send("Animation.setTiming", request, nil)
	return
}

type SeekAnimationsRequest struct {
	Animations	[]string	`json:"animations"`// List of animation ids to seek.
	CurrentTime	float32		`json:"currentTime"`// Set the current time of each animation.
}

func (obj *Animation) SeekAnimations(request *SeekAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.seekAnimations", request, nil)
	return
}

type ReleaseAnimationsRequest struct {
	Animations []string `json:"animations"`// List of animation ids to seek.
}

func (obj *Animation) ReleaseAnimations(request *ReleaseAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.releaseAnimations", request, nil)
	return
}

type ResolveAnimationRequest struct {
	AnimationId string `json:"animationId"`// Animation id.
}

func (obj *Animation) ResolveAnimation(request *ResolveAnimationRequest) (response struct {
	RemoteObject types.Runtime_RemoteObject `json:"remoteObject"`// Corresponding remote object.
}, err error) {
	err = obj.conn.Send("Animation.resolveAnimation", request, &response)
	return
}
