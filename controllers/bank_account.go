package controllers

import (
    "net/http"

    "my-gin-app/project/config"
    "my-gin-app/project/models"

    "github.com/gin-gonic/gin"
)

func CreateBankAccount(c *gin.Context) {
    var req models.CreateBankAccountRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("user_id")

    var existing models.BankAccount
    if err := config.DB.Where("user_id = ? AND bank_account_number = ?", userID, req.BankAccountNumber).First(&existing).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "account number already exists"})
        return
    }

    account := models.BankAccount{
        BankName:          req.BankName,
        BankAccountName:   req.BankAccountName,
        BankAccountNumber: req.BankAccountNumber,
        UserID:            userID,
    }

    if err := config.DB.Create(&account).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "account added successfully", "data": account})
}

func ListBankAccounts(c *gin.Context) {
    userID := c.GetUint("user_id")
    var accounts []models.BankAccount

    if err := config.DB.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "success", "data": accounts})
}

func UpdateBankAccount(c *gin.Context) {
    var req models.UpdateBankAccountRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("user_id")
    id := c.Param("id")

    var account models.BankAccount
    if err := config.DB.First(&account, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
        return
    }

    if account.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    updates := map[string]interface{}{}
    if req.BankName != "" {
        updates["bank_name"] = req.BankName
    }
    if req.BankAccountName != "" {
        updates["bank_account_name"] = req.BankAccountName
    }
    if req.BankAccountNumber != "" {
        updates["bank_account_number"] = req.BankAccountNumber
    }

    if len(updates) > 0 {
        if err := config.DB.Model(&account).Updates(updates).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "account updated successfully", "data": account})
}


func DeleteBankAccount(c *gin.Context) {
    userID := c.GetUint("user_id")
    id := c.Param("id")

    var account models.BankAccount
    if err := config.DB.First(&account, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
        return
    }

    if account.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    if err := config.DB.Delete(&account).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
}
