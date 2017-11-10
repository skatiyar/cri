package domstorage

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type DOMStorage struct {
	conn cri.Connector
}

func New(conn cri.Connector) *DOMStorage {
	return &DOMStorage{conn}
}
func (obj *DOMStorage) Enable() (err error) {
	err = obj.conn.Send("DOMStorage.enable", nil, nil)
	return
}
func (obj *DOMStorage) Disable() (err error) {
	err = obj.conn.Send("DOMStorage.disable", nil, nil)
	return
}

type ClearRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) Clear(request *ClearRequest) (err error) {
	err = obj.conn.Send("DOMStorage.clear", request, nil)
	return
}

type GetDOMStorageItemsRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
}

func (obj *DOMStorage) GetDOMStorageItems(request *GetDOMStorageItemsRequest) (response struct {
	Entries []types.DOMStorage_Item `json:"entries"`
}, err error) {
	err = obj.conn.Send("DOMStorage.getDOMStorageItems", request, &response)
	return
}

type SetDOMStorageItemRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
	Value     string                     `json:"value"`
}

func (obj *DOMStorage) SetDOMStorageItem(request *SetDOMStorageItemRequest) (err error) {
	err = obj.conn.Send("DOMStorage.setDOMStorageItem", request, nil)
	return
}

type RemoveDOMStorageItemRequest struct {
	StorageId types.DOMStorage_StorageId `json:"storageId"`
	Key       string                     `json:"key"`
}

func (obj *DOMStorage) RemoveDOMStorageItem(request *RemoveDOMStorageItemRequest) (err error) {
	err = obj.conn.Send("DOMStorage.removeDOMStorageItem", request, nil)
	return
}
