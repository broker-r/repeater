package list_handler

import (
	"fmt"
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"

	"github.com/fatih/color"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.List {
		words, err := storage.GetWords()
		if err != nil {
			logger.Debug("Error when getting words from the database", prettylog.PrettyError(err))
			os.Exit(1)
		}

		fmt.Print(color.CyanString("Output [all words]\n"))
		PrintWords(words)
	}
}

func PrintWords(words []storage.Word) {
	fmt.Printf("%-20s | %-20s | %-20s | %-30s|\n", "NAME", "TRANSLATION", "REPEAT COUNTER", "LAST REPEAT")
	for _, word := range words {
		fmt.Printf("%-20s | %-20s | %-20d | %-30s|\n", word.Name, word.Translation, word.Repeat_counter, word.Last_repeat)
	}
	fmt.Println()
}
