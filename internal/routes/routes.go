package routes

import (
	"dota-api/internal"
	"dota-api/internal/db"
	"dota-api/internal/handlers"
	"github.com/gorilla/mux"
)

func InitializeRoutes(logger *internal.Logger, db *db.DB) *mux.Router {
	servMux := mux.NewRouter()
	dota := handlers.NewDota(logger, db)
	servMux.HandleFunc("/heroes", dota.GetAllHeroes).Methods("GET")
	return servMux
}
