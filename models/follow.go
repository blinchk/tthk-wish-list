package models

type Follow struct {
	ID           uint64 `json:"id"`
	UserFrom     uint64 `json:"user_from"`
	UserTo       uint64 `json:"user_to"`
	CreationTime string `json:"creation_time"`
}
