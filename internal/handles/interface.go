package handles

import "LibraryService/internal/models"

type LybraryService interface {
	SearchBooks(query string) (models.Book, error)
	SearchAuthors(query string) (models.Author, error)
}
