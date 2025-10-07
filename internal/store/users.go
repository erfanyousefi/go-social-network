package store

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
        INSERT INTO user (firstname, lastname, email, password)
        VALUES($1, $2, $3, $4) RETURNING id, created_at
    `
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) Delete(ctx context.Context, id uint) error {
	return nil
}

func (s *UserStore) Get(ctx context.Context, id uint) (*User, error) {
	return nil, nil
}
func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}
func (s *UserStore) CreateAndActive(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	return nil
}
