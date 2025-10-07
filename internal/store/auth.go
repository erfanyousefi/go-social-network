package store

import (
	"context"
	"database/sql"
	"time"
)

type AuthTokenStore struct {
	db *sql.DB
}

func (s *AuthTokenStore) Create(ctx context.Context, exp time.Duration, token string) error {
	return nil
}
