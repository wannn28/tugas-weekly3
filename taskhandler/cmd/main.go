package main

import (
	"log"
	"net/http"

	"TaskHandler/pkg/config"
	"TaskHandler/pkg/controllers"
	"TaskHandler/pkg/routes"
)

func main() {
	// Menghubungkan ke database
	dbConn := config.InitializeDB()
	controllers.Database = dbConn // Set database di controller

	// Menyiapkan rute
	router := routes.ConfigureRoutes()

	log.Println("Server is running on http://localhost:8080")
	// Jalankan server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
