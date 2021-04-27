package models

type Follow struct {
	ID           uint64 `json:"id"`
	UserFrom     string `json:"user_from"`
	UserTo       string `json:"user_to"`
	CreationTime string `json:"creation_time"`
}
