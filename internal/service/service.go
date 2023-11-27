package service

import "AdExchangerGo/internal/model"

type Service struct {
	Authorization
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
}
