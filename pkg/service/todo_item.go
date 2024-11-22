package service

import (
	"github.com/nhassl3/todo-app/entity"
	"github.com/nhassl3/todo-app/pkg/repository"
)

type TodoItemService struct {
	repos repository.TodoItem
}

func NewTodoItemService(repos repository.TodoItem) *TodoItemService {
	return &TodoItemService{repos: repos}
}

func (h *TodoItemService) CreateItem(userId int, item entity.TodoItem) (int, error) {
	return h.repos.CreateItem(userId, item)
}

func (h *TodoItemService) GetAllItems(userId int) ([]entity.TodoItem, error) {
	return h.repos.GetAllItems(userId)
}

func (h *TodoItemService) GetListByIdItem(userId, id int) (entity.TodoItem, error) {
	return h.repos.GetListByIdItem(userId, id)
}

func (h *TodoItemService) DeleteItem(userId, id int) (int, error) {
	return h.repos.DeleteItem(userId, id)
}

func (h *TodoItemService) UpdateItem(userId, listId int, input entity.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return h.repos.UpdateItem(userId, listId, input)
}
