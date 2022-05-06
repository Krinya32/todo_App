package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krinya32/todoApp"
)

type Authorization interface {
	CreateUser(user todoApp.User) (int, error)
	GetUser(username, password string) (todoApp.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
