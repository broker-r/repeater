package main

import (
	"log/slog"
	"os"
	"repeater/internal/prettylog"
)

func main() {
	logger := setupLogger()
	logger.Debug("Debug mode enabled")
}

func setupLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	prettyHandler := prettylog.NewPrettyHandler(os.Stdout, opts)
	return slog.New(prettyHandler)
}
