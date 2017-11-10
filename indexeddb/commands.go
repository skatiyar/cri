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
	SecurityOrigin string `json:"securityOrigin"`
}

func (obj *IndexedDB) RequestDatabaseNames(request *RequestDatabaseNamesRequest) (response struct {
	DatabaseNames []string `json:"databaseNames"`
}, err error) {
	err = obj.conn.Send("IndexedDB.requestDatabaseNames", request, &response)
	return
}

type RequestDatabaseRequest struct {
	SecurityOrigin string `json:"securityOrigin"`
	DatabaseName   string `json:"databaseName"`
}

func (obj *IndexedDB) RequestDatabase(request *RequestDatabaseRequest) (response struct {
	DatabaseWithObjectStores types.IndexedDB_DatabaseWithObjectStores `json:"databaseWithObjectStores"`
}, err error) {
	err = obj.conn.Send("IndexedDB.requestDatabase", request, &response)
	return
}

type RequestDataRequest struct {
	SecurityOrigin  string                    `json:"securityOrigin"`
	DatabaseName    string                    `json:"databaseName"`
	ObjectStoreName string                    `json:"objectStoreName"`
	IndexName       string                    `json:"indexName"`
	SkipCount       int                       `json:"skipCount"`
	PageSize        int                       `json:"pageSize"`
	KeyRange        *types.IndexedDB_KeyRange `json:"keyRange,omitempty"`
}

func (obj *IndexedDB) RequestData(request *RequestDataRequest) (response struct {
	ObjectStoreDataEntries []types.IndexedDB_DataEntry `json:"objectStoreDataEntries"`
	HasMore                bool                        `json:"hasMore"`
}, err error) {
	err = obj.conn.Send("IndexedDB.requestData", request, &response)
	return
}

type ClearObjectStoreRequest struct {
	SecurityOrigin  string `json:"securityOrigin"`
	DatabaseName    string `json:"databaseName"`
	ObjectStoreName string `json:"objectStoreName"`
}

func (obj *IndexedDB) ClearObjectStore(request *ClearObjectStoreRequest) (err error) {
	err = obj.conn.Send("IndexedDB.clearObjectStore", request, nil)
	return
}

type DeleteDatabaseRequest struct {
	SecurityOrigin string `json:"securityOrigin"`
	DatabaseName   string `json:"databaseName"`
}

func (obj *IndexedDB) DeleteDatabase(request *DeleteDatabaseRequest) (err error) {
	err = obj.conn.Send("IndexedDB.deleteDatabase", request, nil)
	return
}
