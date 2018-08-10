/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package serviceworker provides commands and events for ServiceWorker domain.
package serviceworker

import (
	"github.com/skatiyar/cri"
	types "github.com/skatiyar/cri/types"
)

// List of commands in ServiceWorker domain
const (
	DeliverPushMessage       = "ServiceWorker.deliverPushMessage"
	Disable                  = "ServiceWorker.disable"
	DispatchSyncEvent        = "ServiceWorker.dispatchSyncEvent"
	Enable                   = "ServiceWorker.enable"
	InspectWorker            = "ServiceWorker.inspectWorker"
	SetForceUpdateOnPageLoad = "ServiceWorker.setForceUpdateOnPageLoad"
	SkipWaiting              = "ServiceWorker.skipWaiting"
	StartWorker              = "ServiceWorker.startWorker"
	StopAllWorkers           = "ServiceWorker.stopAllWorkers"
	StopWorker               = "ServiceWorker.stopWorker"
	Unregister               = "ServiceWorker.unregister"
	UpdateRegistration       = "ServiceWorker.updateRegistration"
)

// List of events in ServiceWorker domain
const (
	WorkerErrorReported       = "ServiceWorker.workerErrorReported"
	WorkerRegistrationUpdated = "ServiceWorker.workerRegistrationUpdated"
	WorkerVersionUpdated      = "ServiceWorker.workerVersionUpdated"
)

type ServiceWorker struct {
	conn cri.Connector
}

// New creates a ServiceWorker instance
func New(conn cri.Connector) *ServiceWorker {
	return &ServiceWorker{conn}
}

type DeliverPushMessageRequest struct {
	Origin         string `json:"origin"`
	RegistrationId string `json:"registrationId"`
	Data           string `json:"data"`
}

func (obj *ServiceWorker) DeliverPushMessage(request *DeliverPushMessageRequest) (err error) {
	err = obj.conn.Send(DeliverPushMessage, request, nil)
	return
}

func (obj *ServiceWorker) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type DispatchSyncEventRequest struct {
	Origin         string `json:"origin"`
	RegistrationId string `json:"registrationId"`
	Tag            string `json:"tag"`
	LastChance     bool   `json:"lastChance"`
}

func (obj *ServiceWorker) DispatchSyncEvent(request *DispatchSyncEventRequest) (err error) {
	err = obj.conn.Send(DispatchSyncEvent, request, nil)
	return
}

func (obj *ServiceWorker) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type InspectWorkerRequest struct {
	VersionId string `json:"versionId"`
}

func (obj *ServiceWorker) InspectWorker(request *InspectWorkerRequest) (err error) {
	err = obj.conn.Send(InspectWorker, request, nil)
	return
}

type SetForceUpdateOnPageLoadRequest struct {
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}

func (obj *ServiceWorker) SetForceUpdateOnPageLoad(request *SetForceUpdateOnPageLoadRequest) (err error) {
	err = obj.conn.Send(SetForceUpdateOnPageLoad, request, nil)
	return
}

type SkipWaitingRequest struct {
	ScopeURL string `json:"scopeURL"`
}

func (obj *ServiceWorker) SkipWaiting(request *SkipWaitingRequest) (err error) {
	err = obj.conn.Send(SkipWaiting, request, nil)
	return
}

type StartWorkerRequest struct {
	ScopeURL string `json:"scopeURL"`
}

func (obj *ServiceWorker) StartWorker(request *StartWorkerRequest) (err error) {
	err = obj.conn.Send(StartWorker, request, nil)
	return
}

func (obj *ServiceWorker) StopAllWorkers() (err error) {
	err = obj.conn.Send(StopAllWorkers, nil, nil)
	return
}

type StopWorkerRequest struct {
	VersionId string `json:"versionId"`
}

func (obj *ServiceWorker) StopWorker(request *StopWorkerRequest) (err error) {
	err = obj.conn.Send(StopWorker, request, nil)
	return
}

type UnregisterRequest struct {
	ScopeURL string `json:"scopeURL"`
}

func (obj *ServiceWorker) Unregister(request *UnregisterRequest) (err error) {
	err = obj.conn.Send(Unregister, request, nil)
	return
}

type UpdateRegistrationRequest struct {
	ScopeURL string `json:"scopeURL"`
}

func (obj *ServiceWorker) UpdateRegistration(request *UpdateRegistrationRequest) (err error) {
	err = obj.conn.Send(UpdateRegistration, request, nil)
	return
}

type WorkerErrorReportedParams struct {
	ErrorMessage types.ServiceWorker_ServiceWorkerErrorMessage `json:"errorMessage"`
}

func (obj *ServiceWorker) WorkerErrorReported(fn func(event string, params WorkerErrorReportedParams, err error) bool) {
	listen, closer := obj.conn.On(WorkerErrorReported)
	go func() {
		defer closer()
		for {
			var params WorkerErrorReportedParams
			if !fn(WorkerErrorReported, params, listen(&params)) {
				return
			}
		}
	}()
}

type WorkerRegistrationUpdatedParams struct {
	Registrations []types.ServiceWorker_ServiceWorkerRegistration `json:"registrations"`
}

func (obj *ServiceWorker) WorkerRegistrationUpdated(fn func(event string, params WorkerRegistrationUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(WorkerRegistrationUpdated)
	go func() {
		defer closer()
		for {
			var params WorkerRegistrationUpdatedParams
			if !fn(WorkerRegistrationUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type WorkerVersionUpdatedParams struct {
	Versions []types.ServiceWorker_ServiceWorkerVersion `json:"versions"`
}

func (obj *ServiceWorker) WorkerVersionUpdated(fn func(event string, params WorkerVersionUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(WorkerVersionUpdated)
	go func() {
		defer closer()
		for {
			var params WorkerVersionUpdatedParams
			if !fn(WorkerVersionUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}
