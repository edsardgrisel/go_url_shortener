package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/edsardgrisel/go_url_shortener/internal/services"
	"github.com/redis/go-redis/v9"
)

func ShortenHandler(db *sql.DB, rdb *redis.Client) http.HandlerFunc {
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
		s_url, err := services.UrlShortener(db, rdb, string(body))
		if err != nil {
			switch {
			case errors.Is(err, services.ErrEmptyURL):
				http.Error(w, "URL cannot be empty", http.StatusBadRequest)
				return
			case errors.Is(err, services.ErrDatabase):
				http.Error(w, "Database failure", http.StatusInternalServerError)
				return
			case errors.Is(err, services.ErrGenerationFailed):
				http.Error(w, "Unable to generate shortened URL", http.StatusInternalServerError)
				return
			default:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, s_url)
	}
}
