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

		fmt.Println(color.CyanString("[REPEAT MODE]"))
		var translation string
		for {
			words, err := storage.GetWords()
			if err != nil {
				logger.Error("Error when getting words from the database", prettylog.PrettyError(err))
				os.Exit(1)
			}

			if len(words) == 0 {
				fmt.Println("You dont have any words in your storage!")
				return
			}

			sorted_words, err := sorter.Sort(words, logger)
			if err != nil {
				logger.Error("Error when sorting words", prettylog.PrettyError(err))
				os.Exit(1)
			}

			r_number := rand.IntN(100)
			var r_level int
			if r_number < 60 {
				r_level = 1
			} else if r_number < 90 {
				r_level = 2
			} else {
				r_level = 3
			}

			for empty := true; empty; {
				if len(sorted_words[r_level]) == 0 {
					if r_level == sorter.LEVELS_COUNT {
						r_level = 1
					} else {
						r_level++
					}
				} else {
					empty = false
				}
			}

			r_wordIndex := rand.IntN(len(sorted_words[r_level]))
			r_word := sorted_words[r_level][r_wordIndex] // random word (random_level + random_index)

			fmt.Printf("Translate word: %s: ", color.HiYellowString(r_word.Name))
			fmt.Scanf("%s", &translation)
			fmt.Printf("Correct answer: %s\n", color.HiGreenString(r_word.Translation))
			if strings.EqualFold(translation, r_word.Translation) {
				fmt.Println(color.GreenString("You answered correctly!"))
				storage.ChangeCounter(r_word.Name, 1)
				storage.UpdateTime(r_word.Name)
			} else {
				fmt.Println(color.RedString("You answered incorrectly."))
			}
			fmt.Println()

			translation = "" // Reset translation
		}
	}
}
