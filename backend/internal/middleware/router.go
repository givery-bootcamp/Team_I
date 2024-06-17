package middleware

import (
	"myapp/internal/external"
	"myapp/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	db := external.DB
	postRepository := repositories.NewPostRepository(db)
	userRepository := repositories.NewUserRepository(db)
	h := NewHandler(postRepository, userRepository)

	app.POST("/signin", h.PostSignin)
	authGroup := app.Group("/auth")
	authGroup.Use(AuthMiddleware)
	authGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "you are authorized"})
	})

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	app.GET("/posts", h.GetPosts)

	app.GET("/posts/:id", h.GetPostById)
}
