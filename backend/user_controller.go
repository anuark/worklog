package main

import (
	"fmt"
	"net/http"
)

// UserCreate .
func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	user := NewUser()
	user.Email = r.PostFormValue("email")
	p := r.PostFormValue("password")
	user.SetPassword(p)
	user.Save(user)
}

// UserList .
func UserList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User list")
}
