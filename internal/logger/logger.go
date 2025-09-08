package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init() {
	if Log == nil {
		Log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
}
