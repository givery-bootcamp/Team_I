package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/external"
	"myapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	external.SetupDB()

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction())
	app.Use(middleware.Cors())
	middleware.SetupRoutes(app)

	app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
