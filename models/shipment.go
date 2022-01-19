package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	From   string  `json:"from"`
	To     string  `json:"to"`
	Price  float64 `json:"price"`
}