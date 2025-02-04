package repositories

type SongSearchParams struct {
	SongId    int
	SongName  string
	GroupId   int
	GroupName string
	StartDate string
	EndDate   string
	Offset    int
	Limit     int
}

func NewSongSearchParams() *SongSearchParams {
	return &SongSearchParams{
		SongId:  -1,
		GroupId: -1,
		Offset:  -1,
		Limit:   -1,
	}
}
