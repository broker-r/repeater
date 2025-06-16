package deleteWord_handler

import (
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.DeleteWord != "" {
		if err := storage.DeleteWord(opts.DeleteWord); err != nil {
			logger.Error("Error when deleting a word from the database", prettylog.PrettyError(err))
			os.Exit(1)
		}
	}
}
