package types

type IndexedDB_DatabaseWithObjectStores struct {
	Name         string                  `json:"name"`
	Version      int                     `json:"version"`
	ObjectStores []IndexedDB_ObjectStore `json:"objectStores"`
}
type IndexedDB_ObjectStore struct {
	Name          string                       `json:"name"`
	KeyPath       IndexedDB_KeyPath            `json:"keyPath"`
	AutoIncrement bool                         `json:"autoIncrement"`
	Indexes       []IndexedDB_ObjectStoreIndex `json:"indexes"`
}
type IndexedDB_ObjectStoreIndex struct {
	Name       string            `json:"name"`
	KeyPath    IndexedDB_KeyPath `json:"keyPath"`
	Unique     bool              `json:"unique"`
	MultiEntry bool              `json:"multiEntry"`
}
type IndexedDB_Key struct {
	Type   string           `json:"type"`
	Number *float32         `json:"number,omitempty"`
	String *string          `json:"string,omitempty"`
	Date   *float32         `json:"date,omitempty"`
	Array  []*IndexedDB_Key `json:"array,omitempty"`
}
type IndexedDB_KeyRange struct {
	Lower     *IndexedDB_Key `json:"lower,omitempty"`
	Upper     *IndexedDB_Key `json:"upper,omitempty"`
	LowerOpen bool           `json:"lowerOpen"`
	UpperOpen bool           `json:"upperOpen"`
}
type IndexedDB_DataEntry struct {
	Key        Runtime_RemoteObject `json:"key"`
	PrimaryKey Runtime_RemoteObject `json:"primaryKey"`
	Value      Runtime_RemoteObject `json:"value"`
}
type IndexedDB_KeyPath struct {
	Type   string   `json:"type"`
	String *string  `json:"string,omitempty"`
	Array  []string `json:"array,omitempty"`
}
