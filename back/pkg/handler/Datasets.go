package handler

import (
	"net/http"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/gin-gonic/gin"
)

func (h *Handler) dataSetGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) dataSetsAllGet(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	data, err := h.service.Datasets.GetAll(userId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (h *Handler) dataSetCreate(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var input grod.Dataset
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Datasets.Create(userId, input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
		"desc":   input.Description,
		"title":  input.Title,
	})
}
func (h *Handler) dataSetDelete(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) dataSetChange(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) dataSetGetRandomRow(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
