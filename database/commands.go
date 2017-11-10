package database

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Database struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Database {
	return &Database{conn}
}
func (obj *Database) Enable() (err error) {
	err = obj.conn.Send("Database.enable", nil, nil)
	return
}
func (obj *Database) Disable() (err error) {
	err = obj.conn.Send("Database.disable", nil, nil)
	return
}

type GetDatabaseTableNamesRequest struct {
	DatabaseId types.Database_DatabaseId `json:"databaseId"`
}

func (obj *Database) GetDatabaseTableNames(request *GetDatabaseTableNamesRequest) (response struct {
	TableNames []string `json:"tableNames"`
}, err error) {
	err = obj.conn.Send("Database.getDatabaseTableNames", request, &response)
	return
}

type ExecuteSQLRequest struct {
	DatabaseId types.Database_DatabaseId `json:"databaseId"`
	Query      string                    `json:"query"`
}

func (obj *Database) ExecuteSQL(request *ExecuteSQLRequest) (response struct {
	ColumnNames []string              `json:"columnNames,omitempty"`
	Values      []interface{}         `json:"values,omitempty"`
	SqlError    *types.Database_Error `json:"sqlError,omitempty"`
}, err error) {
	err = obj.conn.Send("Database.executeSQL", request, &response)
	return
}
