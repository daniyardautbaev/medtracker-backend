package logger

import (
	"log"
)

func Info(message string) {
	log.Printf("INFO: %s", message)
}
func Error(message string) {
	log.Printf("ERROR: %s", message)
}
