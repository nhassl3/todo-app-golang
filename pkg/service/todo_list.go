package service

import (
	"github.com/nhassl3/todo-app/entity"
	"github.com/nhassl3/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list entity.Todos) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]entity.Todos, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetListById(userId, id int) (entity.Todos, error) {
	return s.repo.GetListById(userId, id)
}

func (s *TodoListService) Delete(userId, id int) (int, error) {
	return s.repo.Delete(userId, id)
}

func (s *TodoListService) Update(userId, listId int, input entity.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
