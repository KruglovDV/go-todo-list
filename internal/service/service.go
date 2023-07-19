package service

import (
	"rest-api-postgres/internal/models"
	"rest-api-postgres/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func New(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
