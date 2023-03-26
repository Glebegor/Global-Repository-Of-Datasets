package handler

import (
	"net/http"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/gin-gonic/gin"
)

type LogRedInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (h *Handler) authReg(c *gin.Context) {
	var input grod.User
	input.Subscribe = "free"

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
		"UserId": input,
	})
}
func (h *Handler) authLog(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
