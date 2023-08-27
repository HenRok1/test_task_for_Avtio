package main

import (
	"github.com/HenRok1/test_task_for_Avito/internal/db"
)

func main() {
	db.InitDatabase()
	api.StartServer()
}
