package middleware

import (
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(app *gin.Engine, db *gorm.DB) {
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

	authGroup.POST("/comments", h.PostComment)
	authGroup.PUT("/comments/:id", h.PutComment)
	authGroup.DELETE("/comments/:id", h.DeleteComment)
}
