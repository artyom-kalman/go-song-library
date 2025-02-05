package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetSongInfo(song *models.NewSong) error {
	url := getApiUrl(song.Name, song.Group)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(&song); err != nil {
		return err
	}

	logger.Debug("Added info to song %v", song)
	return nil
}

func getApiUrl(songName string, groupName string) string {
	return fmt.Sprintf("http://localhost:3030/info?group=%s&song=%s", songName, groupName)
}
