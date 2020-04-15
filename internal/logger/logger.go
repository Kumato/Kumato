package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	blue   = "\033[34m"
	green  = "\033[97;32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	reset  = "\033[0m"
)

var debug bool

func EnableDebug() {
	debug = true
}

func init() {
	if m := os.Getenv("KUMATO_DEV_MODE"); m == "1" || m == "on" {
		debug = true
	}
}

func Default(v ...interface{}) {
	if debug {
		logger := log.New(os.Stdout, "[DEFA] ", log.Ldate|log.Ltime|log.Lshortfile)
		s := fmt.Sprintln(v...)
		logger.Output(2, "\n"+s[:len(s)-1]+reset)
	}
}

func Info(v ...interface{}) {
	if debug {
		logger := log.New(os.Stdout, blue+"[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
		s := fmt.Sprintln(v...)
		logger.Output(2, "\n"+s[:len(s)-1]+reset)
	}
}

func Warn(v ...interface{}) {
	logger := log.New(os.Stdout, yellow+"[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	s := fmt.Sprintln(v...)
	logger.Output(2, "\n"+s[:len(s)-1]+reset)
}

func Fatal(v ...interface{}) {
	logger := log.New(os.Stdout, red+"[FATA] ", log.Ldate|log.Ltime|log.Lshortfile)
	s := fmt.Sprintln(v...)
	logger.Output(2, "\n"+s[:len(s)-1]+reset)
}
