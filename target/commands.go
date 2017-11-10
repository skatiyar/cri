package target

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Target struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Target {
	return &Target{conn}
}

type SetDiscoverTargetsRequest struct {
	Discover bool `json:"discover"`
}

func (obj *Target) SetDiscoverTargets(request *SetDiscoverTargetsRequest) (err error) {
	err = obj.conn.Send("Target.setDiscoverTargets", request, nil)
	return
}

type SetAutoAttachRequest struct {
	AutoAttach             bool `json:"autoAttach"`
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

func (obj *Target) SetAutoAttach(request *SetAutoAttachRequest) (err error) {
	err = obj.conn.Send("Target.setAutoAttach", request, nil)
	return
}

type SetAttachToFramesRequest struct {
	Value bool `json:"value"`
}

func (obj *Target) SetAttachToFrames(request *SetAttachToFramesRequest) (err error) {
	err = obj.conn.Send("Target.setAttachToFrames", request, nil)
	return
}

type SetRemoteLocationsRequest struct {
	Locations []types.Target_RemoteLocation `json:"locations"`
}

func (obj *Target) SetRemoteLocations(request *SetRemoteLocationsRequest) (err error) {
	err = obj.conn.Send("Target.setRemoteLocations", request, nil)
	return
}

type SendMessageToTargetRequest struct {
	Message   string                  `json:"message"`
	SessionId *types.Target_SessionID `json:"sessionId,omitempty"`
	TargetId  *types.Target_TargetID  `json:"targetId,omitempty"`
}

func (obj *Target) SendMessageToTarget(request *SendMessageToTargetRequest) (err error) {
	err = obj.conn.Send("Target.sendMessageToTarget", request, nil)
	return
}

type GetTargetInfoRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

func (obj *Target) GetTargetInfo(request *GetTargetInfoRequest) (response struct {
	TargetInfo types.Target_TargetInfo `json:"targetInfo"`
}, err error) {
	err = obj.conn.Send("Target.getTargetInfo", request, &response)
	return
}

type ActivateTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

func (obj *Target) ActivateTarget(request *ActivateTargetRequest) (err error) {
	err = obj.conn.Send("Target.activateTarget", request, nil)
	return
}

type CloseTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

func (obj *Target) CloseTarget(request *CloseTargetRequest) (response struct {
	Success bool `json:"success"`
}, err error) {
	err = obj.conn.Send("Target.closeTarget", request, &response)
	return
}

type AttachToTargetRequest struct {
	TargetId types.Target_TargetID `json:"targetId"`
}

func (obj *Target) AttachToTarget(request *AttachToTargetRequest) (response struct {
	SessionId types.Target_SessionID `json:"sessionId"`
}, err error) {
	err = obj.conn.Send("Target.attachToTarget", request, &response)
	return
}

type DetachFromTargetRequest struct {
	SessionId *types.Target_SessionID `json:"sessionId,omitempty"`
	TargetId  *types.Target_TargetID  `json:"targetId,omitempty"`
}

func (obj *Target) DetachFromTarget(request *DetachFromTargetRequest) (err error) {
	err = obj.conn.Send("Target.detachFromTarget", request, nil)
	return
}
func (obj *Target) CreateBrowserContext() (response struct {
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}, err error) {
	err = obj.conn.Send("Target.createBrowserContext", nil, &response)
	return
}

type DisposeBrowserContextRequest struct {
	BrowserContextId types.Target_BrowserContextID `json:"browserContextId"`
}

func (obj *Target) DisposeBrowserContext(request *DisposeBrowserContextRequest) (response struct {
	Success bool `json:"success"`
}, err error) {
	err = obj.conn.Send("Target.disposeBrowserContext", request, &response)
	return
}

type CreateTargetRequest struct {
	Url                     string                         `json:"url"`
	Width                   *int                           `json:"width,omitempty"`
	Height                  *int                           `json:"height,omitempty"`
	BrowserContextId        *types.Target_BrowserContextID `json:"browserContextId,omitempty"`
	EnableBeginFrameControl *bool                          `json:"enableBeginFrameControl,omitempty"`
}

func (obj *Target) CreateTarget(request *CreateTargetRequest) (response struct {
	TargetId types.Target_TargetID `json:"targetId"`
}, err error) {
	err = obj.conn.Send("Target.createTarget", request, &response)
	return
}
func (obj *Target) GetTargets() (response struct {
	TargetInfos []types.Target_TargetInfo `json:"targetInfos"`
}, err error) {
	err = obj.conn.Send("Target.getTargets", nil, &response)
	return
}
