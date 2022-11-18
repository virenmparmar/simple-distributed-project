package web

import (
	"fmt"
	"net/http"

	"github.com/simple-distributed-project/web/model/controller/controller"
)

func StartService() {
	http.HandleFunc("/", *controller.Home)
	http.HandleFunc("/login", *controller.login)
	http.HandleFunc("/register", *controller.register)

	err := http.ListenAndServe(":9000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}
