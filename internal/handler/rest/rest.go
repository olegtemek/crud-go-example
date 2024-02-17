package rest

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/olegtemek/crud-go/internal/config"
	"github.com/olegtemek/crud-go/internal/handler/rest/middleware/logger"
	"github.com/olegtemek/crud-go/internal/service"
)

type Handler struct {
	Log      *slog.Logger
	Cfg      *config.Config
	Handler  *chi.Mux
	Services *service.Service
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		Log:      log,
		Cfg:      cfg,
		Handler:  chi.NewRouter(),
		Services: services,
	}
}

func (h *Handler) Init() *http.Server {
	h.Handler.Use(middleware.RequestID)
	h.Handler.Use(middleware.URLFormat)
	h.Handler.Use(logger.New(h.Log))
	h.Handler.Use(middleware.Recoverer)
	h.Handler.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	//Init info
	h.NewUserHandler()
	h.NewPostHandler()

	handler := h.Handler

	srv := &http.Server{
		Addr:        h.Cfg.Address,
		ReadTimeout: h.Cfg.Timeout,
		Handler:     handler,
	}

	return srv
}
