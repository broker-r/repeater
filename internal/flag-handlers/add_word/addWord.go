package addWord_handler

import (
	"fmt"
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"

	gtranslate "github.com/gilang-as/google-translate"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) error {
	if opts.AddWord != "" {
		var translation string // final translation

		value := gtranslate.Translate{
			Text: opts.AddWord,
			From: "en",
			To:   "ru",
		}

		translated, err := gtranslate.Translator(value)
		if err != nil {
			logger.Error("Error when translation text", prettylog.PrettyError(err))
			return err
		}

		fmt.Printf("Enter translation for '%s' (Default = '%s'): ", opts.AddWord, translated.Text)
		fmt.Scanf("%s", &translation)

		if translation == "" || translation == " " {
			translation = translated.Text
		}

		if err := storage.AddWord(opts.AddWord, translation); err != nil {
			logger.Error("Error when adding a word to the database", prettylog.PrettyError(err))
			os.Exit(1)
		}

		fmt.Printf("Word added successfully: '%s' ('%s')\n", opts.AddWord, translation)
	}

	return nil
}
