package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nhassl3/todo-app/entity"
)

func (h *Handler) createItem(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid item id", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}

	var input entity.TodoItem
	if err = c.BindJSON(&input); err != nil {
		h.logger.Error("error in converting JSON", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idItem, err := h.services.TodoItem.CreateItem(id, listId, input)
	if err != nil {
		h.logger.Error("error in creating todo-item", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"idItem": idItem,
		"idUser": id,
		"status": "ok",
	}
	c.JSON(http.StatusOK, response)
	h.logger.Info("todo-item successfully created", slog.Any("response", response))
}

func (h *Handler) getAllItems(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid item id", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}

	items, err := h.services.TodoItem.GetAllItems(id, listId)
	if err != nil {
		h.logger.Error("error in getting todo-items", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
	h.logger.Info("successfully getting all records for user", slog.Any("Data", items))

}

func (h *Handler) getItemById(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid item id", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, "Invalid item id")
		return
	}

	item, err := h.services.TodoItem.GetByIdItem(id, listId)
	if err != nil {
		h.logger.Error("error in getting item", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
	h.logger.Info("successfully getting record for user", slog.Any("response", item))
}

func (h *Handler) updateItem(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error in getting id from query", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input entity.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		h.logger.Error("error in converting JSON", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoItem.UpdateItem(id, itemId, input); err != nil {
		h.logger.Error("error in updating item", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "ok"})
	h.logger.Info("successfully update item", slog.Int("itemId", itemId))
}

func (h *Handler) deleteItem(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		h.logger.Error("User not found in context", slog.String("error", err.Error()))
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("error in converting id to int", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoItem.DeleteItem(id, itemId); err != nil {
		h.logger.Error("error deleting todo", slog.String("error", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	h.logger.Info("successfully deleting todo", slog.Int("itemId", itemId))
}
