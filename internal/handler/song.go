package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"Test_task/models"
	// "fmt"
)

// AddSongs добавление новой песни.
// @Summary Add a new song
// @Description Add a new song to the database
// @Accept json
// @Param song body models.Song true "Song data"
// @Success 201 {object}  nil "OK"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /songs [post]
func (h *Handler) AddSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var newSong models.Song

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal([]byte(body), &newSong)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// fmt.Println(newSong.Song)
	// fmt.Println(newSong.GroupName)

	err = h.services.CreateSong(newSong)
	if err != nil {
		// добавить проверку на 404 код
		http.Error(w, "Error while adding song", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetSongs получение списка песен с фильтрацией.
// @Summary Get songs
// @Description Get a list of songs with optional filters
// @Produce json
// @Param song query string false "Song name"
// @Param group_name query string false "Group name"
// @Param release_date query string false "Release date"
// @Param text query string false "Text"
// @Param link query string false "Link"
// @Param last_id query string false "Last ID"
// @Param limit query int false "Limit"
// @Success 200 {array} models.Song
// @Failure 500 {object} models.ErrorResponse
// @Router /songs [get]
func (h *Handler) GetSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := []string{"song", "group_name", "release_date", "text", "link", "last_id", "limit"}
	values := make(map[string]string)

	for _, param := range params {
		values[param] = r.URL.Query().Get(param)
	}

	filter := models.Song{
		Song:        values["song"],
		GroupName:   values["group_name"],
		ReleaseDate: values["release_date"],
		Text:        values["text"],
		Link:        values["link"],
	}

	songs, err := h.services.GetAll(filter, values["last_id"], values["limit"])
	if err != nil {
		http.Error(w, "Error while adding song", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(songs); err != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
		return
	}
}

// SongCouplets получение куплета песни по ID.
// @Summary Get song couplets by ID
// @Description Get couplets of a song by its ID
// @Produce json
// @Param id path string true "Song ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Couplet
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /songs/{id}/couplets [get]
func (h *Handler) SongCouplets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := chi.URLParam(r, "id") //спрятать роутер
	if id == "" {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
	}

	offsetStr := r.URL.Query().Get("offset")
	offset := 0
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			http.Error(w, "Invalid offset number", http.StatusBadRequest)
			return
		}
	}

	songText, err := h.services.GetCouplet(id)
	if err != nil {
		http.Error(w, "Error while retrieving song text", http.StatusInternalServerError)
		return
	}

	couplet := splitCouplets(songText, limit)

	if offset >= len(couplet) {
		couplet = []models.Couplet{}
	} else if offset+limit > len(couplet) {
		couplet = couplet[offset:]
	} else {
		couplet = couplet[offset : offset+limit]
	}

	if err := json.NewEncoder(w).Encode(couplet); err != nil {
		http.Error(w, "Error while encoding JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func splitCouplets(text string, limit int) []models.Couplet {
	parts := strings.Split(text, "/br")
	verses := make([]models.Couplet, 0)

	for i, part := range parts {
		if i >= limit {
			break
		}
		verses = append(verses, models.Couplet{Number: i + 1, Text: strings.TrimSpace(part)})
	}
	return verses
}

// EditSong обновление информации о песне.
// @Summary Update song by ID
// @Description Update song by ID
// @Produce json
// @Param id path string true "Song ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /songs/{id} [patch]
func (h *Handler) EditSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := chi.URLParam(r, "id") //спрятать роутер
	if id == "" {
		http.Error(w, "Invalid song ID", http.StatusBadRequest)
		return
	}

	var updatedSong models.Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.services.UpdateSong(id, updatedSong)
	if err != nil {
		http.Error(w, "Error while updating song", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteSong удаление песни по ID.
// @Summary Delete song by ID
// @Description Delete song by ID
// @Produce json
// @Param id path string true "Song ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /songs/{id} [delete]
func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := chi.URLParam(r, "id") //спрятать роутер
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := h.services.DeleteSong(id)
	if err != nil {
		// проверка на 404 код
		http.Error(w, "Error while adding song", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
