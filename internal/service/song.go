package service

import (
	_ "github.com/lib/pq"
	"Test_task/models"
	"Test_task/repository"
	"strconv"
)

type SongService struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) *SongService{
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(song models.Song) error {
	return s.repo.CreateSong(song)
}

func (s *SongService) GetAll(filter models.Song, last string, limit string) ([]*models.Song, error) {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	return s.repo.GetAll(filter, last, limitInt)
}

func (s *SongService) GetCouplet(id string) (string, error) {
	return s.repo. GetText(id)
}

func (s *SongService) UpdateSong(id string, song models.Song) error {
	return s.repo.UpdateSong(id, song)
}

func (s *SongService) DeleteSong(id string) error {
	return s.repo.DeleteSong(id)
}
