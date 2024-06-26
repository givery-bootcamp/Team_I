package middleware

import (
	"myapp/internal/external"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	db := external.DB
	postRepository := repositories.NewPostRepository(db)
	userRepository := repositories.NewUserRepository(db)
	commentRepository := repositories.NewCommentRepository(db)
	h := NewHandler(postRepository, userRepository, commentRepository)

	app.POST("/signin", h.PostSignin)

	authGroup := app.Group("/")
	authGroup.Use(AuthMiddleware)

	app.POST("/signout", h.PostSignout)

	authGroup.GET("/user", h.GetUser)

	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	app.GET("/posts", h.GetPosts)

	app.GET("/posts/:id", h.GetPostById)

	// app.DELETE("/posts/:id", h.DeletePost)
	authGroup.DELETE("/posts/:id", h.DeletePost)
	authGroup.POST("/posts", h.PostPost)

	authGroup.PUT("/posts/:id", h.PutPostById)
}
