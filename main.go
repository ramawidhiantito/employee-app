package main

import (
	"employee_app/database"
	"employee_app/internal/app/models"
	"employee_app/internal/app/routes"
	"employee_app/pkg/server"
)

func main() {
	// Initialize the database connection
	db, err := database.InitDB() // You should define an InitDB function in your database package
	if err != nil {
		// Handle the error
		panic(err)
	}
	defer db.Close()

	//Run AutoMigration
	db.AutoMigrate(models.Employee{})

	r := server.New()
	routes.SetupRoutes(r, db)

	err = server.Run(r, ":8080")
	if err != nil {
		panic(err)
	}
}
