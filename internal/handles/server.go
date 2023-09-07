package handles

import (
	"LibraryService/internal/api"
	"context"
)

type LibraryRPC struct {
	s LybraryService
	api.UnimplementedLibraryServer
}

// NewLibraryRPC creates a new instance of LibraryRPC.
//
// It takes a LybraryService as a parameter.
// It returns a pointer to LibraryRPC.
func NewLibraryRPC(s LybraryService) *LibraryRPC {
	return &LibraryRPC{s: s}
}

// SearchBooks searches for books based on a given query.
//
// ctx: The context.Context object for the request.
// in: The api.SearchBooksRequest object containing the query information.
// Returns a *api.SearchBooksResponse and an error.
func (r *LibraryRPC) SearchBooks(ctx context.Context, in *api.SearchBooksRequest) (*api.SearchBooksResponse, error) {
	//Выполнение запроса
	book, err := r.s.SearchBooks(in.Query)
	//Формирование ответа
	res := &api.SearchBooksResponse{
		Book:   book.Name,
		Author: book.Author,
	}
	return res, err
}

// SearchAuthors searches for authors based on the given query.
//
// ctx: The context of the request.
// in: The request object containing the query.
//
// *api.SearchAuthorsResponse: The response containing the author and their books.
// error: An error if any occurred during the search.
func (r *LibraryRPC) SearchAuthors(ctx context.Context, in *api.SearchAuthorsRequest) (*api.SearchAuthorsResponse, error) {
	//Выполнение запроса
	author, err := r.s.SearchAuthors(in.Query)
	//Формирование ответа
	res := &api.SearchAuthorsResponse{
		Author: author.Name,
		Books:  author.Books,
	}
	return res, err
}
