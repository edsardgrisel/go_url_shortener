package main

import (
	"net/http"
	"github.com/edsardgrisel/go_url_shortener/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)

	http.ListenAndServe(":8090", nil)
}
