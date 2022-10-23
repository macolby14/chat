package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Define our struct
type authenticationMiddleware struct {
	store *sessions.CookieStore
	uf *UserFactory
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Init() {

	/* gob setup
	  	avoid securecookie: error - caused by: securecookie: error - caused by: gob: type not registered for interface: main.User
		from gorilla/sessions dependency
	*/
	gob.Register(User{})

	amw.uf = newUserFactory()

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

		if _, ok := session.Values["user"]; !ok {
			session.Values["user"] = amw.uf.createUser()

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		log.Printf("User:%v\n",session.Values["user"])

		next.ServeHTTP(w,r)
    })
}