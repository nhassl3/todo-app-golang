package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nhassl3/todo-app/entity"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	var input entity.Todos
	if err = c.BindJSON(&input); err != nil {
		h.logger.Error("error in converting JSON", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idTodo, err := h.services.TodoList.Create(id, input)
	if err != nil {
		h.logger.Error("error in creating todo", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"idTodo": idTodo,
		"idUser": id,
		"status": "ok",
	}

	c.JSON(http.StatusOK, response)
	h.logger.Info("todo successfully created", slog.Any("response", response))
}

func (h *Handler) getAllLists(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	lists, err := h.services.TodoList.GetAll(id)
	if err != nil {
		h.logger.Error("error when get all records for this user", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"Data": lists,
	}
	c.JSON(http.StatusOK, response)
	h.logger.Info("successfully getting all records for user", slog.Any("response", response))
}

func (h *Handler) getListById(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	idTodo, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error in getting id from query", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.GetListById(id, idTodo)
	if err != nil {
		h.logger.Error("error getting todo by id", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
	h.logger.Info("successfully getting all records for user", slog.Any("response", list))
}

func (h *Handler) updateList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	idTodo, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error in getting id from query", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input entity.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		h.logger.Error("error in converting JSON", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Update(id, idTodo, input); err != nil {
		h.logger.Error("error in updating todo", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "ok",
	})
	h.logger.Info("successfully update todo", slog.String("Status", "ok"))
}

func (h *Handler) deleteList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	idTodo, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error in converting id to int", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idList, err := h.services.TodoList.Delete(id, idTodo)
	if err != nil {
		h.logger.Error("error deleting todo", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"idTodo": idList,
	})
	h.logger.Info("successfully deleting todo", slog.Int("id", idList))
}
