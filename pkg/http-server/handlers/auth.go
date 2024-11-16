package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nhassl3/todo-app/entity"
	"log/slog"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error("error to convert data in JSON", slog.String("errorResponse", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		h.logger.Error("error to create user", slog.String("errorResponse", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"id":       id,
		"name":     input.Name,
		"username": input.Username,
	}

	c.JSON(http.StatusOK, response)
	h.logger.Info("successfully created user", slog.Any("response", response))
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error("error to convert data in JSON", slog.String("errorResponse", err.Error()))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		h.logger.Error("error to create user", slog.String("errorResponse", err.Error()))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"token":    token,
		"username": input.Username,
	}

	c.JSON(http.StatusOK, response)
	h.logger.Info("successfully logged in", slog.Any("response", response))
}
