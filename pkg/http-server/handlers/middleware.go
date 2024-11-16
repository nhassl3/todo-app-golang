package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId "
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}

	headersParts := strings.Split(header, " ")
	if len(headersParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headersParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return 0, errors.New("user not found in context")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "User id have a invalid type")
		return 0, errors.New("user id have a invalid type")
	}
	return idInt, nil
}
