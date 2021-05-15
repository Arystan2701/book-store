package repository

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user book_store.User) (int, error)
	GetUser(email, password string) (book_store.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
