package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
)

// Middleware .
type Middleware func(f http.HandlerFunc) http.HandlerFunc

// Chain .
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Before .
func Before() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var ctx context.Context
			f(w, r.WithContext(ctx))
		}
	}
}

// func After() Middleware {
// 	return func(f http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			f(w, r)
// 		}
// 	}
// }

// Log .
func Log() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx := r.Context()
			user := UserFromContext(ctx)
			id := user.AuthKey

			defer func() { log.Printf("%s [%s] %s", time.Since(start), id, r.URL.Path) }()

			// Call next middelware/handler in chain
			f(w, r)
		}
	}
}

// Auth .
func Auth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authKey := r.Header["Authorization"]
			if authKey == nil {
				http.Error(w, "No Bearer Token.", http.StatusForbidden)
				return
			}

			key := strings.Split(authKey[0], " ")[1]
			q := datastore.NewQuery("User").Filter("authKey =", key)
			authUser := []User{}
			if _, err := dsClient.GetAll(dsCtx, q, &authUser); err != nil {
				log.Fatal(err)
			}

			if len(authUser) == 0 {
				http.Error(w, "Wrong authentication token.", http.StatusForbidden)
				return
			}

			ctx := r.Context()
			ctx = UserNewContext(ctx, authUser[0])

			// Call next middelware/handler in chain
			f(w, r.WithContext(ctx))
		}
	}
}

// Method .
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call next middelware
			f(w, r)
		}
	}
}
