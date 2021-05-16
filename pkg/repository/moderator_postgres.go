package repository

import (
	"database/sql"
	"fmt"
	book_store "github.com/Arystan2701/book-store"
	"github.com/jmoiron/sqlx"
)

type ModeratorPostgre struct {
	db *sqlx.DB
}

func NewModeratorPostgre(db *sqlx.DB) *ModeratorPostgre {
	return &ModeratorPostgre{db: db}
}

func (m *ModeratorPostgre) GetModerators() ([]book_store.Moderator, error) {
	var moderators []book_store.Moderator
	query := fmt.Sprintf("SELECT id, email, role FROM  %s", moderatorTable)
	err := m.db.Select(&moderators, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return moderators, nil
		}
		return nil, err
	}
	return moderators, err
}

func (m *ModeratorPostgre) CreateModerator(request book_store.CreateModeratorInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash,role) values ($1, $2, $3) RETURNING id", moderatorTable)
	row := m.db.QueryRow(query, request.Email, request.Password, request.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
