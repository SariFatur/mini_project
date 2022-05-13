package model

type payment_method struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
