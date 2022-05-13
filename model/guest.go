package model

type Guest struct {
	IdGuest      int    `json:"idGuest"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone_number int    `json:"phone_number"`
	Email        string `json:"email"`
}

func (Guest) TableName() string {
	return "guests"
}
