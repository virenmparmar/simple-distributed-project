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

func UpdateUser(user userModel.User) error {
	resultChan := make(chan bool)
	go userstorage.SaveUser(user, resultChan)
	saved := <-resultChan
	if !saved {
		return fmt.Errorf("User not saved")
	}
	return nil
}

func GetUsersToFollow(username string) ([]string, error) {
	users := make([]string, 0)
	for _, user := range userstorage.Users {
		if user.Username != username {
			users = append(users, user.Username)
		}
	}
	return users, nil
}

func GetUsersToUnfollow(username string) ([]string, error) {
	users := make([]string, 0)
	currentUser, err := FindUser(username)
	if err != nil {
		return users, err
	}
	for _, user := range currentUser.Followers {
		users = append(users, user.Username)
	}
	return users, nil
}

func FollowUser(username string, userToFollow string) error {
	currentUser, err := FindUser(username)
	if err != nil {
		return err
	}
	user, err := FindUser(userToFollow)
	if err != nil {
		return fmt.Errorf("User to follow does not exist")
	}
	currentUser.Followers = append(currentUser.Followers, user)
	err = UpdateUser(currentUser)
	if err != nil {
		return err
	}
	return nil
}
