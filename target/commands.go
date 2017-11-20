/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package target provides commands and events for Target domain.
package target

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// Supports additional targets discovery and allows to attach to them.
type Target struct {
	conn cri.Connector
}

// New creates a Target instance
func New(conn cri.Connector) *Target {
	return &Target{conn}
}

type SetDiscoverTargetsRequest struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

// Controls whether to discover available targets and notify via <code>targetCreated/targetInfoChanged/targetDestroyed</code> events.
func (obj *Target) SetDiscoverTargets(request *SetDiscoverTargetsRequest) (err error) {
	err = obj.conn.Send("Target.setDiscoverTargets", request, nil)
	return
}

type SetAutoAttachRequest struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`
	// Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

// Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (obj *Target) SetAutoAttach(request *SetAutoAttachRequest) (err error) {
	err = obj.conn.Send("Target.setAutoAttach", request, nil)
	return
}

type SetAttachToFramesRequest struct {
	// Whether to attach to frames.
	Value bool `json:"value"`
}

func (obj *Target) SetAttachToFrames(request *SetAttachToFramesRequest) (err error) {
	err = obj.conn.Send("Target.setAttachToFrames", request, nil)
	return
}

type SetRemoteLocationsRequest struct {
	// List of remote locations.
	Locations []types.Target_RemoteLocation `json:"locations"`
}

// Enables target discovery for the specified locations, when <code>setDiscoverTargets</code> was set to <code>true</code>.
func (obj *Target) SetRemoteLocations(request *SetRemoteLocationsRequest) (err error) {
	err = obj.conn.Send("Target.setRemoteLocations", request, nil)
	return
}

type SendMessageToTargetRequest struct {
	Message string `json:"message"`
	// Identifier of the session.
	SessionId *types.Target_SessionID `json:"sessionId,omitempty"`
	// Deprecated.
	TargetId *types.Target_TargetID `json:"targetId,omitempty"`
}

// Sends protocol message over session with given id.
func (obj *Target) SendMessageToTarget(request *SendMessageToTargetRequest) (err error) {
	err = obj.conn.Send("Target.sendMessageToTarget", request, nil)
	return
}

type GetTargetInfoRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

type GetTargetInfoResponse struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Returns information about a target.
func (obj *Target) GetTargetInfo(request *GetTargetInfoRequest) (response GetTargetInfoResponse, err error) {
	err = obj.conn.Send("Target.getTargetInfo", request, &response)
	return
}

type ActivateTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

// Activates (focuses) the target.
func (obj *Target) ActivateTarget(request *ActivateTargetRequest) (err error) {
	err = obj.conn.Send("Target.activateTarget", request, nil)
	return
}

type CloseTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

type CloseTargetResponse struct {
	Success bool `json:"success"`
}

// Closes the target. If the target is a page that gets closed too.
func (obj *Target) CloseTarget(request *CloseTargetRequest) (response CloseTargetResponse, err error) {
	err = obj.conn.Send("Target.closeTarget", request, &response)
	return
}

type AttachToTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

type AttachToTargetResponse struct {
	// Id assigned to the session.
	SessionId types.Target_SessionID `json:"sessionId"`
}

// Attaches to the target with given id.
func (obj *Target) AttachToTarget(request *AttachToTargetRequest) (response AttachToTargetResponse, err error) {
	err = obj.conn.Send("Target.attachToTarget", request, &response)
	return
}

type DetachFromTargetRequest struct {
	// Session to detach.
	SessionId *types.Target_SessionID `json:"sessionId,omitempty"`
	// Deprecated.
	TargetId *types.Target_TargetID `json:"targetId,omitempty"`
}

// Detaches session with given id.
func (obj *Target) DetachFromTarget(request *DetachFromTargetRequest) (err error) {
	err = obj.conn.Send("Target.detachFromTarget", request, nil)
	return
}

type CreateBrowserContextResponse struct {
	// The id of the context created.
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}

// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
func (obj *Target) CreateBrowserContext() (response CreateBrowserContextResponse, err error) {
	err = obj.conn.Send("Target.createBrowserContext", nil, &response)
	return
}

type DisposeBrowserContextRequest struct {
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}

type DisposeBrowserContextResponse struct {
	Success bool `json:"success"`
}

// Deletes a BrowserContext, will fail of any open page uses it.
func (obj *Target) DisposeBrowserContext(request *DisposeBrowserContextRequest) (response DisposeBrowserContextResponse, err error) {
	err = obj.conn.Send("Target.disposeBrowserContext", request, &response)
	return
}

type CreateTargetRequest struct {
	// The initial URL the page will be navigated to.
	Url string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width *int `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height *int `json:"height,omitempty"`
	// The browser context to create the page in (headless chrome only).
	BrowserContextId *types.Target_BrowserContextID `json:"browserContextId,omitempty"`
	// Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
	// NOTE Experimental
	EnableBeginFrameControl *bool `json:"enableBeginFrameControl,omitempty"`
}

