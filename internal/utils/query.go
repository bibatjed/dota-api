package utils

import (
	"net/http"
)

func GenerateFilter(r *http.Request) string {
	search := r.URL.Query().Get("search")

	if search == "" {
		return ""
	}

	return search
}