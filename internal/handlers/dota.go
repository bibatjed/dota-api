package handlers

import (
	"dota-api/internal"
	"dota-api/internal/db"
	"dota-api/internal/models"
	"dota-api/internal/utils"
	"encoding/json"
	"net/http"
)

type Dota struct {
	*internal.Logger
	*db.DB
}

type Response struct {
	//Pagination `json:"pagination"`
	Result []models.Hero `json:"result"`
}

func NewDota(logger *internal.Logger, db *db.DB) *Dota {
	return &Dota{
		logger,
		db,
	}
}

func (d *Dota) GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	pagination := utils.GeneratePagination(r)
	heroes := d.DB.GetAllHeroes(pagination)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{Result: heroes})
}
