package models

type Suggestion struct {
	Wishes       []Wish `json:"wishes,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
