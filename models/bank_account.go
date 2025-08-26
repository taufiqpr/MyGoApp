package models

import "time"

type BankAccount struct {
    ID                uint      `gorm:"primaryKey" json:"bankAccountId"`
    BankName          string    `gorm:"size:15;not null" json:"bankName"`
    BankAccountName   string    `gorm:"size:50;not null" json:"bankAccountName"`
    BankAccountNumber string    `gorm:"uniqueIndex:idx_user_bank_number"`
    UserID            uint      `gorm:"not null" json:"userId"`
    CreatedAt         time.Time `json:"createdAt"`
    UpdatedAt         time.Time `json:"updatedAt"`
}

type CreateBankAccountRequest struct {
    BankName          string `json:"bankName" binding:"required,min=3,max=15"`
    BankAccountName   string `json:"bankAccountName" binding:"required,min=3,max=50"`
    BankAccountNumber string `json:"bankAccountNumber" binding:"required,min=5,max=30"`
}

type UpdateBankAccountRequest struct {
    BankName          string `json:"bankName" binding:"omitempty,min=3,max=15"`
    BankAccountName   string `json:"bankAccountName" binding:"omitempty,min=3,max=50"`
    BankAccountNumber string `json:"bankAccountNumber" binding:"omitempty,min=5,max=30"`
}
