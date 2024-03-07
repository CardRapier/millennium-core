package logger

import (
	"fmt"
	"log"
	"os"
)
type Logger struct {
	Info *log.Logger
	Error *log.Logger
}

func (l *Logger) Init(serviceName string) {
    l.Info = log.New(os.Stdout, fmt.Sprint("(",serviceName ,") INFO: "), log.Ldate|log.Ltime)
    l.Error = log.New(os.Stdout, fmt.Sprint("(", serviceName, ") ERROR: "), log.Ldate|log.Ltime)
}