package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nhassl3/todo-app/pkg/service"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
)

type MainHandler struct {
	services *service.Service
	logger   *slog.Logger
}

func NewMainHandler(services *service.Service, logger *slog.Logger) *MainHandler {
	return &MainHandler{services: services, logger: logger}
}

func (h *MainHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET(":id", h.getListById)
			lists.PUT(":id", h.updateList)
			lists.DELETE(":id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	router.Use(sloggin.New(h.logger))

	return router
}
