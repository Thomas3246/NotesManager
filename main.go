package main

import (
	"fmt"
	"net/http"
	"testTask/src/handle"
	"testTask/src/service"
)

func main() {

	// Добавить куки

	loggedUser := service.Login("FirstUser", "Password1")
	fmt.Print(loggedUser)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/notes/", handle.MainHandler)
	http.HandleFunc("/notes/login", handle.LoginHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	service.Check(err)
}
