package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id       int
	Name     string
	Password string
}

var UserById = make(map[int]*User)
var UserByName = make(map[string][]*User)

func loginMemory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		user := &User{Id: 1, Name: r.Form.Get("username"), Password: r.Form.Get("password")}
		store(user)
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "Welcome login, Hello %s", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "Password error, please try agin!")
		}
	}
}
func store(user *User) {
	UserById[user.Id] = user
	UserByName[user.Name] = append(UserByName[user.Name], user)
}
func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(UserById[1])
	r.ParseForm()
	for _, user := range UserByName[r.Form.Get("username")] {
		fmt.Fprintf(w, " %v", user)
	}
}
func main() {
	http.HandleFunc("/login", loginMemory)
	http.HandleFunc("/info", userInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
