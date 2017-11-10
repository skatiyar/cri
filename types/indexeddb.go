package types

type IndexedDB_DatabaseWithObjectStores struct {
	Name		string			`json:"name"`// Database name.
	Version		int			`json:"version"`// Database version.
	ObjectStores	[]IndexedDB_ObjectStore	`json:"objectStores"`// Object stores in this database.
}
type IndexedDB_ObjectStore struct {
	Name		string				`json:"name"`// Object store name.
	KeyPath		IndexedDB_KeyPath		`json:"keyPath"`// Object store key path.
	AutoIncrement	bool				`json:"autoIncrement"`// If true, object store has auto increment flag set.
	Indexes		[]IndexedDB_ObjectStoreIndex	`json:"indexes"`// Indexes in this object store.
}
type IndexedDB_ObjectStoreIndex struct {
	Name		string			`json:"name"`// Index name.
	KeyPath		IndexedDB_KeyPath	`json:"keyPath"`// Index key path.
	Unique		bool			`json:"unique"`// If true, index is unique.
	MultiEntry	bool			`json:"multiEntry"`// If true, index allows multiple entries for a key.
}
type IndexedDB_Key struct {
	Type	string			`json:"type"`// Key type.
	Number	*float32		`json:"number,omitempty"`// Number value.
	String	*string			`json:"string,omitempty"`// String value.
	Date	*float32		`json:"date,omitempty"`// Date value.
	Array	[]*IndexedDB_Key	`json:"array,omitempty"`// Array value.
}
type IndexedDB_KeyRange struct {
	Lower		*IndexedDB_Key	`json:"lower,omitempty"`// Lower bound.
	Upper		*IndexedDB_Key	`json:"upper,omitempty"`// Upper bound.
	LowerOpen	bool		`json:"lowerOpen"`// If true lower bound is open.
	UpperOpen	bool		`json:"upperOpen"`// If true upper bound is open.
}
type IndexedDB_DataEntry struct {
	Key		Runtime_RemoteObject	`json:"key"`// Key object.
	PrimaryKey	Runtime_RemoteObject	`json:"primaryKey"`// Primary key object.
	Value		Runtime_RemoteObject	`json:"value"`// Value object.
}
type IndexedDB_KeyPath struct {
	Type	string		`json:"type"`// Key path type.
	String	*string		`json:"string,omitempty"`// String value.
	Array	[]string	`json:"array,omitempty"`// Array value.
}
