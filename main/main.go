package main

import (
	"bile-go-server/core/handle"
	"bile-go-server/core/login"
	"bile-go-server/core/register"
	"net/http"
)

func main() {
	http.HandleFunc("/register", register.RegisterHandle)
	http.HandleFunc("/login", login.LoginHandle)
	http.ListenAndServe("127.0.0.1:8080", new(handle.AuthMidServer))
 
}
