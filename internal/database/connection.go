package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = ""
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "url_shortener"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("Failed to open sql database %w", err)
	}
	return db, nil
}
