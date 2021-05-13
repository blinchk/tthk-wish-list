package models

type Like struct {
	ID           int    `json:"id,omitempty"`
	Wish         int    `json:"wish,omitempty"`
	Comment      int    `json:"comment,omitempty"`
	User         User   `json:"user,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
