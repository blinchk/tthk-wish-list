package models

type User struct {
	ID          uint64 `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username" validate:"required,min=2,max=49,email"`
	Password    string `json:"password" validate:"required,min=2,max=50"`
	AccessToken string `json:"access_token"`
}
