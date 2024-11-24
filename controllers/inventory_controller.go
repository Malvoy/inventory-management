package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	DB *sql.DB
}

func NewInventoryController(db *sql.DB) *InventoryController {
	return &InventoryController{DB: db}
}

// Melihat tingkat stok untuk suatu produk
func (i *InventoryController) GetStock(c *gin.Context) {
	productID := c.Query("product_id")

	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var quantity int
	err := i.DB.QueryRow("SELECT quantity FROM inventory WHERE product_id = ?", productID).Scan(&quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found for the product"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID, "quantity": quantity})
}

// Memperbarui tingkat stok (menambah/mengurangi stok)
func (i *InventoryController) UpdateStock(c *gin.Context) {
	var request struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Quantity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must not be zero"})
		return
	}

	_, err := i.DB.Exec(
		"UPDATE inventory SET quantity = quantity + ? WHERE product_id = ?",
		request.Quantity, request.ProductID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
}
