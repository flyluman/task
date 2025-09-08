package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func GetLogger() *slog.Logger {
	if logger == nil {
		Init()
	}
	return logger
}
