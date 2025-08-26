package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"my-gin-app/project/config"
	"my-gin-app/project/models"

	"github.com/gin-gonic/gin"
)

func Healths(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func CreateProduct(c *gin.Context) {
    var req models.CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("user_id")

    product := models.Product{
        Name:           req.Name,
        Price:          req.Price,
        ImageURL:       req.ImageURL,
        Stock:          req.Stock,
        Condition:      req.Condition,
        Tags:           req.Tags,
        IsPurchaseable: req.IsPurchaseable,
        UserID:         userID,
    }

    if err := config.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "product added successfully", "data": product})
}

func UpdateProduct(c *gin.Context) {
    var req models.UpdateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var product models.Product
    if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
        return
    }

    updates := map[string]interface{}{}

    if req.Name != "" {
        updates["name"] = req.Name
    }
    if req.Price != nil {
        updates["price"] = *req.Price
    }
    if req.ImageURL != "" {
        updates["image_url"] = req.ImageURL
    }
    if req.Stock != nil {
        updates["stock"] = *req.Stock
    }
    if req.Condition != "" {
        updates["condition"] = req.Condition
    }
    if req.Tags != "" {
        updates["tags"] = req.Tags
    }
    if req.IsPurchaseable != nil {
        updates["is_purchaseable"] = *req.IsPurchaseable
    }

    if len(updates) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
        return
    }

    if err := config.DB.Model(&product).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "product updated successfully", "data": product})
}

func DeleteProduct(c *gin.Context) {
    result := config.DB.Delete(&models.Product{}, c.Param("id"))

    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
        return
    }

    if result.RowsAffected == 0 {
        
        c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

func UpdateStock(c *gin.Context) {
    var req models.UpdateStockRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload", "detail": err.Error()})
        return
    }

    userID := c.GetUint("user_id")
    productID := c.Param("id")

    var product models.Product
    if err := config.DB.First(&product, productID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
        return
    }

    
    if product.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
        return
    }

    product.Stock = req.Stock
    if err := config.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update stock"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "stock updated successfully"})
}

func ListProductsWithFilter(c *gin.Context) {
	
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	tags := c.QueryArray("tags")
	search := c.Query("search")
	condition := c.Query("condition")
	showEmptyStock := c.Query("showEmptyStock") == "true"

	var products []models.Product
	query := config.DB.Model(&models.Product{})

	if len(tags) > 0 {
		for _, tag := range tags {
			query = query.Where("tags ILIKE ?", "%"+tag+"%")
		}
	}
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	if condition != "" {
		query = query.Where("condition = ?", condition)
	}
	if !showEmptyStock {
		query = query.Where("stock > 0")
	}

	var total int64
	query.Count(&total)

	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    products,
		"meta": gin.H{
			"limit":  limit,
			"offset": offset,
			"total":  total,
		},
	})
}

func GetProductDetail(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	var seller models.User
	if err := config.DB.First(&seller, product.UserID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load seller"})
		return
	}

	var accounts []models.BankAccount
	if err := config.DB.Where("user_id = ?", seller.ID).Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load bank accounts"})
		return
	}

	var purchaseCount int64
	config.DB.Model(&models.Payment{}).Where("product_id = ?", product.ID).Count(&purchaseCount)

	var productSoldTotal int64
	config.DB.Model(&models.Payment{}).Where("seller_id = ?", seller.ID).Count(&productSoldTotal)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": gin.H{
			"product": gin.H{
				"productId":      product.ID,
				"name":           product.Name,
				"price":          product.Price,
				"imageUrl":       product.ImageURL,
				"stock":          product.Stock,
				"condition":      product.Condition,
				"tags":           strings.Split(product.Tags, ","),
				"isPurchaseable": product.IsPurchaseable,
				"purchaseCount":  purchaseCount,
			},
			"seller": gin.H{
				"name":             seller.Name,
				"productSoldTotal": productSoldTotal,
				"bankAccounts":     accounts,
			},
		},
	})
}
