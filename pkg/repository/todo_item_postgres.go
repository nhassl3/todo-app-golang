package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nhassl3/todo-app/entity"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(userId int, list entity.TodoItem) (int, error) {
	
}

func (r *TodoItemPostgres) GetAllItems(userId int) ([]entity.TodoItem, error) {
	var lists []entity.TodoItem
	return lists, nil
}

func (r *TodoItemPostgres) GetListByIdItem(userId, id int) (entity.TodoItem, error) {
	var list entity.TodoItem
	return list, nil
}

func (r *TodoItemPostgres) DeleteItem(userId, id int) (int, error) {
	var idTodo int
	return idTodo, nil
}

func (r *TodoItemPostgres) UpdateItem(userId, listId int, input entity.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	_, _, _ = setValues, args, argId // TODO: PLUG

	return nil
}
