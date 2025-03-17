package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (songRepo *SongRepo) UpdateSong(song *models.UpdateSongRequestBody, ctx context.Context) (*models.Song, error) {
	if song.Id <= 0 {
		return nil, errors.New("Invalid song ID")
	}

	if song.Name != "" {
		_, err := songRepo.updateSongName(song.Id, song.Name, ctx)
		if err != nil {
			return nil, err
		}
	}

	if song.ReleaseDate != "" {
		_, err := songRepo.updateSongReleaseDate(song.Id, song.ReleaseDate, ctx)
		if err != nil {
			return nil, err
		}
	}

	if song.Link != "" {
		_, err := songRepo.updateSongLink(song.Id, song.Link, ctx)
		if err != nil {
			return nil, err
		}
	}

	updatedSong, err := songRepo.GetSongById(song.Id, ctx)
	if err != nil {
		return nil, err
	}

	return updatedSong, nil
}

func (songRepo *SongRepo) updateSongName(songId int, name string, ctx context.Context) (*models.Song, error) {
	return songRepo.updateSongField(songId, "name", name, ctx)
}

func (songRepo *SongRepo) updateSongReleaseDate(songId int, releaseDate string, ctx context.Context) (*models.Song, error) {
	return songRepo.updateSongField(songId, "releaseDate", releaseDate, ctx)
}

func (songRepo *SongRepo) updateSongLink(songId int, link string, ctx context.Context) (*models.Song, error) {
	return songRepo.updateSongField(songId, "link", link, ctx)
}

func (songRepo *SongRepo) updateSongGroupId(songId int, groupId int, ctx context.Context) (*models.Song, error) {
	return songRepo.updateSongField(songId, "groupId", string(groupId), ctx)
}

func (songRepo *SongRepo) updateSongField(songId int, fieldName string, fieldValue string, ctx context.Context) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET %s = '%s' WHERE id = %d RETURNING id, name, group_id, release_date, link",
		fieldName, fieldValue, songId,
	)

	queryResult, err := songRepo.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	isSongUpdated := queryResult.Next()
	if !isSongUpdated {
		return nil, ErrSongNotFound
	}

	var updatedSong models.Song
	queryResult.Scan(
		&updatedSong.Id,
		&updatedSong.Name,
		&updatedSong.GroupId,
		&updatedSong.ReleaseDate,
		&updatedSong.Link,
	)

	return &updatedSong, nil
}
