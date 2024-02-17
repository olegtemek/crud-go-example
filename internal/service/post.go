package service

import (
	"log/slog"

	"github.com/olegtemek/crud-go/internal/repository"
)

type PostServiceImpl struct {
	log        *slog.Logger
	repository *repository.PostRepository
}

func NewPostService(log *slog.Logger, repository *repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		log:        log,
		repository: repository,
	}
}

func (ur *PostServiceImpl) Create() {
	// create
	(*ur.repository).Create()
}
func (ur *PostServiceImpl) GetOne() {
	// GetOne
	(*ur.repository).GetOne()
}
func (ur *PostServiceImpl) Delete() {
	// Delete
	(*ur.repository).Delete()
}

func (ur *PostServiceImpl) Update() {
	(*ur.repository).Update()
}
