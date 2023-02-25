package models

type User struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Photo           string `json:"photo"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}
