package model

type Transactions struct {
	Id               int            `gorm: "primaryKey" json: "id"`
	Total_price      string         `json: "total_price"`
	Status           string         `json: "status"`
	IdPayment_method int            `json:"-"`
	IdReservation    int            `json:"-"`
	Payment_method   Payment_method `json:"payment_method" gorm:"foreignKey:IdPayment_method;references:Id"`
	Reservation      Reservation    `json:"reservation" gorm:"foreignKey:IdReservation;references:Id"`
}

func (Transactions) TableName() string {
	return "transactions"
}
