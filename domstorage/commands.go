/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package domstorage provides commands and events for DOMStorage domain.
package domstorage

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in DOMStorage domain
const (
	Enable               = "DOMStorage.enable"
	Disable              = "DOMStorage.disable"
	Clear                = "DOMStorage.clear"
	GetDOMStorageItems   = "DOMStorage.getDOMStorageItems"
	SetDOMStorageItem    = "DOMStorage.setDOMStorageItem"
	RemoveDOMStorageItem = "DOMStorage.removeDOMStorageItem"
)

// List of events in DOMStorage domain
const (
	DomStorageItemsCleared = "DOMStorage.domStorageItemsCleared"
	DomStorageItemRemoved  = "DOMStorage.domStorageItemRemoved"
	DomStorageItemAdded    = "DOMStorage.domStorageItemAdded"
	DomStorageItemUpdated  = "DOMStorage.domStorageItemUpdated"
)

// Query and modify DOM storage.
type DOMStorage struct {
	conn cri.Connector
}

// New creates a DOMStorage instance
func New(conn cri.Connector) *DOMStorage {
	return &DOMStorage{conn}
}

// Enables storage tracking, storage events will now be delivered to the client.
func (obj *DOMStorage) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (obj *DOMStorage) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type ClearRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) Clear(request *ClearRequest) (err error) {
	err = obj.conn.Send(Clear, request, nil)
	return
}

type GetDOMStorageItemsRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

type GetDOMStorageItemsResponse struct {
	Entries []types.DOMStorage_Item `json:"entries"`
}

func (obj *DOMStorage) GetDOMStorageItems(request *GetDOMStorageItemsRequest) (response GetDOMStorageItemsResponse, err error) {
	err = obj.conn.Send(GetDOMStorageItems, request, &response)
	return
}

type SetDOMStorageItemRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
	Value     string                     `json:"value"`
}

func (obj *DOMStorage) SetDOMStorageItem(request *SetDOMStorageItemRequest) (err error) {
	err = obj.conn.Send(SetDOMStorageItem, request, nil)
	return
}

type RemoveDOMStorageItemRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
}

func (obj *DOMStorage) RemoveDOMStorageItem(request *RemoveDOMStorageItemRequest) (err error) {
	err = obj.conn.Send(RemoveDOMStorageItem, request, nil)
	return
}

type DomStorageItemsClearedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) DomStorageItemsCleared(fn func(params *DomStorageItemsClearedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DomStorageItemsCleared, closeChn)
	go func() {
		for {
			params := DomStorageItemsClearedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type DomStorageItemRemovedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
}

func (obj *DOMStorage) DomStorageItemRemoved(fn func(params *DomStorageItemRemovedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DomStorageItemRemoved, closeChn)
	go func() {
		for {
			params := DomStorageItemRemovedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type DomStorageItemAddedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
	NewValue  string                     `json:"newValue"`
}

func (obj *DOMStorage) DomStorageItemAdded(fn func(params *DomStorageItemAddedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DomStorageItemAdded, closeChn)
	go func() {
		for {
			params := DomStorageItemAddedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type DomStorageItemUpdatedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
	OldValue  string                     `json:"oldValue"`
	NewValue  string                     `json:"newValue"`
}

func (obj *DOMStorage) DomStorageItemUpdated(fn func(params *DomStorageItemUpdatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DomStorageItemUpdated, closeChn)
	go func() {
		for {
			params := DomStorageItemUpdatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}
