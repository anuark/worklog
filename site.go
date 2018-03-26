package main

import (
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

// TaskCreate action for new task.
func TaskCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	task := NewTask()
	val := r.PostFormValue("desc")
	if len(val) > 1 {
		task.Description = r.PostFormValue("desc")
		task.Save()
	}
}

// UserCreate .
func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	user := NewUser()
	user.Email = r.PostFormValue("email")
	p := r.PostFormValue("password")
	password, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	user.Password = string(password)
	user.Save()
}

// Authenticate authenticate
func UserAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	pass := r.PostFormValue("password")
	email := r.PostFormValue("email")
	_ = email

	q := datastore.NewQuery("User").Filter("email =", email)
	users := make([]User, 0)
	if _, err := dsClient.GetAll(dsCtx, q, &users); err != nil {
		log.Fatal(err)
	}

	if len(users) > 0 {
		if err := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(pass)); err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "User not registered", http.StatusBadRequest)
	}

	h := md4.New()
	io.WriteString(h, time.Now().String())
	hash := h.Sum(nil)
	users[0].AuthKey = fmt.Sprintf("%x", hash)
	users[0].Save()
	fmt.Fprint(w, users[0].AuthKey)
}
