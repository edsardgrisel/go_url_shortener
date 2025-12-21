package services

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func exists(db *sql.DB, s_url string) bool {
	ctx := context.Background()
	query := "SELECT COUNT(*) FROM urls WHERE shortened_url = ?"
	var count int
	err := db.QueryRowContext(ctx, query, s_url).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func UrlShortener(db *sql.DB, url string) (string, error) {
	if url == "" {
		return "", ErrEmptyURL
	}
	salt := ""
	maxAttemps := 100
	var s_url string
	for maxAttemps > 0{
		hasher := sha1.New()
		hasher.Write([]byte(url + salt))
		hash := hasher.Sum(nil)
		s_url = hex.EncodeToString(hash)[:10]
		if !exists(db, s_url) {
			break
		}
		salt = strconv.Itoa(rand.Intn(1000000))
		maxAttemps--
	}
	if s_url == "" {
		return "", ErrGenerationFailed
	}
	ctx := context.Background()
	query := "INSERT INTO urls (shortened_url, original_url) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, s_url, url)
	if err != nil {
		return "", ErrDatabase
	}
	return s_url, nil
}
