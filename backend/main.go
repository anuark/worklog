package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"cloud.google.com/go/datastore"
)

var dsClient *datastore.Client
var dsCtx context.Context

func main() {
	inputDate := flag.String("since", time.Now().Format("2006-01-02"), "Generate pdf data since the input date.")
	generateFlag := flag.Bool("generate", false, "To generate invoice pdf.")
	flag.Parse()

	dsCtx = context.Background()

	var err error
	dsClient, err = datastore.NewClient(dsCtx, "worklog-191500")
	if err != nil {
		log.Fatal(err)
	}

	if *generateFlag {
		_ = inputDate
		dateRange, err := time.Parse("2006-01-02", *inputDate)
		if err != nil {
			log.Fatal(err)
		}
		_ = dateRange

		GeneratePdf(dateRange)
	} else {
		r := mux.NewRouter()
		r.Use(Cors, JSONContentType, Auth, Log)

		// Site
		r.HandleFunc("/", Index).Methods("GET")
		r.HandleFunc("/auth", Authenticate).Methods("POST", "OPTIONS")

		// User
		s := r.PathPrefix("/users").Subrouter()
		s.HandleFunc("", UserList).Methods("GET", "OPTIONS")
		s.HandleFunc("/", UserCreate).Methods("POST", "OPTIONS")

		// Tasks
		s = r.PathPrefix("/tasks").Subrouter()
		s.HandleFunc("", TaskList).Methods("GET", "OPTIONS")
		s.HandleFunc("", TaskCreate).Methods("POST", "OPTIONS")
		s.HandleFunc("/{taskId}", TaskView).Methods("GET", "OPTIONS")
		s.HandleFunc("/{taskId}", TaskUpdate).Methods("PUT", "OPTIONS")
		s.HandleFunc("/{taskId}", TaskDelete).Methods("DELETE", "OPTIONS")

		// Invoice
		s = r.PathPrefix("/invoices").Subrouter()
		s.HandleFunc("", InvoiceList).Methods("GET", "OPTIONS")
		s.HandleFunc("", InvoiceCreate).Methods("POST", "OPTIONS")
		s.HandleFunc("/{invoiceId}", InvoiceDelete).Methods("DELETE", "OPTIONS")

		log.Fatal(http.ListenAndServe("localhost:8000", r))
	}
}
