package service

import (
	"LibraryService/config"
	"LibraryService/internal/db"
	"LibraryService/internal/models"
	"LibraryService/internal/storage"
	"reflect"
	"testing"
)

func TestLibraryService_SearchBooks(t *testing.T) {
	type args struct {
		query string
	}
	cfg := config.DB{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		Db:       "test",
		Timeout:  10,
	}
	db, _ := db.NewSqlDB(&cfg)
	storage := storage.NewStorage(db)
	tests := []struct {
		name    string
		s       *LibraryService
		args    args
		want    models.Book
		wantErr bool
	}{
		{
			name: "SearchBooks",
			s:    NewService(storage),
			args: args{
				query: "test",
			},
			want: models.Book{
				Name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SearchBooks(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("LibraryService.SearchBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LibraryService.SearchBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLibraryService_SearchAuthors(t *testing.T) {
	type args struct {
		query string
	}
	cfg := config.DB{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		Db:       "test",
		Timeout:  10,
	}
	db, _ := db.NewSqlDB(&cfg)
	storage := storage.NewStorage(db)
	tests := []struct {
		name    string
		s       *LibraryService
		args    args
		want    models.Author
		wantErr bool
	}{
		{
			name: "SearchAuthors",
			s:    NewService(storage),
			args: args{
				query: "test",
			},
			want: models.Author{
				Name: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SearchAuthors(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("LibraryService.SearchAuthors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LibraryService.SearchAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}
