package models

type User struct {
	ID               uint64 `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	AccessToken      string `json:"access_token"`
	RegistrationTime string `json:"registration_time"`
}
