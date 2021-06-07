package models

type Gift struct {
	ID           int    `json:"id,omitempty"`
	Wish         int    `json:"wish,omitempty"`
	User         User   `json:"user,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
