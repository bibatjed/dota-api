package routes

import (
	"dota-api/internal"
	"dota-api/internal/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes(logger *internal.Logger) *mux.Router {
	servMux := mux.NewRouter()
	dota := handlers.NewDota(logger)
	servMux.HandleFunc("/heroes", dota.GetAllHeroes)
	return servMux
}
