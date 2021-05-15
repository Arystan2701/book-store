package service

import "github.com/Arystan2701/book-store/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
