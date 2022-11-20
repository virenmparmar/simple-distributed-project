package model

import (
	"container/list"
	"sync"
)

type User struct {
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Token     string     `json:"token"`
	Followers *list.List `json:"followers"`
}

var UserMux = &sync.Mutex{}
