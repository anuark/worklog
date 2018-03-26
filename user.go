package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

// User .
type User struct {
	Key      *datastore.Key `datastore:"__key__"`
	Email    string         `datastore:"email"`
	Password string         `datastore:"password"`
	Created  time.Time      `datastore:"created"`
	AuthKey  string         `datastore:"authKey"`
}

// NewUser .
func NewUser() *User {
	return &User{
		Created: time.Now(),
	}
}

// Save .
func (e *User) Save() {
	var k *datastore.Key
	if e.Key == nil {
		k = datastore.IncompleteKey("User", nil)
	} else {
		k = e.Key
	}

	if _, err := dsClient.Put(dsCtx, k, e); err != nil {
		log.Fatal(err)
	}

	log.Printf("Saved User %#v", e)
}

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// userKey is the context key for the user.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
var userKey key = 0

func UserNewContext(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func UserFromContext(ctx context.Context) User {
	return ctx.Value(userKey).(User)
}
