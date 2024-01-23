package main

import (
	"fmt"
	"insta/pkg/handler"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.LoginHandler)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
