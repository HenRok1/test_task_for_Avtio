package server

import (
	"github.com/HenRok1/test_task_for_Avito/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/create-segment", handler.CreateSegment)
	router.DELETE("/delete-segment/:slug", handler.DeleteSegment)
	router.POST("/add-user-segments/:userID", handler.AddUserSegments)
	router.GET("/user-segments/:userID", handler.GetUserSegments)
}
