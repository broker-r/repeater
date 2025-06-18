package storage

import (
	"database/sql"
	"log"
	"log/slog"
	"repeater/internal/prettylog"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db     *sql.DB
	logger *slog.Logger
}

type Word struct {
	Name           string
	Translation    string
	Repeat_counter int
	Last_repeat    string
}

func NewStorage(dbpath string, logger *slog.Logger) *Storage {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal("Failed to load database", err)
	}

	storage := Storage{
		db:     db,
		logger: logger,
	}

	if err := storage.CreateTables(); err != nil {
		log.Fatal("Failed to create tables", err)
	}

	return &storage
}

func (s *Storage) CreateTables() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS word (
    name text primary_key UNIQUE,
	translation text NOT NULL,
    repeat_counter integer DEFAULT 0,
    last_repeat timestamp DEFAULT CURRENT_TIMESTAMP
)`)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddWord(name string, translation string) error {
	stmt, err := s.db.Prepare(`INSERT INTO word (name, translation) VALUES (?, ?);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, translation)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteWord(name string) error {
	stmt, err := s.db.Prepare(`DELETE FROM word WHERE name = ?;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	s.logger.Debug("Word deleted: " + name)

	return nil
}

func (s *Storage) DeleteAllWords() error {
	_, err := s.db.Exec(`DELETE FROM word;`)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetWords() ([]Word, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM word;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []Word

	for rows.Next() {
		var word Word
		var last_repeat string
		err := rows.Scan(&word.Name, &word.Translation, &word.Repeat_counter, &last_repeat)
		if err != nil {
			return nil, err
		}
		repeat_time, err := time.Parse(time.RFC3339, last_repeat)
		word.Last_repeat = repeat_time.Format("2006-01-02 15:04:05")
		if err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func (s *Storage) ChangeCounter(name string, value int) error {
	stmt, err := s.db.Prepare("UPDATE word SET repeat_counter = repeat_counter + ? WHERE name = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(value, name)
	if err != nil {
		s.logger.Debug("Error when changing repeat_counter", prettylog.PrettyError(err))
		return err
	}

	return nil
}

func (s *Storage) UpdateTime(name string) error {
	stmt, err := s.db.Prepare("UPDATE word SET last_repeat = ? WHERE name = ?;")
	if err != nil {
		return err
	}

	current_time := time.Now().Format("2006-01-02 15:04:05")

	_, err = stmt.Exec(current_time, name)
	if err != nil {
		s.logger.Debug("Error when updating last_repeat", prettylog.PrettyError(err))
		return err
	}

	return nil
}
