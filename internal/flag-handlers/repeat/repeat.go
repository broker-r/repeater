package repeat_handler

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"repeater/internal/options"
	"repeater/internal/prettylog"
	"repeater/internal/sorter"
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

		sorted_words, err := sorter.Sort(words, logger)
		if err != nil {
			logger.Error("Error when sorting words", prettylog.PrettyError(err))
			os.Exit(1)
		}

		r_level := rand.IntN(4)
		for notEmpty := false; notEmpty; {
			if len(sorted_words[r_level]) == 0 {
				if r_level == sorter.LEVELS_COUNT {
					r_level = 1
				} else {
					r_level++
				}
			} else {
				notEmpty = true
			}
		}
		r_wordIndex := rand.IntN(len(sorted_words[r_level]))
		r_word := sorted_words[r_level][r_wordIndex]
		_ = r_word

		/////// DO NOT PUSH IN MAIN.

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
			if strings.EqualFold(translation, words[i].Translation) {
				fmt.Println(color.RedString("You answered incorrectly."))
			} else {
				fmt.Println(color.GreenString("You answered correctly!"))
				storage.ChangeCounter(words[i].Name, 1)
				storage.UpdateTime(words[i].Name)
			}
			fmt.Println()

			if i == len(words)-1 {
				// пересчет мапы
				i = 0
			} else {
				i++
			}

			translation = ""
		}

	}
}
