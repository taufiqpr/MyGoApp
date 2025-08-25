package models

import "time"

type Payment struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	ProductID            uint      `gorm:"not null" json:"productId"`
	BuyerID              uint      `gorm:"not null" json:"buyerId"`
	SellerID             uint      `gorm:"not null" json:"sellerId"`
	BankAccountID        uint      `gorm:"not null" json:"bankAccountId"`
	PaymentProofImageUrl string    `gorm:"not null" json:"paymentProofImageUrl"`
	Quantity             int       `gorm:"not null" json:"quantity"`
	CreatedAt            time.Time `json:"createdAt"`
}

type BuyRequest struct {
	BankAccountID        uint   `json:"bankAccountId" binding:"required"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl" binding:"required,url"`
	Quantity             int    `json:"quantity" binding:"required,min=1"`
}
