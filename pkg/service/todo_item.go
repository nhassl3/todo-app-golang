package service

import (
	"github.com/nhassl3/todo-app/entity"
	"github.com/nhassl3/todo-app/pkg/repository"
)

type TodoItemService struct {
	repos    repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repos repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repos: repos, listRepo: listRepo}
}

func (h *TodoItemService) CreateItem(userId, listId int, item entity.TodoItem) (int, error) {
	_, err := h.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return h.repos.CreateItem(listId, item)
}

func (h *TodoItemService) GetAllItems(userId, listId int) ([]entity.TodoItem, error) {
	return h.repos.GetAllItems(userId, listId)
}

func (h *TodoItemService) GetByIdItem(userId, itemId int) (entity.TodoItem, error) {
	return h.repos.GetByIdItem(userId, itemId)
}

func (h *TodoItemService) DeleteItem(userId, id int) error {
	return h.repos.DeleteItem(userId, id)
}

func (h *TodoItemService) UpdateItem(userId, itemId int, input entity.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return h.repos.UpdateItem(userId, itemId, input)
}
