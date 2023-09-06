package handles

import "LibraryService/internal/models"

type LybraryService interface {
	SearchBooks(string) (models.Book, error)
	SearchAuthors(string) (models.Author, error)
}
