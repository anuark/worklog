package main

import (
	"fmt"

	"cloud.google.com/go/datastore"
)

// ModelInterface .
type ModelInterface interface {
	Save()
	Delete()
	Load()
}

// Model .
type Model struct {
	Key         *datastore.Key `datastore:"__key__" json:"-"`
	AncestorKey *datastore.Key `datastore:"-" json:"-"`
	ID          int64          `datastore:"-" json:"id"`
	Kind        string         `datastore:"-" json:"-"`
}

// Save .
func (m *Model) Save(model interface{}) {
	var k *datastore.Key
	if m.Key == nil {
		k = datastore.IncompleteKey(m.Kind, m.AncestorKey)
	} else {
		k = m.Key
	}

	k, err := dsClient.Put(dsCtx, k, model)
	if err != nil {
		fmt.Println(err)
	}

	m.ID = k.ID
}

// Delete .
func (m *Model) Delete() {
	if err := dsClient.Delete(dsCtx, m.Key); err != nil {
		fmt.Println(err)
	}
}

// Load .
func (m *Model) Load(ps []datastore.Property) error {
	// Load I as usual.
	return datastore.LoadStruct(m, ps)
}

// SaveProps .
func (m *Model) SaveProps(ps []datastore.Property) ([]datastore.Property, error) {
	// Load I as usual.
	return datastore.SaveStruct(m)
}
