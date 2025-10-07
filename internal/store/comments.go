package store

import (
	"context"
	"database/sql"
	"time"
)

type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"postId"`
	UserID    uint      `json:"userId"`
	Text      string    `json:"text"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	return nil
}
func (s *CommentStore) GetByPostID(ctx context.Context, postID uint) (*Comment, error) {
	return nil, nil
}
