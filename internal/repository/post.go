package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepositoryImpl struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewPostRepository(log *slog.Logger, db *pgxpool.Pool) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		log: log,
		db:  db,
	}
}

func (ur *PostRepositoryImpl) Create() {
	// create
}
func (ur *PostRepositoryImpl) GetOne() {
	// GetOne
}
func (ur *PostRepositoryImpl) Delete() {
	// Delete
}

func (ur *PostRepositoryImpl) Update() {
	// Update
}
