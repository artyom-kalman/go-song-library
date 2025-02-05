package main

import (
	"fmt"
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/handlers"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
)

// TODO

// Все поля при добавлении песни
//
// Endpoint на изменение песни

func main() {
	logger.InitLogger()

	err := config.LoadConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	serverConfig, err := config.GetServerConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	databaseConfig, err := config.GetDBConfig()
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	err = db.InitDatabase(databaseConfig)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	err = db.RunMigration()
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}

	http.HandleFunc("/song", handlers.SongHandler)
	http.HandleFunc("/songs", handlers.GetSongsHandler)
	http.HandleFunc("/lyrics", handlers.GetLyricsHandler)

	logger.Logger.Info(fmt.Sprintf("Server is running on %s", serverConfig.Port))

	err = http.ListenAndServe(serverConfig.Port, nil)
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}
