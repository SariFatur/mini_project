package model

import "time"

type Reservation struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	Id_guest      int       `json:"-"`
	Id_room       int       `json:"-"`
	Guest         Guest     `json:"idGuest" gorm:"foreignKey:Id_guest;references:id"`
	Room          Room      `json:"idRoom" gorm:"foreignKey:Id_room;references:id"`
	Checkin_date  time.Time `json:"checkin_date"`
	Checkout_date time.Time `json:"checout_date"`
}

func (Reservation) TableName() string {
	return "reservation"
}
