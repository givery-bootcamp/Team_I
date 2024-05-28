package controllers

import (
	"errors"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	u := usecases.NewPostUsecase(ctx)
	result, err := u.Execute()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
