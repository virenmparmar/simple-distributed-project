package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	userrepo "github.com/simple-distributed-project/web/auth/repository"
	bcrypt "golang.org/x/crypto/bcrypt"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

func RegisterService(r *http.Request) error {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	user, err := userrepo.FindUser(strings.TrimSpace(username))
	if user.Username == username {
		return fmt.Errorf("User already exists")
	}
	if err.Error() != "User does not exist" {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return fmt.Errorf("Error While Hashing Password, Try Again")
	}
	password = string(hash)

	err = userrepo.CreateUser(strings.TrimSpace(username), password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func LoginService(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := userrepo.FindUser(strings.TrimSpace(username))
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("Incorrect Password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return fmt.Errorf("Error While Signing Token")
	}
	user.Token = tokenString
	err = userrepo.UpdateUser(user)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) error {
	username, err := GetUserNameFromToken(r)
	if err != nil {
		return err
	}
	user, err := userrepo.FindUser(strings.TrimSpace(username))
	if err != nil {
		return err
	}
	user.Token = ""
	err = userrepo.UpdateUser(user)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	return nil
}

func GetUsersToFollow(r *http.Request) ([]string, error) {
	username, err := GetUserNameFromToken(r)
	users, err := userrepo.GetUsersToFollow(strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersToUnfollow(r *http.Request) ([]string, error) {
	username, err := GetUserNameFromToken(r)
	users, err := userrepo.GetUsersToUnfollow(strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}
	return users, nil
}

func StartFollowing(r *http.Request) error {
	r.ParseForm()
	username, err := GetUserNameFromToken(r)
	if err != nil {
		return err
	}
	following := r.FormValue("startfollow")
	fmt.Println("Following: ", following)
	fmt.Println("Username: ", username)
	err = userrepo.FollowUser(strings.TrimSpace(username), strings.TrimSpace(following))
	if err != nil {
		return err
	}
	return nil
}

func StopFollowing(r *http.Request) error {
	r.ParseForm()
	username, err := GetUserNameFromToken(r)
	if err != nil {
		return err
	}
	following := r.FormValue("stopfollow")
	fmt.Println("Following: ", following)
	fmt.Println("Username: ", username)
	err = userrepo.UnfollowUser(strings.TrimSpace(username), strings.TrimSpace(following))
	if err != nil {
		return err
	}
	return nil
}

func Tweet(r *http.Request) error {
	r.ParseForm()
	username, err := GetUserNameFromToken(r)
	if err != nil {
		return err
	}
	user, err := userrepo.FindUser(strings.TrimSpace(username))
	if err != nil {
		return err
	}
	tweet := r.FormValue("tweet")
	err = userrepo.SaveTweet(user, strings.TrimSpace(tweet))
	if err != nil {
		return err
	}
	return nil
}

func GetUserNameFromToken(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}

	tokenString := c.Value
	token, tokenerr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || tokenerr != nil {
		return "", fmt.Errorf("Invalid Token")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username, nil
}
