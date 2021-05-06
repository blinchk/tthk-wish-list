package models

type Follow struct {
	ID           int    `json:"id,omitempty"`
	UserFrom     int    `json:"user_from"`
	UserTo       int    `json:"user_to"`
	CreationTime string `json:"creation_time"`
}
