package model

type Guest struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type GuestResponse struct {
	Id    int    `json: "id" form: "id"`
	Email string `json: "email" form: "email"`
	Name  string `json: "name" form: "name"`
	Token string `json: "token" form: "token"`
}

type GuestCustom struct { // table user_custom
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (Guest) TableName() string {
	return "guest"
}
