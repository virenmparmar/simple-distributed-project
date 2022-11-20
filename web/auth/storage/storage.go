package storage

import (
	userModel "github.com/simple-distributed-project/web/auth/model"
)

var Users = make(map[string]userModel.User)

func ReturnUser(username string, resultChan chan userModel.User, errChan chan bool) {
	userModel.UserMux.Lock()
	user, exists := Users[username]
	userModel.UserMux.Unlock()
	resultChan <- user
	errChan <- exists
}

func SaveUser(user userModel.User, resultChan chan bool) {
	userModel.UserMux.Lock()
	Users[user.Username] = user
	userModel.UserMux.Unlock()
	resultChan <- true
}
