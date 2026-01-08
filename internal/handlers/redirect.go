package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/edsardgrisel/go_url_shortener/internal/services"
	"github.com/redis/go-redis/v9"
)

func RedirectHandler(db *sql.DB, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		s_url := strings.TrimPrefix(req.URL.Path, "/r/")
		o_url, err := services.GetOriginalUrl(db, rdb, s_url)
		if err != nil {
			switch {
			case errors.Is(err, services.ErrEmptyURL):
				http.Error(w, "URL cannot be empty", http.StatusBadRequest)
				return
			case errors.Is(err, services.ErrDatabase):
				http.Error(w, "Database failure", http.StatusInternalServerError)
				return
			case errors.Is(err, services.ErrURLNotFound):
				http.Error(w, "URL not found", http.StatusNotFound)
				return
			default:
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, req, o_url, http.StatusFound)
	}
}
