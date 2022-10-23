package main

import "github.com/gorilla/sessions"

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
	u := &UserFactory{}
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

	for _, noun := range nouns {
		for _, adj := range adjectives {
			u.users = append(u.users, newUser(noun+adj))
		}
	}

	return u
}

func (uf *UserFactory) createUser() *User {
	ind := uf.i
	uf.i++
	return uf.users[ind]
}
