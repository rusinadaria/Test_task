package handler

import (
	"github.com/go-chi/chi/v5"
	"Test_task/internal/service"
	docs "Test_task/docs"
	"github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.Get("/swagger/*", httpSwagger.WrapHandler) 

    router.Route("/", func(r chi.Router) {
		router.Post("/songs", h.AddSong)
		router.Get("/songs", h.GetSongs)
		router.Get("/songs/{id}/verses", h.SongVerse)
		router.Patch("/songs/{id}", h.EditSong)
		router.Delete("/songs/{id}", h.DeleteSong)
    })
	return router
}

