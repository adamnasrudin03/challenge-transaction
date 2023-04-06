package entity

type Transaction struct {
	ID        uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Customer  string  `gorm:"type:varchar(100);" json:"customer"`
	Quantity  uint64  `gorm:"default:0" json:"quantity"`
	Price     float64 `gorm:"default:0" json:"price"`
	Timestime string  `gorm:"type:varchar(100)" json:"timestime"`
}
