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

    if req.BankName != "" {
        account.BankName = req.BankName
    }
    if req.BankAccountName != "" {
        account.BankAccountName = req.BankAccountName
    }
    if req.BankAccountNumber != "" {
        account.BankAccountNumber = req.BankAccountNumber
    }

    if err := config.DB.Save(&account).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
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
