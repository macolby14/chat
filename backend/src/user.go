package main

type UserFactory struct {
	users 		[]string
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
			u.users = append(u.users, noun+adj)
		}
	}

	return u
}

func (uf *UserFactory) getUserName() string {
	ind := uf.i
	uf.i++
	return uf.users[ind]
}