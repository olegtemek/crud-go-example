package service

import (
	"log/slog"

	"github.com/olegtemek/crud-go/internal/repository"
)

type UserService interface {
	Create()
	GetOne()
	Delete()
}
type PostService interface {
	Create()
	GetOne()
	Delete()
	Update()
}

type Service struct {
	UserService
	PostService
}

func NewService(log *slog.Logger, repos *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(log, &repos.UserRepository),
		PostService: NewPostService(log, &repos.PostRepository),
	}
}
