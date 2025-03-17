package repositories

import (
	"context"
	"fmt"
)

func (songRepo *SongRepo) IsGroupExist(groupName string, ctx context.Context) bool {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM groups WHERE name = '%s');", groupName)

	queryResult, err := songRepo.conn.QueryContext(ctx, query)
	if err != nil {
		return false
	}

	var exists bool
	queryResult.Next()
	queryResult.Scan(&exists)

	return exists
}
