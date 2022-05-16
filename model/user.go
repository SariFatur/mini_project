package model

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `json: username`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UserCustom struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (UserCustom) TableName() string {
	return "users"
}
