package main

import (
	"github.com/HenRok1/test_task_for_Avito/internal/db"
	handlers "github.com/HenRok1/test_task_for_Avito/internal/handler"
	"github.com/HenRok1/test_task_for_Avito/pkg/api"
)

func main() {

	segmentHandler := &handlers.SegmentHandler{}

	db.InitDB("")
	api.StartServer(segmentHandler)
}
