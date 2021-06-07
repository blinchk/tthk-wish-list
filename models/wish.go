package models

type Wish struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	User         User   `json:"user,omitempty"`
	Hidden       bool   `json:"hidden,omitempty"`
	Liked        bool   `json:"liked"`
	Likes        int    `json:"likes,omitempty"`
	CreationTime string `json:"creationTime,omitempty"`
}
