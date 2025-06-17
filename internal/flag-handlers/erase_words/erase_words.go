package erasewords_handler

import (
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.EraseWords {
		if err := storage.DeleteAllWords(); err != nil {
			logger.Error("Error when deleting all words from the database", prettylog.PrettyError(err))
			os.Exit(1)
		}
	}
}
