package internal

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func InitializeLogger() *Logger {
	return &Logger{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.LstdFlags),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.LstdFlags),
	}
}
