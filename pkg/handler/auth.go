package handler

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/gin-gonic/gin"
)

func (h *Handler) userSignUp(c *gin.Context) {
	var request book_store.User
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}

	id, err := h.services.CreateUser(request)
	if err != nil {
		internalError(c, err)
		return
	}
	created(c, id)
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) userSignIn(c *gin.Context) {
	var request signInInput
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}

	token, err := h.services.GenerateToken(request.Email, request.Password)
	if err != nil {
		internalError(c, err)
		return
	}
	success(c, gin.H{
		"access_token": token,
	})
}

func (h *Handler) moderatorSignIn(c *gin.Context) {

}
