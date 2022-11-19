package web

import (
	"fmt"
	"net/http"

	"github.com/simple-distributed-project/web/controller"
)

func StartService() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)

	err := http.ListenAndServe(":9000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}
