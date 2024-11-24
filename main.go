package main

import (
	"database/sql"
	"fmt"
	"inventory-management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database configuration
	dbUser := "root"
	dbPassword := "Roma123!"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "inventory_management"

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Database connection established!")

	// Ensure upload directory exists
	uploadDir := "uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create uploads directory: %v", err)
		}
	}

	// Initialize Gin
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, db)

	// Run the server
	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
