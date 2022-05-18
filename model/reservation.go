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

// type Jadwal struct {
// 	Id           int        `gorm:"column:id" json:"id"`
// 	Day          string     `gorm:"column:hari" json:"hari"`
// 	Time         string     `gorm:"column:waktu" json:"time"`
// 	IdMatakuliah int        `json:"-"`
// 	NidDosen     int        `json:"-"`
// 	IdRuangan    int        `json:"-"`
// 	Matakuliah   Matakuliah `json:"matakuliah" gorm:"foreignKey:IdMatakuliah;references:ID"`
// 	Dosen        Dosen      `json:"dosen" gorm:"foreignKey:NidDosen;references:NID"`
// 	Ruangan      Ruangan    `json:"ruangan" gorm:"foreignKey:IdRuangan;references:ID"`
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	DeletedAt    gorm.DeletedAt `gorm:"index"`
// }
