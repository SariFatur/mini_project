package model

type Payment_method struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
