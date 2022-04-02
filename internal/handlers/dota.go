package handlers

import (
	"dota-api/internal"
	"dota-api/internal/db"
	"encoding/json"
	"net/http"
)

type Dota struct {
	*internal.Logger
	*db.DB
}

func NewDota(logger *internal.Logger, db *db.DB) *Dota {
	return &Dota{
		logger,
		db,
	}
}

//TODO: add query filters
func (d *Dota) GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	heroes := d.DB.GetAllHeroes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
