package service

import (
	"github.com/nhassl3/todo-app/entity"
	"github.com/nhassl3/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list entity.Todos) (int, error)
	GetAll(userId int) ([]entity.Todos, error)
	GetListById(userId, id int) (entity.Todos, error)
	Delete(userId, id int) (int, error)
	Update(userId, itemId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userId, listId int, item entity.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]entity.TodoItem, error)
	GetByIdItem(userId, itemId int) (entity.TodoItem, error)
	DeleteItem(userId, id int) error
	UpdateItem(userId, itemId int, input entity.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
