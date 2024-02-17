package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) NewUserHandler() {
	h.Handler.Route("/project", func(r chi.Router) {

		r.Post("/", h.createUser)
		r.Get("/{id}", h.getOneUser)
		r.Delete("/{id}", h.deleteUser)

	})
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	h.Services.UserService.Create()
	//Create
}
func (h *Handler) getOneUser(w http.ResponseWriter, r *http.Request) {
	h.Services.UserService.GetOne()
	//GetOne
}
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	h.Services.UserService.Delete()
	//Delete
}
