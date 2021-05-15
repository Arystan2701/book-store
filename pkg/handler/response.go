package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseError struct {
	Code    int32
	Message string
}

func (b *BaseError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", b.Code, b.Message)
}

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func badRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, processError(err))
}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, processError(err))
}

func processError(err error) interface{} {
	var result interface{}
	if e, ok := err.(*BaseError); ok {
		result = Response{
			Code:    e.Code,
			Message: e.Message,
		}
	} else if err != nil {
		result = gin.H{"message": err.Error(), "code": 8888}
	} else {
		result = gin.H{"message": "error", "code": 8888}
	}
	return result
}

func success(c *gin.Context, data interface{}) {
	response := Response{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func created(c *gin.Context, data interface{}) {
	response := Response{
		Code:    0,
		Message: "Created",
		Data:    data,
	}
	c.JSON(http.StatusCreated, response)
}

func unauthorized(c *gin.Context, err *BaseError) {
	c.JSON(http.StatusUnauthorized, processError(err))
}

func forbidden(c *gin.Context, err *BaseError) {
	c.JSON(http.StatusForbidden, processError(err))
}

func errorResponse(c *gin.Context, err *BaseError) {
	c.JSON(http.StatusForbidden, processError(err))
}
