package main

import (
	"net/http"
	"testTask/src/handle"
	"testTask/src/service"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/notes/", handle.MainHandler)
	http.HandleFunc("/notes/login/", handle.LoginHandlerPage)
	http.HandleFunc("/notes/logout/", handle.LogoutHandler)
	http.HandleFunc("/notes/add/", handle.AddNoteHandler)
	http.HandleFunc("/notes/confirm/", handle.ConfirmNoteHandler)
	http.HandleFunc("/notes/mynotes/", handle.MyNotesHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	service.Check(err)
}
