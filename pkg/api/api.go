package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	// Создаем экземпляры обработчиков
	segmentHandler := &handler.SegmentHandler{}
	// userHandler := &handlers.UserHandler{} // Добавьте обработчик для пользователей

	// Определяем роуты
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
