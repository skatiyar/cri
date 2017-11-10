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
	PlaybackRate float32 `json:"playbackRate"`
}, err error) {
	err = obj.conn.Send("Animation.getPlaybackRate", nil, &response)
	return
}

type SetPlaybackRateRequest struct {
	PlaybackRate float32 `json:"playbackRate"`
}

func (obj *Animation) SetPlaybackRate(request *SetPlaybackRateRequest) (err error) {
	err = obj.conn.Send("Animation.setPlaybackRate", request, nil)
	return
}

type GetCurrentTimeRequest struct {
	Id string `json:"id"`
}

func (obj *Animation) GetCurrentTime(request *GetCurrentTimeRequest) (response struct {
	CurrentTime float32 `json:"currentTime"`
}, err error) {
	err = obj.conn.Send("Animation.getCurrentTime", request, &response)
	return
}

type SetPausedRequest struct {
	Animations []string `json:"animations"`
	Paused     bool     `json:"paused"`
}

func (obj *Animation) SetPaused(request *SetPausedRequest) (err error) {
	err = obj.conn.Send("Animation.setPaused", request, nil)
	return
}

type SetTimingRequest struct {
	AnimationId string  `json:"animationId"`
	Duration    float32 `json:"duration"`
	Delay       float32 `json:"delay"`
}

func (obj *Animation) SetTiming(request *SetTimingRequest) (err error) {
	err = obj.conn.Send("Animation.setTiming", request, nil)
	return
}

type SeekAnimationsRequest struct {
	Animations  []string `json:"animations"`
	CurrentTime float32  `json:"currentTime"`
}

func (obj *Animation) SeekAnimations(request *SeekAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.seekAnimations", request, nil)
	return
}

type ReleaseAnimationsRequest struct {
	Animations []string `json:"animations"`
}

func (obj *Animation) ReleaseAnimations(request *ReleaseAnimationsRequest) (err error) {
	err = obj.conn.Send("Animation.releaseAnimations", request, nil)
	return
}

type ResolveAnimationRequest struct {
	AnimationId string `json:"animationId"`
}

func (obj *Animation) ResolveAnimation(request *ResolveAnimationRequest) (response struct {
	RemoteObject types.Runtime_RemoteObject `json:"remoteObject"`
}, err error) {
	err = obj.conn.Send("Animation.resolveAnimation", request, &response)
	return
}
