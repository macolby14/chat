package main

import (
	"encoding/gob"
	"log"
	"net/http"
)

// Define our struct
type authenticationMiddleware struct {
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
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "session")
		if err != nil {
			log.Println("Could not decode existing session. Creating new session")
		}

		if _, ok := session.Values["user"]; !ok {
			// need to dereference the pointer for storage in the session
			session.Values["user"] = *amw.uf.createUser()

			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		next.ServeHTTP(w,r)
    })
}

func getUserFromRequest(r *http.Request) *User{
	session, err := store.Get(r, "session")
	if err != nil {
		log.Println("Could not decode existing session. Creating new session")
	}

	var user interface{}
	var ok bool
 	if user, ok = session.Values["user"]; !ok {
		log.Fatalf("No user value assigned to session")
	}

	var userTyped User = (user).(User)
	return &userTyped
}