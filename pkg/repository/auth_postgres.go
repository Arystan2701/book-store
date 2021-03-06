package repository

import (
	"fmt"
	book_store "github.com/Arystan2701/book-store"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user book_store.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.UserName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(email, password string) (book_store.User, error) {
	var user book_store.User
	query := fmt.Sprintf("SELECT id FROM  %s WHERE email=$1 AND password_hash=$2", userTable)
	err := a.db.Get(&user, query, email, password)
	return user, err
}

func (a *AuthPostgres) GetModerator(email, password string) (book_store.Moderator, error) {
	var moderator book_store.Moderator
	query := fmt.Sprintf("SELECT id,email, role FROM  %s WHERE email=$1 AND password_hash=$2", moderatorTable)
	err := a.db.Get(&moderator, query, email, password)
	return moderator, err
}
