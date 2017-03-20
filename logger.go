package main

import (
	"io"

	"github.com/Sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Reset(outStream io.Writer, errStream io.Writer)
}

type SpmLogger struct {
	OutLogger *logrus.Logger
	ErrLogger *logrus.Logger
}

func NewSpmLogger(outStream io.Writer, errStream io.Writer) *SpmLogger {
	outLogger := logrus.New()
	outLogger.Out = outStream
	errLogger := logrus.New()
	errLogger.Out = errStream
	return &SpmLogger{
		OutLogger: outLogger,
		ErrLogger: errLogger,
	}
}

func (l *SpmLogger) Info(args ...interface{}) {
	l.OutLogger.Info(args...)
}

func (l *SpmLogger) Infof(format string, args ...interface{}) {
	l.OutLogger.Infof(format, args...)
}

func (l *SpmLogger) Warning(args ...interface{}) {
	l.OutLogger.Warning(args...)
}

func (l *SpmLogger) Warningf(format string, args ...interface{}) {
	l.OutLogger.Warningf(format, args...)
}

func (l *SpmLogger) Error(args ...interface{}) {
	l.OutLogger.Error(args...)
}

func (l *SpmLogger) Errorf(format string, args ...interface{}) {
	l.OutLogger.Errorf(format, args...)
}

func (l *SpmLogger) Reset(outStream io.Writer, errStream io.Writer) {
	l.OutLogger.Out = outStream
	l.ErrLogger.Out = errStream
}

type NullWriter struct {
	Logger
}

func (w *NullWriter) Write(b []byte) (int, error) {
	return 0, nil
}
