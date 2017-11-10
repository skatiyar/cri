package indexeddb

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type IndexedDB struct {
	conn cri.Connector
}

func New(conn cri.Connector) *IndexedDB {
	return &IndexedDB{conn}
}
func (obj *IndexedDB) Enable() (err error) {
	err = obj.conn.Send("IndexedDB.enable", nil, nil)
	return
}
func (obj *IndexedDB) Disable() (err error) {
	err = obj.conn.Send("IndexedDB.disable", nil, nil)
	return
}

type RequestDatabaseNamesRequest struct {
	SecurityOrigin string `json:"securityOrigin"`// Security origin.
}

func (obj *IndexedDB) RequestDatabaseNames(request *RequestDatabaseNamesRequest) (response struct {
	DatabaseNames []string `json:"databaseNames"`// Database names for origin.
}, err error) {
	err = obj.conn.Send("IndexedDB.requestDatabaseNames", request, &response)
	return
}

type RequestDatabaseRequest struct {
	SecurityOrigin	string	`json:"securityOrigin"`// Security origin.
	DatabaseName	string	`json:"databaseName"`// Database name.
}

func (obj *IndexedDB) RequestDatabase(request *RequestDatabaseRequest) (response struct {
	DatabaseWithObjectStores types.IndexedDB_DatabaseWithObjectStores `json:"databaseWithObjectStores"`// Database with an array of object stores.
}, err error) {
	err = obj.conn.Send("IndexedDB.requestDatabase", request, &response)
	return
}

type RequestDataRequest struct {
	SecurityOrigin	string				`json:"securityOrigin"`// Security origin.
	DatabaseName	string				`json:"databaseName"`// Database name.
	ObjectStoreName	string				`json:"objectStoreName"`// Object store name.
	IndexName	string				`json:"indexName"`// Index name, empty string for object store data requests.
	SkipCount	int				`json:"skipCount"`// Number of records to skip.
	PageSize	int				`json:"pageSize"`// Number of records to fetch.
	KeyRange	*types.IndexedDB_KeyRange	`json:"keyRange,omitempty"`// Key range.
}

func (obj *IndexedDB) RequestData(request *RequestDataRequest) (response struct {
	ObjectStoreDataEntries	[]types.IndexedDB_DataEntry	`json:"objectStoreDataEntries"`// Array of object store data entries.
	HasMore			bool				`json:"hasMore"`// If true, there are more entries to fetch in the given range.
}, err error) {
	err = obj.conn.Send("IndexedDB.requestData", request, &response)
	return
}

type ClearObjectStoreRequest struct {
	SecurityOrigin	string	`json:"securityOrigin"`// Security origin.
	DatabaseName	string	`json:"databaseName"`// Database name.
	ObjectStoreName	string	`json:"objectStoreName"`// Object store name.
}

func (obj *IndexedDB) ClearObjectStore(request *ClearObjectStoreRequest) (err error) {
	err = obj.conn.Send("IndexedDB.clearObjectStore", request, nil)
	return
}

type DeleteDatabaseRequest struct {
	SecurityOrigin	string	`json:"securityOrigin"`// Security origin.
	DatabaseName	string	`json:"databaseName"`// Database name.
}

func (obj *IndexedDB) DeleteDatabase(request *DeleteDatabaseRequest) (err error) {
	err = obj.conn.Send("IndexedDB.deleteDatabase", request, nil)
	return
}
