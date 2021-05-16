package handler

import (
	book_store "github.com/Arystan2701/book-store"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getModerators(c *gin.Context) {
	moderators, err := h.services.GetModerators()
	if err != nil {
		internalError(c, err)
		return
	}
	success(c, moderators)
}

func (h *Handler) createModerator(c *gin.Context) {
	var request book_store.CreateModeratorInput
	if err := c.ShouldBindJSON(&request); err != nil {
		badRequest(c, err)
		return
	}

	moderator, err := h.services.CreateModerator(request)
	if err != nil {
		internalError(c, err)
		return
	}
	created(c, moderator)
}

func (h *Handler) updateModeratorByID(c *gin.Context) {

}
