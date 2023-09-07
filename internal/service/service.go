package service

import (
	"LibraryService/internal/models"
)

type LibraryService struct {
	storage LibraryStorage
}

// NewService creates a new instance of LibraryService.
//
// It takes a LibraryStorage as a parameter and returns a pointer to the newly created LibraryService.
func NewService(storage LibraryStorage) *LibraryService {

	return &LibraryService{
		storage: storage,
	}
}

// SearchBooks searches for books by a given query.
//
// It takes a query string as a parameter and returns a single book and an error.
func (s *LibraryService) SearchBooks(query string) (models.Book, error) {
	res := models.Book{
		Name: query,
	}
	//Получение данных
	authors, err := s.storage.GetAuthors(query)
	if err != nil {
		return res, err
	}
	//Обработка данных
	res.Author = authors
	return res, nil
}

// SearchAuthors searches for authors based on a query string.
//
// query: The query string to search for.
// Returns: An instance of the models.Author struct and an error, if any.
func (s *LibraryService) SearchAuthors(query string) (models.Author, error) {
	res := models.Author{
		Name: query,
	}
	//Получение данных
	books, err := s.storage.GetBooks(query)
	if err != nil {
		return res, err
	}
	//Обработка данных
	res.Books = books
	return res, nil
}
