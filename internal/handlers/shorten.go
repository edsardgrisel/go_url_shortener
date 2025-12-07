package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/edsardgrisel/go_url_shortener/internal/services"
)

func ShortenHandler(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	}
	s_url := services.Url_shortener(string(body))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, s_url)
}
