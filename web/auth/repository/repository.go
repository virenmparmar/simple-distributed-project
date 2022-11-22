package repository

import (
	"fmt"

	uuid "github.com/google/uuid"

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
	currentUser, err := FindUser(username)
	if err != nil {
		return users, err
	}
	for _, user := range userstorage.Users {
		if user.Username != username && !contains(currentUser.Followers, user.Username) {
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
		users = append(users, user)
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
	currentUser.Followers = append(currentUser.Followers, user.Username)
	err = UpdateUser(currentUser)
	fmt.Println("I am here")
	fmt.Println(currentUser.Followers)
	if err != nil {
		return err
	}
	return nil
}

func UnfollowUser(username string, userToUnfollow string) error {
	currentUser, err := FindUser(username)
	if err != nil {
		return err
	}
	user, err := FindUser(userToUnfollow)
	if err != nil {
		return fmt.Errorf("User to unfollow does not exist")
	}
	currentUser.Followers = remove(currentUser.Followers, user.Username)
	err = UpdateUser(currentUser)
	if err != nil {
		return err
	}
	return nil
}

func SaveTweet(user userModel.User, tweet string) error {
	resultChan := make(chan bool)
	id := uuid.New().String()
	tweetModel := userModel.Tweet{Id: id, Username: user.Username, Tweet: tweet}
	go userstorage.SaveTweet(user, tweetModel, resultChan)
	saved := <-resultChan
	if !saved {
		return fmt.Errorf("Tweet not saved")
	}
	return nil
}

func remove(s []string, r string) []string {
	fmt.Println(s)
	fmt.Println(r)
	var result []string
	for _, a := range s {
		if a != r {
			result = append(result, a)
		}
	}
	return result
}
func contains(s []string, e string) bool {
	fmt.Println("I am inside contains")
	fmt.Println(s)
	fmt.Println(e)
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetTweets(username string) []userModel.Tweet {
	// users := make([]string, 0)
	// currentUser, err := FindUser(username)
	// if err != nil {
	// 	return users, err
	// }
	// for _, user := range currentUser.Followers {
	// 	users = append(users, user)
	// }
	// return users, nil
	resultChan := make(chan []userModel.Tweet)
	go userstorage.ReturnTweets(username, resultChan)
	return <-resultChan
}
