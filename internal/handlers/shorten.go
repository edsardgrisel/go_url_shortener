package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/edsardgrisel/go_url_shortener/internal/services"
)

func ShortenHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		s_url := services.UrlShortener(db, string(body))
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, s_url)
	}
}
