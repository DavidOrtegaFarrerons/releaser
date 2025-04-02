package httpserver

import (
	"log"
	"net/http"
)

func Start() error {
	router := NewRouter()
	log.Println("Server running on :8080")
	return http.ListenAndServe(":8080", router)
}
