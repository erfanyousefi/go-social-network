package store

import (
	"context"
	"database/sql"
	"fmt"
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
	query := "DELETE FROM user WHERE id = $1"
	if err := s.db.QueryRowContext(ctx, query, id).Err(); err != nil {
		return err
	}
	return nil

}

func (s *UserStore) Get(ctx context.Context, id uint) (*User, error) {
	var user User
	query := "SELECT * FROM user WHERE id = $1"
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found by id: %d", id)
		}
		return nil, fmt.Errorf("user by id: %d, %v", id, err)
	}
	return &user, nil
}
func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := "SELECT * FROM user WHERE email = $1"
	row := s.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("user not found by email: %s", email)
		}
		return &user, err
	}
	return &user, nil

}
