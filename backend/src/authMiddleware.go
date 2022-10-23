package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Define our struct
type authenticationMiddleware struct {
	store *sessions.CookieStore
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Init() {

	// session setup
	amw.store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := amw.store.Get(r, "session")
		if err != nil {
			log.Println("Could not decode existing session. Creating new session")
		}


		session.Values["userId"] = "test-user-id"
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w,r)
    })
}