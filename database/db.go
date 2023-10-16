package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import the PostgreSQL driver
)

// InitDB initializes a database connection and returns a pointer to the DB instance.
func InitDB() (*gorm.DB, error) {
	// Read database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	// Open a connection to the database
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Set the table name singularization to false (optional)
	db.SingularTable(true)

	// Automigrate your models here if needed
	// For example, if you have an Employee model:
	// db.AutoMigrate(&Employee{})

	return db, nil
}
