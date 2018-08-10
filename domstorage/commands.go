/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package domstorage provides commands and events for DOMStorage domain.
package domstorage

import (
	types "github.com/SKatiyar/cri/types"
	"github.com/skatiyar/cri"
)

// List of commands in DOMStorage domain
const (
	Clear                = "DOMStorage.clear"
	Disable              = "DOMStorage.disable"
	Enable               = "DOMStorage.enable"
	GetDOMStorageItems   = "DOMStorage.getDOMStorageItems"
	RemoveDOMStorageItem = "DOMStorage.removeDOMStorageItem"
	SetDOMStorageItem    = "DOMStorage.setDOMStorageItem"
)

// List of events in DOMStorage domain
const (
	DomStorageItemAdded    = "DOMStorage.domStorageItemAdded"
	DomStorageItemRemoved  = "DOMStorage.domStorageItemRemoved"
	DomStorageItemUpdated  = "DOMStorage.domStorageItemUpdated"
	DomStorageItemsCleared = "DOMStorage.domStorageItemsCleared"
)

// Query and modify DOM storage.
type DOMStorage struct {
	conn cri.Connector
}

// New creates a DOMStorage instance
func New(conn cri.Connector) *DOMStorage {
	return &DOMStorage{conn}
}

type ClearRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) Clear(request *ClearRequest) (err error) {
	err = obj.conn.Send(Clear, request, nil)
	return
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (obj *DOMStorage) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables storage tracking, storage events will now be delivered to the client.
func (obj *DOMStorage) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
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

type RemoveDOMStorageItemRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
}

func (obj *DOMStorage) RemoveDOMStorageItem(request *RemoveDOMStorageItemRequest) (err error) {
	err = obj.conn.Send(RemoveDOMStorageItem, request, nil)
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

type DomStorageItemAddedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
	NewValue  string                     `json:"newValue"`
}

func (obj *DOMStorage) DomStorageItemAdded(fn func(event string, params DomStorageItemAddedParams, err error) bool) {
	listen, closer := obj.conn.On(DomStorageItemAdded)
	go func() {
		defer closer()
		for {
			var params DomStorageItemAddedParams
			if !fn(DomStorageItemAdded, params, listen(&params)) {
				return
			}
		}
	}()
}

type DomStorageItemRemovedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
}

func (obj *DOMStorage) DomStorageItemRemoved(fn func(event string, params DomStorageItemRemovedParams, err error) bool) {
	listen, closer := obj.conn.On(DomStorageItemRemoved)
	go func() {
		defer closer()
		for {
			var params DomStorageItemRemovedParams
			if !fn(DomStorageItemRemoved, params, listen(&params)) {
				return
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

func (obj *DOMStorage) DomStorageItemUpdated(fn func(event string, params DomStorageItemUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(DomStorageItemUpdated)
	go func() {
		defer closer()
		for {
			var params DomStorageItemUpdatedParams
			if !fn(DomStorageItemUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type DomStorageItemsClearedParams struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) DomStorageItemsCleared(fn func(event string, params DomStorageItemsClearedParams, err error) bool) {
	listen, closer := obj.conn.On(DomStorageItemsCleared)
	go func() {
		defer closer()
		for {
			var params DomStorageItemsClearedParams
			if !fn(DomStorageItemsCleared, params, listen(&params)) {
				return
			}
		}
	}()
}
