package main

import (
	"time"
)

type Message struct{
	Text []byte
	Time int64
	User *User
}

func newMessage(text []byte, user *User) *Message{
	return &Message{
		Text: text,
		Time: time.Now().Unix(),
		User: user,
	}
}