package routes

import (
	"database/sql"
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	productController := controllers.NewProductController(db)
	inventoryController := controllers.NewInventoryController(db)
	orderController := controllers.NewOrderController(db)

	// Routes for products
	r.POST("/products", productController.CreateProduct)
	r.GET("/products", productController.GetProducts)
	r.GET("/products/:id", productController.GetProductByID)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.POST("/:id/image", productController.UploadImage)
	r.GET("/:id/image", productController.DownloadImage)


	// Routes for inventory
	r.GET("/inventory", inventoryController.GetStock)
	r.PUT("/inventory", inventoryController.UpdateStock)

	// Routes for orders
	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders/:id", orderController.GetOrderByID)
}
