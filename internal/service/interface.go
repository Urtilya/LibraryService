package service

type LibraryStorage interface {
	GetBooks(name string) ([]string, error)
	GetAuthors(name string) ([]string, error)
}
