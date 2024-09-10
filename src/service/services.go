package service

import (
	"log"
	"os"
)

func Check(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

func OpenLog() {
	file, err := os.OpenFile("testTask/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Check(err)
	log.SetOutput(file)
}
