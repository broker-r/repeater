package addWord_handler

import (
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.AddWord != "" {
		if err := storage.AddWord(opts.AddWord); err != nil {
			logger.Error("Error when adding a word to the database", prettylog.PrettyError(err))
			os.Exit(1)
		}
	}
}
