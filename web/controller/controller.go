package controller

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("github.com//web/template/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {
}
