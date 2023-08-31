package handler

import (
	"dynamic-segmentation-service/pkg/service"
	"github.com/gin-gonic/gin"

	_ "dynamic-segmentation-service/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", h.createSegment)
			segments.DELETE("/:id", h.deleteSegment)
		}

		userSegments := api.Group("user-segments")
		{
			userSegments.PUT("/", h.updateUserSegments)
			userSegments.GET("/users/:userId", h.getUserSegments)
		}

		users := api.Group("/users")
		{
			users.POST("/", h.createUser)
		}
	}

	return router
}
