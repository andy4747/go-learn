package core

import (
	"log"
	"os"
)

// creating custom logger
var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("extras/logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Logger in golang
func Logger() {
	first := 10
	second := 20

	if first != second {
		// default logger
		// log.Fatalf("%v and %v are not same", first, second)

		// using custom logger that logs in a file
		InfoLogger.Fatalf("%v and %v are not same", first, second)
	}

}
