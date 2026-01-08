package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/edsardgrisel/go_url_shortener/internal/database"
	"github.com/edsardgrisel/go_url_shortener/internal/handlers"
)

const (
	DefaultPort = "8090"
)

var db *sql.DB

func main() {
	var err error

	db, err = database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	rdb, err := database.ConnectRedis()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer rdb.Close()

	http.HandleFunc("/shorten", handlers.ShortenHandler(db, rdb))
	http.HandleFunc("/r/", handlers.RedirectHandler(db, rdb))
	err = http.ListenAndServe(":" + DefaultPort, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
