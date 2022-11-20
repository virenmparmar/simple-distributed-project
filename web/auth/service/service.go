package service

import (
	"fmt"
	"net/http"

	userrepo "github.com/simple-distributed-project/web/auth/repository"
)

func RegisterService(r *http.Request) error {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)
	_, err := userrepo.FindUser(username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = userrepo.CreateUser(username, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
