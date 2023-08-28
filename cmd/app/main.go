package main

import (
	"github.com/HenRok1/test_task_for_Avito/internal/db"
	handlers "github.com/HenRok1/test_task_for_Avito/internal/handler"
	"github.com/HenRok1/test_task_for_Avito/pkg/api"
)

func main() {

	segmentHandler := &handlers.SegmentHandler{}

	db.InitDB("host=localhost port=5432 dbname=my_db user=postgres password=803800 sslmode=disable")
	api.StartServer(segmentHandler)
}
