package service

import (
	"log/slog"

	"github.com/olegtemek/crud-go/internal/repository"
)

type UserServiceImpl struct {
	log        *slog.Logger
	repository *repository.UserRepository
}

func NewUserService(log *slog.Logger, repository *repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		log:        log,
		repository: repository,
	}
}

func (ur *UserServiceImpl) Create() {
	// create
	(*ur.repository).Create()
}
func (ur *UserServiceImpl) GetOne() {
	// GetOne
	(*ur.repository).GetOne()
}
func (ur *UserServiceImpl) Delete() {
	// Delete
	(*ur.repository).Delete()
}
