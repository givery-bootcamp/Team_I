package controllers

import (
	"errors"
	"fmt"
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

type PostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var ErrUserInfoNotFound = fmt.Errorf("user info not found")

func PostPost(ctx *gin.Context, usecase *usecases.CreatePostUsecase) {
	var post PostRequest
	if err := ctx.ShouldBindJSON(&post); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	userInfoAny, exists := ctx.Get("userInfo")
	if !exists {
		fmt.Println("User info not found")
		handleError(ctx, http.StatusBadRequest, ErrUserInfoNotFound)
		return
	}
	userInfo := userInfoAny.(map[string]any)
	userId := userInfo["Id"].(int)
	result, err := usecase.Execute(userId, post.Title, post.Body)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
