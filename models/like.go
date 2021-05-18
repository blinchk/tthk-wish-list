package models

type Like struct {
	ID             int    `json:"id,omitempty"`
	Connection     int    `json:"connection,omitempty"`
	ConnectionType string `json:"connection_type,omitempty"`
	User           User   `json:"user,omitempty"`
	CreationTime   string `json:"creation_time,omitempty"`
}
