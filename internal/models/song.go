package models

type Song struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	GroupId     int    `json:"groupId"`
	GroupName   string `json:"groupName"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}
