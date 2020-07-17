package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login....")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("127.0.0.1:8800", nil)
	if err != nil {
		fmt.Println("htpp listen failed")
	}
}
