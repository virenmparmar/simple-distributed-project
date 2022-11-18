package web

import (
	"fmt"
	"net/http"
)

func StartService() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	err := http.ListenAndServe(":9000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	} else {
		fmt.Println("Server starts at: localhost:8000")
	}
}
