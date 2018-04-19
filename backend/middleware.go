package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/iterator"

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
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var id string
		if user, ok := UserFromContext(r.Context()); ok {
			id = strconv.Itoa(int(user.ID))
		} else {
			id = "-"
		}

		defer func() { log.Printf("%s [%s] (%s) %s", time.Since(start), id, r.Method, r.URL.Path) }()

		next.ServeHTTP(w, r)
	})
}

var skipUrls = []string{
	"/",
	"/auth",
}

// Auth .
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip urls that don't need authentication
		for _, v := range skipUrls {
			if r.URL.Path == v {
				next.ServeHTTP(w, r)
				return
			}
		}

		authKey := r.Header["Authorization"]
		if authKey == nil {
			http.Error(w, "No Bearer Token.", http.StatusForbidden)
			return
		}

		var key []string
		if key = strings.Split(authKey[0], " "); len(key) < 2 {
			http.Error(w, "Wrong Bearer token format.", http.StatusBadRequest)
			return
		}

		q := datastore.NewQuery("User").Filter("authKey =", key[1])
		var user User
		for t := dsClient.Run(r.Context(), q); ; {
			k, err := t.Next(&user)
			if err == iterator.Done {
				break
			}
			if err != nil {
				panic(err)
			}
			user.ID = k.ID
			break
		}

		if user.Key == nil {
			http.Error(w, "Wrong authentication token.", http.StatusForbidden)
			return
		}

		ctx := UserNewContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Cors .
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")

		w.Header().Add("Access-Control-Expose-Headers", "X-Total-Count")
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Methods", "POST, DELETE, PUT, PATCH")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// JSONContentType All responses will be json Content-Type
func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
