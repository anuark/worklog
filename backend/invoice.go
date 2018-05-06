package main

import "time"

// Invoice .
type Invoice struct {
	Name    string    `datastore:"name" json:"name"`
	Path    string    `datastore:"path" json:"path"`
	Created time.Time `datastore:"created" json:"created"`

	Model
}

// NewInvoice .
func NewInvoice() *Invoice {
	i := &Invoice{
		Created: time.Now(),
	}

	i.Kind = "Invoice"
	return i
}
