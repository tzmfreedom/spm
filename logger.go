package main

import (
	"github.com/Sirupsen/logrus"
	"io"
)

type Logger struct {
	OutLogger *logrus.Logger
	ErrLogger *logrus.Logger
}

func NewLogger(outStream io.Writer, errStream io.Writer) *Logger {
	outLogger := logrus.New()
	outLogger.Out = outStream
	errLogger := logrus.New()
	errLogger.Out = errStream
	return &Logger{
		OutLogger: outLogger,
		ErrLogger: errLogger,
	}
}

func (l *Logger) Info(args ...interface{}) {
	l.OutLogger.Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.OutLogger.Infof(format, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.OutLogger.Warning(args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.OutLogger.Warningf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.OutLogger.Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.OutLogger.Errorf(format, args...)
}
