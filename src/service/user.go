package service

import "errors"

type User struct {
	userId    int
	userLogin string
}

func (u User) GetLogin() string {
	return u.userLogin
}

func (u User) GetNotes() (notesList []string) {
	database := OpenDB("database.db")
	defer database.Close()
	rows, err := database.Query(`SELECT NoteText
								FROM notes
								WHERE UserId = ?`, u.GetId())
	Check(err)
	defer rows.Close()
	var note string
	for rows.Next() {
		err = rows.Scan(&note)
		Check(err)
		notesList = append(notesList, note)
	}
	return notesList
}

func (u User) GetId() int {
	return u.userId
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

	return user, err
}

func (u User) Logout() User {
	var UnLoggedUser User
	UnLoggedUser.userId = 0
	UnLoggedUser.userLogin = ""
	return (UnLoggedUser)
}

func (u *User) AddNote(newNote string) {
	database := OpenDB("database.db")
	defer database.Close()
	query := `INSERT INTO notes (UserId, NoteText) 
			  VALUES(?, ?)`
	_, err := database.Exec(query, u.GetId(), newNote)
	Check(err)
}
