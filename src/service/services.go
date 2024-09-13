package service

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Response struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	Sug  []string `json:"s"`
}

func GetFix(newNote string) (result []Response, err error) {
	URL := "https://speller.yandex.net/services/spellservice.json/checkText"
	par := url.Values{}
	par.Add("text", newNote)

	query, err := http.Get(URL + "?" + par.Encode())
	if err != nil {
		return result, err
	}

	defer query.Body.Close()
	jsonResult, err := io.ReadAll(query.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(jsonResult, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

func FixNote(newNote string, fixedWords []Response) (fixedText string) {
	fixedText = newNote
	for _, sug := range fixedWords {
		if len(sug.Sug) > 0 {
			fixedText = strings.Replace(fixedText, sug.Word, sug.Sug[0], -1)
		}
	}

	return fixedText
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func OpenDB(path string) *sql.DB {
	database, err := sql.Open("sqlite3", path)
	Check(err)

	return (database)
}

func ClearLog() {
	file, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_TRUNC, 0644)
	Check(err)
	defer file.Close()
}
