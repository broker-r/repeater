package main

import (
	"fmt"
	"log/slog"
	"os"
	"repeater/internal/prettylog"
	"repeater/internal/storage"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Repeat     bool   `short:"r" long:"repeat" description:"Start repeating all words"`
	Count      int    `short:"c" long:"count" description:"Start repeating selected number of words"`
	List       bool   `short:"l" long:"list" description:"Print all words in storage"`
	AddWord    string `short:"a" long:"add" description:"Add a word to storage"`
	DeleteWord string `short:"d" long:"delete" description:"Delete a word from storage"`
}

func main() {
	flags.Parse(&opts)

	logger := setupLogger()
	logger.Debug("Debug mode enabled")
	logger.Debug("Options", "options", opts)

	storage := storage.NewStorage("./db.sqlite3", logger)

	if opts.AddWord != "" {
		if err := storage.AddWord(opts.AddWord); err != nil {
			logger.Error("Error", prettylog.PrettyError(err))
		}
	}

	if opts.List {
		words, err := storage.GetWords()
		if err != nil {
			logger.Error("Error", prettylog.PrettyError(err))
			os.Exit(1)
		}
		fmt.Printf("%-10s | %-15s | %-20s|\n", "NAME", "REPEAT COUNTER", "LAST REPEAT")
		for _, word := range words {
			fmt.Printf("%-10s | %-15d | %-20s|\n", word.Name, word.Repeat_counter, word.Last_repeat)
		}
	}

	if opts.AddWord != "" {
		storage.AddWord(opts.AddWord)
	}

	if opts.DeleteWord != "" {
		storage.RemoveWord(opts.DeleteWord)
	}
}

func setupLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	prettyHandler := prettylog.NewPrettyHandler(os.Stdout, opts)
	return slog.New(prettyHandler)
}
