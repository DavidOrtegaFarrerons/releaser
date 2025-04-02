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

	dtoResponse := make(map[string]TableTicketDTO)
	for key, value := range response {
		dtoResponse[key] = ToTableTicketDTO(value)
	}

	w.Header().Set("Content-Type", "application/json")

	if len(dtoResponse) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(dtoResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
