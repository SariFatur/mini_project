package model

import "time"

type Reservation struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	IdGuest  int       `gorm:"foreignKey: GuestID" json:"idGuest"`
	IdRoom   int       `gorm:"foreignKey: RoomID" json:"idRoom"`
	Checkin  time.Time `json:"checkin"`
	Checkout time.Time `json:"checout"`
}

func (Reservation) TableName() string {
	return "reservation"
}
