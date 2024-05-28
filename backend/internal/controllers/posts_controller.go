package controllers

import (
	"errors"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	u := usecases.NewListPostUsecase(ctx)
	result, err := u.Execute()
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}

func GetPostById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "not an integer")
		return
	}
	// ctx.String(http.StatusOK, greetings[id])

	u := usecases.NewGetPostByIdUsecase(ctx)
	result, err := u.Execute(id)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
