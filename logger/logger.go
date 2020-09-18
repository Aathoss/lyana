package logger

import (
	"io"
	"log"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(logFile, os.Stdout)

	DebugLogger = log.New(mw, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
