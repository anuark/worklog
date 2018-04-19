package main

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	Email    string    `datastore:"email"`
	Password string    `datastore:"password"`
	Created  time.Time `datastore:"created"`
	AuthKey  string    `datastore:"authKey"`

	Model
}

// NewUser .
func NewUser() *User {
	u := &User{
		Created: time.Now(),
	}
	u.Kind = "User"

	return u
}

// Save .
// func (e *User) Save() {

// }

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// userKey is the context key for the user.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
var userKey key = 1

// UserNewContext .
func UserNewContext(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// UserFromContext .
func UserFromContext(ctx context.Context) (User, bool) {
	v, ok := ctx.Value(userKey).(User)
	return v, ok
}

// SetPassword .
func (u *User) SetPassword(password string) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = string(pass)
}

// func (u *User) getModel() User {
// 	return u
// }
