package controllers

import (
	"errors"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context, usecase *usecases.ListPostUsecase) {
	result, err := usecase.Execute()
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else if result != nil {
		ctx.JSON(http.StatusOK, result)
	} else {
		handleError(ctx, http.StatusNotFound, errors.New("not found"))
	}
}

func GetPostById(ctx *gin.Context, usecase *usecases.GetPostByIdUsecase) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	// ctx.String(http.StatusOK, greetings[id])

	result, err := usecase.Execute(id)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else if result != nil {
		ctx.JSON(http.StatusOK, result)
	} else {
		handleError(ctx, http.StatusNotFound, errors.New("not found"))
	}
}

func DeletePost(ctx *gin.Context, usecase *usecases.DeletePostUsecase) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	// ctx.String(http.StatusOK, greetings[id])

	err = usecase.Execute(id)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "success"})
	}
}