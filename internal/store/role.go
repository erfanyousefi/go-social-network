package store

import (
	"context"
	"database/sql"
)

type RoleStore struct {
	db *sql.DB
}

func (s *RoleStore) GetByName(ctx context.Context, name string) (any, error) {
	return nil, nil
}
