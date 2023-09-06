package storage

import (
	"LibraryService/internal/models"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) GetBooks() (models.Book, error) {
	_, err := s.db.Query("SELECT * FROM books")
	return models.Book{}, err
}
func (s *Storage) GetAuthors() (models.Author, error) {
	_, err := s.db.Query("SELECT * FROM authors")
	return models.Author{}, err
}
