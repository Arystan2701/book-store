package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		unauthorized(c, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		unauthorized(c, "invalid auth header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		unauthorized(c, err.Error())
		return
	}

	c.Set(userContext, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userContext)
	if !ok {
		internalError(c, &BaseError{Code: 101, Message: "user id not found"})
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		internalError(c, &BaseError{Code: 102, Message: "user id is of invalid type"})
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}
