package service

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/Arystan2701/book-store/pkg/repository"
)

type Authorization interface {
	CreateUser(user book_store.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Moderator interface {
	GetModerators() ([]book_store.Moderator, error)
	CreateModerator(request book_store.CreateModeratorInput) (int, error)
}

type Service struct {
	Authorization
	Moderator
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Moderator:     NewModeratorService(repos.Moderator),
	}
}
