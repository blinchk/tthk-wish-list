package models

type Follow struct {
	ID           int    `json:"id,omitempty"`
	UserFrom     int    `json:"user_from,omitempty"`
	UserTo       int    `json:"user_to,omitempty"`
	CreationTime string `json:"creation_time,omitempty"`
}
