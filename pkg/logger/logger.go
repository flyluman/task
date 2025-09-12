package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type stdLogger struct {
	l *slog.Logger
}

var logger *stdLogger

func NewStdLogger() Logger {
	if logger == nil {
		logger = &stdLogger{l: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
	}
	return logger
}

func (s *stdLogger) Info(msg string, kv ...interface{})  { s.l.Info(msg, kv...) }
func (s *stdLogger) Error(msg string, kv ...interface{}) { s.l.Error(msg, kv...) }
