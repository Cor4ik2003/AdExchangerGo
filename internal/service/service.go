package service

import (
	"AdExchangerGo/internal/model"
	"AdExchangerGo/internal/repository"
)

type Service struct {
	Authorization
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

func NewService(repos *repository.Repository) *Service {
	return nil
}
