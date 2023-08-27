package api

import (
	"github.com/gin-gonic/gin"
)

type SegmentHandlerInterface interface {
	CreateSegment(c *gin.Context)
	// Добавьте другие методы обработчика сегментов
}

func StartServer(segmentHandler SegmentHandlerInterface) {
	router := gin.Default()

	api := router.Group("/api")
	{
		segments := api.Group("/segments")
		{
			segments.POST("/", segmentHandler.CreateSegment)
			// Добавьте другие роуты для сегментов
		}

		// Добавьте аналогичные группы и роуты для пользователей
		// users := api.Group("/users")
		// {
		//     users.POST("/", userHandler.CreateUser)
		//     // Добавьте другие роуты для пользователей
		// }
	}

	router.Run(":8080")
}
