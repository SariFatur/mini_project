package model

type Room struct {
	IdRoom int    `json:"idRoom" form:"idRoom"`
	Type   string `json:"type" form:"type"`
	Price  int    `json:"price" form:"type"`
}

func (Room) TableName() string {
	return "rooms"
}
