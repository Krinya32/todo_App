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
	Create(userId int, list todoApp.TodoList) (int, error)
	GetAll(userId int) ([]todoApp.TodoList, error)
	GetById(userId, listId int) (todoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoApp.UpdateListInput) error
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
		TodoList:      NewTodoListPostgres(db),
	}
}
