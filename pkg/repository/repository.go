package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nhassl3/todo-app/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type TodoList interface {
	Create(userId int, list entity.Todos) (int, error)
	GetAll(userId int) ([]entity.Todos, error)
	GetListById(userId, id int) (entity.Todos, error)
	Delete(userId, id int) (int, error)
	Update(userId, listId int, input entity.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(listId int, item entity.TodoItem) (int, error)
	GetAllItems(userId, listId int) ([]entity.TodoItem, error)
	GetByIdItem(userId, itemId int) (entity.TodoItem, error)
	DeleteItem(userId, id int) error
	UpdateItem(userId, itemId int, input entity.UpdateItemInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
