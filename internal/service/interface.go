package service

import "LibraryService/internal/models"

type LibraryStorage interface {
	GetBooks() (models.Book, error)
	GetAuthors() (models.Author, error)
}
