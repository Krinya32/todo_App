package service

import (
	"github.com/krinya32/todoApp"
	"github.com/krinya32/todoApp/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todoApp.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todoApp.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todoApp.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todoApp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return nil
	}
	return s.repo.Update(userId, listId, input)
}
