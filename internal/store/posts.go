package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID               uint      `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short_description"`
	Content          string    `json:"content"`
	AuthorID         uint      `json:"authorId"`
	Image            string    `json:"image"`
	Tags             []string  `json:"tags"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
        INSERT INTO post (title, short_description, content, authorId, image, tags) 
        VALUES($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at
    `
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.ShortDescription,
		post.Content,
		post.AuthorID,
		post.Image,
		pq.Array(post.Tags),
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostStore) Update(ctx context.Context, id uint, post *Post) error {
	return nil
}
func (s *PostStore) Delete(ctx context.Context, id uint) error {
	return nil
}
func (s *PostStore) Get(ctx context.Context) (*Post, error) {
	return nil, nil
}
func (s *PostStore) GetByID(ctx context.Context, id uint) (*Post, error) {
	return nil, nil
}
func (s *PostStore) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	return nil, nil
}
