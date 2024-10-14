package service

import (
	"Test_task/models"
	"Test_task/repository"
)

type SongRepo interface {
	CreateSong(song models.Song) error
    GetAll(filter models.Song, last string, limit string) ([]*models.Song, error)
	GetCouplet(id string) (string, error)
    UpdateSong(id string, song models.Song) error
    DeleteSong(id string) error
	// findSong(song models.Song) models.Song
}

type Service struct {
	SongRepo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SongRepo:  NewSongService(repos.SongRepository),
	}
}