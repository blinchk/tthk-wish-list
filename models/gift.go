package models

type Gift struct {
	ID           int    `json:"id,omitempty"`
	Wish         Wish   `json:"wish"`
	User         User   `json:"user"`
	Title        string `json:"title,omitempty"`
	Link         string `json:"link,omitempty"`
	Booked       bool   `json:"booked,omitempty"`
	UserBooked   User   `json:"user_booked"`
	CreationTime string `json:"creation_time,omitempty"`
}
