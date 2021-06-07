package models

type Gift struct {
	ID           int    `json:"id,omitempty"`
	Wish         Wish   `json:"wish,omitempty"`
	User         User   `json:"user,omitempty"`
	Link         string `json:"link, omtempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
