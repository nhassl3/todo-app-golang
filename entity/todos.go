package entity

import (
	"errors"
)

type Todos struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int `json:"id"`
	UserId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	TitleItem       *string `json:"title"`
	DescriptionItem *string `json:"description"`
	Done            *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.TitleItem == nil && i.DescriptionItem == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
