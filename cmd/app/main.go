package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	handlers "github.com/HenRok1/test_task_for_Avito/internal/handler"
	"github.com/HenRok1/test_task_for_Avito/internal/repositories"
	"github.com/HenRok1/test_task_for_Avito/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=803800 sslmode=disable host=localhost port=5432 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	segmentRepo := repositories.NewSegmentRepository(db)
	userRepo := repositories.NewUserRepository(db)

	segmentService := services.NewSegmentService(segmentRepo)
	userService := services.NewUserService(userRepo)

	segmentHandler := handlers.NewSegmentHandler(segmentService)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()

	segments := router.Group("/segments")
	{
		segments.POST("/", segmentHandler.CreateSegment)
		segments.DELETE("/:name", segmentHandler.DeleteSegment)
	}

	users := router.Group("/users")
	{
		users.POST("/:id/:name", userHandler.AddUserSegments)
	}

	router.Run(":8080")
}
