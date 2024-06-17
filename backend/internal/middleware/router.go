package middleware

import (
	"myapp/internal/external"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	db := external.DB
	postRepository := repositories.NewPostRepository(db)
	h := NewHandler(postRepository)

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	app.GET("/posts", h.GetPosts)

	app.GET("/posts/:id", h.GetPostById)
}
