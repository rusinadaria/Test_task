package repository

import (
	"Test_task/models"
	"database/sql"
)

type SongRepository interface {
	CreateSong(song models.Song) error
    GetAll(filter models.Song, lastId string, limit int) ([]*models.Song, error)
	GetText(id string) (string, error)
    UpdateSong(id string, song models.Song) error
    DeleteSong(id string) error
	// findSong(song models.Song) models.Song
}

type Repository struct {
	SongRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		SongRepository:  NewPostgresSongRepo(db),
	}
}