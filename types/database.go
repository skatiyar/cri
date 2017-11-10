package types

type Database_DatabaseId string
type Database_Database struct {
	Id	Database_DatabaseId	`json:"id"`// Database ID.
	Domain	string			`json:"domain"`// Database domain.
	Name	string			`json:"name"`// Database name.
	Version	string			`json:"version"`// Database version.
}
type Database_Error struct {
	Message	string	`json:"message"`// Error message.
	Code	int	`json:"code"`// Error code.
}
