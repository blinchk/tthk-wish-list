package models

type User struct {
	ID               uint64 `json:"id,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	RegistrationTime string `json:"registration_time,omitempty"`
}
