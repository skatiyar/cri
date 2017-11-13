/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package database

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

type Database struct {
	conn cri.Connector
}

// New creates a Database instance
func New(conn cri.Connector) *Database {
	return &Database{conn}
}

// Enables database tracking, database events will now be delivered to the client.
func (obj *Database) Enable() (err error) {
	err = obj.conn.Send("Database.enable", nil, nil)
	return
}

// Disables database tracking, prevents database events from being sent to the client.
func (obj *Database) Disable() (err error) {
	err = obj.conn.Send("Database.disable", nil, nil)
	return
}

type GetDatabaseTableNamesRequest struct {
	DatabaseId types.Database_DatabaseId `json:"databaseId"`
}

type GetDatabaseTableNamesResponse struct {
	TableNames []string `json:"tableNames"`
}

func (obj *Database) GetDatabaseTableNames(request *GetDatabaseTableNamesRequest) (response GetDatabaseTableNamesResponse, err error) {
	err = obj.conn.Send("Database.getDatabaseTableNames", request, &response)
	return
}

type ExecuteSQLRequest struct {
	DatabaseId types.Database_DatabaseId `json:"databaseId"`
	Query      string                    `json:"query"`
}

type ExecuteSQLResponse struct {
	ColumnNames []string              `json:"columnNames,omitempty"`
	Values      []interface{}         `json:"values,omitempty"`
	SqlError    *types.Database_Error `json:"sqlError,omitempty"`
}

func (obj *Database) ExecuteSQL(request *ExecuteSQLRequest) (response ExecuteSQLResponse, err error) {
	err = obj.conn.Send("Database.executeSQL", request, &response)
	return
}

type AddDatabaseParams struct {
	Database types.Database_Database `json:"database"`
}

func (obj *Database) AddDatabase(fn func(params *AddDatabaseParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On("Database.addDatabase", closeChn)
	go func() {
		for {
			params := AddDatabaseParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				closeChn <- struct{}{}
				close(closeChn)
				break
			}
		}
	}()
}
