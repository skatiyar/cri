package types

type Database_DatabaseId string
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
type Database_Error struct {
	// Error message.
	Message string `json:"message"`
	// Error code.
	Code int `json:"code"`
}
