package models

type Wish struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	User         User   `json:"user"`
	Hidden       bool   `json:"hidden"`
	CreationTime string `json:"creationTime"`
}
