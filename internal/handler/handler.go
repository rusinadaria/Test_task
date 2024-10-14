package handler

import (
	"github.com/go-chi/chi/v5"
	"Test_task/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()

    router.Route("/", func(r chi.Router) {
		router.Post("/songs", h.AddSong)
		router.Get("/songs", h.GetSongs)
		router.Get("/songs/{id}/couplets", h.SongCouplets)
		router.Patch("/songs/{id}", h.EditSong)
		router.Delete("/songs/{id}", h.DeleteSong)
    })
	return router
}

