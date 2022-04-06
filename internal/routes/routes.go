package routes

import (
	"dota-api/internal"
	"dota-api/internal/db"
	"dota-api/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func InitializeRoutes(logger *internal.Logger, db *db.DB) http.Handler {
	servMux := mux.NewRouter()
	dota := handlers.NewDota(logger, db)
	servMux.HandleFunc("/heroes", dota.GetAllHeroes).Methods("GET")

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(servMux)
	return handler
}
