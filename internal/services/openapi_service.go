package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/internal/models"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

func GetSongInfo(song *models.NewSong) error {
	url := getApiUrl(song.Name, song.Group)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error getting song info")
	}

	if err := json.NewDecoder(resp.Body).Decode(&song); err != nil {
		return err
	}

	logger.Debug("Added info to song %v", song)
	return nil
}

func getApiUrl(songName string, groupName string) string {
	baseUrl := fmt.Sprintf("%s/info", config.GetOpenAPIURL())
	params := url.Values{}
	params.Add("song", songName)
	params.Add("group", groupName)

	return fmt.Sprintf("%s?%s", baseUrl, params.Encode())
}
