package handler

import (
	"github.com/Arystan2701/book-store/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.userSignUp)
		auth.POST("/sign-in", h.userSignIn)
		auth.POST("/management", h.moderatorSignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		categories := api.Group("/categories")
		{
			categories.GET("", h.getCategories)
		}
		moderators := api.Group("/moderators")
		{
			moderators.GET("", h.getModerators)
			moderators.POST("", h.createModerator)
			moderators.PUT("/:moderatorID", h.updateModeratorByID)
		}
	}
	return router
}
