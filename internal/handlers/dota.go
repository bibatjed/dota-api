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

type PaginationResponse struct {
	Count int `json:"count"`
	Page  int `json:"page"`
}

type Response struct {
	Pagination PaginationResponse `json:"pagination"`
	Result     []models.Hero      `json:"result"`
}

func NewDota(logger *internal.Logger, db *db.DB) *Dota {
	return &Dota{
		logger,
		db,
	}
}

func (d *Dota) GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	pagination := utils.GeneratePagination(r)
	search := utils.GenerateFilter(r)
	heroes := make(chan []models.Hero)
	countHeroes := make(chan int)
	go func() {
		heroes <- d.DB.GetAllHeroes(pagination, search)
	}()

	go func() {
		countHeroes <- d.DB.GetAllHeroesCount()
	}()

	apiResponse := Response{}
	for i := 0; i < 2; i++ {
		select {
		case s1 := <-heroes:
			apiResponse.Result = s1
		case s2 := <-countHeroes:
			totalPages := utils.CalculatePages(s2, pagination.Limit)
			apiResponse.Pagination = PaginationResponse{Page: totalPages, Count: s2}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(apiResponse)
}
