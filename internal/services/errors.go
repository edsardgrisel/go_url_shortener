package services

import "errors"

var (
	ErrEmptyURL = errors.New("URL cannot be empty")
	ErrURLNotFound = errors.New("URL not found")
	ErrDatabase = errors.New("database connection or query failure")
	ErrGenerationFailed = errors.New("failed to generate unique URL")
)