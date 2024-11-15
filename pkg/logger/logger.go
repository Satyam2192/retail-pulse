package logger

import (
    "log"
    "os"
)

type Logger struct {
    *log.Logger
}

func New() *Logger {
    return &Logger{
        Logger: log.New(os.Stdout, "", log.LstdFlags),
    }
}

func (l *Logger) Infof(format string, v ...interface{}) {
    l.Printf("INFO: "+format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
    l.Printf("ERROR: "+format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
    l.Logger.Fatal(v...)
}