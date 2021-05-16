package repository

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user book_store.User) (int, error)
	GetUser(email, password string) (book_store.User, error)
	GetModerator(email, password string) (book_store.Moderator, error)
}

type Moderator interface {
	GetModerators() ([]book_store.Moderator, error)
	CreateModerator(request book_store.CreateModeratorInput) (int, error)
}

type Repository struct {
	Authorization
	Moderator
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Moderator:     NewModeratorPostgre(db),
	}
}
