package handler

import (
	"net/http"
	"strconv"

	grod "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ItemsGetAll(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	data, err := h.service.DatasetItems.GetAll(userId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (h *Handler) ItemsGet(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	data, err := h.service.Datasets.GetById(userId, datasetId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (h *Handler) ItemCreate(c *gin.Context) {
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
func (h *Handler) ItemChange(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.Datasets.Delete(userId, datasetId); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
func (h *Handler) ItemDelete(c *gin.Context) {
	userId, err := GetUserById(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input grod.UpdateDataset
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Datasets.Update(userId, datasetId, input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
