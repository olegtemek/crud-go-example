package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewUserRepository(log *slog.Logger, db *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		log: log,
		db:  db,
	}
}

func (ur *UserRepositoryImpl) Create() {
	// create
}
func (ur *UserRepositoryImpl) GetOne() {
	// GetOne
}
func (ur *UserRepositoryImpl) Delete() {
	// Delete
}
