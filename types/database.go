package types

type Database_DatabaseId string
type Database_Database struct {
	Id      Database_DatabaseId `json:"id"`
	Domain  string              `json:"domain"`
	Name    string              `json:"name"`
	Version string              `json:"version"`
}
type Database_Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
