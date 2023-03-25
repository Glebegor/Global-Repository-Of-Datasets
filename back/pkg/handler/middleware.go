package handler

import "github.com/gin-gonic/gin"

func (h *Handler) userIndentity(c *gin.Context) {
	c.Set("UserID", "none")
}
