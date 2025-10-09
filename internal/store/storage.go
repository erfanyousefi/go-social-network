package store

import (
	"context"
	"database/sql"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, id uint, post *Post) error
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context) (*Post, error)
	GetByID(ctx context.Context, id uint) (*Post, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	// GetUserFeed(ctx context.Context, userID uint, fq any) (any, any, error)
}
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	GetByPostID(ctx context.Context, postID uint) (*Comment, error)
}
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id uint) error
}
type FollowerRepository interface {
	Follow(ctx context.Context, userID, followerID uint) error
	Unfollow(ctx context.Context, userID, followerID uint) error
}
type AuthTokenRepository interface {
	Create(ctx context.Context, exp time.Duration, token string) error
}

type RoleRepository interface {
	GetByName(ctx context.Context, name string) (any, error)
}

type Storage struct {
	Posts      PostRepository
	Users      UserRepository
	Comments   CommentRepository
	Followers  FollowerRepository
	AuthTokens AuthTokenRepository
	Roles      RoleRepository
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:      &PostStore{db},
		Users:      &UserStore{db},
		Comments:   &CommentStore{db},
		Followers:  &FollowerStore{db},
		AuthTokens: &AuthTokenStore{db},
		Roles:      &RoleStore{db},
	}

}
