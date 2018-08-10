/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package target provides commands and events for Target domain.
package target

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in Target domain
const (
	ActivateTarget         = "Target.activateTarget"
	AttachToTarget         = "Target.attachToTarget"
	AttachToBrowserTarget  = "Target.attachToBrowserTarget"
	CloseTarget            = "Target.closeTarget"
	ExposeDevToolsProtocol = "Target.exposeDevToolsProtocol"
	CreateBrowserContext   = "Target.createBrowserContext"
	GetBrowserContexts     = "Target.getBrowserContexts"
	CreateTarget           = "Target.createTarget"
	DetachFromTarget       = "Target.detachFromTarget"
	DisposeBrowserContext  = "Target.disposeBrowserContext"
	GetTargetInfo          = "Target.getTargetInfo"
	GetTargets             = "Target.getTargets"
	SendMessageToTarget    = "Target.sendMessageToTarget"
	SetAutoAttach          = "Target.setAutoAttach"
	SetDiscoverTargets     = "Target.setDiscoverTargets"
	SetRemoteLocations     = "Target.setRemoteLocations"
)

// List of events in Target domain
const (
	AttachedToTarget          = "Target.attachedToTarget"
	DetachedFromTarget        = "Target.detachedFromTarget"
	ReceivedMessageFromTarget = "Target.receivedMessageFromTarget"
	TargetCreated             = "Target.targetCreated"
	TargetDestroyed           = "Target.targetDestroyed"
	TargetCrashed             = "Target.targetCrashed"
	TargetInfoChanged         = "Target.targetInfoChanged"
)

// Supports additional targets discovery and allows to attach to them.
type Target struct {
	conn cri.Connector
}

// New creates a Target instance
func New(conn cri.Connector) *Target {
	return &Target{conn}
}

type ActivateTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

// Activates (focuses) the target.
func (obj *Target) ActivateTarget(request *ActivateTargetRequest) (err error) {
	err = obj.conn.Send(ActivateTarget, request, nil)
	return
}

type AttachToTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// NOTE Experimental
	Flatten *bool `json:"flatten,omitempty"`
}

type AttachToTargetResponse struct {
	// Id assigned to the session.
	SessionId types.Target_SessionID `json:"sessionId"`
}

// Attaches to the target with given id.
func (obj *Target) AttachToTarget(request *AttachToTargetRequest) (response AttachToTargetResponse, err error) {
	err = obj.conn.Send(AttachToTarget, request, &response)
	return
}

type AttachToBrowserTargetResponse struct {
	// Id assigned to the session.
	SessionId types.Target_SessionID `json:"sessionId"`
}

// Attaches to the browser target, only uses flat sessionId mode.
func (obj *Target) AttachToBrowserTarget() (response AttachToBrowserTargetResponse, err error) {
	err = obj.conn.Send(AttachToBrowserTarget, nil, &response)
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
	err = obj.conn.Send(CloseTarget, request, &response)
	return
}

type ExposeDevToolsProtocolRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
	// Binding name, 'cdp' if not specified.
	BindingName *string `json:"bindingName,omitempty"`
}

// Inject object to the target's main frame that provides a communication channel with browser target.  Injected object will be available as `window[bindingName]`.  The object has the follwing API: - `binding.send(json)` - a method to send messages over the remote debugging protocol - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
func (obj *Target) ExposeDevToolsProtocol(request *ExposeDevToolsProtocolRequest) (err error) {
	err = obj.conn.Send(ExposeDevToolsProtocol, request, nil)
	return
}

type CreateBrowserContextResponse struct {
	// The id of the context created.
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}

// Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
func (obj *Target) CreateBrowserContext() (response CreateBrowserContextResponse, err error) {
	err = obj.conn.Send(CreateBrowserContext, nil, &response)
	return
}

type GetBrowserContextsResponse struct {
	// An array of browser context ids.
	BrowserContextIds []types.Target_BrowserContextID `json:"browserContextIds"`
}

// Returns all browser contexts created with `Target.createBrowserContext` method.
func (obj *Target) GetBrowserContexts() (response GetBrowserContextsResponse, err error) {
	err = obj.conn.Send(GetBrowserContexts, nil, &response)
	return
}

type CreateTargetRequest struct {
	// The initial URL the page will be navigated to.
	Url string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width *int `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height *int `json:"height,omitempty"`
	// The browser context to create the page in.
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
	err = obj.conn.Send(CreateTarget, request, &response)
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
	err = obj.conn.Send(DetachFromTarget, request, nil)
	return
}

type DisposeBrowserContextRequest struct {
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}

// Deletes a BrowserContext. All the belonging pages will be closed without calling their beforeunload hooks.
func (obj *Target) DisposeBrowserContext(request *DisposeBrowserContextRequest) (err error) {
	err = obj.conn.Send(DisposeBrowserContext, request, nil)
	return
}

type GetTargetInfoRequest struct {
	TargetId *types.Target_TargetID `json:"targetId,omitempty"`
}

type GetTargetInfoResponse struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Returns information about a target.
func (obj *Target) GetTargetInfo(request *GetTargetInfoRequest) (response GetTargetInfoResponse, err error) {
	err = obj.conn.Send(GetTargetInfo, request, &response)
	return
}

type GetTargetsResponse struct {
	// The list of targets.
	TargetInfos []types.Target_TargetInfo `json:"targetInfos"`
}

// Retrieves a list of available targets.
func (obj *Target) GetTargets() (response GetTargetsResponse, err error) {
	err = obj.conn.Send(GetTargets, nil, &response)
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
	err = obj.conn.Send(SendMessageToTarget, request, nil)
	return
}

type SetAutoAttachRequest struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`
	// Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger` to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
	// Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// NOTE Experimental
	Flatten *bool `json:"flatten,omitempty"`
}

// Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (obj *Target) SetAutoAttach(request *SetAutoAttachRequest) (err error) {
	err = obj.conn.Send(SetAutoAttach, request, nil)
	return
}

type SetDiscoverTargetsRequest struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

// Controls whether to discover available targets and notify via `targetCreated/targetInfoChanged/targetDestroyed` events.
func (obj *Target) SetDiscoverTargets(request *SetDiscoverTargetsRequest) (err error) {
	err = obj.conn.Send(SetDiscoverTargets, request, nil)
	return
}

type SetRemoteLocationsRequest struct {
	// List of remote locations.
	Locations []types.Target_RemoteLocation `json:"locations"`
}

// Enables target discovery for the specified locations, when `setDiscoverTargets` was set to `true`.
func (obj *Target) SetRemoteLocations(request *SetRemoteLocationsRequest) (err error) {
	err = obj.conn.Send(SetRemoteLocations, request, nil)
	return
}

type AttachedToTargetParams struct {
	// Identifier assigned to the session used to send/receive messages.
	SessionId          types.Target_SessionID  `json:"sessionId"`
	TargetInfo         types.Target_TargetInfo `json:"targetInfo"`
	WaitingForDebugger bool                    `json:"waitingForDebugger"`
}

// Issued when attached to target because of auto-attach or `attachToTarget` command.
// NOTE Experimental
func (obj *Target) AttachedToTarget(fn func(event string, params AttachedToTargetParams, err error) bool) {
	listen, closer := obj.conn.On(AttachedToTarget)
	go func() {
		defer closer()
		for {
			var params AttachedToTargetParams
			if !fn(AttachedToTarget, params, listen(&params)) {
				return
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

// Issued when detached from target for any reason (including `detachFromTarget` command). Can be issued multiple times per target if multiple sessions have been attached to it.
// NOTE Experimental
func (obj *Target) DetachedFromTarget(fn func(event string, params DetachedFromTargetParams, err error) bool) {
	listen, closer := obj.conn.On(DetachedFromTarget)
	go func() {
		defer closer()
		for {
			var params DetachedFromTargetParams
			if !fn(DetachedFromTarget, params, listen(&params)) {
				return
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

// Notifies about a new protocol message received from the session (as reported in `attachedToTarget` event).
func (obj *Target) ReceivedMessageFromTarget(fn func(event string, params ReceivedMessageFromTargetParams, err error) bool) {
	listen, closer := obj.conn.On(ReceivedMessageFromTarget)
	go func() {
		defer closer()
		for {
			var params ReceivedMessageFromTargetParams
			if !fn(ReceivedMessageFromTarget, params, listen(&params)) {
				return
			}
		}
	}()
}

type TargetCreatedParams struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Issued when a possible inspection target is created.
func (obj *Target) TargetCreated(fn func(event string, params TargetCreatedParams, err error) bool) {
	listen, closer := obj.conn.On(TargetCreated)
	go func() {
		defer closer()
		for {
			var params TargetCreatedParams
			if !fn(TargetCreated, params, listen(&params)) {
				return
			}
		}
	}()
}

type TargetDestroyedParams struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

// Issued when a target is destroyed.
func (obj *Target) TargetDestroyed(fn func(event string, params TargetDestroyedParams, err error) bool) {
	listen, closer := obj.conn.On(TargetDestroyed)
	go func() {
		defer closer()
		for {
			var params TargetDestroyedParams
			if !fn(TargetDestroyed, params, listen(&params)) {
				return
			}
		}
	}()
}

type TargetCrashedParams struct {
	TargetId types.Target_TargetID `json:"targetId"`
	// Termination status type.
	Status string `json:"status"`
	// Termination error code.
	ErrorCode int `json:"errorCode"`
}

// Issued when a target has crashed.
func (obj *Target) TargetCrashed(fn func(event string, params TargetCrashedParams, err error) bool) {
	listen, closer := obj.conn.On(TargetCrashed)
	go func() {
		defer closer()
		for {
			var params TargetCrashedParams
			if !fn(TargetCrashed, params, listen(&params)) {
				return
			}
		}
	}()
}

type TargetInfoChangedParams struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}

// Issued when some information about a target has changed. This only happens between `targetCreated` and `targetDestroyed`.
func (obj *Target) TargetInfoChanged(fn func(event string, params TargetInfoChangedParams, err error) bool) {
	listen, closer := obj.conn.On(TargetInfoChanged)
	go func() {
		defer closer()
		for {
			var params TargetInfoChangedParams
			if !fn(TargetInfoChanged, params, listen(&params)) {
				return
			}
		}
	}()
}
