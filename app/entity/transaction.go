package entity

import "time"

type Transaction struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Customer  string    `gorm:"type:varchar(100);default:''" json:"customer"`
	Quantity  uint64    `gorm:"type:bigint(20);default:0" json:"quantity"`
	Price     float64   `gorm:"type:double;default:0" json:"price"`
	Timestime time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"timestime"`
}
