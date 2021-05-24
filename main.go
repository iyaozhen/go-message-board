package main

import (
	"log"
	"mydemo/handler"
	"mydemo/handler/user"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/user/login", userhandler.Login)
	http.HandleFunc("/user/signup", userhandler.SingUp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
