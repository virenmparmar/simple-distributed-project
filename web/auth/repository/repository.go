package repository

import (
	"fmt"

	userModel "github.com/simple-distributed-project/web/auth/model"
	userstorage "github.com/simple-distributed-project/web/auth/storage"
)

func FindUser(username string) (user userModel.User, err error) {
	resultChan := make(chan userModel.User)
	errChan := make(chan bool)
	go userstorage.ReturnUser(username, resultChan, errChan)
	user = <-resultChan
	exists := <-errChan
	if !exists {
		err = fmt.Errorf("User does not exist")
	}
	return user, err
}

func CreateUser(username string, password string) error {
	resultChan := make(chan bool)
	user := userModel.User{Username: username, Password: password}
	go userstorage.SaveUser(user, resultChan)
	saved := <-resultChan
	if !saved {
		return fmt.Errorf("User not saved")
	}
	return nil
}
