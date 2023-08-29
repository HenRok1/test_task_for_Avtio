package api

import (
	handler "github.com/HenRok1/test_task_for_Avito/internal/handler"
	"github.com/gin-gonic/gin"
)

func StartServer(segmentHandler handler.SegmentHandlerInterface, userHandler handler.UserHandlerInterface) {
	router := gin.Default()

	// api := router.Group("/api")
	// {
	// 	segments := api.Group("/segments")
	// 	{
	// 		segments.POST("/", segmentHandler.CreateSegment)
	// 		// Добавьте другие роуты для сегментов
	// 	}

	// 	// Добавьте аналогичные группы и роуты для пользователей
	// 	users := api.Group("/users")
	// 	{
	// 		users.POST("/", userHandler.CreateUser)
	// 		// Добавьте другие роуты для пользователей
	// 	}
	// }

	segments := router.Group("/segments")
	{
		segments.POST("/", segmentHandler.CreateSegment)
		segments.DELETE("/:name", segmentHandler.DeleteSegment)
	}

	users := router.Group("/users")
	{
		users.POST("/", userHandler.AddUserSegments)
		users.DELETE("/:user_id/:segment_name", userHandler.RemoveUserSegments)
	}

	router.Run(":8080")
}
