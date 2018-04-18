package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	user := NewUser()
	for t := dsClient.Run(r.Context(), q); ; {
		t.Next(user)
		break
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(t.Password)); err != nil {
		http.Error(w, "Username or password is invalid.", http.StatusBadRequest)
		return
	}

	h := md4.New()
	io.WriteString(h, time.Now().String())
	hash := h.Sum(nil)
	user.AuthKey = fmt.Sprintf("%x", hash)
	user.Save(user)

	fmt.Fprint(w, "{\"token\": \""+user.AuthKey+"\"}")
}
