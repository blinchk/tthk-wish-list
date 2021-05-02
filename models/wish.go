package models

type Wish struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        User   `json:"user"`
	Hidden      bool   `json:"hidden"`
}
