package models

type User struct {
	ID               int    `json:"id,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	Email            string `json:"email,omitempty"`
	Password         string `json:"password,omitempty"`
	AccessToken      string `json:"access_token,omitempty"`
	RegistrationTime string `json:"registration_time,omitempty"`
}

func (user *User) DeleteSensitiveData() {
	user.Email = ""
	user.AccessToken = ""
	user.Password = ""
	user.RegistrationTime = ""
}
