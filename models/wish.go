package models

type Wish struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        uint64 `json:"user"`
	Hidden      bool   `json:"hidden"`
}
