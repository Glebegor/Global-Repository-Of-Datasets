package handler

import (
	service "github.com/Glebegor/Global-Repository-Of-Datasets/tree/master/back/pkg/service"
	gin "github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORS())

	auth := router.Group("/auth")
	{
		auth.POST("/reg", h.authReg)
		auth.POST("/log", h.authLog)
	}

	api := router.Group("/api", h.userIndentity)
	{
		database := api.Group("/datasets")
		{
			database.GET("/", h.dataSetsAllGet)
			database.POST("/", h.dataSetCreate)

			database.GET("/:id", h.dataSetGet)
			database.GET("/:id/random", h.dataSetGetRandomRow)
			database.PUT("/:id", h.dataSetChange)
			database.DELETE("/:id", h.dataSetDelete)

			databaseId := database.Group("/:id")
			{
				items := databaseId.Group("/items")
				{
					items.GET("/", h.ItemsGetAll)
					items.GET("/:item_id", h.ItemsGet)
					items.POST("/", h.ItemCreate)
					items.PUT("/:item_id", h.ItemChange)
					items.DELETE("/:item_id", h.ItemDelete)
				}
			}

		}

		subscribes := api.Group("/subscribes")
		{
			common := subscribes.Group("/common")
			{
				common.POST("/buy", h.buyCommon)
				common.POST("/buy-free", h.buyFreeCommon)
				common.POST("/canceling", h.unSubCommon)
			}
			premium := subscribes.Group("/premium")
			{
				premium.POST("/buy", h.buyPremium)
				premium.POST("/canceling", h.unSubPremium)
			}
		}
	}
	return router
}
