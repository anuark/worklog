package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InvoiceList .
func InvoiceList(w http.ResponseWriter, r *http.Request) {

}

// InvoiceCreate .
func InvoiceCreate(w http.ResponseWriter, r *http.Request) {
	invoice := NewInvoice()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&invoice)
	if err != nil {
		panic(err)
	}

	user, _ := UserFromContext(r.Context())
	invoice.AncestorKey = user.Key

	invoice.Save(invoice)
	json, err := json.Marshal(invoice)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(json))
}

// InvoiceDelete .
func InvoiceDelete(w http.ResponseWriter, r *http.Request) {

}
