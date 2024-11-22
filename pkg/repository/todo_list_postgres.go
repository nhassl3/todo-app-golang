package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nhassl3/todo-app/entity"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list entity.Todos) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err = row.Scan(&id); err != nil {
		return Rollback[int](tx, err)
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		fmt.Println(2)
		return Rollback[int](tx, err)
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]entity.Todos, error) {
	var lists []entity.Todos
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id=$1", todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetListById(userId, id int) (entity.Todos, error) {
	var list entity.Todos
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
		INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id=$1 AND ul.list_id=$2`, todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, id)
	return list, err
}

func (r *TodoListPostgres) Delete(userId, id int) (int, error) {
	var idTodo int
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND ul.list_id=$2 RETURNING tl.id",
		todoListsTable, usersListsTable)
	row := r.db.QueryRow(query, userId, id)
	if err := row.Scan(&idTodo); err != nil {
		return -1, err
	}
	return idTodo, nil
}

func (r *TodoListPostgres) Update(userId, listId int, input entity.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
