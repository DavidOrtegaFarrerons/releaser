package task

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"release-handler/config"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString(config.CorsDomain))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight request --> Pending refactor to middleware
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var input CreateTaskInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.Service.AddTask(
		input.PrId,
		input.ReleaseId,
		input.Type,
		input.Content,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("task created successfully"))
}

func (h *Handler) GetTasksByIds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", viper.GetString(config.CorsDomain))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight request --> Pending refactor to middleware
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var input GetTasksByIdInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	tasks, err := h.Service.GetTasksByIdsAndType(
		input.PrIds,
		input.Type,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		fmt.Printf("error while sending tasks:, %v\n", err)
	}

}
