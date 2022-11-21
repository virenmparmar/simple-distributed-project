package model

import (
	"sync"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Followers []User `json:"followers"`
}

var UserMux = &sync.Mutex{}
