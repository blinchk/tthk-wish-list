package models

type Comment struct {
	ID           int    `json:"id,omitempty"`
	Content      string `json:"content,omitempty"`
	Wish         int    `json:"wish,omitempty"`
	User         User   `json:"user,omitempty"`
	Parent       int    `json:"comment,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
