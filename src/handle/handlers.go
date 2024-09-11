package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"testTask/src/service"
)

var user service.User

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	htmlfile, err := template.ParseFiles("templates/mainView.html")
	service.Check(err)
	err = htmlfile.Execute(writer, nil)
	service.Check(err)
}

func LoginHandlerPage(writer http.ResponseWriter, request *http.Request) {
	var output string
	var err error

	fmt.Println(request.Method)

	if request.Method == http.MethodPost {
		login := request.FormValue("login")
		password := request.FormValue("password")

		user, err = service.Login(login, password)
		if err != nil {
			output = "Неверый логин или пароль"
		} else {
			http.Redirect(writer, request, "/notes/", http.StatusFound)
		}
	}

	htmlfile, err := template.ParseFiles("templates/loginView.html")
	service.Check(err)
	err = htmlfile.Execute(writer, output)
	service.Check(err)
}

func MyNotesHandler(writer http.ResponseWriter, request *http.Request) {
	htmlfile, err := template.ParseFiles("templates/myNotesView.html")
	service.Check(err)

	if user.GetLogin() == "" {
		var answer []string
		nonAuth := "Вы не авторизованы"
		answer = append(answer, nonAuth)

		err = htmlfile.Execute(writer, answer)
		service.Check(err)
	} else {
		myNotes := user.GetNotes()
		err = htmlfile.Execute(writer, myNotes)
		service.Check(err)
	}
}

func AddNoteHandler(writer http.ResponseWriter, request *http.Request) {
	htmlfile, err := template.ParseFiles("templates/addView.html")
	service.Check(err)
	err = htmlfile.Execute(writer, user.GetLogin())
	service.Check(err)
}

func ConfirmNoteHandler(writer http.ResponseWriter, request *http.Request) {
	newNote := request.FormValue("newNoteInput")
	if newNote == "" {
		http.Redirect(writer, request, "/notes/add/", http.StatusFound)
	} else {
		fmt.Println(newNote)
		// Добавить фичу яндекса
		// Добавить новую заметку через u.AddNote()
		http.Redirect(writer, request, "/notes/", http.StatusFound)
	}
}

func LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	user = user.Logout()
	http.Redirect(writer, request, "/notes/login/", http.StatusFound)
}
