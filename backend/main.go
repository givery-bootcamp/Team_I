package main

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/external"
	"myapp/internal/middleware"
	"net/http"

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

	authGroup := app.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)
	authGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "you are authorized"})
	})

	app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
}
