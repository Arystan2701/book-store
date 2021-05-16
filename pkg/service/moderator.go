package service

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/Arystan2701/book-store/pkg/repository"
)

type ModeratorService struct {
	repo repository.Moderator
}

func NewModeratorService(repo repository.Moderator) *ModeratorService {
	return &ModeratorService{repo: repo}
}

func (m *ModeratorService) GetModerators() ([]book_store.Moderator, error) {
	return m.repo.GetModerators()
}
func (m *ModeratorService) CreateModerator(request book_store.CreateModeratorInput) (int, error) {
	request.Password = GeneratePasswordHash(request.Password)
	return m.repo.CreateModerator(request)
}
