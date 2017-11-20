/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package animation provides commands and events for Animation domain.
package animation

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Animation domain
const (
	Enable            = "Animation.enable"
	Disable           = "Animation.disable"
	GetPlaybackRate   = "Animation.getPlaybackRate"
	SetPlaybackRate   = "Animation.setPlaybackRate"
	GetCurrentTime    = "Animation.getCurrentTime"
	SetPaused         = "Animation.setPaused"
	SetTiming         = "Animation.setTiming"
	SeekAnimations    = "Animation.seekAnimations"
	ReleaseAnimations = "Animation.releaseAnimations"
	ResolveAnimation  = "Animation.resolveAnimation"
)

// List of events in Animation domain
const (
	AnimationCreated  = "Animation.animationCreated"
	AnimationStarted  = "Animation.animationStarted"
	AnimationCanceled = "Animation.animationCanceled"
)

type Animation struct {
	conn cri.Connector
}

// New creates a Animation instance
func New(conn cri.Connector) *Animation {
	return &Animation{conn}
}

// Enables animation domain notifications.
func (obj *Animation) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disables animation domain notifications.
func (obj *Animation) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type GetPlaybackRateResponse struct {
	// Playback rate for animations on page.
	PlaybackRate float32 `json:"playbackRate"`
}

// Gets the playback rate of the document timeline.
func (obj *Animation) GetPlaybackRate() (response GetPlaybackRateResponse, err error) {
	err = obj.conn.Send(GetPlaybackRate, nil, &response)
	return
}

type SetPlaybackRateRequest struct {
	// Playback rate for animations on page
	PlaybackRate float32 `json:"playbackRate"`
}

// Sets the playback rate of the document timeline.
func (obj *Animation) SetPlaybackRate(request *SetPlaybackRateRequest) (err error) {
	err = obj.conn.Send(SetPlaybackRate, request, nil)
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
	err = obj.conn.Send(GetCurrentTime, request, &response)
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
	err = obj.conn.Send(SetPaused, request, nil)
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
	err = obj.conn.Send(SetTiming, request, nil)
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
	err = obj.conn.Send(SeekAnimations, request, nil)
	return
}

type ReleaseAnimationsRequest struct {
	// List of animation ids to seek.
	Animations []string `json:"animations"`
}

// Releases a set of animations to no longer be manipulated.
func (obj *Animation) ReleaseAnimations(request *ReleaseAnimationsRequest) (err error) {
	err = obj.conn.Send(ReleaseAnimations, request, nil)
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
	err = obj.conn.Send(ResolveAnimation, request, &response)
	return
}

type AnimationCreatedParams struct {
	// Id of the animation that was created.
	Id string `json:"id"`
}

// Event for each animation that has been created.
func (obj *Animation) AnimationCreated(fn func(params *AnimationCreatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AnimationCreated, closeChn)
	go func() {
		for {
			params := AnimationCreatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type AnimationStartedParams struct {
	// Animation that was started.
	Animation types.Animation_Animation `json:"animation"`
}

// Event for animation that has been started.
func (obj *Animation) AnimationStarted(fn func(params *AnimationStartedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AnimationStarted, closeChn)
	go func() {
		for {
			params := AnimationStartedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type AnimationCanceledParams struct {
	// Id of the animation that was cancelled.
	Id string `json:"id"`
}

// Event for when an animation has been cancelled.
func (obj *Animation) AnimationCanceled(fn func(params *AnimationCanceledParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AnimationCanceled, closeChn)
	go func() {
		for {
			params := AnimationCanceledParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
