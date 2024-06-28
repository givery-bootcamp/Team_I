package main

import (
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/external/database"
	"myapp/internal/external/database/seed"
	"myapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.SetupDB()

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction())
	app.Use(middleware.Cors())
	middleware.SetupRoutes(app, database.DB)

	err := database.DB.AutoMigrate(&database.User{}, &database.Post{}, &database.Comment{}, &database.Intentions{})
	if err != nil {
		log.Printf("AutoMigration failed: %v", err)
		return
	}

	if err = seed.SeedDatabase(database.DB); err != nil {
		log.Fatal("failed to seeding database: ", err)
	}

	err = app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
	if err != nil {
		log.Fatalf("Failed to run the app: %v", err)
	}
}
