package model

type Room struct {
	Id    int    `gorm:"primaryKey" json:"id" form:"id"`
	Type  string `json:"type" form:"type"`
	Price string `json:"price" form:"type"`
}

func (Room) TableName() string {
	return "room"
}
