package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	opts := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	Logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))

	Logger.Info("Initialized logger")
}
