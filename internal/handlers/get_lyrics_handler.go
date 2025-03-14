package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/internal/repositories"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// @Summary Get song lyrics
// @Description Returns lyrics for a song
// @Tags lyrics
// @Produce json
// @Param songid query int true "Song ID"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} models.Lyrics
// @Failure 400 {object} string "Song ID is required"
// @Failure 400 {object} string "Invalid arguments"
// @Failure 405 {object} string "Method not allowed"
// @Failure 500 {object} string "Internal server error"
// @Router /lyrics [get]
func HandleGetLyricsRequest(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request to get lyrics")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryParams, err := getLyricsQueryParamsFromRequest(r)
	if err != nil {
		logger.Error("Error parsing params request: %v", err)
		if err == ErrSongIdRequired {
			http.Error(w, "Song ID is required", http.StatusBadRequest)
		} else {
			http.Error(w, "Invalid arguments", http.StatusBadRequest)
		}
		return
	}
	logger.Debug("GetLyrics query params: %+v", queryParams)

	lyrics, err := getLyrics(queryParams)
	if err != nil {
		logger.Error("Error processing request: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	logger.Debug("GetLyrics response: %+v", lyrics)

	err = json.NewEncoder(w).Encode(lyrics)
	if err != nil {
		logger.Error("Error encoding response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Logger.Info("Successfully processed GetLyrics request")
}

func getLyricsQueryParamsFromRequest(r *http.Request) (*repositories.LyricsQueryParams, error) {
	requestParams := r.URL.Query()
	result := repositories.LyricsQueryParams{}

	songIdParam := requestParams.Get("songid")
	if songIdParam == "" {
		return nil, ErrSongIdRequired
	}

	songId, err := strconv.Atoi(songIdParam)
	if err != nil {
		return nil, err
	}
	result.SongId = songId

	offsetParam := requestParams.Get("offset")
	if offsetParam != "" {
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			return nil, err
		}

		result.Offset = offset
	}

	limitParam := requestParams.Get("limit")
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			return nil, err
		}

		result.Limit = limit
	}

	return &result, nil
}

func getLyrics(searchParams *repositories.LyricsQueryParams) ([]*models.Lyrics, error) {
	repo := repositories.NewSongRepo(db.Database())

	lyrics, err := repo.GetLyrics(searchParams)
	if err != nil {
		return nil, err
	}

	return lyrics, nil
}
