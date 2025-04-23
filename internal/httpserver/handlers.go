package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"release-handler/internal/release"
	"release-handler/internal/scm/azure"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))

	if err != nil {
		fmt.Println("There has been an error in the healthcheck: ", err)
	}
}

func ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5174")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	tickets := release.MergeTickets()

	dtoResponse := make([]TableTicketDTO, 0, len(tickets))
	for _, ticket := range tickets {
		dtoResponse = append(dtoResponse, ToTableTicketDTO(ticket))
	}

	if len(dtoResponse) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(dtoResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func SetAutoCompletePullRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5174")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not Allowed, Methods allowed are: "+http.MethodPost, http.StatusMethodNotAllowed)
		return
	}

	var req AutoCompleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	client := azure.NewClient()
	resp, err := client.SetAutocompletionInPullRequest(req.PullRequestId)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

type AutoCompleteRequest struct {
	PullRequestId int `json:"pullRequestId"`
}
