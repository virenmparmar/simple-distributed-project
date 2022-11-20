package controller

import (
	"fmt"
	"html/template"
	"net/http"

	userService "github.com/simple-distributed-project/web/auth/service"
)

func Register(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("web/template/register.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method != "POST" {
		t.Execute(w, nil)
		return
	} else {
		fmt.Println("Register")
		err := userService.RegisterService(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
}
