package types

type IndexedDB_DatabaseWithObjectStores struct {
	// Database name.
	Name string `json:"name"`
	// Database version.
	Version int `json:"version"`
	// Object stores in this database.
	ObjectStores []IndexedDB_ObjectStore `json:"objectStores"`
}
type IndexedDB_ObjectStore struct {
	// Object store name.
	Name string `json:"name"`
	// Object store key path.
	KeyPath IndexedDB_KeyPath `json:"keyPath"`
	// If true, object store has auto increment flag set.
	AutoIncrement bool `json:"autoIncrement"`
	// Indexes in this object store.
	Indexes []IndexedDB_ObjectStoreIndex `json:"indexes"`
}
type IndexedDB_ObjectStoreIndex struct {
	// Index name.
	Name string `json:"name"`
	// Index key path.
	KeyPath IndexedDB_KeyPath `json:"keyPath"`
	// If true, index is unique.
	Unique bool `json:"unique"`
	// If true, index allows multiple entries for a key.
	MultiEntry bool `json:"multiEntry"`
}
type IndexedDB_Key struct {
	// Key type.
	Type string `json:"type"`
	// Number value.
	Number *float32 `json:"number,omitempty"`
	// String value.
	String *string `json:"string,omitempty"`
	// Date value.
	Date *float32 `json:"date,omitempty"`
	// Array value.
	Array []*IndexedDB_Key `json:"array,omitempty"`
}
type IndexedDB_KeyRange struct {
	// Lower bound.
	Lower *IndexedDB_Key `json:"lower,omitempty"`
	// Upper bound.
	Upper *IndexedDB_Key `json:"upper,omitempty"`
	// If true lower bound is open.
	LowerOpen bool `json:"lowerOpen"`
	// If true upper bound is open.
	UpperOpen bool `json:"upperOpen"`
}
type IndexedDB_DataEntry struct {
	// Key object.
	Key Runtime_RemoteObject `json:"key"`
	// Primary key object.
	PrimaryKey Runtime_RemoteObject `json:"primaryKey"`
	// Value object.
	Value Runtime_RemoteObject `json:"value"`
}
type IndexedDB_KeyPath struct {
	// Key path type.
	Type string `json:"type"`
	// String value.
	String *string `json:"string,omitempty"`
	// Array value.
	Array []string `json:"array,omitempty"`
}
