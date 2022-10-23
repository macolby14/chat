package main

import (
	"strconv"

	"github.com/gorilla/sessions"
)

type User struct {
	Name string
	Sessions []*sessions.Session
}

func newUser(name string) *User{
	return &User{
		Name: name,
		Sessions: make([]*sessions.Session, 0),
	}
}


type UserFactory struct {
	users 		[]*User
	i 			int
}

func newUserFactory() *UserFactory {
	uf := &UserFactory{}
	nouns := []string{
		"giraffe",
		"cheetah",
		"elephant",
		"lion",
		"baboon",
	}
	adjectives := []string{
		"purple",
		"red",
		"fast",
		"smart",
		"clever",
	}

	ct := 0

	for _, noun := range nouns {
		for _, adj := range adjectives {
			uf.users = append(uf.users, newUser(noun+"-"+adj+"-"+strconv.Itoa(ct)))
			ct++
		}
	}

	return uf
}

func (uf *UserFactory) createUser() *User {
	ind := uf.i
	uf.i++
	return uf.users[ind]
}
