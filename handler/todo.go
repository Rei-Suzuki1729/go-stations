package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

type TODOHandler struct {
	TODOService *service.TODOService
}

func NewTODOHandler(service *service.TODOService) *TODOHandler {
	return &TODOHandler{TODOService: service}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func (h *TODOHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTODORequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if req.Subject == "" {
		http.Error(w, "Subject cannot be empty", http.StatusBadRequest)
		return
	}

	todo, err := h.TODOService.CreateTODO(r.Context(), req.Subject, req.Description)
	if err != nil {
		http.Error(w, "Failed to create TODO", http.StatusInternalServerError)
		return
	}

	res := model.CreateTODOResponse{TODO: *todo}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
