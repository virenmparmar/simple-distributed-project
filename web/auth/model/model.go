package model

import (
	"sync"
)

type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Token     string   `json:"token"`
	Followers []string `json:"followers"`
}

type Tweet struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Tweet    string `json:"tweet"`
}

var UserMux = &sync.Mutex{}

var TweetMux = &sync.Mutex{}
