package models

type Pool struct {
	ID     uint64 `json:"id"`
	UserID string `json:"userid"`
	WishID string `json:"wishid"`
}
