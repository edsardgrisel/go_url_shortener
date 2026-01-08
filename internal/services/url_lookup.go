package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetOriginalUrl(db *sql.DB, rdb *redis.Client, s_url string) (string, error) {
	if s_url == "" {
		return "", ErrEmptyURL
	}
	ctx := context.Background()
	cacheKey := "url:" + s_url
	val, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Printf("Cache hit")
		return val, nil
	}
	query := "SELECT original_url FROM urls WHERE shortened_url = ?"
	var o_url string
	err = db.QueryRowContext(ctx, query, s_url).Scan(&o_url)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrURLNotFound
		}
		return "", ErrDatabase
	}
	rdb.Set(ctx, cacheKey, o_url, time.Hour)
	return o_url, nil
}
