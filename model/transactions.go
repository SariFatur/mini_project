package model

type Transactions struct {
	Id               int    `gorm: "primaryKey" json: "id"`
	Total_price      string `json: "total_price"`
	Status           string `json: "status"`
	IdPayment_method int    `gorm: "foreignKey: IdPayment_method" json: "idPayment_method"`
	IdReservation    int    `gorm: "foreignKey: IdReservation" json: "idReservsation"`
}

func (Transactions) TableName() string {
	return "transactions"
}
