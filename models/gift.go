package models

type Gift struct {
	ID           int    `json:"id,omitempty"`
	Wish         Wish   `json:"wish"`
	User         User   `json:"user"`
	Title        string `json:"title,omitempty"`
	Link         string `json:"link,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
