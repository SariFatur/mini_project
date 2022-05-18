package model

import "time"

type Reservation struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	IdGuest       int       `json:"-"`
	IdRoom        int       `json:"-"`
	Guest         Guest     `json:"idGuest" gorm:"foreignKey:IdGuest;references:id"`
	Room          Room      `json:"idRoom" gorm:"foreignKey:IdRoom;references:id"`
	Checkin_date  time.Time `json:"checkin_date"`
	Checkout_date time.Time `json:"checout_date"`
}

func (Reservation) TableName() string {
	return "reservation"
}
