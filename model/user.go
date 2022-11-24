package model

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
