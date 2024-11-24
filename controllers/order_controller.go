package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	DB *sql.DB
}

func NewOrderController(db *sql.DB) *OrderController {
	return &OrderController{DB: db}
}

// Membuat pesanan baru
func (o *OrderController) CreateOrder(c *gin.Context) {
	var order struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := o.DB.Exec(
		"INSERT INTO orders (product_id, quantity) VALUES (?, ?)",
		order.ProductID, order.Quantity,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

// Mengambil detail pesanan berdasarkan ID
func (o *OrderController) GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")

	var order struct {
		ID        int    `json:"id"`
		ProductID int    `json:"product_id"`
		Quantity  int    `json:"quantity"`
		OrderDate string `json:"order_date"`
	}

	err := o.DB.QueryRow(
		"SELECT id, product_id, quantity, order_date FROM orders WHERE id = ?",
		orderID,
	).Scan(&order.ID, &order.ProductID, &order.Quantity, &order.OrderDate)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
		}
		return
	}

	c.JSON(http.StatusOK, order)
}
