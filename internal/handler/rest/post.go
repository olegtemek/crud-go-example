package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) NewPostHandler() {
	h.Handler.Route("/project", func(r chi.Router) {

		r.Post("/", h.createPost)
		r.Get("/{id}", h.getOnePost)
		r.Patch("/{id}", h.updatePost)
		r.Delete("/{id}", h.deletePost)

	})
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	h.Services.PostService.Create()
	//Create
}
func (h *Handler) getOnePost(w http.ResponseWriter, r *http.Request) {
	h.Services.PostService.GetOne()
	//GetOne
}
func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	h.Services.PostService.Delete()
	//Delete
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	h.Services.PostService.Update()
	//Delete
}
