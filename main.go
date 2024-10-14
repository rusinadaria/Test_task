package main

import (
	"Test_task/repository"
	// "fmt"
	"Test_task/internal/service"
	"log/slog"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
)

// @title Songs API
// @version 1.0
// @description This is a sample API for managing songs.
// @host localhost:8080
// @BasePath /songs
func main() {
	// godotenv.Load()
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("Unable to open a file for writing")
	}

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(f, opts))
	logger.Info("Info message")

	repository.ConnectDatabase(logger)

	router := chi.NewRouter()
	router.Post("/songs", service.SongsAddPost) // Create a new song
	router.Get("/songs", service.SongsGet) // Get all songs
	router.Delete("/songs/{id}", service.SongsIdDelete) // Delete a song by ID
	router.Patch("/songs/{id}", service.SongsIdEditPatch) // Update a song by ID
	router.Get("/songs/{id}/couplets", service.SongsIdCoupletGet) // Get couplets of a song by ID

	err = http.ListenAndServe(os.Getenv("PORT"), router)
	if err != nil {
		logger.Error("failed start server")
		panic(err)
	}
}
