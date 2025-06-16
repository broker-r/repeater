package main

import (
	"log/slog"
	"os"
	"repeater/internal/prettylog"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Repeat     bool   `short:"r" long:"repeat" description:"Start repeating all words"`
	Count      int    `short:"c" long:"count" description:"Start repeating selected number of words"`
	List       bool   `short:"l" long:"list" description:"Print all words in storage"`
	AddWord    string `short:"a" long:"add" description:"Add a word to storage"`
	RemoveWord string `short:"r" long:"remove" description:"Remove a word from storage"`
}

func main() {
	flags.Parse(&opts)

	logger := setupLogger()
	logger.Debug("Debug mode enabled")
	logger.Debug("Options", "options", opts)
}

func setupLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	prettyHandler := prettylog.NewPrettyHandler(os.Stdout, opts)
	return slog.New(prettyHandler)
}
