package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) authReg(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) authLog(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
