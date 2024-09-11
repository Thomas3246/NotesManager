package service

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	userId    int
	userLogin string
	notes     []string
}

func (u User) Logout() User {
	var UnLoggedUser User
	UnLoggedUser.userId = 0
	UnLoggedUser.userLogin = ""
	UnLoggedUser.notes = nil
	return (UnLoggedUser)
}

func (u User) GetLogin() string {
	return u.userLogin
}

func (u User) GetNotes() []string {
	return u.notes
}

func (u User) GetId() int {
	return u.userId
}

func (u *User) AddNote(newNote string) {
	database := OpenDB("database.db")
	defer database.Close()
	query := `INSERT INTO notes (UserId, NoteText) 
			  VALUES(?, ?)`
	_, err := database.Exec(query, u.GetId(), newNote)
	Check(err)
	defer database.Close()
	fmt.Println("added")
	u.notes = append(u.notes, newNote)
}

func OpenDB(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)
	Check(err)

	return (database)
}

func Login(login string, password string) (user User, err error) {
	database := OpenDB("database.db")
	defer database.Close()

	rows, err := database.Query(`SELECT UserId, UserLogin 
								 FROM users
								 WHERE UserLogin = ? AND UserPassword = ?`, login, password)
	Check(err)

	if !rows.Next() {
		return user, errors.New("неверный логин или пароль")
	}

	err = rows.Scan(&user.userId, &user.userLogin)
	Check(err)
	defer rows.Close()

	rows, err = database.Query(`SELECT NoteText
								FROM notes
								WHERE UserId = ?`, user.userId)
	Check(err)
	defer rows.Close()
	var note string
	for rows.Next() {
		err = rows.Scan(&note)
		Check(err)
		user.notes = append(user.notes, note)
	}

	return user, err
}
