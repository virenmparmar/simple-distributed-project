package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/simple-distributed-project/web/auth/service"
	userService "github.com/simple-distributed-project/web/auth/service"
)

var BASE_DIR = os.Getenv("BASE_DIR")

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

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
	err := userService.Logout(w, r)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
	return
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	fmt.Println(r.Method)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	username, err := userService.GetUserNameFromToken(r)
	log.Println(username, err)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	fmt.Println("Profile")
	m := make(map[string]string)
	m["username"] = username
	t, err := template.ParseFiles("web/template/profile.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, m)
}

func Follow(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/follow.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == "GET" {
		fmt.Println("Follow in get method")
		users, err := userService.GetUsersToFollow(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m := make(map[string]interface{})
			m["Users"] = users
			t.Execute(w, m)
			return
		}
	} else {
		fmt.Print("Follow in post method")
		err := userService.StartFollowing(r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			http.Redirect(w, r, "/follow", http.StatusFound)
			return
		} else {
			m["Success"] = "Started Following"
			http.Redirect(w, r, "/follow", http.StatusFound)
			return
		}
	}
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unfollow")
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/unfollow.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.Method == "GET" {
		users, err := userService.GetUsersToUnfollow(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m := make(map[string]interface{})
			m["Users"] = users
			t.Execute(w, m)
			return
		}
	} else {
		fmt.Println("Unfollow in post method")
		err := userService.StopFollowing(r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m["Success"] = "Started Following"
			http.Redirect(w, r, "/unfollow", http.StatusFound)
			return
		}
	}
}

func Tweet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tweet")
	m := make(map[string]string)
	t, err := template.ParseFiles("web/template/profile.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	users, err := userService.GetUserNameFromToken(r)
	log.Println(users, err)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if r.Method == "GET" {
		t.Execute(w, m)
		return
	} else {
		fmt.Println("Tweet in post method")
		err := userService.Tweet(r)
		if err != nil {
			fmt.Println(err)
			m["Error"] = err.Error()
			t.Execute(w, m)
			return
		} else {
			m["Success"] = "Tweeted"
			t.Execute(w, m)
			return
		}
	}
}

func redirectToLogin(w http.ResponseWriter) {
	t, _ := template.ParseFiles("login.gtpl")
	m := map[string]interface{}{}
	m["Error"] = "Please login to continue!"
	m["Success"] = nil
	log.Println("Please login to continue")
	t.Execute(w, m)
}

func Feed(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("web/template/feed.html")
	m := map[string]interface{}{}
	users, erro := userService.GetUserNameFromToken(r)
	log.Println(users, erro)
	if erro != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	loginerr, err, feed := service.FeedService(r)

	if loginerr != nil {
		redirectToLogin(w)
		return
	} else if err != "" {
		m["Error"] = err
		m["Success"] = nil
		log.Println(err)
		t.Execute(w, m)
		return
	} else {
		m["Error"] = nil
		m["Success"] = nil
		m["Feed"] = feed
		log.Println("Feed Succesfull")
		t.Execute(w, m)
		return
	}
}
