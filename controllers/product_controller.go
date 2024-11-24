package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	DB *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{DB: db}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		ImagePath   string  `json:"image_path"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := p.DB.Exec(
		"INSERT INTO products (name, description, price, category, image_path) VALUES (?, ?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Category, product.ImagePath,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (p *ProductController) GetProducts(c *gin.Context) {
	category := c.Query("category")

	var rows *sql.Rows
	var err error

	if category != "" {
		rows, err = p.DB.Query("SELECT id, name, description, price, category, image_path FROM products WHERE category = ?", category)
	} else {
		rows, err = p.DB.Query("SELECT id, name, description, price, category, image_path FROM products")
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	var products []map[string]interface{}
	for rows.Next() {
		var id int
		var name, description, category, imagePath string
		var price float64
		if err := rows.Scan(&id, &name, &description, &price, &category, &imagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse products"})
			return
		}
		products = append(products, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"price":       price,
			"category":    category,
			"image_path":  imagePath,
		})
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	row := p.DB.QueryRow("SELECT id, name, description, price, category, image_path FROM products WHERE id = ?", id)

	var product struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		ImagePath   string  `json:"image_path"`
	}

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.ImagePath)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		}
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		ImagePath   string  `json:"image_path"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := p.DB.Exec(
		"UPDATE products SET name = ?, description = ?, price = ?, category = ?, image_path = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.Category, product.ImagePath, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := p.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (pc *ProductController) UploadImage(c *gin.Context) {
	productID := c.Param("id")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	filePath := fmt.Sprintf("uploads/%s-%s", productID, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	_, err = pc.DB.Exec("UPDATE products SET image_path = ? WHERE id = ?", filePath, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product image path"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file_path": filePath})
}


func (pc *ProductController) DownloadImage(c *gin.Context) {
	productID := c.Param("id")

	var imagePath string
	err := pc.DB.QueryRow("SELECT image_path FROM products WHERE id = ?", productID).Scan(&imagePath)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product image"})
		}
		return
	}

	c.File(imagePath)
}


