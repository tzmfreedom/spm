package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
)

type Logger struct {
}

func NewLogger(outStream *os.File) *Logger {
	log.SetOutput(outStream)
	return &Logger{}
}

func (l *Logger)Info(args ...interface{}) {
	log.Info(args...)
}

func (l *Logger)Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func (l *Logger)Warning(args ...interface{}) {
	log.Warning(args...)
}

func (l *Logger)Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

func (l *Logger)Error(args ...interface{}) {
	log.Error(args...)
}

func (l *Logger)Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}