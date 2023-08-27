package main

import (
	"fmt"

	"github.com/HenRok1/test_task_for_Avito/internal/db"
	"github.com/HenRok1/test_task_for_Avito/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db.InitDB("host=localhost port=5432 dbname=my_db user=postgres password=803800 sslmode=disable")

	r := gin.Default()
	server.SetupRoutes(r)

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
