package main

import (
	"log"

	"cloud.google.com/go/datastore"
)

// ModelInterface .
type ModelInterface interface {
	Get()
	GetAll()
	Save()
	Delete()
}

// Model .
type Model struct {
	Key  *datastore.Key `datastore:"__key__" json:"key"`
	Kind string         `datastore:"-"`
}

// Get .
func (m Model) Get(keyID int64, model interface{}) {
	k := datastore.IDKey(m.Kind, keyID, nil)
	if err := dsClient.Get(dsCtx, k, model); err != nil {
		log.Fatal(err)
	}
}

// Save .
func (m *Model) Save(model interface{}) {
	var k *datastore.Key
	if m.Key == nil {
		k = datastore.IncompleteKey(m.Kind, nil)
	} else {
		k = m.Key
	}

	if _, err := dsClient.Put(dsCtx, k, model); err != nil {
		log.Fatal(err)
	}
}

// Delete .
func (m *Model) Delete() {
	if err := dsClient.Delete(dsCtx, m.Key); err != nil {
		log.Fatal(err)
	}
}
