package models

type Song struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	GroupId int    `json:"groupId"`
}
