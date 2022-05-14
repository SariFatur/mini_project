package model

type Guest struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func (Guest) TableName() string {
	return "guest"
}
