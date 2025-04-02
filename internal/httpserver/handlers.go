package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"release-handler/internal/release"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))

	if err != nil {
		fmt.Println("There has been an error in the healthcheck: " + err.Error())
	}
}

func ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	response := release.MergeTickets()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
