package service

import (
	"Test_task/models"
	"Test_task/repository"
)

type SongRepo interface {
	CreateSong(song models.Song) error
    GetAll(filter models.Song, last string, limit string) ([]*models.Song, error)
	GetVerse(id string, limit int, offset int) ([]models.Verse, error)
    UpdateSong(id string, song models.Song) error
    DeleteSong(id string) error
}

type Service struct {
	SongRepo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SongRepo:  NewSongService(repos.SongRepository),
	}
}