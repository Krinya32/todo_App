package service

import (
	"github.com/krinya32/todoApp"
	"github.com/krinya32/todoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoApp.User) (int, error)
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
