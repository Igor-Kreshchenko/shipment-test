package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	SenderName           string  `json:"senderName" validate:"required,max=30,alpha"`
	SenderEmail          string  `json:"senderEmail"  validate:"required,email"`
	SenderAddress        string  `json:"senderAddress"  validate:"required,max=100"`
	SenderCountryCode    string  `json:"senderCountryCode" validate:"required,alpha"`
	RecipientName        string  `json:"recipientName" validate:"required,max=30,alpha"`
	RecipientEmail       string  `json:"recipientEmail" validate:"required,email"`
	RecipientAddress     string  `json:"recipientAddress" validate:"required,max=100"`
	RecipientCountryCode string  `json:"recipientCountryCode" validate:"required,alpha"`
	Weight               float64 `json:"weight" validate:"required,max=1000"`
	Price                float64 `json:"price"`
}
