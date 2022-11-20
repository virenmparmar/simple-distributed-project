package web

import (
	"fmt"
	"net/http"

	auth_controller "github.com/simple-distributed-project/web/auth/controller"
)

func StartService() {
	http.HandleFunc("/", auth_controller.Home)
	http.HandleFunc("/login", auth_controller.Login)
	http.HandleFunc("/register", auth_controller.Register)

	err := http.ListenAndServe(":9000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}
