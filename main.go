package main

import (
	"net/http"

	"github.com/artyom-kalman/go-song-library/internal/config"
	"github.com/artyom-kalman/go-song-library/internal/db"
	"github.com/artyom-kalman/go-song-library/internal/handlers"
	"github.com/artyom-kalman/go-song-library/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/artyom-kalman/go-song-library/docs"
)

// TODO
// 1. Миграции не реализованы. Механизмы миграции подразумевают хранение информации о том, что уже было применено. Удаление schema_migrations при старте это не позволяет

func init() {
	logger.InitLogger()
}

// @title Song Library API
// @version 1.0
// @description API for managing a song library

// @host localhost:3030
// @BasePath /
// @schemes http
func main() {
	err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	serverConfig, err := config.GetServerConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	databaseConfig, err := config.GetDBConfig()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = db.ConnectToDatabase(databaseConfig)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = db.RunMigration()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	http.HandleFunc("/song", handlers.HandleSongRequest)
	http.HandleFunc("/songs", handlers.HandleGetSongRequest)
	http.HandleFunc("/lyrics", handlers.HandleGetLyricsRequest)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	logger.Info("Server is running on %s", serverConfig.Port)

	err = http.ListenAndServe(serverConfig.Port, nil)
	if err != nil {
		logger.Error(err.Error())
	}
}
