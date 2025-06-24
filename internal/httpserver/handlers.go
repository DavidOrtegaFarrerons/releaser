package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"release-handler/config"
	"release-handler/internal/release"
	"release-handler/internal/scm/azure"
	"sync"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))

	if err != nil {
		fmt.Println("There has been an error in the healthcheck: ", err)
	}
}

func ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString(config.CorsDomain))
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	releaseDate := r.PathValue("releaseName")

	if releaseDate == "" {
		releaseDate = viper.GetString(config.JiraDefaultRelease)
	}

	releaseName := "Release/" + releaseDate

	tickets := release.MergeTickets(releaseName)

	tableTickets := make([]TableTicketDTO, 0, len(tickets))
	for _, ticket := range tickets {
		tableTickets = append(tableTickets, ToTableTicketDTO(ticket))
	}

	if len(tableTickets) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := ToResponse(tableTickets, releaseDate)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func SetAutoCompletePullRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString(config.CorsDomain))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed, use POST", http.StatusMethodNotAllowed)
		return
	}

	var req autoCompleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	var ids []int
	if req.PullRequestID != nil {
		ids = append(ids, *req.PullRequestID)
	}
	ids = append(ids, req.PullRequestIDs...)
	if len(ids) == 0 {
		http.Error(w, "provide pullRequestId or pullRequestIds", http.StatusBadRequest)
		return
	}

	client := azure.NewClient()
	results := make([]autoCompleteResult, len(ids))

	var wg sync.WaitGroup
	wg.Add(len(ids))
	for i, id := range ids {
		go func(idx, prID int) {
			defer wg.Done()
			res, err := client.SetAutocompletionInPullRequest(prID)
			if err != nil {
				results[idx] = autoCompleteResult{PullRequestID: prID, Error: err.Error()}
				return
			}
			results[idx] = autoCompleteResult{PullRequestID: prID, Result: res}
		}(i, id)
	}
	wg.Wait()

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(results)
}

type autoCompleteRequest struct {
	PullRequestID  *int  `json:"pullRequestId,omitempty"`
	PullRequestIDs []int `json:"pullRequestIds,omitempty"`
}

type autoCompleteResult struct {
	PullRequestID int         `json:"pullRequestId"`
	Result        interface{} `json:"result,omitempty"`
	Error         string      `json:"error,omitempty"`
}
