package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User_id struct {
	Id int `json: "id" binding:"required"`
}

func (h *Handler) buyCommon(c *gin.Context) {
	var userId User_id
	if err := c.BindJSON(&userId); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.BuyCommon(userId.Id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
func (h *Handler) buyFreeCommon(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) unSubCommon(c *gin.Context) {
	var userId User_id
	if err := c.BindJSON(&userId); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.UnSubCommon(userId.Id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
func (h *Handler) buyPremium(c *gin.Context) {
	var userId User_id
	if err := c.BindJSON(&userId); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.BuyPremium(userId.Id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
func (h *Handler) unSubPremium(c *gin.Context) {
	var userId User_id
	if err := c.BindJSON(&userId); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.UnSubPremium(userId.Id)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
