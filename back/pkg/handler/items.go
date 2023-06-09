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
	datasetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	data, err := h.service.DatasetItems.GetAll(userId, datasetId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
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
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	data, err := h.service.DatasetItems.ItemsGet(userId, datasetId, itemId)
	if err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}
func (h *Handler) ItemCreate(c *gin.Context) {
	var input grod.DatasetItem
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
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid request")
		return
	}
	if err := h.service.DatasetItems.Create(userId, datasetId, input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}
func (h *Handler) ItemChange(c *gin.Context) {
	var input grod.UpdateDatasetItem
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
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	if err := h.service.DatasetItems.ItemsUpdate(userId, datasetId, itemId, input); err != nil {
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
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.service.DatasetItems.ItemsDelete(userId, datasetId, itemId); err != nil {
		newResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "200",
	})
}
