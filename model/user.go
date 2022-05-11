package model

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID       int    `json: "id" form: "id"`
	Email    string `json: "email" form: "email"`
	Username string `json: "username" form: "username"`
	Token    string `json: "token" form: "token"`
}
