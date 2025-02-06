package repositories

import "fmt"

func (songRepo *SongRepo) IsGroupExist(groupName string) bool {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM songs WHERE group_name = '%s');", groupName)

	var exists bool
	queryResult, err := songRepo.conn.Query(query)
	if err != nil {
		return false
	}

	queryResult.Next()
	queryResult.Scan(&exists)

	return exists
}
