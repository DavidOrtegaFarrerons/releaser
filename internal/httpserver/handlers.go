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
		fmt.Println("There has been an error in the healthcheck: ", err)
	}
}

func ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	tickets := release.MergeTickets()

	dtoResponse := make([]TableTicketDTO, 0, len(tickets))
	for _, ticket := range tickets {
		dtoResponse = append(dtoResponse, ToTableTicketDTO(ticket))
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
