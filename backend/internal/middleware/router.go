package middleware

import (
	"myapp/internal/controllers"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/posts", func(ctx *gin.Context) {
		db := usecases.DB(ctx)
		postRepository := repositories.NewPostRepository(db)
		usecase := usecases.NewListPostUsecase(postRepository)
		controllers.GetPosts(ctx, usecase)
	})

	app.GET("/posts/:id", func(ctx *gin.Context) {
		db := usecases.DB(ctx)
		postRepository := repositories.NewPostRepository(db)
		getPostByIdUsecase := usecases.NewGetPostByIdUsecase(postRepository)
		controllers.GetPostById(ctx, getPostByIdUsecase)
	})
}
