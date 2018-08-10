/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package types

// Unique identifier of Database object.
type Database_DatabaseId string

// Database object.
type Database_Database struct {
	// Database ID.
	Id Database_DatabaseId `json:"id"`
	// Database domain.
	Domain string `json:"domain"`
	// Database name.
	Name string `json:"name"`
	// Database version.
	Version string `json:"version"`
}

// Database error.
type Database_Error struct {
	// Error message.
	Message string `json:"message"`
	// Error code.
	Code int `json:"code"`
}
