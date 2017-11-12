/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/


package serviceworker

import (
    "github.com/SKatiyar/cri"
)

type ServiceWorker struct {
	conn cri.Connector
}

// New creates a ServiceWorker instance
func New(conn cri.Connector) *ServiceWorker {
	return &ServiceWorker{conn}
}

func (obj *ServiceWorker) Enable() (err error) {
	err = obj.conn.Send("ServiceWorker.enable", nil, nil)
	return
}


func (obj *ServiceWorker) Disable() (err error) {
	err = obj.conn.Send("ServiceWorker.disable", nil, nil)
	return
}


type UnregisterRequest struct {
	ScopeURL string `json:"scopeURL"`
}


func (obj *ServiceWorker) Unregister(request *UnregisterRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.unregister", request, nil)
	return
}


type UpdateRegistrationRequest struct {
	ScopeURL string `json:"scopeURL"`
}


func (obj *ServiceWorker) UpdateRegistration(request *UpdateRegistrationRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.updateRegistration", request, nil)
	return
}


type StartWorkerRequest struct {
	ScopeURL string `json:"scopeURL"`
}


func (obj *ServiceWorker) StartWorker(request *StartWorkerRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.startWorker", request, nil)
	return
}


type SkipWaitingRequest struct {
	ScopeURL string `json:"scopeURL"`
}


func (obj *ServiceWorker) SkipWaiting(request *SkipWaitingRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.skipWaiting", request, nil)
	return
}


type StopWorkerRequest struct {
	VersionId string `json:"versionId"`
}


func (obj *ServiceWorker) StopWorker(request *StopWorkerRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.stopWorker", request, nil)
	return
}


func (obj *ServiceWorker) StopAllWorkers() (err error) {
	err = obj.conn.Send("ServiceWorker.stopAllWorkers", nil, nil)
	return
}


type InspectWorkerRequest struct {
	VersionId string `json:"versionId"`
}


func (obj *ServiceWorker) InspectWorker(request *InspectWorkerRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.inspectWorker", request, nil)
	return
}


type SetForceUpdateOnPageLoadRequest struct {
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}


func (obj *ServiceWorker) SetForceUpdateOnPageLoad(request *SetForceUpdateOnPageLoadRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.setForceUpdateOnPageLoad", request, nil)
	return
}


type DeliverPushMessageRequest struct {
	Origin string `json:"origin"`
	RegistrationId string `json:"registrationId"`
	Data string `json:"data"`
}


func (obj *ServiceWorker) DeliverPushMessage(request *DeliverPushMessageRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.deliverPushMessage", request, nil)
	return
}


type DispatchSyncEventRequest struct {
	Origin string `json:"origin"`
	RegistrationId string `json:"registrationId"`
	Tag string `json:"tag"`
	LastChance bool `json:"lastChance"`
}


func (obj *ServiceWorker) DispatchSyncEvent(request *DispatchSyncEventRequest) (err error) {
	err = obj.conn.Send("ServiceWorker.dispatchSyncEvent", request, nil)
	return
}
