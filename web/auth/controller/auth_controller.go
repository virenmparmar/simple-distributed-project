package controller

import (
	"fmt"
	"html/template"
	"net/http"

	userService "github.com/simple-distributed-project/web/auth/service"
	util "github.com/simple-distributed-project/web/auth/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/register.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == "GET" {
		t.Execute(w, m)
		return
	} else {
		fmt.Println("Register")
		err := userService.RegisterService(r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("no error")
	if r.Method == "GET" {
		t.Execute(w, m)
		return
	} else {
		fmt.Println("Login")
		err := userService.LoginService(w, r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	fmt.Println(r.Method)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Profile")
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/profile.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, m)
}

func Follow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Follow")
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/follow.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == "GET" {
		users, err := userService.GetUsersToFollow(r)
		data := util.ConvertUserstoHTML(users)
		fmt.Print(data)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m["Users"] = data
			t.Execute(w, m)
			return
		}
	} else {
		fmt.Println("Follow")
		err := userService.FollowService(r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m["Success"] = "Followed"
			http.Redirect(w, r, "/follow", http.StatusFound)
			return
		}
	}
}
