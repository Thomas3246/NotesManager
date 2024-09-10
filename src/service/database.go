package service

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserId    int
	UserLogin string
	Notes     []string
}

func OpenDB(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)
	Check(err)
	return (database)
}

func Login(login string, password string) User {
	database := OpenDB("database.db")

	rows, err := database.Query(`SELECT UserId, UserLogin 
								 FROM users
								 WHERE UserLogin = ? AND UserPassword = ?`, login, password)
	Check(err)

	var user User
	rows.Next()
	err = rows.Scan(&user.UserId, &user.UserLogin)
	Check(err)

	rows, err = database.Query(`SELECT NoteText
								FROM notes
								WHERE UserId = ?`, user.UserId)
	Check(err)
	var note string
	for rows.Next() {
		err = rows.Scan(&note)
		Check(err)
		user.Notes = append(user.Notes, note)
	}

	return (user)
}
