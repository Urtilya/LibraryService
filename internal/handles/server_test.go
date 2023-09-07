package handles

import (
	"LibraryService/config"
	"LibraryService/internal/api"
	"LibraryService/internal/db"
	"LibraryService/internal/service"
	"LibraryService/internal/storage"
	"context"
	"reflect"
	"testing"
)

func TestLibraryRPC_SearchBooks(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *api.SearchBooksRequest
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
	service := service.NewService(storage)
	tests := []struct {
		name    string
		r       *LibraryRPC
		args    args
		want    *api.SearchBooksResponse
		wantErr bool
	}{
		{
			name: "SearchBooks",
			r:    NewLibraryRPC(service),
			args: args{
				ctx: context.Background(),
				in: &api.SearchBooksRequest{
					Query: "test",
				},
			},
			want: &api.SearchBooksResponse{
				Book:   "test",
				Author: []string{"test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.SearchBooks(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("LibraryRPC.SearchBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LibraryRPC.SearchBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLibraryRPC_SearchAuthors(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *api.SearchAuthorsRequest
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
	service := service.NewService(storage)
	tests := []struct {
		name    string
		r       *LibraryRPC
		args    args
		want    *api.SearchAuthorsResponse
		wantErr bool
	}{
		{
			name: "SearchAuthors",
			r:    NewLibraryRPC(service),
			args: args{
				ctx: context.Background(),
				in: &api.SearchAuthorsRequest{
					Query: "test",
				},
			},
			want: &api.SearchAuthorsResponse{
				Author: "test",
				Books:  []string{"test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.SearchAuthors(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("LibraryRPC.SearchAuthors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LibraryRPC.SearchAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}
