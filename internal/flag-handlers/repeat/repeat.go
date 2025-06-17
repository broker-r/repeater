package repeat_handler

import (
	"fmt"
	"log/slog"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/storage"
	"strings"

	"github.com/fatih/color"
)

func Handle(opts *options.Opts, logger *slog.Logger, storage *storage.Storage) {
	if opts.Repeat {
		words, err := storage.GetWords()
		if err != nil {
			logger.Error("Error when getting words from the database", prettylog.PrettyError(err))
			os.Exit(1)
		}

		if len(words) == 0 {
			fmt.Println("You dont have any words in your storage!")
			return
		}

		fmt.Println(color.CyanString("[REPEAT MODE]"))

		var translation string

		for i := 0; i < len(words); {
			fmt.Printf("Translate word: %s: ", color.HiYellowString(words[i].Name))
			fmt.Scanf("%s", &translation)
			fmt.Printf("Correct answer: %s\n", color.GreenString(words[i].Translation))
			if strings.ToLower(translation) != strings.ToLower(words[i].Translation) {
				fmt.Println(color.RedString("You answered incorrectly."))
			} else {
				fmt.Println(color.GreenString("You answered correctly!"))
				storage.ChangeCounter(words[i].Name, 1)
				storage.UpdateTime(words[i].Name)
			}
			fmt.Println()

			if i == len(words)-1 {
				i = 0
			} else {
				i++
			}

			translation = ""
		}

	}
}
