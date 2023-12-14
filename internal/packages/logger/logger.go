package logger

import (
	"log"
	"os"
)

type logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

var loggerInstance *logger

func Init() {
	loggerInstance = &logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func Info(message string) {
	loggerInstance.infoLogger.Println(message)
}

func Error(err error) {
	if (err != nil) {
		loggerInstance.errorLogger.Println(err)
	}
}
