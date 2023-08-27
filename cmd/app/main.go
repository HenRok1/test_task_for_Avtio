package main

import (
	"github.com/HenRok1/test_task_for_Avito/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.SetupRoutes(r)
	r.Run(":8080")
}
