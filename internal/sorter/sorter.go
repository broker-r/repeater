package sorter

import (
	"log/slog"
	"repeater/internal/storage"
	"time"
)

const (
	LEVELS_COUNT = 3
	DEFAULT_SIZE = 5

	FIRST_LEVEL_HOURS  = 24 * 7
	SECOND_LEVEL_HOURS = 24 * 3

	FIRST_LEVEL_TOPCOUNTER  = 8
	SECOND_LEVEL_TOPCOUNTER = 20
)

// приоритеты повторения слова:
// 1 - высокий   -  слова из этого приоритета выпадают с шансом 60%
// 2 - средний   -  слова из этого приоритета выпадают с шансом 30%
// 3 - низский   -  слова из этого приоритета выпадают с шансом 10%

// распределение по приоритетам:
// 1: до 8 повторений ИЛИ прошло уже 7 дней или больше
// 2: от 8 до 20 повторений ИЛИ прошло от 3 до 7 дней
// 3: все остальные слова

// когда достаем слово:
// рандомно выбираем приоритет -> если приоритет пуст, переходим к следующему приоритету (+1) (если 1 пустой -> идем на 2) (2 -> 3) (3 -> 1)
// достаем рандомное слово из приоритета

func Sort(words []storage.Word, logger *slog.Logger) (map[int][]storage.Word, error) {
	sorted_words := make(map[int][]storage.Word, LEVELS_COUNT)

	for k := range sorted_words {
		sorted_words[k] = make([]storage.Word, DEFAULT_SIZE)
	}

	for _, word := range words {
		last_repeat, err := time.Parse("2006-01-02 15:04:05", word.Last_repeat)
		if err != nil {
			return nil, err
		}
		hours := time.Since(last_repeat).Hours()

		if word.Repeat_counter < FIRST_LEVEL_TOPCOUNTER || hours >= FIRST_LEVEL_HOURS {
			sorted_words[1] = append(sorted_words[1], word)
		} else if word.Repeat_counter < SECOND_LEVEL_TOPCOUNTER || hours >= SECOND_LEVEL_HOURS {
			sorted_words[2] = append(sorted_words[2], word)
		} else {
			sorted_words[3] = append(sorted_words[3], word)
		}
	}

	logger.Debug("SORTED MAP", "sorted_words", sorted_words)

	return sorted_words, nil
}
