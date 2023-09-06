package handles

import (
	"LibraryService/internal/api"
	"LibraryService/internal/service"
	"context"
)

type LibraryRPC struct {
	s *service.LibraryService
}

func NewLibraryRPC(s *service.LibraryService) *LibraryRPC {
	return &LibraryRPC{s: s}
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
