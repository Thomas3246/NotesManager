package service

import (
	"os"
	"time"
)

func (u User) LogLogin() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)
	defer file.Close()
	stringToLog := time.Now().Format(time.RFC3339) + " | User " + u.GetLogin() + " has logged in\n\n"
	file.WriteString(stringToLog)
}

func (u User) LogLogout() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)
	defer file.Close()
	stringToLog := time.Now().Format(time.RFC3339) + " | User " + u.GetLogin() + " has logged out\n\n"
	file.WriteString(stringToLog)
}

func (u User) FixedNoteLog(nonFixed string, fixed string) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)
	defer file.Close()
	stringToLog := time.Now().Format(time.RFC3339) + " | User " + u.GetLogin() + " | sent request with string << " + nonFixed + " >> \nFixed to << " + fixed + " >>\n\n"
	file.WriteString(stringToLog)
}

func (u User) AddedNoteLog(newNote string) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)
	defer file.Close()
	stringToLog := time.Now().Format(time.RFC3339) + " | User " + u.GetLogin() + " has added new note << " + newNote + " >>\n\n"
	file.WriteString(stringToLog)
}
