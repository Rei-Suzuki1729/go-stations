package router

import (
	"database/sql"
	"net/http"
	"github.com/TechBowl-japan/go-stations/handler"
)

import (
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	// Add the /healthz endpoint
	mux.Handle("/healthz", handler.NewHealthzHandler())

	// Create a new TODOService instance using the todoDB instance
	todoService := service.NewTODOService(todoDB)

	// Pass the todoService instance to the NewTODOHandler function
	mux.Handle("/todos", handler.NewTODOHandler(todoService))
	return mux
}
