package service

import (
	"LibraryService/internal/models"
)

type LibraryService struct {
}

func (s *LibraryService) SearchBooks(string) (models.Book, error) {
	return models.Book{}, nil
}

func (s *LibraryService) SearchAuthors(string) (models.Author, error) {
	return models.Author{}, nil
}
