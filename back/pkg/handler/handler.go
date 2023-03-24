package handler

import (
	service "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/service"
	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/reg", h.authReg)
			auth.POST("/log", h.authLog)
		}
	}
	return router
}
