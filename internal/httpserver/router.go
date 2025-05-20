package httpserver

import (
	"net/http"
	"release-handler/internal/task"
)

func NewRouter(taskHandler *task.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", HealthHandler)
	mux.HandleFunc("/api/release", ReleaseHandler)
	mux.HandleFunc("/api/release/{releaseName}", ReleaseHandler)
	mux.HandleFunc("/api/set-autocomplete", SetAutoCompletePullRequestHandler)

	mux.HandleFunc("/api/task", taskHandler.AddTask)

	return mux
}
