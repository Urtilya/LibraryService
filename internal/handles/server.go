package handles

import (
	"LibraryService/internal/api"
	"context"
)

type LibraryRPC struct {
	s LybraryService
}

func NewLibraryRPC(s *LybraryService) *LibraryRPC {
	return &LibraryRPC{s: *s}
}

func (r *LibraryRPC) SearchBooks(ctx context.Context, in *api.SearchBooksRequest) (*api.SearchBooksResponse, error) {

	book, err := r.s.SearchBooks(in.Query)
	res := &api.SearchBooksResponse{
		Book:   book.Name,
		Author: book.Author,
	}
	return res, err
}

func (r *LibraryRPC) SearchAuthors(ctx context.Context, in *api.SearchAuthorsRequest) (*api.SearchAuthorsResponse, error) {
	author, err := r.s.SearchAuthors(in.Query)
	res := &api.SearchAuthorsResponse{
		Author: author.Name,
		Books:  author.Books,
	}
	return res, err
}
