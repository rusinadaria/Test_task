package service

import (
	_ "github.com/lib/pq"
	"Test_task/models"
	"Test_task/repository"
	"strconv"
	"strings"
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

func (s *SongService) GetVerse(id string, limit int, offset int) ([]models.Verse, error) {
	songText, err := s.repo.GetText(id)
	if err != nil {
		return nil, err
	}
	verse := splitVerse(songText, limit)

	if offset >= len(verse) {
		verse = []models.Verse{}
	} else if offset+limit > len(verse) {
		verse = verse[offset:]
	} else {
		verse = verse[offset : offset+limit]
	}
	return verse, nil
}

func splitVerse(text string, limit int) []models.Verse {
	parts := strings.Split(text, "/br")
	verses := make([]models.Verse, 0)

	for i, part := range parts {
		if i >= limit {
			break
		}
		verses = append(verses, models.Verse{Number: i + 1, Text: strings.TrimSpace(part)})
	}
	return verses
}

func (s *SongService) UpdateSong(id string, song models.Song) error {
	return s.repo.UpdateSong(id, song)
}

func (s *SongService) DeleteSong(id string) error {
	return s.repo.DeleteSong(id)
}
