package services

import (
	"context"
	"database/sql"
)

func GetOriginalUrl(db *sql.DB, s_url string) (string, error) {
	if s_url == "" {
		return "", ErrEmptyURL
	}
	ctx := context.Background()
	query := "SELECT original_url FROM urls WHERE shortened_url = ?"
	var o_url string
	err := db.QueryRowContext(ctx, query, s_url).Scan(&o_url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrURLNotFound
		}
		return "", ErrDatabase
	}
	return o_url, nil
}
