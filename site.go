package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
)

// Index Action for listing tasks
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "Hello world")
}

type authStruct struct {
	Email    string
	Password string
}

// Authenticate authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var t authStruct
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	q := datastore.NewQuery("User").Filter("email =", t.Email)
	users := make([]User, 0)
	if _, err := dsClient.GetAll(dsCtx, q, &users); err != nil {
		log.Fatal(err)
	}

	if len(users) > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(t.Password)); err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Username or password is invalid.", http.StatusBadRequest)
		return
	}

	h := md4.New()
	io.WriteString(h, time.Now().String())
	hash := h.Sum(nil)
	users[0].AuthKey = fmt.Sprintf("%x", hash)
	users[0].Save()

	fmt.Fprint(w, "{\"token\": \""+users[0].AuthKey+"\"}")
}
