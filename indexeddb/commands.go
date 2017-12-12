/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package indexeddb provides commands and events for IndexedDB domain.
package indexeddb

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in IndexedDB domain
const (
	ClearObjectStore         = "IndexedDB.clearObjectStore"
	DeleteDatabase           = "IndexedDB.deleteDatabase"
	DeleteObjectStoreEntries = "IndexedDB.deleteObjectStoreEntries"
	Disable                  = "IndexedDB.disable"
	Enable                   = "IndexedDB.enable"
	RequestData              = "IndexedDB.requestData"
	RequestDatabase          = "IndexedDB.requestDatabase"
	RequestDatabaseNames     = "IndexedDB.requestDatabaseNames"
)

type IndexedDB struct {
	conn cri.Connector
}

// New creates a IndexedDB instance
func New(conn cri.Connector) *IndexedDB {
	return &IndexedDB{conn}
}

type ClearObjectStoreRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
}

// Clears all entries from an object store.
func (obj *IndexedDB) ClearObjectStore(request *ClearObjectStoreRequest) (err error) {
	err = obj.conn.Send(ClearObjectStore, request, nil)
	return
}

type DeleteDatabaseRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

// Deletes a database.
func (obj *IndexedDB) DeleteDatabase(request *DeleteDatabaseRequest) (err error) {
	err = obj.conn.Send(DeleteDatabase, request, nil)
	return
}

type DeleteObjectStoreEntriesRequest struct {
	SecurityOrigin  string `json:"securityOrigin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
	// Range of entry keys to delete
	KeyRange types.IndexedDB_KeyRange `json:"keyRange"`
}

// Delete a range of entries from an object store
func (obj *IndexedDB) DeleteObjectStoreEntries(request *DeleteObjectStoreEntriesRequest) (err error) {
	err = obj.conn.Send(DeleteObjectStoreEntries, request, nil)
	return
}

// Disables events from backend.
func (obj *IndexedDB) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

// Enables events from backend.
func (obj *IndexedDB) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

type RequestDataRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
	// Object store name.
	ObjectStoreName string `json:"objectStoreName"`
	// Index name, empty string for object store data requests.
	IndexName string `json:"indexName"`
	// Number of records to skip.
	SkipCount int `json:"skipCount"`
	// Number of records to fetch.
	PageSize int `json:"pageSize"`
	// Key range.
	KeyRange *types.IndexedDB_KeyRange `json:"keyRange,omitempty"`
}

type RequestDataResponse struct {
	// Array of object store data entries.
	ObjectStoreDataEntries []types.IndexedDB_DataEntry `json:"objectStoreDataEntries"`
	// If true, there are more entries to fetch in the given range.
	HasMore bool `json:"hasMore"`
}

// Requests data from object store or index.
func (obj *IndexedDB) RequestData(request *RequestDataRequest) (response RequestDataResponse, err error) {
	err = obj.conn.Send(RequestData, request, &response)
	return
}

type RequestDatabaseRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
	// Database name.
	DatabaseName string `json:"databaseName"`
}

type RequestDatabaseResponse struct {
	// Database with an array of object stores.
	DatabaseWithObjectStores types.IndexedDB_DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}

// Requests database with given name in given frame.
func (obj *IndexedDB) RequestDatabase(request *RequestDatabaseRequest) (response RequestDatabaseResponse, err error) {
	err = obj.conn.Send(RequestDatabase, request, &response)
	return
}

type RequestDatabaseNamesRequest struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

type RequestDatabaseNamesResponse struct {
	// Database names for origin.
	DatabaseNames []string `json:"databaseNames"`
}

// Requests database names for given security origin.
func (obj *IndexedDB) RequestDatabaseNames(request *RequestDatabaseNamesRequest) (response RequestDatabaseNamesResponse, err error) {
	err = obj.conn.Send(RequestDatabaseNames, request, &response)
	return
}
