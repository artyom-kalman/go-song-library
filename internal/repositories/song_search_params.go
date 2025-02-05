package repositories

type SongQueryParams struct {
	SongId    int
	SongName  string
	GroupId   int
	GroupName string
	StartDate string
	EndDate   string
	Offset    int
	Limit     int
}

func NewSongQueryParams() *SongQueryParams {
	return &SongQueryParams{
		SongId:  -1,
		GroupId: -1,
		Offset:  -1,
		Limit:   -1,
	}
}
