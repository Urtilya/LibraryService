package storage

import (
	"LibraryService/config"
	"LibraryService/internal/db"
	"reflect"
	"testing"
)

func TestStorage_GetBooks(t *testing.T) {
	type args struct {
		name string
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
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "GetBooks",
			s:    NewStorage(db),
			args: args{
				name: "test",
			},
			want: []string{
				"test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBooks(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetAuthors(t *testing.T) {
	type args struct {
		name string
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
	tests := []struct {
		name    string
		s       *Storage
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "GetAuthors",
			s:    NewStorage(db),
			args: args{
				name: "test",
			},
			want: []string{
				"test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAuthors(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetAuthors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetAuthors() = %v, want %v", got, tt.want)
			}
		})
	}
}
