package models

type Gift struct {
	ID           int    `json:"id,omitempty"`
	Wish         int    `json:"wish,omitempty"`
	Follow       int    `json:"follow,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
