package logger

import (
	"log"
	"os"
)

func LoggerProvider() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	return logger
}
