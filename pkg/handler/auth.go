package handler

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/gin-gonic/gin"
)

// @Summary UserSignUp
// @Tags auth
// @Description create user account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body book_store.User true "account info"
// @Success 200,201 {integer} integer 1
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-up [post]
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

// @Summary UserSignIn
// @Tags auth
// @Description authorization user account
// @ID authorization-account
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign in form"
// @Success 200,201 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /auth/sign-in [post]
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
	var request signInInput
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}

	token, err := h.services.GenerateModeratorToken(request.Email, request.Password)
	if err != nil {
		internalError(c, err)
		return
	}
	success(c, gin.H{
		"access_token": token,
	})
}
