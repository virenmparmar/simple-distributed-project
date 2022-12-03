package web

import (
	"fmt"
	"net/http"

	auth_controller "github.com/simple-distributed-project/web/auth/controller"
	userrepo "github.com/simple-distributed-project/web/auth/repository"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func StartService() {
	http.HandleFunc("/", auth_controller.Home)
	http.HandleFunc("/login", auth_controller.Login)
	http.HandleFunc("/register", auth_controller.Register)
	http.HandleFunc("/profile", auth_controller.Profile)
	http.HandleFunc("/follow", auth_controller.Follow)
	http.HandleFunc("/unfollow", auth_controller.Unfollow)
	http.HandleFunc("/tweet", auth_controller.Tweet)
	http.HandleFunc("/feed", auth_controller.Feed)
	http.HandleFunc("/logout", auth_controller.Logout)
	password := "aa"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		fmt.Println("Error While Hashing Password, Try Again")
	}
	password = string(hash)

	err = userrepo.CreateUser("virenmparmar@gmail.com", password)
	err = userrepo.CreateUser("vmp2018@nyu.edu", password)
	err = http.ListenAndServe(":8000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}
