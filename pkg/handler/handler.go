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
		auth.POST("sign-up", h.userSignUp)
		auth.POST("sign-in", h.userSignUp)
		auth.POST("management", h.moderatorSignIn)
	}
	return router
}