type CreateTargetResponse struct {
	// The id of the page opened.
	TargetId types.Target_TargetID `json:"targetId"`
}

// Creates a new page.
func (obj *Target) CreateTarget(request *CreateTargetRequest) (response CreateTargetResponse, err error) {
	err = obj.conn.Send("Target.createTarget", request, &response)
	return
}

type GetTargetsResponse struct {
	// The list of targets.
	TargetInfos []types.Target_TargetInfo `json:"targetInfos"`
}

// Retrieves a list of available targets.
func (obj *Target) GetTargets() (response GetTargetsResponse, err error) {
	err = obj.conn.Send("Target.getTargets", nil, &response)
	return
}

type TargetCreatedParams struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Issued when a possible inspection target is created.
func (obj *Target) TargetCreated(fn func(params *TargetCreatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.targetCreated", closeChn)
	go func() {
		for {
			params := TargetCreatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type TargetInfoChangedParams struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Issued when some information about a target has changed. This only happens between <code>targetCreated</code> and <code>targetDestroyed</code>.
func (obj *Target) TargetInfoChanged(fn func(params *TargetInfoChangedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.targetInfoChanged", closeChn)
	go func() {
		for {
			params := TargetInfoChangedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type TargetDestroyedParams struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

// Issued when a target is destroyed.
func (obj *Target) TargetDestroyed(fn func(params *TargetDestroyedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.targetDestroyed", closeChn)
	go func() {
		for {
			params := TargetDestroyedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type AttachedToTargetParams struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionId          types.Target_SessionID  `json:"sessionId"`
	TargetInfo         types.Target_TargetInfo `json:"targetInfo"`
	WaitingForDebugger bool                    `json:"waitingForDebugger"`
}

// Issued when attached to target because of auto-attach or <code>attachToTarget</code> command.
func (obj *Target) AttachedToTarget(fn func(params *AttachedToTargetParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.attachedToTarget", closeChn)
	go func() {
		for {
			params := AttachedToTargetParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type DetachedFromTargetParams struct {
	// Detached session identifier.
	SessionId types.Target_SessionID `json:"sessionId"`
	// Deprecated.
	TargetId *types.Target_TargetID `json:"targetId,omitempty"`
}

// Issued when detached from target for any reason (including <code>detachFromTarget</code> command). Can be issued multiple times per target if multiple sessions have been attached to it.
func (obj *Target) DetachedFromTarget(fn func(params *DetachedFromTargetParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.detachedFromTarget", closeChn)
	go func() {
		for {
			params := DetachedFromTargetParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ReceivedMessageFromTargetParams struct {
	// Identifier of a session which sends a message.
	SessionId types.Target_SessionID `json:"sessionId"`
	Message   string                 `json:"message"`
	// Deprecated.
	TargetId *types.Target_TargetID `json:"targetId,omitempty"`
}

// Notifies about a new protocol message received from the session (as reported in <code>attachedToTarget</code> event).
func (obj *Target) ReceivedMessageFromTarget(fn func(params *ReceivedMessageFromTargetParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Target.receivedMessageFromTarget", closeChn)
	go func() {
		for {
			params := ReceivedMessageFromTargetParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
