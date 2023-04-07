package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) buyCommon(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.BuyCommon(userId)
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
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.UnSubCommon(userId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
func (h *Handler) buyPremium(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.BuyPremium(userId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
func (h *Handler) unSubPremium(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	status, err := h.service.Subscribe.UnSubPremium(userId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
