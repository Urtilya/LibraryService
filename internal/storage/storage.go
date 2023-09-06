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
	//select  `Authors`.name from `Books` join `Relevant` on `Relevant`.id_book =`Books`.id join `Authors` on `Authors`.id = `Relevant`.id_author WHERE `Books`.name = 'День черных звезд'
	_, err := s.db.Query("SELECT * FROM books")
	return models.Book{}, err
}
func (s *Storage) GetAuthors() (models.Author, error) {
	//select  `Books`.name from `Authors` join `Relevant` on `Relevant`.id_author = `Authors`.id join `Books` on `Books`.id = `Relevant`.id_book WHERE `Authors`.name = 'Иар Эльтеррус'
	_, err := s.db.Query("SELECT * FROM authors")
	return models.Author{}, err
}
