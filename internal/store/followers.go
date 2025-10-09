package store

import (
	"context"
	"database/sql"
	"fmt"
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
	var follower Follower
	findQuery := "SELECT * FROM follower WHERE userId = $1 AND followerId = $2"
	createQuery := "INSERT INTO follower(userId, followerId) VALUES($1, $2) RETURNING id"

	row := s.db.QueryRowContext(ctx, findQuery, userID, followerID)
	if err := row.Scan(&follower.ID); err != nil {
		if err == sql.ErrNoRows {
			if err := s.db.QueryRowContext(ctx, createQuery, userID, followerID).Scan(&follower.ID); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("you already followed this account: %d", followerID)

}
func (s *FollowerStore) Unfollow(ctx context.Context, userID, followerID uint) error {
	var follower Follower
	findQuery := "SELECT * FROM follower WHERE userId = $1 AND followerId = $2"
	deleteQuery := "DELETE FROM follower id = $1"

	row := s.db.QueryRowContext(ctx, findQuery, userID, followerID)
	if err := row.Scan(&follower.ID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("not found this user in your followers")
		}
		return fmt.Errorf("can n't unfollow %v", err)
	}
	if err := s.db.QueryRowContext(ctx, deleteQuery, follower.ID); err != nil {
		return fmt.Errorf("can n't unfollow %v", err)
	}
	return nil
}
