package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create()
	GetOne()
	Delete()
	//more...
}

type PostRepository interface {
	Create()
	GetOne()
	Delete()
	Update()
	//more...
}

type Repository struct {
	UserRepository
	PostRepository
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(log, db),
		PostRepository: NewPostRepository(log, db),
	}
}
