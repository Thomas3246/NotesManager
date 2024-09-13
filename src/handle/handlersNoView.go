package handle

import (
	"net/http"
	"testTask/src/service"
)

func ConfirmNoteHandler(writer http.ResponseWriter, request *http.Request) {
	newNote := request.FormValue("newNoteInput")
	if newNote == "" {
		http.Redirect(writer, request, "/notes/add/", http.StatusFound)
	} else {
		fixedWords, err := service.GetFix(newNote)
		service.Check(err)
		fixedText := service.FixNote(newNote, fixedWords)
		user.FixedNoteLog(newNote, fixedText)

		user.AddNote(fixedText)
		user.AddedNoteLog(fixedText)

		http.Redirect(writer, request, "/notes/", http.StatusFound)
	}
}

func LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	if user.GetLogin() != "" {
		user.LogLogout()
	}
	user = user.Logout()
	http.Redirect(writer, request, "/notes/login/", http.StatusFound)
}
