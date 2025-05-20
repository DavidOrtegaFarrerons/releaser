package httpserver

import (
	"log"
	"net/http"
	"release-handler/internal/db"
	"release-handler/internal/task"
)

func Start() error {
	dbClient, err := db.NewPostgresClient()
	if err != nil {
		return err
	}
	taskRepo := task.NewPostgresRepository(dbClient)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)
	router := NewRouter(taskHandler)
	log.Println("Server running on :8080")
	return http.ListenAndServe(":8080", router)
}
