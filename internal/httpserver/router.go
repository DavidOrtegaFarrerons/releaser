package httpserver

import "net/http"

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", HealthHandler)
	mux.HandleFunc("/api/release", ReleaseHandler)
	mux.HandleFunc("/api/release/{releaseName}", ReleaseHandler)
	mux.HandleFunc("/api/set-autocomplete", SetAutoCompletePullRequestHandler)

	return mux
}
