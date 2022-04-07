package utils

import (
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int
}

func calculateOffset(pagination *Pagination) {
	pagination.Offset = pagination.Limit * (pagination.Page - 1)
}

func CalculatePages(page, limit int) int {
	return int(math.Ceil(float64(page / limit)))
}

func GeneratePagination(r *http.Request) Pagination {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	//add default pagination
	pagination := Pagination{
		Page:  1,
		Limit: 10,
	}

	if page != "" {
		p, _ := strconv.Atoi(page)
		pagination.Page = p
	}

	if limit != "" {
		l, _ := strconv.Atoi(limit)
		pagination.Limit = l
	}

	calculateOffset(&pagination)
	return pagination
}
