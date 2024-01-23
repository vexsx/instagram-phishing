package main

import (
	"insta/pkg/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.LoginHandler)
	http.ListenAndServe(":8080", nil)
}
