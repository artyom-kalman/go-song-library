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

// @Summary Gets song library with filters
// @Description Gets songs by name, id, release date, etc.
// @Tags songs
// @Produce json
// @Param songid query int false "Song ID"
// @Param songname query string false "Song Name"
// @Param groupid query int false "Group ID"
// @Param groupname query string false "Group Name"
// @Param releasedate-start query string false "Release date start in YYYY-MM-DD format"
// @Param releasedate-end query string false "Release date end in YYYY-MM-DD format"
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Success 200 {array} models.Song
// @Failure 400 {object} string "Invalid query parameters"
// @Failure 405 {object} string "Method not allowed"
// @Failure 500 {object} string "Internal server error"
// @Router /songs [get]
func HandleGetSongRequest(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received getsongs request")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	searchParams, err := getSongQueryParams(r)
	if err != nil {
		logger.Error("Error parsing query parameters: %v", err)
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}
	logger.Debug("Getsong parameters: %+v", searchParams)

	songs, err := getSongs(searchParams)
	if err != nil {
		logger.Error("Error getting songs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(songs); err != nil {
		logger.Error("Error encoding songs: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	logger.Info("Successfully served getsongs request")
}

func getSongQueryParams(r *http.Request) (*repositories.SongQueryParams, error) {
	queryParams := r.URL.Query()
	searchParams := repositories.NewSongQueryParams()

	searchParams.SongName = queryParams.Get("songname")
	searchParams.GroupName = queryParams.Get("groupname")
	searchParams.StartDate = queryParams.Get("releasedate-start")
	searchParams.EndDate = queryParams.Get("releasedate-end")

	groupIdParam := queryParams.Get("groupid")
	if groupIdParam != "" {
		groupId, err := strconv.Atoi(groupIdParam)
		if err != nil {
			return nil, err
		}
		searchParams.GroupId = groupId
	}

	songIdParam := queryParams.Get("songid")
	if songIdParam != "" {
		songId, err := strconv.Atoi(songIdParam)
		if err != nil {
			return nil, err
		}
		searchParams.SongId = songId
	}

	offsetParam := queryParams.Get("offset")
	if offsetParam != "" {
		offset, err := strconv.Atoi(offsetParam)
		if err != nil {
			return nil, err
		}
		searchParams.Offset = offset
	}

	limitParam := queryParams.Get("limit")
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			return nil, err
		}
		searchParams.Limit = limit
	}

	return searchParams, nil
}

func getSongs(searchParams *repositories.SongQueryParams) ([]*models.Song, error) {
	songRepo := repositories.NewSongRepo(db.Database())

	songs, err := songRepo.GetSongs(searchParams)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
