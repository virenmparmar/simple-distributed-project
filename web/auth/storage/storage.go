package storage

import (
	userModel "github.com/simple-distributed-project/web/auth/model"
)

var Users = make(map[string]userModel.User)
var Tweets = make(map[string][]userModel.Tweet)

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

func SaveTweet(user userModel.User, tweet userModel.Tweet, resultChan chan bool) {
	userModel.TweetMux.Lock()
	Tweets[user.Username] = append(Tweets[tweet.Username], tweet)
	userModel.TweetMux.Unlock()
	resultChan <- true
}

func ReturnTweets(username string, resultChan chan []userModel.Tweet) {
	userModel.TweetMux.Lock()
	tweets := Tweets[username]
	userModel.TweetMux.Unlock()
	resultChan <- tweets
}
