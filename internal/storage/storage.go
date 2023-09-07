package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

// NewStorage creates a new instance of the Storage struct.
//
// It takes a pointer to a sqlx.DB object as a parameter.
// The function returns a pointer to a Storage object.
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// GetBooks retrieves a list of author names of books with the given name.
//
// It takes a string parameter 'name' representing the name of the book.
// It returns a slice of strings containing the author names and an error.
func (s *Storage) GetBooks(name string) ([]string, error) {
	//Создание запроса
	query := "select  `Authors`.name from `Books` join `Relevant` on `Relevant`.id_book =`Books`.id join `Authors` on `Authors`.id = `Relevant`.id_author WHERE `Books`.name = '%s'"
	//Запрос
	row, err := s.db.Query(fmt.Sprintf(query, name))
	//Парсинг результата
	var str string
	var res []string
	for row.Next() {
		row.Scan(&str)
		res = append(res, str)
	}
	return res, err
}

// GetAuthors retrieves a list of book names written by authors with the given name.
//
// name: the name of the author
// []string: a slice of book names
// error: an error if the query fails
func (s *Storage) GetAuthors(name string) ([]string, error) {
	//Создание запроса
	query := "select  `Books`.name from `Authors` join `Relevant` on `Relevant`.id_author = `Authors`.id join `Books` on `Books`.id = `Relevant`.id_book WHERE `Authors`.name = '%s'"
	//Запрос
	row, err := s.db.Query(fmt.Sprintf(query, name))
	//Парсинг результата
	var str string
	var res []string
	for row.Next() {
		row.Scan(&str)
		res = append(res, str)
	}
	return res, err
}
