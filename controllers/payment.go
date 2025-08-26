package controllers

import (
	"net/http"

	"my-gin-app/project/config"
	"my-gin-app/project/models"

	"github.com/gin-gonic/gin"
)

func BuyProduct(c *gin.Context) {
	productID := c.Param("id")
	buyerID := c.GetUint("user_id")

	var req models.BuyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if product.Stock < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient stock"})
		return
	}

	var account models.BankAccount
	if err := config.DB.First(&account, req.BankAccountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bank account not found"})
		return
	}
	if account.UserID != product.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank account not belongs to seller"})
		return
	}

	payment := models.Payment{
		ProductID:            product.ID,
		BuyerID:              buyerID,
		SellerID:             product.UserID,
		BankAccountID:        account.ID,
		PaymentProofImageUrl: req.PaymentProofImageUrl,
		Quantity:             req.Quantity,
	}
	if err := config.DB.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create payment"})
		return
	}

	product.Stock -= req.Quantity
	config.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"message": "payment processed successfully",
		"data":    payment,
	})
}
