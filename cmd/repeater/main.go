package main

import (
	"fmt"
	"log/slog"
	"os"
	addWord_handler "repeater/internal/flag-handlers/add_word"
	count_handler "repeater/internal/flag-handlers/count"
	deleteWord_handler "repeater/internal/flag-handlers/delete_word"
	list_handler "repeater/internal/flag-handlers/list"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"

	"github.com/jessevdk/go-flags"
)

func main() {
	fmt.Println() // for beautiful output

	var opts options.Opts
	flags.Parse(&opts)

	logger := setupLogger()
	storage := storage.NewStorage("./db.sqlite3", logger)

	addWord_handler.Handle(&opts, logger, storage)

	deleteWord_handler.Handle(&opts, logger, storage)

	count_handler.Handle(&opts, logger, storage)

	list_handler.Handle(&opts, logger, storage)
}

func setupLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	prettyHandler := prettylog.NewPrettyHandler(os.Stdout, opts)
	return slog.New(prettyHandler)
}
