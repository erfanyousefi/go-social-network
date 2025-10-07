package store

import (
	"context"
	"database/sql"
)

type Follower struct {
	ID         uint `json:"id"`
	UserID     uint `json:"userId"`
	FollowerID uint `json:"followerId"`
}
type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, userID, followerID uint) error {
	return nil
}
func (s *FollowerStore) Unfollow(ctx context.Context, userID, followerID uint) error {
	return nil
}
