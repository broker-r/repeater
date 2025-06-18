package count_handler

import (
	"fmt"
	"log/slog"
	"os"
	list_handler "repeater/internal/flag-handlers/list"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"

	"github.com/fatih/color"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.Count > 0 {
		words, err := storage.GetWords()
		if err != nil {
			logger.Debug("Error when getting words from the database", prettylog.PrettyError(err))
			os.Exit(1)
		}

		if len(words) == 0 {
			return
		}

		maxCount := len(words)
		var count int

		if opts.Count > maxCount {
			count = maxCount
		} else {
			count = opts.Count
		}

		words = words[:count]
		output := fmt.Sprintf("Output [by quantity: %d]\n", count)
		fmt.Print(color.CyanString(output))
		list_handler.PrintWords(words)
	}
}
