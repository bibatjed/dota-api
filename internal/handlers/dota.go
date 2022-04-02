package handlers

import (
	"dota-api/internal"
	"net/http"
)

type Dota struct {
	*internal.Logger
}

func NewDota(logger *internal.Logger) *Dota {
	return &Dota{
		logger,
	}
}

func (d *Dota) GetAllHeroes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sample"))
}
