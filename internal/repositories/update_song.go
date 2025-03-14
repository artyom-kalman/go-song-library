package repositories

import (
	"errors"
	"fmt"

	"github.com/artyom-kalman/go-song-library/internal/models"
)

func (songRepo *SongRepo) UpdateSong(song *models.UpdateSongRequestBody) (*models.Song, error) {
	if song.Id <= 0 {
		return nil, errors.New("Invalid song ID")
	}

	if song.Name != "" {
		_, err := songRepo.updateSongName(song.Id, song.Name)
		if err != nil {
			return nil, err
		}
	}

	if song.ReleaseDate != "" {
		_, err := songRepo.updateSongReleaseDate(song.Id, song.ReleaseDate)
		if err != nil {
			return nil, err
		}
	}

	if song.Link != "" {
		_, err := songRepo.updateSongLink(song.Id, song.Link)
		if err != nil {
			return nil, err
		}
	}

	updatedSong, err := songRepo.GetSongById(song.Id)
	if err != nil {
		return nil, err
	}

	return updatedSong, nil
}

func (songRepo *SongRepo) updateSongName(songId int, name string) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET name = '%s' WHERE id = %d RETURNING id, name, group_id, release_date, link",
		name, songId,
	)
	fmt.Println(query)

	queryResult, err := songRepo.conn.Query(query)
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

func (songRepo *SongRepo) updateSongReleaseDate(songId int, releaseDate string) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET release_date = '%s' WHERE id = %d RETURNING id, name, group_id, release_date, link",
		releaseDate, songId,
	)

	queryResult, err := songRepo.conn.Query(query)
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

func (songRepo *SongRepo) updateSongLink(songId int, link string) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET link = '%s' WHERE id = %d RETURNING id, name, group_id, release_date, link",
		link, songId,
	)

	queryResult, err := songRepo.conn.Query(query)
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

func (songRepo *SongRepo) updateSongGroupId(songId int, groupId int) (*models.Song, error) {
	query := fmt.Sprintf(
		"UPDATE songs SET group_id = %d WHERE id = %d RETURNING id, name, group_id, release_date, link",
		groupId, songId,
	)

	queryResult, err := songRepo.conn.Query(query)
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
