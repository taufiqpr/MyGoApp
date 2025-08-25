package models

import "time"

type ProductCondition string

const (
    ConditionNew    ProductCondition = "new"
    ConditionSecond ProductCondition = "second"
)

type Product struct {
    ID             uint              `gorm:"primaryKey" json:"id"`
    Name           string            `gorm:"size:60;not null" json:"name"`
    Price          int               `gorm:"not null" json:"price"`
    ImageURL       string            `gorm:"not null" json:"imageUrl"`
    Stock          int               `gorm:"not null" json:"stock"`
    Condition      ProductCondition  `gorm:"type:varchar(10);not null" json:"condition"`
    Tags           string            `gorm:"type:text" json:"tags"`
    IsPurchaseable bool              `gorm:"not null" json:"isPurchaseable"`
    UserID         uint              `gorm:"not null" json:"userId"`
    CreatedAt      time.Time         `json:"createdAt"`
    UpdatedAt      time.Time         `json:"updatedAt"`
}

type CreateProductRequest struct {
    Name           string            `json:"name" binding:"required,min=5,max=60"`
    Price          int               `json:"price" binding:"required,min=0"`
    ImageURL       string            `json:"imageUrl" binding:"required,url"`
    Stock          int               `json:"stock" binding:"required,min=0"`
    Condition      ProductCondition  `json:"condition" binding:"required,oneof=new second"`
    Tags           string          `json:"tags"`
    IsPurchaseable bool              `json:"isPurchaseable" binding:"required"`
}

type UpdateProductRequest struct {
    Name           string            `json:"name" binding:"omitempty,min=5,max=60"`
    Price          *int              `json:"price" binding:"omitempty,min=0"`
    ImageURL       string            `json:"imageUrl" binding:"omitempty,url"`
    Stock          *int              `json:"stock" binding:"omitempty,min=0"`
    Condition      ProductCondition  `json:"condition" binding:"omitempty,oneof=new second"`
    Tags           string          `json:"tags"`
    IsPurchaseable *bool             `json:"isPurchaseable"`
}

type UpdateStockRequest struct {
    Stock int `json:"stock" binding:"required,min=0"`
}

